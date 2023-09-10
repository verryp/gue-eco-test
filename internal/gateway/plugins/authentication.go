package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	PluginName      = "authentication"
	ConfigAuthURL   = "auth_url"
	GatewayName     = "krakend"
	TokenBearerType = "Bearer"

	HeaderAuthorization = "Authorization"
	HeaderGateway       = "X-Gateway"
	HeaderForwardedFor  = "X-Forwarded-For"
	HeaderContentType   = "Content-Type"
	HeaderContentJson   = "application/json"
	HeaderClientID      = "X-Client-Id"
	HeaderClientName    = "X-Client-Name"
	HeaderUserID        = "X-User-Id"
	HeaderUserName      = "X-User-Name"
	HeaderUserEmail     = "X-User-Email"

	PathURLV1Authorize = "/auth/v1/authorize"

	ResponseSuccess     = "success"
	ResponseAuthFailure = "error"

	HttpTimeout = 10 * time.Second
)

var (
	ExcludedPaths = []string{
		"/health",
	}

	AllowedTokenTypes = []string{TokenBearerType}
)

type (
	Authorization struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}

	Data struct {
		Client Client `json:"client"`
		User   *User  `json:"user"`
	}

	Client struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	User struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	Response struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Data    *Data  `json:"data,omitempty"`
	}
)

func init() {
	fmt.Println(fmt.Sprintf("plugin: %s loaded!", PluginName))
}

func main() {}

// HandlerRegisterer :nodoc:
var HandlerRegisterer = registerer(PluginName)

type registerer string

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(PluginName, r.registerHandlers)
}

func (r registerer) registerHandlers(ctx context.Context, extra map[string]interface{}, handler http.Handler) (http.Handler, error) {
	err := checkEnvironment(extra)
	if err != nil {
		panic(err.Error())
	}

	authenticationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if InArray(r.URL.Path, ExcludedPaths) {
			handler.ServeHTTP(w, r)
			return
		}

		if _, ok := r.Header[HeaderAuthorization]; !ok {
			http.Error(w, "access denied", http.StatusForbidden)
			return
		}

		authorization, status := getAuthorization(r.Header.Get(HeaderAuthorization), r.URL.Path)
		if ResponseSuccess != status {
			failureResponse(w, status)
			return
		}

		if authorization != nil && authorization.Type == TokenBearerType && authorization.Value != "" {
			forwardedIP := r.Header.Get(HeaderForwardedFor)

			authResponse, respStatus := validateToken(extra, authorization.Value, forwardedIP)
			if ResponseSuccess != respStatus {
				failureResponse(w, status)
				return
			}

			authData := authResponse.Data
			r.Header.Set(HeaderContentType, HeaderContentJson)
			r.Header.Set(HeaderClientID, strconv.Itoa(authData.Client.ID))
			r.Header.Set(HeaderClientName, authData.Client.Name)

			if authData.User != nil {
				r.Header.Set(HeaderUserID, authData.User.ID)
				r.Header.Set(HeaderUserName, authData.User.Name)
				r.Header.Set(HeaderUserEmail, authData.User.Email)
			}
		}

		r.Header.Set(HeaderGateway, GatewayName)

		handler.ServeHTTP(w, r)
	})

	return authenticationHandler, nil
}

func checkEnvironment(extra map[string]interface{}) error {
	_, ok := extra[ConfigAuthURL]
	if !ok {
		message := fmt.Sprintf("Incorrect authentication configuration URL")
		return errors.New(message)
	}

	return nil
}

func getAuthorization(authorization, urlPath string) (*Authorization, string) {
	authorizationIndex := strings.SplitN(authorization, " ", 2)

	if urlPath == PathURLV1Authorize {
		if len(authorizationIndex[0]) == 32 {
			return nil, ResponseSuccess
		}

		return nil, ResponseAuthFailure
	}

	if len(authorizationIndex) <= 1 || len(authorizationIndex) > 2 {
		return nil, ResponseAuthFailure
	}

	authType := strings.ToLower(authorizationIndex[0])
	authValue := authorizationIndex[1]
	isTypeAllowed := findSlice(AllowedTokenTypes, authType)

	if !isTypeAllowed {
		return nil, ResponseAuthFailure
	}

	if authValue == `` {
		return nil, ResponseAuthFailure
	}

	auth := &Authorization{
		Type:  authType,
		Value: authValue,
	}

	return auth, ResponseSuccess
}

func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

func findSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func authenticationFailureResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	bodyString, _ := json.Marshal(Response{
		Name:    ResponseAuthFailure,
		Message: "authentication failure",
	})
	_, _ = fmt.Fprintln(w, string(bodyString))
	return
}

func failureResponse(w http.ResponseWriter, status string) {
	switch status {
	case ResponseAuthFailure:
		authenticationFailureResponse(w)
		return
	}
}

func validateToken(extra map[string]interface{}, token, forwardedIP string) (*Response, string) {
	client := &http.Client{Timeout: HttpTimeout}

	BaseURL, _ := extra[ConfigAuthURL]
	authenticationURL := fmt.Sprintf("%s/%s", BaseURL, "/v1/auth/token")
	req, err := http.NewRequest(http.MethodPost, authenticationURL, nil)
	if err != nil {
		return nil, ResponseAuthFailure
	}
	req.Header.Set(HeaderContentType, HeaderContentJson)
	req.Header.Set(HeaderAuthorization, fmt.Sprintf("%s %s", TokenBearerType, token))
	req.Header.Set(HeaderForwardedFor, forwardedIP)

	resp, err := client.Do(req)
	if err != nil {
		return nil, ResponseAuthFailure
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ResponseAuthFailure
	}

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ResponseAuthFailure
	}

	stringBody := string(respBodyBytes)
	data := &Response{}
	err = json.Unmarshal([]byte(stringBody), data)
	if err != nil || data.Data == nil {
		return nil, ResponseAuthFailure
	}

	return data, ResponseSuccess
}

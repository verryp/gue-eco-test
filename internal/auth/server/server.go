package server

import (
	"fmt"

	"github.com/verryp/gue-eco-test/internal/auth/authenticator"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/driver"
	"github.com/verryp/gue-eco-test/internal/auth/handler"
	"github.com/verryp/gue-eco-test/internal/auth/model"
	"github.com/verryp/gue-eco-test/internal/auth/repository"
	"github.com/verryp/gue-eco-test/internal/auth/service"
	"github.com/verryp/gue-eco-test/pkg/generator"
	"gopkg.in/gorp.v2"
)

func Start() {
	conf, err := common.NewConfig()
	if err != nil {
		panic(err)
	}

	logger := common.NewLogger(&conf.Log)

	opt := &common.Option{
		Config: conf,
		Log:    logger,
	}

	db, err := driver.NewMySQLDatabase(conf.DB)
	if err != nil {
		logger.Err(err).Msg("DB connection error")
		panic(err)
	}

	initDB(db)

	repos := wiringServerRepository(&repository.Option{
		Option: opt,
		DB:     db,
	})

	svc := wiringServerService(&service.Option{
		Option:        opt,
		Repository:    repos,
		Authenticator: authenticator.NewJWTAuthenticator(opt, repos),
	})

	handlers := &handler.Option{
		Option:  opt,
		Service: svc,
	}

	svr := NewRouter(conf, handlers)

	addr := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)
	logger.Info().Msgf("server is serving at %s", addr)

	err = svr.Route().Listen(addr)
	if err != nil {
		logger.Err(err).Msgf("Server failed to serve at %s", addr)
	}
}

func wiringServerRepository(opt *repository.Option) *repository.Repository {
	userRepo := repository.NewUserRepo(opt)
	clientRepo := repository.NewClientRepo(opt)

	return &repository.Repository{
		User:   userRepo,
		Client: clientRepo,
	}
}

func wiringServerService(opt *service.Option) *service.Service {
	// bootstrapping
	generator.New(1)

	healthCheck := service.NewHealthCheckService(opt)
	auth := service.NewAuthService(opt)

	return &service.Service{
		HealthCheck: healthCheck,
		Auth:        auth,
	}
}

func initDB(db *gorp.DbMap) {
	db.AddTableWithName(model.User{}, "users").SetKeys(false, "id")
	db.AddTableWithName(model.Client{}, "clients").SetKeys(true, "id")
}

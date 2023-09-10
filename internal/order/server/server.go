package server

import (
	"fmt"

	"github.com/verryp/gue-eco-test/internal/order/common"
	"github.com/verryp/gue-eco-test/internal/order/connector/product"
	"github.com/verryp/gue-eco-test/internal/order/driver"
	"github.com/verryp/gue-eco-test/internal/order/handler"
	"github.com/verryp/gue-eco-test/internal/order/model"
	"github.com/verryp/gue-eco-test/internal/order/repository"
	"github.com/verryp/gue-eco-test/internal/order/service"
	"github.com/verryp/gue-eco-test/pkg/generator"
	"github.com/verryp/gue-eco-test/pkg/httpclient"
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

	// clients
	productRest := httpclient.NewRestClient(conf.Dependency.Product.BaseURL)
	productAPI := product.NewProductAPI(conf, productRest)

	svc := wiringServerService(&service.Option{
		Option:     opt,
		Repository: repos,
		ProductAPI: productAPI,
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
	orderRepo := repository.NewOrderRepo(opt)
	orderDetailRepo := repository.NewOrderDetailRepo(opt)

	return &repository.Repository{
		Order:       orderRepo,
		OrderDetail: orderDetailRepo,
	}
}

func wiringServerService(opt *service.Option) *service.Service {
	// bootstrapping
	generator.New(1)

	healthCheck := service.NewHealthCheckService(opt)
	cartSvc := service.NewCartService(opt)
	orderSvc := service.NewOrderService(opt)

	return &service.Service{
		HealthCheck: healthCheck,
		Cart:        cartSvc,
		Order:       orderSvc,
	}
}

func initDB(db *gorp.DbMap) {
	db.AddTableWithName(model.Order{}, "orders").SetKeys(false, "id")
	db.AddTableWithName(model.OrderHistory{}, "order_histories").SetKeys(false, "id")
	db.AddTableWithName(model.OrderDetail{}, "order_details").SetKeys(true, "id")
}

package server

import (
	"fmt"

	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/driver"
	"github.com/verryp/gue-eco-test/internal/product/handler"
	"github.com/verryp/gue-eco-test/internal/product/model"
	"github.com/verryp/gue-eco-test/internal/product/repository"
	"github.com/verryp/gue-eco-test/internal/product/service"
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
		Option:     opt,
		Repository: repos,
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
	itemRepo := repository.NewItemsRepo(opt)
	itemQuotaRepo := repository.NewItemQuotaRepo(opt)

	return &repository.Repository{
		Item:      itemRepo,
		ItemQuota: itemQuotaRepo,
	}
}

func wiringServerService(opt *service.Option) *service.Service {
	// bootstrapping
	generator.New(1)

	healthCheck := service.NewHealthCheckService(opt)
	items := service.NewItemService(opt)

	return &service.Service{
		HealthCheck: healthCheck,
		Item:        items,
	}
}

func initDB(db *gorp.DbMap) {
	db.AddTableWithName(model.Item{}, "items").SetKeys(false, "id")
	db.AddTableWithName(model.Quota{}, "item_quotas").SetKeys(true, "id")
}

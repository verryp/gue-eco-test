package server

import (
	"fmt"

	"github.com/brainlabs/snowflake"
	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/driver"
	"github.com/verryp/gue-eco-test/internal/product/handler"
	"github.com/verryp/gue-eco-test/internal/product/model"
	"github.com/verryp/gue-eco-test/internal/product/repository"
	"github.com/verryp/gue-eco-test/internal/product/service"
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
	return &repository.Repository{
		Item: itemRepo,
	}
}

func wiringServerService(opt *service.Option) *service.Service {
	// bootstrapping
	sf, _ := snowflake.NewNode(0)

	healthCheck := service.NewHealthCheckService(opt)
	items := service.NewItemService(opt, sf.Generate().Int64())

	return &service.Service{
		HealthCheck: healthCheck,
		Item:        items,
	}
}

func initDB(db *gorp.DbMap) {
	db.AddTableWithName(model.Item{}, "items").SetKeys(false, "id")
}

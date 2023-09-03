package driver

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/consts"
	"gopkg.in/gorp.v2"
)

func NewMySQLDatabase(opt common.DB) (*gorp.DbMap, error) {
	param := url.Values{}
	param.Add("parseTime", "True")
	param.Add("loc", consts.TimezoneAsiaJakarta)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Port,
		opt.Name,
		param.Encode(),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	dbMap := &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	return dbMap, nil
}

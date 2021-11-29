package data

import (
	"github.com/jmoiron/sqlx"
	"myblog/internal/conf"
)

// Data .
type Data struct {
	db  *sqlx.DB
}

func NewData(conf *conf.Data) (*Data, error) {
	database, err := sqlx.Open(
		conf.Database.Driver,
		conf.Database.Source,
		)
	if err != nil {
		return nil, err
	}
	return &Data{db: database}, nil
}
package api

import (
	"company/db"
	"company/db/pg"
	util "company/libgo"
)

var c *Client

type Client struct {
	Pg *pg.PgCompany

	CompanyDb *db.CompanyDbC
}

func Init() *Client {
	c = &Client{}
	c.Pg = pg.InitPgCompany(util.MustOsGetEnv("DB_URL"))

	c.CompanyDb = db.Init(
		c.Pg,
	)

	return c
}

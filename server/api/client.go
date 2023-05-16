package api

import (
	"company/db"
	"company/db/pg"
)

var c *Client

type Client struct {
	Pg *pg.PgCompany

	CompanyDb *db.CompanyDbC
}

func Init() *Client {
	c = &Client{}
	c.Pg = pg.InitPgCompany("host=localhost port=5432 user=mslm password=mslm dbname=company connect_timeout=2")

	c.CompanyDb = db.Init(
		c.Pg,
	)

	return c
}

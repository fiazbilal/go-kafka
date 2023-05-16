package api

import (
	"company/db/pg"
)

var c *Client

type Client struct {
	Pg *pg.PgCompany
}

func Init() *Client {
	c = &Client{}
	c.Pg = pg.InitPgCompany("host=localhost port=5432 user=mslm password=mslm dbname=company connect_timeout=2")

	return c
}

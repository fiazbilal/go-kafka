package db

import (
	"company/db/pg"
)

type CompanyDbC struct {
	Pg *pg.PgCompany
}

func Init(
	Pg *pg.PgCompany,
) *CompanyDbC {
	return &CompanyDbC{
		Pg: Pg,
	}
}

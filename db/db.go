package db

import "github.com/google/uuid"

type CompanyCreate struct {
	Id          uuid.UUID
	Name        string
	Description string
	Employees   int
	Registered  bool
	Type        string
}

func (c *CompanyDbC) CreateCompany(company *CompanyCreate) error {
	_, err := c.Pg.Exec(
		`INSERT INTO company (
            id,
            name,
            description,
            employees,
            registered,
            type
        ) VALUES (
            $1,
            $2,
            $3,
            $4,
            $5,
            $6
        )`,
		company.Id,
		company.Name,
		company.Description,
		company.Employees,
		company.Registered,
		company.Type,
	)
	return err
}

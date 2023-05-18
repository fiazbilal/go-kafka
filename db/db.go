package db

import (
	"github.com/google/uuid"
)

type CompanyType string

var (
	COMPANY_TYPE_NULL           CompanyType = ""
	COMPANY_TYPE_CORPORATIONS   CompanyType = "CORPORATIONS"
	COMPANY_TYPE_NONPROFIT      CompanyType = "NONPROFIT"
	COMPANY_TYPE_COOPERATIVE    CompanyType = "COOPERATIVE"
	COMPANY_TYPE_SOLE           CompanyType = "SOLE"
	COMPANY_TYPE_PROPRIETORSHIP CompanyType = "PROPRIETORSHIP"
)

// Safely scans db value of type *MJobStatus.
func ScanCompanyType(src *CompanyType) (dest CompanyType) {
	if src == nil {
		dest = COMPANY_TYPE_NULL
		return
	}
	dest = *src
	return
}

type CompanyCreateTup struct {
	Name        string
	Description string
	Employees   int
	Registered  bool
	Type        string
}

func (c *CompanyDbC) CreateCompany(company *CompanyCreateTup) (uuid.UUID, error) {
	var id uuid.UUID
	params := []any{company.Name, company.Description, company.Employees, company.Registered, company.Type}
	qStr := `
	INSERT INTO companies (
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
		$5
	) RETURNING id`

	err := c.Pg.QueryRow(qStr, params...).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (c *CompanyDbC) DeleteCompany(id uuid.UUID) error {
	_, err := c.Pg.Exec("DELETE FROM companies WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

type CompanyInfoTup struct {
	Id          uuid.UUID
	Name        string
	Description string
	Employees   int
	Registered  bool
	Type        string
}

func (c *CompanyDbC) GetCompanyById(id uuid.UUID) (*CompanyInfoTup, error) {
	company := &CompanyInfoTup{}
	err := c.Pg.QueryRow("SELECT id, name, description, employees, registered, type FROM companies WHERE id=$1", id).Scan(&company.Id, &company.Name, &company.Description, &company.Employees, &company.Registered, &company.Type)
	if err != nil {
		return nil, err
	}

	return company, nil
}

type CompanyUpdateTup struct {
	Id          uuid.UUID
	Name        string
	Description string
	Employees   int
	Registered  bool
	Type        string
}

func (c *CompanyDbC) UpdateCompany(company CompanyUpdateTup) error {
	_, err := c.Pg.Exec(
		`UPDATE companies SET
            name=$2,
            description=$3,
            employees=$4,
            registered=$5,
            type=$6
        WHERE id=$1`,
		company.Id,
		company.Name,
		company.Description,
		company.Employees,
		company.Registered,
		company.Type,
	)

	return err
}

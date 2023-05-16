package api

import (
	"company/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type CompanyUpdateReq struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Desc          string    `json:"description,omitempty"`
	NoOfEmployees int       `json:"employees"`
	Registered    bool      `json:"registered"`
	Type          string    `json:"type"`
}

func CompanyUpdate(req *Req, resp *Resp) {
	// Parse req body.
	defer req.Body.Close()
	rawBody, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("failed to parse req body: %v", err)
		resp.Send(RC_E_NO_BODY)
		return
	}

	body := &CompanyUpdateReq{}
	if err := json.Unmarshal(rawBody, body); err != nil {
		fmt.Println("failed to parse JSON object: %v", err)
		resp.Send(RC_E_MALFORMED)
		return
	}

	// Get the current company record from the database.
	company, err := c.CompanyDb.GetCompanyById(body.Id)
	if err != nil {
		fmt.Println("failed to get company by Id=%v: %v", body.Id, err)
		resp.Send(http.StatusInternalServerError)
		return
	}

	companyTup := db.CompanyUpdateTup{
		Id:          company.Id,
		Name:        company.Name,
		Description: company.Description,
		Employees:   company.Employees,
		Registered:  company.Registered,
		Type:        company.Type,
	}

	// Update the company fields with the new values.
	companyTup.Id = body.Id
	if body.Name != "" {
		companyTup.Name = body.Name
	}
	if body.Desc != "" {
		companyTup.Description = body.Desc
	}
	if body.NoOfEmployees != 0 {
		companyTup.Employees = body.NoOfEmployees
	}
	if body.Registered != company.Registered {
		companyTup.Registered = body.Registered
	}
	if body.Type != company.Type {
		companyTup.Type = body.Type
	}

	// Update the company record in the database.
	err = c.CompanyDb.UpdateCompany(companyTup)
	if err != nil {
		fmt.Println("failed to update company by Id=%v: %v", body.Id, err)
		resp.Send(http.StatusInternalServerError)
		return
	}

	// Send.
	resp.Send(RC_COMPANY_UPDATE)
}

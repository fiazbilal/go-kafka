package api

import (
	"company/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type CompanyCreateReq struct {
	Name          string `json:"name"`
	Desc          string `json:"description,omitempty"`
	NoOfEmployees int    `json:"employees"`
	Registered    bool   `json:"registered"`
	Type          string `json:"type"`
}

type CompanyCreateResp struct {
	Id uuid.UUID `json:"id"`
}

func CompanyCreate(req *Req, resp *Resp) {
	// Parse req body.
	defer req.Body.Close()
	rawBody, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("failed to parse req body: %v", err)
		resp.Send(RC_E_NO_BODY)
		return
	}

	body := &CompanyCreateReq{}
	if err := json.Unmarshal(rawBody, body); err != nil {
		fmt.Println("failed to parse JSON object: %v", err)
		resp.Send(RC_E_MALFORMED)
		return
	}

	companyId := uuid.New()
	companyTup := &db.CompanyCreateTup{
		Id:          companyId,
		Name:        body.Name,
		Description: body.Desc,
		Employees:   body.NoOfEmployees,
		Registered:  body.Registered,
		Type:        body.Type,
	}

	// Company create.
	err = c.CompanyDb.CreateCompany(companyTup)
	if err != nil {
		fmt.Println("failed to add company by Id=%v: %v", companyId, err)
		resp.Send(http.StatusInternalServerError)
		return
	}

	// Send.
	resp.SendData(RC_COMPANY_CREATE, &CompanyCreateResp{
		Id: companyId,
	})
}

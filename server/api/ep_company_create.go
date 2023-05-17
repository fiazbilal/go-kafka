package api

import (
	"company/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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
		fmt.Printf("failed to parse req body: %v\n", err)
		resp.Send(RC_E_NO_BODY)
		return
	}

	body := &CompanyCreateReq{}
	if err := json.Unmarshal(rawBody, body); err != nil {
		fmt.Printf("failed to parse JSON object: %v\n", err)
		resp.Send(RC_E_MALFORMED)
		return
	}

	companyTypeStr := strings.ToUpper(body.Type)
	var companyType db.CompanyType
	if companyTypeStr == "CORPORATIONS" {
		companyType = db.COMPANY_TYPE_CORPORATIONS
	} else if companyTypeStr == "NONPROFIT" {
		companyType = db.COMPANY_TYPE_NONPROFIT
	} else if companyTypeStr == "COOPERATIVE" {
		companyType = db.COMPANY_TYPE_COOPERATIVE
	} else if companyTypeStr == "SOLE" {
		companyType = db.COMPANY_TYPE_SOLE
	} else if companyTypeStr == "PROPRIETORSHIP" {
		companyType = db.COMPANY_TYPE_PROPRIETORSHIP
	} else {
		resp.Send(RC_E_COMPANY_CREATE_INVALID_TYPE)
		return
	}

	companyId := uuid.New()
	companyTup := &db.CompanyCreateTup{
		Id:          companyId,
		Name:        body.Name,
		Description: body.Desc,
		Employees:   body.NoOfEmployees,
		Registered:  body.Registered,
		Type:        string(companyType),
	}

	// Company create.
	err = c.CompanyDb.CreateCompany(companyTup)
	if err != nil {
		fmt.Printf("failed to add company by Id=%v: %v\n", companyId, err)
		resp.Send(http.StatusInternalServerError)
		return
	}

	// Send.
	resp.SendData(RC_COMPANY_CREATE, &CompanyCreateResp{
		Id: companyId,
	})
}

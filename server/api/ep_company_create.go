package api

import (
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

	fmt.Printf("I am here")
	// Parse req body.
	defer req.Body.Close()
	rawBody, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Errorf("failed to parse req body: %v", err)
		resp.Send(RC_E_NO_BODY)
		return
	}

	body := &CompanyCreateReq{}
	if err := json.Unmarshal(rawBody, body); err != nil {
		fmt.Errorf("failed to parse JSON object: %v", err)
		resp.Send(RC_E_MALFORMED)
		return
	}

	companyId := uuid.New()

	// Send.
	resp.SendStatus(RC_COMPANY_CREATE, &CompanyCreateResp{
		Id: companyId,
	}, http.StatusAccepted)
}

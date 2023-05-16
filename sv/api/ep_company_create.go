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

func CompanyCreate(resp http.ResponseWriter, req *http.Request) {
	// Parse req body.
	defer req.Body.Close()
	rawBody, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Errorf("failed to parse req body: %v", err)
		resp.Send(RC_E_NO_BODY)
		return
	}

	body := &MJobCreateReq{}
	if err := json.Unmarshal(rawBody, body); err != nil {
		fmt.Errorf("failed to parse JSON object: %v", err)
		resp.Send(RC_E_MALFORMED)
		return
	}

	jobId := uuid.New()
	jobTup := &mjob_db.MJobTup{
		Id:     jobId,
		Name:   body.Name,
		Status: string(mjob_db.MJOB_STATUS_READY),
	}

	if err := c.MJobDb.JobCreate(jobTup); err != nil {
		fmt.Errorf("failed to create job: %v", err)
		resp.Send(http.StatusInternalServerError)
		return
	}

	// Send.
	resp.SendStatus(RC_MJOB_CREATE, &MJobCreateResp{
		JobId: jobId,
	}, http.StatusAccepted)
}

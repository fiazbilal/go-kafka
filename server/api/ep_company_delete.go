package api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func CompanyDelete(req *Req, resp *Resp) {
	// get query params.
	qVals := req.URL.Query()

	// Parse out id variable.
	companyIdStr := qVals.Get("uuid")
	companyId, err := uuid.Parse(companyIdStr)
	if err != nil {
		fmt.Println(
			"failed to parse companyIdStr=%s: %v",
			companyIdStr, err,
		)
		resp.Send(http.StatusBadRequest)
		return
	}

	// Company create.
	err = c.CompanyDb.DeleteCompany(companyId)
	if err != nil {
		fmt.Println("failed to delete company by Id=%v: %v", companyId, err)
		resp.Send(http.StatusInternalServerError)
		return
	}

	// Send.
	resp.Send(RC_COMPANY_DELETE)
}

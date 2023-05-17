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
		fmt.Printf(
			"failed to parse companyIdStr=%s: %v\n",
			companyIdStr, err,
		)
		resp.Send(http.StatusBadRequest)
		return
	}

	// Company create.
	err = c.CompanyDb.DeleteCompany(companyId)
	if err != nil {
		fmt.Printf("failed to delete company by Id=%v: %v\n", companyId, err)
		resp.Send(http.StatusInternalServerError)
		return
	}

	// Send.
	resp.Send(RC_COMPANY_DELETE)
}

package api

import "company/server"

const (
	// Reserved
	RC_E_NO_BODY   server.RespCode = 2999
	RC_E_MALFORMED server.RespCode = 2998
	RC_E_RATELIMIT server.RespCode = 2997

	// POST /api/v1/company/create
	RC_COMPANY_CREATE                server.RespCode = 1000
	RC_E_COMPANY_CREATE_INVALID_TYPE server.RespCode = 2000

	// DELETE /api/v1/company/delete
	RC_COMPANY_DELETE server.RespCode = 1000

	// GET /api/v1/company/info
	RC_COMPANY_INFO server.RespCode = 1000

	// PATCH /api/v1/company/update
	RC_COMPANY_UPDATE                server.RespCode = 1000
	RC_E_COMPANY_UPDATE_INVALID_TYPE server.RespCode = 2000
)

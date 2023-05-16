package api

import "company/server"

const (
	// Reserved
	RC_E_NO_BODY   server.RespCode = 2999
	RC_E_MALFORMED server.RespCode = 2998
	RC_E_RATELIMIT server.RespCode = 2997

	// POST /api/v1/company/create
	RC_COMPANY_CREATE server.RespCode = 1000
)

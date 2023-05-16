package api

const (
	// Reserved
	RC_E_NO_BODY   sv.RespCode = 2999
	RC_E_MALFORMED sv.RespCode = 2998
	RC_E_RATELIMIT sv.RespCode = 2997

	// POST /api/v1/company/create
	RC_COMPANY_CREATE sv.RespCode = 1000
)

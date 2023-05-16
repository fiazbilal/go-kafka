package server

import (
	"net/http"
)

type Req struct {
	*http.Request
}

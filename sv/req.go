package sv

import (
	"net/http"
)

type Req struct {
	*http.Request
}

package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Req struct {
	sv.Req
}

type Resp struct {
	sv.Resp
}

func Main() {
	listenUrl := "localhost:8000"
	srvr := http.NewServeMux()
	srvr.HandleFunc("/", MainHandler)
	log.Printf("listening on %s", listenUrl)
	log.Fatal(http.ListenAndServe(listenUrl, srvr))
}

func MainHandler(w http.ResponseWriter, httpReq *http.Request) {
	req := &Req{Req: sv.Req{Request: httpReq}}
	resp := &Resp{Resp: sv.Resp{ResponseWriter: w}}

	url := req.URL.Path
	if url != "/" {
		url = strings.TrimRight(req.URL.Path, "/")
	}

	fmt.Printf("%v %v\n", req.Method, url)

	// Routing & authorization.
	// revive:disable
	if req.Method == "GET" {
	} else if req.Method == "POST" {
		if url == "/api/v1/company/create" {
			CompanyCreate(req, resp)
			return
		} else {
			resp.Send(http.StatusNotFound)
			return
		}
	} else if req.Method == "PATCH" {
	} else if req.Method == "DELETE" {
	} else {
		resp.Send(http.StatusMethodNotAllowed)
		return
	}
}

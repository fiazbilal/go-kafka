package api

import (
	"company/server"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Req struct {
	server.Req
}

type Resp struct {
	server.Resp
}

func Main(init bool) {
	if init {
		flag.Parse()
		c = Init()
	}

	listenUrl := "localhost:8000"
	srvr := http.NewServeMux()
	srvr.HandleFunc("/", MainHandler)
	log.Printf("listening on %s", listenUrl)
	log.Fatal(http.ListenAndServe(listenUrl, srvr))
}

func MainHandler(w http.ResponseWriter, httpReq *http.Request) {
	req := &Req{Req: server.Req{Request: httpReq}}
	resp := &Resp{Resp: server.Resp{ResponseWriter: w}}

	url := req.URL.Path
	if url != "/" {
		url = strings.TrimRight(req.URL.Path, "/")
	}

	fmt.Println(req.Method, url)

	// Routing & authorization.
	// revive:disable
	if req.Method == "GET" {
		if url == "/api/v1/company/info" {
			CompanyInfo(req, resp)
			return
		} else {
			resp.Send(http.StatusNotFound)
			return
		}
	} else if req.Method == "POST" {
		if url == "/api/v1/company/create" {
			CompanyCreate(req, resp)
			return
		} else {
			resp.Send(http.StatusNotFound)
			return
		}
	} else if req.Method == "PATCH" {
		if url == "/api/v1/company/update" {
			CompanyUpdate(req, resp)
			return
		} else {
			resp.Send(http.StatusNotFound)
			return
		}
	} else if req.Method == "DELETE" {
		if url == "/api/v1/company/delete" {
			CompanyDelete(req, resp)
			return
		} else {
			resp.Send(http.StatusNotFound)
			return
		}
	} else {
		resp.Send(http.StatusMethodNotAllowed)
		return
	}
}

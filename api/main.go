package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Main() {
	listenUrl := "localhost:8000"
	srvr := http.NewServeMux()
	srvr.HandleFunc("/", MainHandler)
	log.Printf("listening on %s", listenUrl)
	log.Fatal(http.ListenAndServe(listenUrl, srvr))
}

func MainHandler(w http.ResponseWriter, req *http.Request) {

	url := req.URL.Path
	if url != "/" {
		url = strings.TrimRight(req.URL.Path, "/")
	}

	fmt.Printf("URL: %s", url)

}

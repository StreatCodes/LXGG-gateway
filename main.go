package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
)

//LXDClient is our global connection to the REST API
var LXDClient http.Client

//LXGGConfig is our global reference to lxgg.toml
var LXGGConfig Config

func main() {
	//Load configuration and setup global vars
	LXGGConfig = loadConfig()
	LXDClient = createClient(LXGGConfig)

	//Setup routes
	r := chi.NewRouter()
	r.Post("/login", loginHandler)
	r.HandleFunc("/", proxyHandler)

	//Start http server
	fmt.Printf("Starting server on %s\n", LXGGConfig.Addr)
	err := http.ListenAndServe(LXGGConfig.Addr, r)
	if err != nil {
		panic(err)
	}
}

func proxyHandler(w http.ResponseWriter, req *http.Request) {
	//Modify request url to include LXD address
	newURL, err := url.Parse(LXGGConfig.LXDAddr + req.URL.String())
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), 500)
		return
	}

	//Modify some settings
	req.RequestURI = ""
	req.URL = newURL
	req.Host = LXGGConfig.LXDAddr

	fmt.Printf("Proxying request %s\n", req.URL.String())

	//Make request
	resp, err := LXDClient.Do(req)
	if err != nil {
		http.Error(w, "ERROR: "+err.Error(), 500)
		return
	}

	//Return LXD response as our response
	io.Copy(w, resp.Body)

	// requestPath := "/"
	// POST
	// response, err := LXDClient.Get("http://unix" + requestPath)
	// GET
	// 	response, err = httpc.Post("http://unix"+flag.Args()[1], "application/octet-stream", strings.NewReader(*post))

}

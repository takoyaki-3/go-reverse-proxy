package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"golang.org/x/crypto/acme/autocert"
	json "github.com/takoyaki-3/go-json"
)

type Domain struct {
	Domain	string	`domain`
	Host		string	`host`
	Scheme	string	`scheme`
}

var mapDomains map[string]Domain

func main() {
	// reverse-proxy
	rp := &httputil.ReverseProxy{Director: func(request *http.Request) {
		host := request.Host

		url := *request.URL
		url.Scheme = mapDomains[host].Scheme
		url.Host = mapDomains[host].Host

		if request.Body != nil {
			buffer, err := ioutil.ReadAll(request.Body)
			if err != nil {
				log.Fatal(err.Error())
			}
			req, err := http.NewRequest(request.Method, url.String(), bytes.NewBuffer(buffer))
			if err != nil {
				log.Fatal(err.Error())
			}
			req.Header = request.Header
			*request = *req	
		} else {
			req, err := http.NewRequest(request.Method, url.String(), nil)
			if err != nil {
				log.Fatal(err.Error())
			}
			req.Header = request.Header
			*request = *req
		}
	}}

	// initialize
	domains := []Domain{}
	if err:=json.LoadFromPath("./conf.json",&domains);err!=nil{
		log.Fatal(err)
	}
	mapDomains = map[string]Domain{}
	domainStrs := []string{}
	for _,domain := range domains{
		mapDomains[domain.Domain] = domain
		domains = append(domains,domain)
	}

	// start server
	log.Fatal(http.Serve(autocert.NewListener(domainStrs...), rp))
}

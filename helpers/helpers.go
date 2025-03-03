package helpers

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

var contentyTypeHeaderJson = []byte("application/json")

func NewReq(url string, body *[]byte, headers *map[string]string, debug ...bool) (*fasthttp.Response, error) {
	readTimeout, _ := time.ParseDuration("1m30s")
	writeTimeout, _ := time.ParseDuration("1m30s")
	maxIdleConnDuration, _ := time.ParseDuration("5m")
	tlsConf := &tls.Config{InsecureSkipVerify: true}
	dial := (&fasthttp.TCPDialer{Concurrency: 100, DNSCacheDuration: time.Hour}).Dial
	client := fasthttp.Client{
		Name:                          "tanda",
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		TLSConfig:                     tlsConf,
		Dial:                          dial,
	}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	req.SetRequestURI(url)
	req.Header.SetContentTypeBytes(contentyTypeHeaderJson)
	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}
	//GET REQUEST.
	if body == nil {
		req.Header.SetMethod(fasthttp.MethodGet)

	} else {
		//POST REQUEST
		req.Header.SetMethod(fasthttp.MethodPost)
		req.SetBodyRaw(*body)
	}
	err := client.Do(req, resp)
	if err != nil {
		fmt.Printf("<ERROR ->>: %s\n", err)

	}
	if len(debug) != 0 {
		if debug[0] {
			fmt.Printf("-------------REQUEST START------------\n")
			if strings.Contains(req.URI().String(), "oauth2/token") {
				fmt.Printf("[URL]: Has cred\n")
			} else {
				fmt.Printf("[URL]: %s\n", req.URI().String())
				// fmt.Printf("[HEADERS]: %s\n", req.Header.String())
				fmt.Printf("[REQ BODY]: %s\n", req.Body())
			}

			fmt.Printf("[CODE]: %d\n", resp.StatusCode())
			// if respnse body contains access_token dont log
			if strings.Contains(string(resp.Body()), "access_token") {
				fmt.Printf("resp.Body() has access_token\n")
			} else {
				fmt.Printf("[RESPONSE]: %s\n", resp.Body())
			}
			fmt.Printf("-------------REQUEST END------------\n")
		}

	}
	if err != nil {
		return nil, err
	}

	// RELEASE RESOURCES.
	fasthttp.ReleaseRequest(req)
	return resp, nil
}

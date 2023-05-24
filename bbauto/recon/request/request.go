package request

import (
	"io"
	"log"
	"net/http"
	"os"
	"recon/config"
	"strings"
)

// request of resource
func Req(c config.ResourceConfig, domain string) {

	// Replace the "{domain}" placeholder with the actual domain
	requestURL := strings.Replace(c.RequestUrl, c.RequestVar, domain, 1)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(1)
	}

	log.Println("client: got response!")
	log.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(1)
	}
	log.Printf("client: response body: %s\n", resBody)

}

package main

import (
	"context"
	"net/http"
	"time"

	"git.sr.ht/~spc/go-log"
	zest "github.com/content-services/zest/release/v3"
)

func connectPulp() {
	configuration := zest.NewConfiguration()
	apiClient := zest.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.StatusRead(context.Background()).Execute()
	if err != nil {
		log.Fatalf("Error when calling `StatusAPI.StatusRead``: %v\n", err)
		log.Fatalf("Full HTTP response: %v\n", r)
	}
	// response from `StatusRead`: StatusResponse
	log.Infof("Response from `StatusAPI.StatusRead`: %v\n", resp)
}

func listUpstreamPulps() {

	ctx2 := context.WithValue(context.Background(), zest.ContextServerIndex, 0)
	timeout := 60 * time.Second
	transport := &http.Transport{ResponseHeaderTimeout: timeout}
	httpClient := http.Client{Transport: transport, Timeout: timeout}

	pulpConfig := zest.NewConfiguration()
	pulpConfig.HTTPClient = &httpClient
	pulpConfig.Servers = zest.ServerConfigurations{zest.ServerConfiguration{
		URL: "http://localhost:8080",
	}}
	client := zest.NewAPIClient(pulpConfig)

	authCtx := context.WithValue(ctx2, zest.ContextBasicAuth, zest.BasicAuth{
		UserName: "admin",
		Password: "changeme",
	})

	resp, r, err := client.UpstreamPulpsAPI.UpstreamPulpsList(authCtx).Execute()
	if err != nil {
		log.Fatalf("Error when calling `StatusAPI.StatusRead``: %v\n", err)
		log.Fatalf("Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsList`: PaginatedUpstreamPulpResponseList
	log.Infof("Response from `StatusAPI.StatusRead`: %v\n", resp)
        log.Infof("Count: %v", resp.GetCount())
        log.Infof("%v", resp.Results)
        for i := 0; i < int(resp.GetCount()); i++ {
                res:= resp.Results[i]
                log.Infof("APIRoot %v", res.ApiRoot)
                log.Infof("BaseURL%v", res.BaseUrl)

        }
        log.Info("Printed everything\n")
}

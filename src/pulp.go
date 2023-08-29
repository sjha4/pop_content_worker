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

func createUpstreamPulp(upstream_pulp_name, upstream_pulp_base_url, upstream_pulp_api_root string) {

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
    upstreamPulp := *zest.NewUpstreamPulp(upstream_pulp_name, upstream_pulp_base_url, upstream_pulp_api_root)

	resp, r, err := client.UpstreamPulpsAPI.UpstreamPulpsCreate(authCtx).UpstreamPulp(upstreamPulp).Execute()
	if err != nil {
		log.Fatalf("Error when calling `UpstreamPulpsAPI.UpstreamPulpsCreate``: %v\n", err)
		log.Fatalf("Full HTTP response: %v\n", r)
	}
	// response from `UpstreamPulpsList`: PaginatedUpstreamPulpResponseList
	log.Infof("Response from `UpstreamPulpsAPI.UpstreamPulpsCreate`: %v\n", resp)
    log.Info("Printed everything\n")
}

func replicateUpstreamPulp() {
    var upstreamPulpHref *string
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
    		log.Fatalf("Error when calling `UpstreamPulpsAPI.UpstreamPulpsList``: %v\n", err)
    		log.Fatalf("Full HTTP response: %v\n", r)
    	}
    	// response from `UpstreamPulpsList`: PaginatedUpstreamPulpResponseList
    	log.Infof("Response from `UpstreamPulpsAPI.UpstreamPulpsList`: %v\n", resp)
        log.Infof("Count: %v", resp.GetCount())
        log.Infof("%v", resp.Results)
        if(int(resp.GetCount()) > 0) {
            upstreamPulpHref = resp.Results[0].PulpHref

        }
        //UpstreamPulpsAPI.UpstreamPulpsReplicate(context.Background(), upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()
        log.Infof("Replicating upstream pulp with href: %v", *upstreamPulpHref)
        upstreamPulp := *zest.NewUpstreamPulp("MainPulp", "10.1.2.22:8080", "/pulp/")
        resp1, r1, err1 := client.UpstreamPulpsAPI.UpstreamPulpsReplicate(authCtx, *upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()
        if err1 != nil {
            		log.Fatalf("Error when calling `UpstreamPulpsAPI.UpstreamPulpsList``: %v\n", err1)
            		log.Fatalf("Full HTTP response: %v\n", r1)
            	}
            	// response from `UpstreamPulpsList`: PaginatedUpstreamPulpResponseList
            	log.Infof("Response from `UpstreamPulpsAPI.UpstreamPulpsList`: %v\n", resp1)

}
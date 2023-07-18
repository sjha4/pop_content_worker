package main

import (
	"context"
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
	limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := zest.NewConfiguration()
    apiClient := zest.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsAPI.UpstreamPulpsList(context.Background()).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        log.Fatalf("Error when calling `UpstreamPulpsAPI.UpstreamPulpsList``: %v\n", err)
        log.Fatalf("Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsList`: PaginatedUpstreamPulpResponseList
    log.Infof("Response from `UpstreamPulpsAPI.UpstreamPulpsList`: %v\n", resp)

}

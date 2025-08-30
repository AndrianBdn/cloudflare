package main

import (
	"os"
	"testing"

	"github.com/libdns/cloudflare"
	"github.com/libdns/libdns/e2e"
)

func TestCloudflareProvider(t *testing.T) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	zoneToken := os.Getenv("CLOUDFLARE_ZONE_TOKEN")
	testZone := os.Getenv("CLOUDFLARE_TEST_ZONE")

	if apiToken == "" || testZone == "" {
		t.Skip("Skipping Cloudflare e2e tests: CLOUDFLARE_API_TOKEN and CLOUDFLARE_TEST_ZONE environment variables must be set")
	}

	provider := &cloudflare.Provider{
		APIToken:  apiToken,
		ZoneToken: zoneToken, // optional
	}

	suite := e2e.NewFullTestSuite(provider, testZone)
	suite.RunFullTests(t)
}

package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sidecar/internal/logger"
	"sidecar/internal/proxy"
	"testing"
)

var (
	cfg = MainConfig{
		Logger: logger.Config{},
		Proxy:  proxy.Config{},
		Server: ServerConfig{},
	}
	ctx = context.Background()
)

func TestServer(t *testing.T) {
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	server := NewServer(&cfg, nil)
	err := server.Initialize(ctx)
	if err != nil {
		t.Error(err)
	}
	router := SetupRouter(server.RESTHandler, server.Config)
	ts := httptest.NewServer(router)
	defer ts.Close()

	tests := []struct {
		description        string
		url                string
		expectedStatusCode int
		initService        func() proxy.Service
	}{
		{
			description:        "404 not found",
			url:                "%s/v2/company",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			description:        "500 internal sever err",
			url:                "%s/company",
			expectedStatusCode: http.StatusInternalServerError, // we are testing this in build process
		},
		{
			description:        "200 ok",
			url:                "%s/company/sd45f768",
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "200 ok",
			url:                "%s/account/acc234234/user",
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "200 ok",
			url:                "%s/account/acc74850",
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "404 not found",
			url:                "%s/account/blocked",
			expectedStatusCode: http.StatusNotFound,
		},

		{
			description:        "200 ok",
			url:                "%s/company/account",
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "200 ok",
			url:                "%s/acc734340",
			expectedStatusCode: http.StatusOK,
		},

		{
			description:        "404 not found",
			url:                "%s/tenant/sj3co3s4",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			description:        "200 ok",
			url:                "%s/tenant/account/blocked",
			expectedStatusCode: http.StatusOK,
		},
		{
			description:        "404 not found",
			url:                "%s/tenant/account/acc23849",
			expectedStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			url := fmt.Sprintf(tt.url, ts.URL)
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			assert.Equal(t, resp.StatusCode, tt.expectedStatusCode)
		})
	}
}

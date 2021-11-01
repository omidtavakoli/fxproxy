package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sidecar/internal/proxy"
	"testing"
	"time"
)

var dummyTime = time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)

func TestHandler_Product(t *testing.T) {
	tests := []struct {
		description        string
		expectedStatusCode int
		expectedData       string
		initService        func() proxy.Service
	}{
		{
			description:        "200 ok",
			expectedStatusCode: http.StatusOK,
			expectedData:       "ok",
			initService: func() proxy.Service {
				mockRep := new(MockProxyService)
				mockRep.On("Company", mock.Anything).Return("ok", nil)
				return mockRep
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			service := tt.initService()
			handler := CreateHandler(service)
			gin.SetMode(gin.TestMode)
			gin.DefaultWriter = ioutil.Discard
			router := gin.Default()
			router.GET("/company", handler.Company)

			req, _ := http.NewRequest("GET", "/company", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.expectedStatusCode)
		})
	}
}

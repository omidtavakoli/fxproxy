package proxy

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
	"time"
)

var dummyTime = time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)

func TestHandler_Product(t *testing.T) {
	tests := []struct {
		description        string
		expectedStatusCode int
		expectedData       string
		initService        func() Service
	}{
		{
			description:        "200 ok",
			expectedStatusCode: http.StatusOK,
			expectedData:       "working",
			initService: func() Service {
				mockRep := new(MockProxyService)
				mockRep.On("Company", mock.Anything).Return("ok", nil)
				return mockRep
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			service := tt.initService()
			ret, err := service.Company()
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, ret, tt.expectedData)
		})
	}
}

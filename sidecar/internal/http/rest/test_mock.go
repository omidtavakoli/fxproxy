package rest

import "github.com/stretchr/testify/mock"

type MockProxyService struct {
	mock.Mock
}

func (m *MockProxyService) Company() (string, error) {
	args := m.Called()
	return "", args.Error(1)
}

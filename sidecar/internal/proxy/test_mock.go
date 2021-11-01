package proxy

import "github.com/stretchr/testify/mock"

type MockProxyService struct {
	mock.Mock
}

func (m *MockProxyService) Company() (string, error) {
	args := m.Called()
	return "working", args.Error(1)
}

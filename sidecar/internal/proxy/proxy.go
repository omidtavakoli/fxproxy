package proxy

import (
	"github.com/go-playground/validator/v10"
	"sidecar/internal/logger"
)

type Service interface {
	//Func()
}

type service struct {
	validate *validator.Validate
	logger   *logger.StandardLogger
	config   *Config
}

func CreateService(
	config *Config,
	logger *logger.StandardLogger,
	validator *validator.Validate) Service {
	return &service{
		validate: validator,
		logger:   logger,
		config:   config,
	}
}

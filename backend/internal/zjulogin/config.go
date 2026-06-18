package zjulogin

import (
	"fmt"
	"strings"
)

type Config struct {
	Username       string
	Password       string
	TYYSSignSecret string
}

func New(cfg Config) (*Auth, error) {
	if strings.TrimSpace(cfg.Username) == "" {
		return nil, fmt.Errorf("ZJU_USERNAME is required")
	}
	if strings.TrimSpace(cfg.Password) == "" {
		return nil, fmt.Errorf("ZJU_PASSWORD is required")
	}

	am, err := NewZJUAM(cfg)
	if err != nil {
		return nil, err
	}

	return &Auth{
		am:     am,
		config: cfg,
	}, nil
}

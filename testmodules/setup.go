package testmodules

import "github.com/TGRZiminiar/go-mc-kafka/config"

func NewTestConfig() *config.Config {
	cfg := config.LoadConfig("../env/test/.env")
	return &cfg
}

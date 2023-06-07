package config

import (
	"log"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	BaseURL      string `env:"GITLAB_BASE_URL" envDefault:"https://gitlab.com"`
	Token        string `env:"GITLAB_API_TOKEN"`
	RunnerStatus string `env:"GITLAB_RUNNER_STATUS" envDefault:"stale"`
	DryRun       bool   `env:"DRY_RUN" envDefault:"true"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	opts := env.Options{RequiredIfNoDef: true}

	// Load config from environment
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		log.Fatal(err)
	}

	return cfg, nil
}

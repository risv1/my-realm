package config

import (
	"log"
	"os"
)

type Env struct {
	GithubToken string `mapstructure:"GITHUB_TOKEN"`
}

func LoadEnv() *Env {
	env := Env{
		GithubToken: os.Getenv("GITHUB_TOKEN"),
	}

	if env.GithubToken == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}

	return &env
}

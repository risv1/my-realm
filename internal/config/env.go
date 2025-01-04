package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	Port           string `mapstructure:"PORT"`
	GithubUsername string `mapstructure:"GITHUB_USERNAME"`
	GithubToken    string `mapstructure:"GITHUB_TOKEN"`
}

func LoadEnv() *Env {
	env := Env{}

	viper.SetConfigFile(".env")
	viper.SetDefault("PORT", "8000")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading .env file: %s\n", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	log.Printf("Environment variables loaded successfully!")
	return &env
}

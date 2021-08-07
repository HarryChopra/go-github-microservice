package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	Config *config
)

type config struct {
	AppUrl          string
	GithubAccessKey string
}

// LoadConfig holds the environment variables in the config.Config variable
func LoadConfig() {
	readEnv()
	Config = &config{
		AppUrl:          os.Getenv("APP_URL"),
		GithubAccessKey: os.Getenv("GH_ACCESS_KEY"),
	}
}

func readEnv() {
	// Read the application environment variable or set if empty
	currentEnvironment, ok := os.LookupEnv("APPENV")
	if !ok {
		if err := os.Setenv("APPENV", "development"); err != nil {
			panic(err)
		}
		currentEnvironment = "development"
	}
	// Absolute path for the .env file, as the execution root changes during tests.
	if err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/harrychopra/go-github-microservice/src/api/.env." + currentEnvironment)); err != nil {
		panic(err)
	}
}

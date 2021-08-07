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

// LoadConfig loads and stores the environment variables in the config.Config object.
func LoadConfig() {
	readEnv()
	Config = &config{
		AppUrl:          os.Getenv("APP_URL"),
		GithubAccessKey: os.Getenv("GH_ACCESS_KEY"),
	}
}

func readEnv() {
	// Read the application environment variable or set it to "development",
	// if empty.
	currentEnvironment, ok := os.LookupEnv("APPENV")
	if !ok {
		if err := os.Setenv("APPENV", "development"); err != nil {
			panic(err)
		}
		currentEnvironment = "development"
	}
	// As the execution root during test isn't the same as the application,
	// provide an absolute path to load the environment variables.

	if err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/harrychopra/go-github-microservice/src/api/.env." + currentEnvironment)); err != nil {
		panic(err)
	}
}

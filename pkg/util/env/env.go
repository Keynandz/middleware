package env

import (
	"flag"
	"log"
	"os"
	"strconv"

	"go-authorization/pkg/file"

	"github.com/joho/godotenv"
)

type env struct{}

func NewEnv() *env {
	return &env{}
}

func (*env) Load() {
	var err error
	mode := flag.String("mode", "dev", "dev, prod or stag")
	flag.Parse()

	os.Setenv("PROJECT_DIR", "go-authorization") // should change this first for new projects
	rootPath := file.GetRootDirectory()

	switch *mode {
	case "prod":
		err = godotenv.Load(rootPath + "/.env.production")
	case "stag ":
		err = godotenv.Load(rootPath + "/.env.staging")
	default:
		err = godotenv.Load(rootPath + "/.env.development")
	}

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func (e *env) GetString(name string) string {
	return os.Getenv(name)
}

func (e *env) GetBool(name string) bool {
	s := e.GetString(name)
	i, err := strconv.ParseBool(s)
	if nil != err {
		return false
	}
	return i
}

func (e *env) GetInt(name string) int {
	s := e.GetString(name)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func (e *env) GetFloat(name string) float64 {
	s := e.GetString(name)
	i, err := strconv.ParseFloat(s, 64)
	if nil != err {
		return 0
	}
	return i
}

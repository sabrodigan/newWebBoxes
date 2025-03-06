package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"reflect"
)

var env Conf1gDto

type Conf1gDto struct {
	port         string
	secretKey    string
	databaseURL  string
	databaseName string
	users        string // Added the users field
	debug string
}

func init() {
	LoadEnvironmentVariable()
	Conf1gEnv()
}

func Conf1gEnv() {
	env = Conf1gDto{
		port:         os.Getenv("PORT"),
		secretKey:    os.Getenv("SECRET_KEY"),
		databaseURL:  os.Getenv("MONGODB_URI"),
		databaseName: os.Getenv("DATABASE_NAME"),
		users:        os.Getenv("USERS"), // Assigning value to the users field
		debug:        os.Getenv("DEBUG"),
	}
}

func LoadEnvironmentVariable() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("ALARM! error loading .env file: %v\n", err)
	}
}

func accessField(key string) (string, error) {
	v := reflect.ValueOf(env)
	t := reflect.TypeOf(env)
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("env is not a struct")
	}
	_, ok := t.FieldByName(key)
	if !ok {
		return "", fmt.Errorf("field %s not found", key)
	}
	f := v.FieldByName(key)
	return f.String(), nil
}

func GetEnvProperty(key string) (string, error) {
	if env.port == "" {
		fmt.Println("ALARM! env (port) is empty")
		Conf1gEnv()
	}

	value, err := accessField(key)
	if err != nil {
		fmt.Printf("WARNING! fetching %s from envDEBUGConfig: %v\n", key, err)
		return value, err
	}
	return value, nil
}
func ConfEnvChecks() {
	env = Conf1gDto{
		port:         os.Getenv("PORT"),
		secretKey:    os.Getenv("SECRET_KEY"),
		databaseURL:  os.Getenv("MONGODB_URI"),
		databaseName: os.Getenv("DATABASE_NAME"),
		users:        os.Getenv("USERS"),
	}
	logEnvVariables()
}

func logEnvVariables() {
	fmt.Printf("PORT: %s\n", env.port)
	fmt.Printf("SECRET_KEY: %s\n", env.secretKey)
	fmt.Printf("MONGODB_URI: %s\n", env.databaseURL)
	fmt.Printf("DATABASE_NAME: %s\n", env.databaseName)
	fmt.Printf("USERS: %s\n", env.users)
}

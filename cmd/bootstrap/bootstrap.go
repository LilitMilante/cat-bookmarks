package bootstrap

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	HttpPort   string
	DbScheme   string
	DbPort     int
	DbName     string
	DbUsername string
	DbPassword string
}

func NewConfig() (*Config, error) {
	var c Config
	if err := parseENV(); err != nil {
		return nil, err
	}

	port := os.Getenv("HTTP_PORT")
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	c.HttpPort = port

	port = os.Getenv("DB_PORT")
	iPort, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	c.DbPort = iPort

	c.DbScheme = os.Getenv("DB_SCHEME")
	c.DbName = os.Getenv("DB_NAME")
	c.DbUsername = os.Getenv("DB_USERNAME")
	c.DbPassword = os.Getenv("DB_PASSWORD")

	return &c, nil
}

func parseENV() error {
	data, err := ioutil.ReadFile("../.env")
	if err != nil {
		return err
	}

	result := strings.Split(string(data), "\n")

	for _, v := range result {
		if v == "" {
			continue
		}
		s := strings.SplitN(v, "=", 2)
		if len(s) != 2 {
			continue
		}
		if err := os.Setenv(s[0], s[1]); err != nil {
			return err
		}
	}

	return nil
}

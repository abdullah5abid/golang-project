package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spongeling/admin-api/internal/errors"
)

type Config struct {
	HttpPort           int    `yaml:"httpPort"`
	DbName             string `yaml:"dbName"`
	DbUser             string `yaml:"dbUser"`
	DbPass             string `yaml:"dbPass"`
	DbHost             string `yaml:"dbHost"`
	DbPort             int    `yaml:"dbPort"`
	DbMigrationsSource string
}

func (c *Config) UpdateFromArguments() {
	flag.IntVar(&c.HttpPort, "port", 6543, "Port for http server")
	flag.StringVar(&c.DbName, "dbName", "", "Database name")
	flag.StringVar(&c.DbUser, "dbUser", "", "Database user's name")
	flag.StringVar(&c.DbPass, "dbPass", "", "Database password")
	flag.StringVar(&c.DbHost, "dbHost", "127.0.0.1", "Database host")
	flag.IntVar(&c.DbPort, "dbPort", 5432, "Database port")
	flag.Parse()

	remainingArgs := flag.Args()
	if len(remainingArgs) > 0 {
		fmt.Println("Unknown arguments:", strings.Join(remainingArgs, " "))
		os.Exit(-1)
	}
}

func (c *Config) UpdateFromEnv() {
	if x, err := strconv.ParseInt(os.Getenv("HTTP_PORT"), 10, 64); err != nil {
		c.HttpPort = int(x)
	}

	if x := os.Getenv("DB_NAME"); x != "" {
		c.DbName = x
	}
	if x := os.Getenv("DB_USER"); x != "" {
		c.DbUser = x
	}
	if x := os.Getenv("DB_PASSWORD"); x != "" {
		c.DbPass = x
	}
	if x := os.Getenv("DB_HOST"); x != "" {
		c.DbHost = x
	}
	if x, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64); err != nil {
		c.DbPort = int(x)
	}
}

func (c *Config) Validate() error {
	if c.HttpPort == 0 {
		return errors.New(errors.InvalidArgument, "empty server port. -port flag required")
	}
	if c.DbUser == "" {
		return errors.New(errors.InvalidArgument, "empty database username. -dbUser flag required")
	}
	if c.DbPass == "" {
		return errors.New(errors.InvalidArgument, "empty database password. -dbPass flag required")
	}
	if c.DbHost == "" {
		return errors.New(errors.InvalidArgument, "empty database hostname. -dbHost flag required")
	}
	if c.DbPort == 0 {
		return errors.New(errors.InvalidArgument, "empty database port. -dbPort flag required")
	}
	if c.DbName == "" {
		return errors.New(errors.InvalidArgument, "empty database name. -dbName flag required")
	}

	return nil
}

func readConfig() (Config, error) {
	var cfg Config

	cfg.UpdateFromArguments()
	cfg.UpdateFromEnv()
	cfg.DbMigrationsSource = "file://db/migrations"

	return cfg, cfg.Validate()
}

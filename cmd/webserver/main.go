package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"log/slog"
	"os"
	"web/iternal/app/webserver"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/webserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	cfg := webserver.NewConfig()
	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := webserver.Start(cfg); err != nil {
		log.Fatal(err)
	}
	log := setupLogger(cfg.LogLevel)
	log.Info("server started")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

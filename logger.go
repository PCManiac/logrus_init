package logrus_init

import (
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

func init() {
	InitLogger()
}

type config struct {
	Output    string `env:"LOG_OUTPUT" envDefault:"stdout"`
	FileName  string `env:"LOG_FILENAME" envDefault:"app.log"`
	Formatter string `env:"LOG_FORMATTER" envDefault:"json"`
	Level     string `env:"LOG_LEVEL" envDefault:"error"`
}

func InitLogger() {
	log := logrus.StandardLogger()

	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Error("log env vars parse error", err)
	}

	logFilename := cfg.FileName
	switch cfg.Output {
	case "file":
		LogOutputFile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(LogOutputFile)
		} else {
			log.SetOutput(os.Stdout)
		}
	case "stderr":
		log.SetOutput(os.Stderr)
	case "stdin":
		log.SetOutput(os.Stdin)
	case "stdout":
		fallthrough
	default:
		log.SetOutput(os.Stdout)
	}

	switch cfg.Formatter {
	case "text":
		log.SetFormatter(&logrus.TextFormatter{})
	case "json":
		fallthrough
	default:
		log.SetFormatter(&logrus.JSONFormatter{})
	}

	if lvl, err := logrus.ParseLevel(cfg.Level); err != nil {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(lvl)
	}
}

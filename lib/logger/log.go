package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.ToSlash(config.Get().Application.Log.Path),
		MaxSize:    config.Get().Application.Log.MaxSize,
		MaxBackups: config.Get().Application.Log.MaxBackup,
		MaxAge:     config.Get().Application.Log.MaxAge,
		Compress:   config.Get().Application.Log.Compress,
	}

	// Fork writing into two outputs
	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	output := zerolog.ConsoleWriter{Out: multiWriter, TimeFormat: time.RFC3339}

	log.Logger = log.Output(output)
	log.Trace().Msg("Zerolog initialized...")
}

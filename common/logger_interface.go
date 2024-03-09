package common

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LoggerInterface interface {
	ErrorLog(ctx context.Context, err error, msg string, additional any)
	WarningLog(ctx context.Context, msg string, additional any)
	InfoLog(ctx context.Context, msg string, additional any)
	DebugLog(ctx context.Context, msg string, additional any)
	FatalLog(ctx context.Context, msg string, additional any)
}

// ZLogger is an implementation of LoggerInterface
type ZLogger struct {
	logger *zerolog.Logger
}

// NewZLogger creates a new instance of MyLogger with the specified output option
func NewZLogger(writeToFile bool, filePath string) *ZLogger {
	var logger ZLogger

	if writeToFile {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			if os.IsNotExist(err) {
				file, err = os.Create(filePath)
				if err != nil {
					log.Error().Err(err).Msg("Error creating log file")
					zerologger := zerolog.New(os.Stdout).With().Timestamp().Logger()
					logger.logger = &zerologger
					return &logger
				}
			} else {
				log.Error().Err(err).Msg("Error opening log file")
				zerologger := zerolog.New(os.Stdout).With().Timestamp().Logger()
				logger.logger = &zerologger
				return &logger
			}
		}
		zerologger := zerolog.New(file).With().Timestamp().Logger()
		logger.logger = &zerologger
	} else {
		zerologger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		logger.logger = &zerologger
	}

	return &logger
}

// ErrorLog implements the ErrorLog method of LoggerInterface
func (l *ZLogger) ErrorLog(ctx context.Context, err error, msg string, additional interface{}) {
	l.logger.Error().Err(err).Msgf(msg+" Additional: %v", additional)
}

// WarningLog implements the WarningLog method of LoggerInterface
func (l *ZLogger) WarningLog(ctx context.Context, msg string, additional interface{}) {
	l.logger.Warn().Msgf(msg+" Additional: %v", additional)
}

// InfoLog implements the InfoLog method of LoggerInterface
func (l *ZLogger) InfoLog(ctx context.Context, msg string, additional interface{}) {
	l.logger.Info().Msgf(msg+" Additional: %v", additional)
}

// DebugLog implements the DebugLog method of LoggerInterface
func (l *ZLogger) DebugLog(ctx context.Context, msg string, additional interface{}) {
	l.logger.Debug().Msgf(msg+" Additional: %v", additional)
}

// FatalLog implements the FatalLog method of LoggerInterface
func (l *ZLogger) FatalLog(ctx context.Context, msg string, additional interface{}) {
	l.logger.Fatal().Msgf(msg+" Additional: %v", additional)
}

package logger

import (
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	logLevelEnvKey = "LOG_LEVEL"
)

// New constructs a new default logger.
func New() *logrus.Entry {
	var (
		logger        *logrus.Logger
		defaultFields logrus.Fields
	)

	once := sync.Once{}
	once.Do(func() {
		output := os.Stdout

		logger = logrus.New()
		logger.SetOutput(output)
		logger.SetLevel(logrus.InfoLevel)

		defaultFields = logrus.Fields{
			"service":         os.Getenv("SERVICE_NAME"),
			"service-version": os.Getenv("SERVICE_VERSION"),
		}

		customLogLevel := os.Getenv(logLevelEnvKey)
		logLevel, err := logrus.ParseLevel(customLogLevel)
		if err != nil {
			logger.Warnf("could not set custom log level: %s, using info level", err.Error())
		} else {
			logger.SetLevel(logLevel)
		}

		// If the formatter is not set,
		logger.Formatter = &logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
			// Enable logging the full timestamp when a TTY is attached instead of just
			// the time passed since beginning of execution.
			// If not set and the run is short, only 0000 is shown.
			FullTimestamp: true,
		}
	})

	return logger.WithFields(defaultFields)
}

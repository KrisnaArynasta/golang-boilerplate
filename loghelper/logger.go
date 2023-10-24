package Loghelper

import (
	"log"
	"path/filepath"

	"os"

	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func WriteLogPlain(text string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	LOG_FILE := filepath.Join(dir, "log.txt")
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println(text)
}

func WriteLog() *zerolog.Logger {
	//return zapLoger() //function need to return ==> *zap.Logger
	return zerologLoger()
}

func zapLoger() *zap.Logger {
	filename := "logs.log"

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}

func zerologLoger() *zerolog.Logger {
	filename := "logs.log" //"C:/log/golang.log",

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	runLogFile, _ := os.OpenFile(
		filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	logger := zerolog.New(multi).With().Timestamp().Logger()
	return &logger
}

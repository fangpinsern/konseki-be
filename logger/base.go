package logger

import (
	"errors"
	"github.com/fangpinsern/konseki-be/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	fileLogger  *log.Logger
)

const (
	InfoFileName  = "info.log"
	ErrorFileName = "error.log"
	FileFileName  = "file.log"
	LogFolderName = "log/"
)

const (
	correlationIDLogKey = "correlationId"
	pathIDLogKey        = "path"
	methodLogKey        = "method"
	fileLogKey          = "file"
)

func InitializeLogger() {
	if _, err := os.Stat("log"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("log", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	newLogger, err := NewLogger(InfoFileName)
	if err != nil {
		panic(err)
	}
	infoLogger = newLogger

	newErrorLogger, err := NewLogger(ErrorFileName)
	if err != nil {
		panic(err)
	}
	errorLogger = newErrorLogger

	newFileLogger, err := NewLogger(FileFileName)
	if err != nil {
		panic(err)
	}
	fileLogger = newFileLogger

}

func NewLogger(fileName string) (*log.Logger, error) {
	newLogger := log.New()
	newLogger.SetFormatter(&log.JSONFormatter{})

	file, err := os.OpenFile(LogFolderName+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		newLogger.Out = file
	} else {
		return nil, err
	}

	return newLogger, nil
}

func LogInternal(c *gin.Context, err error, msg string) {
	fields := log.Fields{}

	if c != nil {
		fields[correlationIDLogKey] = util.GetCorrelationID(c)
		fields[pathIDLogKey] = util.GetPath(c)
		fields[methodLogKey] = util.GetMethod(c)
	}

	_, fileName, lineNo, ok := runtime.Caller(1)
	if ok {
		fields[fileLogKey] = fileName + ":" + strconv.Itoa(lineNo)
	}

	if err != nil {
		fields["error"] = err.Error()
		errorLogger.WithFields(fields).Error(msg)
		return
	}

	infoLogger.WithFields(fields).Info(msg)
}
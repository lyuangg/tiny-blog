package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"tiny-blog/configs"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logger is logrus
var Logger = logrus.New()

// InitLogger is Logger Middleware
func InitLogger() {

	fileName := configs.Conf.Log.Path

	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logWriter, _ := rotatelogs.New(
		fileName+"-%Y-%m-%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Duration(configs.Conf.Log.Days)*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Logger.Out = src
	level, _ := logrus.ParseLevel(configs.Conf.Log.Level)
	Logger.SetLevel(level)
	Logger.AddHook(lfshook.NewHook(
		writeMap,
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	))
}

// LoggerMiddleware is Logger Middleware
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		requestBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		c.Next()

		usedTime := time.Since(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		resBody := ""
		if c.Writer.Header().Get("Content-Type") == "application/json; charset=utf-8" {
			resBody = blw.body.String()
		}
		errorMsg := c.Errors.ByType(gin.ErrorTypePrivate).String()

		// 日志格式
		Logger.WithFields(logrus.Fields{
			"http_code":  statusCode,
			"used_time":  fmt.Sprintf("%v", usedTime),
			"client_ip":  clientIP,
			"req_method": reqMethod,
			"req_uri":    reqURI,
			"request":    string(requestBody),
			"response":   resBody,
			"error_msg":  errorMsg,
		}).Info()

	}
}

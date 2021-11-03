package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
)

type Config struct {
	Path         string
	Level        string
	ReportCaller bool
	Mode         string
}

func New(c *Config) *log.Logger {
	log.SetFormatter(&nested.Formatter{
		// HideKeys:        true,
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        true,
		NoFieldsColors:  true,
		//FieldsOrder:     []string{"name", "age"},
	})

	// 日志文件
	f := c.Path
	writer, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic("open log file fail!")
	}

	var writes io.Writer

	//正式环境输出到日志文件,
	//开发环境输出到控制台和日志文件
	//其他只输出到控制台
	if c.Mode == ReleaseMode {
		writes = io.MultiWriter(writer)
	} else if c.Mode == DebugMode {
		writes = io.MultiWriter(writer, os.Stdout)
	} else {
		writes = io.MultiWriter(os.Stdout)
		defer writer.Close()
	}

	log.SetOutput(writes)

	log.SetReportCaller(c.ReportCaller)

	level, err2 := log.ParseLevel(c.Level)
	if err2 != nil {
		level = log.DebugLevel
	}
	log.SetLevel(level)

	return log.StandardLogger()
}

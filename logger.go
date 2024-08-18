package logger

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type logSyntax uint32

const (
	DateTime logSyntax = 1 << iota
	Loglevel
	ShortFileName
	LongFileName
	LogMessage
)

type LogFileConfigs struct {
	Directory string
	Filename  string
	Stdout    bool
	Include   logSyntax
}

type Logger struct {
	DEBUG *log.Logger
	INFO  *log.Logger
	WARN  *log.Logger
	ERROR *log.Logger
	TRACE *log.Logger
}

func NewLogger(config *LogFileConfigs) (*Logger, error) {
	l := &Logger{}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	multiWriter := io.MultiWriter(os.Stdout)

	if config != nil {
		if config.Filename == "" {
			return nil, errors.New("filename is required")
		}

		path := filepath.Join(wd, config.Directory, config.Filename)

		err = os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			return nil, err
		}

		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}

		if config.Stdout {
			multiWriter = io.MultiWriter(os.Stdout, file)
		} else {
			multiWriter = io.MultiWriter(file)
		}

		l.INFO = log.New(multiWriter, generatePrefix(config.Include, "INFO"), 0)
		l.WARN = log.New(multiWriter, generatePrefix(config.Include, "WARN"), 0)
		l.ERROR = log.New(multiWriter, generatePrefix(config.Include, "ERROR"), 0)
		l.DEBUG = log.New(multiWriter, generatePrefix(config.Include, "DEBUG"), 0)
		l.TRACE = log.New(multiWriter, generatePrefix(config.Include, "TRACE"), 0)
	} else {
		flag := log.Lmsgprefix | log.LstdFlags | log.Lshortfile
		l.INFO = log.New(multiWriter, "INFO ", flag)
		l.WARN = log.New(multiWriter, "WARN ", flag)
		l.ERROR = log.New(multiWriter, "ERROR ", flag)
		l.DEBUG = log.New(multiWriter, "DEBUG ", flag)
		l.TRACE = log.New(multiWriter, "TRACE ", flag)
	}

	return l, nil
}

func generatePrefix(syntax logSyntax, level string) string {
	prefix := ""

	if syntax&DateTime != 0 {
		prefix += time.Now().Format("2006-01-02 15:04:05") + " "
	}

	if syntax&Loglevel != 0 {
		prefix += level + " "
	}

	if syntax&(ShortFileName|LongFileName) != 0 {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			if syntax&ShortFileName != 0 {
				file = filepath.Base(file)
			}
			prefix += fmt.Sprintf("%s:%d ", file, line)
		}
	}

	return prefix
}

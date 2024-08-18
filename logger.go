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

// logSyntax is a bitmask used to define the format of log messages.
type logSyntax uint32

const (
	DateTime      logSyntax = 1 << iota // Include date and time in the log entry.
	Loglevel                            // Include log level in the log entry.
	ShortFileName                       // Include short (basename) file name in the log entry.
	LongFileName                        // Include full file path in the log entry.
)

// LogFileConfigs encapsulates the configuration options for the Logger.
// Directory: Directory path where the log file will be created.
// Filename: Name of the log file.
// Stdout: Whether to also log to standard output.
// Include: Bitmask to define which syntax elements to include in the log entry.
type LogFileConfigs struct {
	Directory string
	Filename  string
	Stdout    bool
	Include   logSyntax
}

// Logger holds loggers for different log levels.
// DEBUG: Logger for debugging information.
// INFO: Logger for general informational messages.
// WARN: Logger for warnings that are not critical.
// ERROR: Logger for errors that require attention.
// TRACE: Logger for detailed trace information, useful for debugging.
type Logger struct {
	DEBUG *log.Logger
	INFO  *log.Logger
	WARN  *log.Logger
	ERROR *log.Logger
	TRACE *log.Logger
}

// NewLogger initializes a Logger based on the provided configuration.
// It creates log files, sets up multi-writer for simultaneous stdout/file logging,
// and assigns loggers for various levels.
// Returns the initialized Logger and an error if any issues are encountered during setup.
func NewLogger(config *LogFileConfigs) (*Logger, error) {
	l := &Logger{}

	// Get the current working directory to construct the log file path.
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Default to logging to standard output.
	multiWriter := io.MultiWriter(os.Stdout)

	// If configuration is provided, set up file logging.
	if config != nil {
		if config.Filename == "" {
			return nil, errors.New("filename is required")
		}

		// Construct the full path for the log file.
		path := filepath.Join(wd, config.Directory, config.Filename)

		// Ensure the directory exists or create it.
		err = os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			return nil, err
		}

		// Open the log file for writing.
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}

		// Setup the writer to write to both stdout and the file if specified.
		if config.Stdout {
			multiWriter = io.MultiWriter(os.Stdout, file)
		} else {
			multiWriter = io.MultiWriter(file)
		}

		// Initialize loggers for each level with the appropriate prefix and multi-writer.
		l.INFO = log.New(multiWriter, generatePrefix(config.Include, "INFO"), 0)
		l.WARN = log.New(multiWriter, generatePrefix(config.Include, "WARN"), 0)
		l.ERROR = log.New(multiWriter, generatePrefix(config.Include, "ERROR"), 0)
		l.DEBUG = log.New(multiWriter, generatePrefix(config.Include, "DEBUG"), 0)
		l.TRACE = log.New(multiWriter, generatePrefix(config.Include, "TRACE"), 0)
	} else {
		// If no configuration is provided, fall back to a basic logger configuration.
		flag := log.Lmsgprefix | log.LstdFlags | log.Lshortfile
		l.INFO = log.New(multiWriter, "INFO ", flag)
		l.WARN = log.New(multiWriter, "WARN ", flag)
		l.ERROR = log.New(multiWriter, "ERROR ", flag)
		l.DEBUG = log.New(multiWriter, "DEBUG ", flag)
		l.TRACE = log.New(multiWriter, "TRACE ", flag)
	}

	return l, nil
}

// generatePrefix generates a log prefix based on the specified syntax and log level.
// syntax: Bitmask that determines which parts of the prefix are included.
// level: The log level string (e.g., "INFO") to include in the prefix.
// Returns the formatted log prefix string.
func generatePrefix(syntax logSyntax, level string) string {
	prefix := ""

	// Include date and time if specified.
	if syntax&DateTime != 0 {
		prefix += time.Now().Format("2006-01-02 15:04:05") + " "
	}

	// Include the log level if specified.
	if syntax&Loglevel != 0 {
		prefix += level + " "
	}

	// Include file name and line number if specified.
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

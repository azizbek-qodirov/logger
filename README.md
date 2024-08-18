# Go Logger Package

[![Go Reference](https://pkg.go.dev/badge/github.com/Azizbek-Qodirov/logger.svg)](https://pkg.go.dev/github.com/Azizbek-Qodirov/logger)
[![Go Report Card](https://goreportcard.com/badge/github.com/Azizbek-Qodirov/logger)](https://goreportcard.com/report/github.com/Azizbek-Qodirov/logger)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

A flexible and easy-to-use logging package for Go applications. This logger provides customizable log levels, output formats, and destinations, making it suitable for a wide range of Go projects.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Usage](#usage)
- [Advanced Examples](#advanced-examples)
- [Performance Considerations](#performance-considerations)
- [Comparison with Other Loggers](#comparison-with-other-loggers)
- [FAQ](#faq)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Features

- Multiple log levels: DEBUG, INFO, WARN, ERROR, TRACE
- Customizable log output format
- File and/or stdout logging
- Configurable log file location
- Thread-safe logging

## Installation

To install the logger package, use `go get`:

```bash
go get -u github.com/Azizbek-Qodirov/logger
```

## Quick Start

Here's a simple example to get you started:

```go
package main

import (
    "github.com/Azizbek-Qodirov/logger"
)

func main() {
    config := &logger.LogFileConfigs{
        Directory: "logs",
        Filename:  "app.log",
        Stdout:    true,
        Include:   logger.DateTime | logger.Loglevel | logger.ShortFileName,
    }

    log, err := logger.NewLogger(config)
    if err != nil {
        panic(err)
    }

    log.INFO.Println("Application started")
    log.DEBUG.Println("This is a debug message")
    log.WARN.Println("This is a warning message")
    log.ERROR.Println("This is an error message")
    log.TRACE.Println("This is a trace message")
}
```

## Configuration

The `LogFileConfigs` struct allows you to customize the logger:

```go
type LogFileConfigs struct {
    Directory string
    Filename  string
    Stdout    bool
    Include   LogSyntax
}
```

- `Directory`: The directory where log files will be stored
- `Filename`: The name of the log file
- `Stdout`: If true, logs will also be written to stdout
- `Include`: A bitmask of `LogSyntax` values to include in the log prefix

The `LogSyntax` type is a bitmask that allows you to customize the log prefix:

```go
const (
    DateTime LogSyntax = 1 << iota
    Loglevel
    ShortFileName
    LongFileName
    LogMessage
)
```

## Usage

### Creating a Logger

```go
config := &logger.LogFileConfigs{
    Directory: "logs",
    Filename:  "app.log",
    Stdout:    true,
    Include:   logger.DateTime | logger.Loglevel | logger.ShortFileName,
}

log, err := logger.NewLogger(config)
if err != nil {
    panic(err)
}
```

### Logging Messages

```go
log.INFO.Println("This is an info message")
log.WARN.Printf("This is a warning message: %s", warningDetails)
log.ERROR.Println("This is an error message")
```

## Advanced Examples

### Logging to Multiple Destinations

You can configure the logger to write to both a file and stdout:

```go
config := &logger.LogFileConfigs{
    Directory: "logs",
    Filename:  "app.log",
    Stdout:    true,
    Include:   logger.DateTime | logger.Loglevel | logger.LongFileName,
}

log, err := logger.NewLogger(config)
if err != nil {
    panic(err)
}

log.INFO.Println("This message will be logged to both file and stdout")
```

### Custom Log Prefix

You can customize the log prefix to include various pieces of information:

```go
config := &logger.LogFileConfigs{
    Directory: "logs",
    Filename:  "app.log",
    Stdout:    true,
    Include:   logger.DateTime | logger.Loglevel | logger.ShortFileName | logger.LogMessage,
}

log, err := logger.NewLogger(config)
if err != nil {
    panic(err)
}

log.INFO.Println("This log entry will include date, time, log level, and short file name in its prefix")
```

## Performance Considerations

- File I/O can be a bottleneck in high-volume logging scenarios. Consider using buffered I/O or asynchronous logging for performance-critical applications.
- Be mindful of the log levels you use in production. Excessive debug or trace logging can impact performance.
- The logger is designed to be thread-safe, but heavy concurrent use may lead to contention. Consider using separate loggers for different components in highly concurrent applications.

## Comparison with Other Loggers

| Feature | This Logger | logrus | zap | standard log |
|---------|-------------|--------|-----|--------------|
| Multiple log levels | ✓ | ✓ | ✓ | ✗ |
| Structured logging | ✗ | ✓ | ✓ | ✗ |
| Performance | Good | Good | Excellent | Basic |
| Customizable output | ✓ | ✓ | ✓ | Limited |
| File rotation | ✗ | ✗ | ✗ | ✗ |

Our logger provides a balance between simplicity and features, making it suitable for many applications without the complexity of more feature-rich loggers.

## FAQ

**Q: Can I use this logger in a production environment?**
A: Yes, the logger is designed to be used in production environments. However, ensure you've properly configured log levels and destinations to avoid performance issues.

**Q: Does this logger support log rotation?**
A: Currently, log rotation is not built into the logger. You may need to implement log rotation externally or use a log management tool.

**Q: Can I create multiple loggers with different configurations?**
A: Yes, you can create multiple logger instances with different configurations to suit various parts of your application.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Please ensure your code adheres to the existing style and passes all tests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to all contributors who have helped shape this project.
- Inspired by the need for a flexible logging solution in Go applications.
- Built with love for the Go community.

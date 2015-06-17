package logger

import (
	"io"
	"log"
	"runtime"
	"strconv"
)

// Log levels
const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var levelNames = []string{
	"TRACE",
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
}

// Logger outputs logs.
type Logger struct {
	level   int
	loggers []*log.Logger
}

// Tracef prints log.
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.printfAt(TRACE, format, v...)
}

// Traceln prints log.
func (l *Logger) Traceln(v ...interface{}) {
	l.printlnAt(TRACE, v...)
}

// Debugf prints log.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.printfAt(DEBUG, format, v...)
}

// Debugln prints log.
func (l *Logger) Debugln(v ...interface{}) {
	l.printlnAt(DEBUG, v...)
}

// Infof prints log.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.printfAt(INFO, format, v...)
}

// Infoln prints log.
func (l *Logger) Infoln(v ...interface{}) {
	l.printlnAt(INFO, v...)
}

// Warnf prints log.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.printfAt(WARN, format, v...)
}

// Warnln prints log.
func (l *Logger) Warnln(v ...interface{}) {
	l.printlnAt(WARN, v...)
}

// Errorf prints log.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.printfAt(ERROR, format, v...)
	l.printStack(ERROR)
}

// Errorln prints log.
func (l *Logger) Errorln(v ...interface{}) {
	l.printlnAt(ERROR, v...)
	l.printStack(ERROR)
}

// Fatalf prints log.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.printfAt(FATAL, format, v...)
	l.printStack(FATAL)
}

// Fatalln prints log.
func (l *Logger) Fatalln(v ...interface{}) {
	l.printlnAt(FATAL, v...)
	l.printStack(FATAL)
}

// printStack prints stack.
func (l *Logger) printStack(level int) {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, true)
	l.loggers[level].Println("Stack:\n" + string(buf[:n]))
}

// printfAt prints log.
func (l *Logger) printfAt(level int, format string, v ...interface{}) {
	if l.level > level {
		return
	}

	_, file, line, ok := runtime.Caller(2)
	if ok {
		format = file + " " + strconv.Itoa(line) + " " + format
	}

	l.loggers[level].Printf(format, v...)
}

// printlnAt prints log.
func (l *Logger) printlnAt(level int, v ...interface{}) {
	if l.level > level {
		return
	}

	_, file, line, ok := runtime.Caller(2)
	if ok {
		v = append([]interface{}{file + " " + strconv.Itoa(line)}, v...)
	}

	l.loggers[level].Println(v...)
}

// New creates and returns a logger.
func New(levelName string, out io.Writer, flag int) *Logger {
	var level int
	for i, name := range levelNames {
		if levelName == name {
			level = i
			break
		}
	}

	l := &Logger{
		level:   level,
		loggers: make([]*log.Logger, len(levelNames)),
	}

	for i, name := range levelNames {
		l.loggers[i] = log.New(out, name+" ", flag)
	}

	return l
}

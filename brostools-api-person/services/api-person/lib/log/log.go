package log

import (
	//nomyne_lib "nomyne/lib"
	//"time"
	//"errors"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
)

var logLevelName = map[int64]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
	4: "CRITICAL",
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	Logf(0, format, args...)
}

// Infof is like Debugf, but at Info level.
func Infof(ctx context.Context, format string, args ...interface{}) {
	Logf(1, format, args...)
}

// Warningf is like Debugf, but at Warning level.
func Warningf(ctx context.Context, format string, args ...interface{}) {
	Logf(2, format, args...)
}

// Errorf is like Debugf, but at Error level.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	Logf(3, format, args...)
}

// Criticalf is like Debugf, but at Critical level.
func Criticalf(ctx context.Context, format string, args ...interface{}) {
	Logf(4, format, args...)
}

func Logf(level int64, format string, args ...interface{}) {
	//if f, ok := ctx.Value(&logOverrideKey).(logOverrideFunc); ok {
	//	f(level, format, args...)
	//	return
	//}
	//c := fromContext(ctx)
	//if c == nil {
	//	panic(errNotAppEngineContext)
	//}
	logf(level, format, args...)
}

func logf(level int64, format string, args ...interface{}) {
	//if c == nil {
	//	panic("not an App Engine context")
	//}
	s := fmt.Sprintf(format, args...)
	s = strings.TrimRight(s, "\n") // Remove any trailing newline characters.
	//c.addLogLine(&logpb.UserAppLogLine{
	//	TimestampUsec: proto.Int64(time.Now().UnixNano() / 1e3),
	//	Level:         &level,
	//	Message:       &s,
	//})
	// Only duplicate log to stderr if not running on App Engine second generation
	//if !IsSecondGen() {
	log.Print(logLevelName[level] + ": " + s)
	//}

}

func WriteFile(format string) {
	f, err := os.OpenFile("./log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(format)
}

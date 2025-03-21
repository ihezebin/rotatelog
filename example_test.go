package rotatelog

import (
	"log"
	"time"
)

// To use lumberjack with the standard library's log package, just pass it into
// the SetOutput function when your application starts.
func Example() {
	log.SetOutput(&Logger{
		Filename:   "/var/log/myapp/foo.log",
		MaxSize:    500, // kb
		MaxBackups: 3,
		MaxAge:     28 * time.Hour,
		Compress:   true, // disabled by default
	})
}

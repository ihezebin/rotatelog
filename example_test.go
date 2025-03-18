package rotatelog

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// To use lumberjack with the standard library's log package, just pass it into
// the SetOutput function when your application starts.
func TestExample(t *testing.T) {
	BackupFilenameSeparator = "."
	pwd, _ := os.Getwd()
	filename := filepath.Join(pwd, "log/foo.log")

	log.SetOutput(&Rotater{
		Filename:   filename,
		MaxSize:    10, // kb
		MaxBackups: 3,
		MaxAge:     60 * time.Minute,
		Compress:   true, // disabled by default
	})

	for i := 0; i < 300000; i++ {
		log.Println("Hello, World!")
	}
}

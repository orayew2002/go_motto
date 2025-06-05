package log

import (
	"fmt"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	logFile := "../../logs/app.log"
	l := InitSlog(true, logFile)

	t.Run("write error log", func(t *testing.T) {
		l.Error("test log message", "detail", fmt.Sprintf("%+v", "you"))
	})

	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	if len(data) == 0 {
		t.Error("log file is empty, expected log output")
	}

	t.Logf("log output:\n%s", data)

	_ = os.Remove(logFile)
}

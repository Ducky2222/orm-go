package log

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	setLogLevel(InfoLevel)
	if infoLog.Writer() != os.Stdout || errorLog.Writer() != os.Stdout {
		t.Fatal("Log test 0 failed")
	}
	setLogLevel(ErrorLevel)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() != os.Stdout {
		t.Fatal("Log test 1 failed")
	}
	setLogLevel(Disabled)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() == os.Stdout {
		t.Fatal("Log test 0 failed")
	}
}
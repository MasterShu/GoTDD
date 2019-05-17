package args

import (
	"testing"
)

func TestArgs(t *testing.T) {
	t.Run("is Boolean args", testArgsBool("-l", true))
	t.Run("is Boolean args", testArgsNumber("-p 8088", 8088))
	t.Run("is Boolean args", testArgsString("-d /usr/hello", "/usr/hello"))
}

func testArgsString(s string, s2 string) func(t *testing.T) {
	return func(t *testing.T) {
		if ArgsParse(s).d != s2 {
			t.Error("is not pass " + s)
		}
	}
}

func testArgsNumber(s string, i int) func(t *testing.T) {
	return func(t *testing.T) {
		if ArgsParse(s).p != i {
			t.Error("is not pass " + s)
		}
	}
}

func testArgsBool(s string, b bool) func(t *testing.T) {
	return func(t *testing.T) {
		if ArgsParse(s).l != b {
			t.Error("is not pass " + s)
		}
	}
}

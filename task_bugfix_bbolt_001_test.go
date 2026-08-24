package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt001SourceContract(t *testing.T) {
    source, err := os.ReadFile("cursor.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(c.stack) == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}

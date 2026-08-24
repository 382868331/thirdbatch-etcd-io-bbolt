package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt009SourceContract(t *testing.T) {
    source, err := os.ReadFile("db.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err := db.munlock(fileSize); err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}

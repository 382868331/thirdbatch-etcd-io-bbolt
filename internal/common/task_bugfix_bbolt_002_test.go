package common

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt002SourceContract(t *testing.T) {
    source, err := os.ReadFile("page.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(b) == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}

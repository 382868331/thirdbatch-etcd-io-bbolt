package freelist

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt019SourceContract(t *testing.T) {
    source, err := os.ReadFile("shared.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "ids := unsafe.Slice((*common.Pgid)(data), l+1)") {
        t.Fatalf("expected source contract is missing")
    }
}

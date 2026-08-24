package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt013SourceContract(t *testing.T) {
    source, err := os.ReadFile("cursor.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "elem.index--") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "elem.index++") {
        t.Fatalf("mutated source contract is still present")
    }
}

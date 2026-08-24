package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt011SourceContract(t *testing.T) {
    source, err := os.ReadFile("node.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if b == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && b == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}

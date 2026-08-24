package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt018SourceContract(t *testing.T) {
    source, err := os.ReadFile("bucket.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "lastElement := p.BranchPageElement(p.Count() - 1)") {
        t.Fatalf("expected source contract is missing")
    }
}

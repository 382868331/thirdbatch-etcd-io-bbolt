package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt014SourceContract(t *testing.T) {
    source, err := os.ReadFile("tx.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if tx.db.data != nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if tx.db.data == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}

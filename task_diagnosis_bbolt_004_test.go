package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBbolt004SourceContract(t *testing.T) {
    source, err := os.ReadFile("bucket.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if b.tx.db == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}

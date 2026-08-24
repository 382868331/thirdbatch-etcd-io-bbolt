package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt003SourceContract(t *testing.T) {
    source, err := os.ReadFile("db.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if (db.batch == nil) || (db.batch != nil && len(db.batch.calls) >= db.MaxBatchSize) {") {
        t.Fatalf("expected source contract is missing")
    }
}

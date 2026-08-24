package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt006SourceContract(t *testing.T) {
    source, err := os.ReadFile("bucket.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return 0, errors.ErrTxClosed") {
        t.Fatalf("expected source contract is missing")
    }
}

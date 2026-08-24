package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBbolt020SourceContract(t *testing.T) {
    source, err := os.ReadFile("tx_check.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer close(ch)") {
        t.Fatalf("expected source contract is missing")
    }
}

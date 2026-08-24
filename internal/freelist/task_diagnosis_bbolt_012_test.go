package freelist

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBbolt012SourceContract(t *testing.T) {
    source, err := os.ReadFile("hashmap.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "end := ids[0]") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "end := ids[1]") {
        t.Fatalf("mutated source contract is still present")
    }
}

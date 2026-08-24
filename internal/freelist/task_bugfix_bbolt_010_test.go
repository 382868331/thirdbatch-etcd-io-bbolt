package freelist

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt010SourceContract(t *testing.T) {
    source, err := os.ReadFile("shared.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "for id := p.Id(); id <= p.Id()+common.Pgid(p.Overflow()); id++ {") {
        t.Fatalf("expected source contract is missing")
    }
}

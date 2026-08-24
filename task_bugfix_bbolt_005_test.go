package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBbolt005SourceContract(t *testing.T) {
    source, err := os.ReadFile("cursor.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "} else if (flags & uint32(common.BucketLeafFlag)) != 0 {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "} else if (flags & uint32(common.BucketLeafFlag)) == 0 {") {
        t.Fatalf("mutated source contract is still present")
    }
}

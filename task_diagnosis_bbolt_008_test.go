package bbolt

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBbolt008SourceContract(t *testing.T) {
    source, err := os.ReadFile("bolt_solaris.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err := unix.Madvise(b, syscall.MADV_RANDOM); err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}

package issue76126

import (
	"sync"
	"testing"
)

func TestIssue76126(t *testing.T) {
	var wg sync.WaitGroup
	wg.Go(func() {
		panic("test")
	})
	wg.Wait()
}

package nyb

import (
	"strings"
	"testing"
)

func TestChangeNick(t *testing.T) {
	ln := "aaaaaaaaaaaaaaaa"
	for i := 1; i <= 4; i++ {
		ln = changeNick(ln)
		if got := strings.Count(ln, "_"); got != i {
			t.Errorf("expecting %d _'s got %d ", i, got)
		}
		if len(ln) != 12+i {
			t.Errorf("expecting lenght %d got %d", 12+i, len(ln))
		}
	}
	if ln := changeNick(ln); len(ln) != 12 {
		t.Errorf("expecting lenght 12 got %d", len(ln))
	}
	n := "a"
	for i := 1; i <= 15; i++ {
		n = changeNick(n)
		c := strings.Count(n, "_")
		if c != i {
			t.Errorf("expecting %d _'s got %d: string: %s", i, c, n)
		}
	}
	if n := changeNick(n); n != "a" {
		t.Error("n wasn't 'a'")
	}
}

package nyb

import (
	"strings"
	"testing"
)

func TestChangeNick(t *testing.T) {
	ln := "h123456789123456"
	sn := changeNick(ln)
	if len(sn) != 13 {
		t.Errorf("Expecting 13 chars got %d chars", len(sn))
	}
	if !strings.HasSuffix(sn, "_") {
		t.Errorf("%s dosen't have _ suffix", sn)
	}
	sn = changeNick(sn)
	if !strings.HasSuffix(sn, "__") {
		t.Errorf("didn't get __ suffix")
	}
	sn = changeNick(sn)
	if !strings.HasSuffix(sn, "___") {
		t.Errorf("didn't get ___ suffix")
	}
	sn = changeNick(sn)
	if !strings.HasSuffix(sn, "____") {
		t.Errorf("didn't get ____ suffix")
	}
	if len(sn) != 16 {
		t.Error("sn not 16 chars")
	}
	sn = changeNick(sn)
	if sn != ln[:12] {
		t.Errorf("didn't get %s, got %s", ln[:12], sn)
	}
	n := "a"
	for i := 1; i <= 15; i++ {
		n = changeNick(n)
		c := strings.Count(n, "_")
		if c != i {
			t.Errorf("expecting %d _'s got %d: string: %s", i, c, n)
		}
	}

}

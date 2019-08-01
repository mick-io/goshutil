package goshutil

import (
	"testing"
)

func Test_File(t *testing.T) {
	tests := []struct{ in, expect string }{
		{"./goshutil.go", "./goshutil.go: ASCII text"},
	}
	for _, test := range tests {
		out, err := File(test.in)
		if err != nil {
			panic(err)
		}
		if out != test.expect {
			t.Errorf("[FAIL] File(%q) == %q expect %q", test.in, out, test.expect)
		}
	}
}

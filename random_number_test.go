package enc

import (
	"testing"
)

func TestNumber(t *testing.T) {
	tests := []struct {
		label     string
		src       uint
		want      uint
		wantPanic string
	}{
		{
			label: "NormalCase:1",
			src:   0,
			want:  1,
		},
		{
			label: "NormalCase:2",
			src:   1,
			want:  1,
		},
		{
			label: "NormalCase:3",
			src:   10,
			want:  10,
		},
		{
			label: "NormalCase:4",
			src:   18,
			want:  18,
		},
		{
			label:     "PanicCase:1",
			src:       19,
			wantPanic: "Do not exceed 18 digits",
		},
	}
	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			if test.wantPanic != "" {
				defer func() {
					err := recover()
					if err != test.wantPanic {
						t.Errorf("\n\n  got: %s\nwantPanic: %s", err, test.wantPanic)
					}
				}()
			}
			n := RandomNumber(test.src)
			got := uint(len(n))
			if got != test.want {
				t.Errorf("\nnumber: %s\n\n  got: %d\nwant: %d", n, got, test.want)
			}
		})
	}
}

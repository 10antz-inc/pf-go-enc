package enc

import (
	"testing"
)

func TestULID(t *testing.T) {
	tests := []struct {
		label string
		want  int
	}{
		{
			label: "NormalCase:1",
			want:  26,
		},
		{
			label: "NormalCase:2",
			want:  26,
		},
		{
			label: "NormalCase:3",
			want:  26,
		},
	}
	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			got := ULID()
			if len(got) != test.want {
				t.Errorf("\n  got: %d\nwant: %d", len(got), test.want)
			}
		})
	}
}

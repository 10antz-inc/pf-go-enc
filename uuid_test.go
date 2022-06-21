package enc

import (
	"testing"
)

func TestUUID(t *testing.T) {
	tests := []struct {
		label string
		want  int
	}{
		{
			label: "NormalCase:1",
			want:  36,
		},
	}
	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			got := UUID()
			if len(got) != test.want {
				t.Errorf("\n  got: %d\nwant: %d", len(got), test.want)
			}
		})
	}
}

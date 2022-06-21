package enc

import (
	"testing"
)

func TestSUUID(t *testing.T) {
	tests := []struct {
		label string
		want  int
	}{
		{
			label: "case:1",
			want:  22,
		},
		{
			label: "case:2",
			want:  22,
		},
		{
			label: "case:3",
			want:  22,
		},
	}
	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			got := SUUID()
			if len(got) != test.want {
				t.Errorf("\n  got: %d\nwant: %d", len(got), test.want)
			}
		})
	}
}

func TestUUIDtoSUUID(t *testing.T) {
	tests := []struct {
		label string
		in    string
		out   string
	}{
		{
			label: "case:1",
			in:    "d52f3060-b1d5-4080-be54-ea6185c0e546",
			out:   "cSUijkCsToFqDGZPXhZ2wf",
		},
		{
			label: "case:2",
			in:    "b941f419-7415-455f-85ba-0f0861b85c64",
			out:   "hb5SKJEu6ksDffe7Yy6oxa",
		},
		{
			label: "case:3",
			in:    "5b2f11cf-6149-481c-9b72-7421b6c86e93",
			out:   "P8QSTZHM3sxoZjuyQGHkEJ",
		},
	}
	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			got, _ := UUIDtoSUUID(test.in)
			if got != test.out {
				t.Errorf("\n  got: %s\nwant: %s", got, test.out)
			}
		})
	}
}

func TestSUUIDtoUUID(t *testing.T) {
	tests := []struct {
		label string
		in    string
		out   string
	}{
		{
			label: "case:1",
			in:    "cSUijkCsToFqDGZPXhZ2wf",
			out:   "d52f3060-b1d5-4080-be54-ea6185c0e546",
		},
		{
			label: "case:2",
			in:    "hb5SKJEu6ksDffe7Yy6oxa",
			out:   "b941f419-7415-455f-85ba-0f0861b85c64",
		},
		{
			label: "case:3",
			in:    "P8QSTZHM3sxoZjuyQGHkEJ",
			out:   "5b2f11cf-6149-481c-9b72-7421b6c86e93",
		},
	}
	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			got, _ := SUUIDtoUUID(test.in)
			if got != test.out {
				t.Errorf("\n  got: %s\nwant: %s", got, test.out)
			}
		})
	}
}

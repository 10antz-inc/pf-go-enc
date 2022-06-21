package enc

import (
	"testing"
)

func TestSHA256(t *testing.T) {
	tests := []struct {
		label string
		src   string
		want  string
	}{
		{
			label: "NormalCase:1",
			src:   "hoge",
			want:  "ecb666d778725ec97307044d642bf4d160aabb76f56c0069c71ea25b1e926825",
		},
		{
			label: "NormalCase:2",
			src:   "fuga",
			want:  "9b9f7060d3290054631e4f2a081601add7307f3f873e9bca60514fb3bddb566c",
		},
	}
	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			got := SHA256(test.src)
			if got != test.want {
				t.Errorf("\nsrc: %s\n\n  got: %s\nwant: %s", test.src, got, test.want)
			}
		})
	}
}

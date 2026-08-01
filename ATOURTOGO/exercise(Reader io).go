package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	rot13_dict := map[byte]byte{
		65: 78, 66: 79, 67: 80, 68: 81, 69: 82, 70: 83, 71: 84, 72: 85, 73: 86,
		74: 87, 75: 88, 76: 89, 77: 90, 78: 65, 79: 66, 80: 67, 81: 68, 82: 69,
		83: 70, 84: 71, 85: 72, 86: 73, 87: 74, 88: 75, 89: 76, 90: 77, 97: 110,
		98: 111, 99: 112, 100: 113, 101: 114, 102: 115, 103: 116, 104: 117,
		105: 118, 106: 119, 107: 120, 108: 121, 109: 122, 110: 97, 111: 98,
		112: 99, 113: 100, 114: 101, 115: 102, 116: 103, 117: 104, 118: 105,
		119: 106, 120: 107, 121: 108, 122: 109,
	}

	for {
		n, err := rot.r.Read(b)
		sliceB := b[:n]
		for i := range sliceB {
			c, ok := rot13_dict[sliceB[i]]
			if ok {
				sliceB[i] = c
			}
		}
		b = sliceB
		return len(b) ,err
	}
}

func main() {
	s := strings.NewReader("Jnuu, orgr fvxu tln o")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

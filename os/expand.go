package os

import (
	"errors"
)

var (
	ErrNoReader = errors.New("no reader")
	ErrNoWriter = errors.New("nil writer")
)

// func ExpandEnv(src io.Reader, dst io.Writer) error {
// 	if src == nil {
// 		return ErrNoReader
// 	}
// 	if dst == nil {
// 		return ErrNoWriter
// 	}
// 	const bufSize = 100
// 	buf := make([]byte, bufSize)
// 	for {
// 		n, err := src.Read(buf)
// 		if n > 0 {

// 		}
// 	}
// 	os.Expand().ErrUnexpectedEOF()
// }

// func Expands(target []byte, mapping func([]byte) []byte) []byte {
// 	var buf []byte
// 	// ${} is all ASCII, so bytes are fine for this operation.
// 	i := 0
// 	for j := 0; j < len(s); j++ {
// 		if s[j] == '$' && j+1 < len(s) {
// 			if buf == nil {
// 				buf = make([]byte, 0, 2*len(s))
// 			}
// 			buf = append(buf, s[i:j]...)
// 			name, w := getShellName(s[j+1:])
// 			if name == "" && w > 0 {
// 				// Encountered invalid syntax; eat the
// 				// characters.
// 			} else if name == "" {
// 				// Valid syntax, but $ was not followed by a
// 				// name. Leave the dollar character untouched.
// 				buf = append(buf, s[j])
// 			} else {
// 				buf = append(buf, mapping(name)...)
// 			}
// 			j += w
// 			i = j + 1
// 		}
// 	}
// 	if buf == nil {
// 		return s
// 	}
// 	return string(buf) + s[i:]
// }

// MapEnv
//   - ${parameter:-word}
//     -
func MapEnv() {

}

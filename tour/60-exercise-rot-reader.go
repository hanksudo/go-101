// ROT13 ("rotate by 13 places")
/* https://en.wikipedia.org/wiki/ROT13
   a~z 97~122
   A~Z 65~90

   A => N
   (A(65) - A(65) + 13) % 26 + A(65) = N(78)
   (N(78) - A(65) + 13) % 26 + A(65) = A(65)

   M => Z
   (M(77) - A(65) + 13) % 26 + A(65) = Z(90)
   (Z(90) - A(65) + 13) % 26 + A(65) = M(77)
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(buffer []byte) (n int, err error) {
	n, err = r.r.Read(buffer)

	for i := 0; i < n; i++ {
		buffer[i] = rot13(buffer[i])
	}
	return
}

func main() {
	secretCode := "Lbh penpxrq gur pbqr!"
	s := strings.NewReader(secretCode)
	r := rot13Reader{s}

	fmt.Println(secretCode)
	io.Copy(os.Stdout, &r)
}

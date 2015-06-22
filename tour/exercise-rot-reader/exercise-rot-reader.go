/* https://en.wikipedia.org/wiki/ROT13
   a~z 97~122
   A~Z 65~90
   (A(65) + 13) % 26 = N(78)
   (Z(90) + 13) - 26 = M(77)
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(buffer []byte) (n int, err error) {
	n, err = r.r.Read(buffer)

	for i := 0; i < n; i++ {
		b := buffer[i]
		switch {
		case 'a' <= b && b <= 'z':
			buffer[i] = (b-'a'+13)%26 + 'a'
		case 'A' <= b && b <= 'Z':
			buffer[i] = (b-'A'+13)%26 + 'A'
		}
	}
	return
}

func main() {
	secretCode := "Lbh penpxrq gur pbqr!"
	fmt.Println(secretCode)
	s := strings.NewReader(secretCode)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

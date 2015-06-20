// http://tour.golang.org/methods/11
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (m MyReader) Read(buffer []byte) (int, error) {
	buffer[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}

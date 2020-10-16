package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

func (erase spaceEraser) Read(sentence []byte) (int, error) {
	n, err := erase.r.Read(sentence)
	counter := 0
	for i := 0; i < n; i++ {
		if sentence[i] != 32 {
			sentence[counter] = sentence[i]
			counter++

		}
	}
	return counter, err
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"

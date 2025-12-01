package utils

import (
	"bufio"
	"io"
	"os"
)

type file struct {
	*os.File
}

type scanner struct {
	*bufio.Scanner
}

func ReadFile(filename string) []string {
	file, err := openFile(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := newScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func openFile(filename string) (*file, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return &file{f}, nil
}

func (f *file) Close() error {
	return f.File.Close()
}

func newScanner(r io.Reader) *scanner {
	return &scanner{bufio.NewScanner(r)}
}

func (s *scanner) Scan() bool {
	return s.Scanner.Scan()
}

func (s *scanner) Text() string {
	return s.Scanner.Text()
}

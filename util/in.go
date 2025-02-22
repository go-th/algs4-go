package util

import (
	"bufio"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Reads in data of various types from standard input, files and URLs.

type In struct {
	reader     io.Reader
	scanner    *bufio.Scanner
	hasScanned bool
	hasNext    bool
}

// Factory method
// default read in lines
func NewIn(uri string) *In {
	return NewInReadLines(uri)
}

func NewInReadWords(uri string) *In {
	r := newReader(uri)

	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanWords)

	return &In{reader: r, scanner: scanner}
}

func NewInReadLines(uri string) *In {
	r := newReader(uri)

	scanner := bufio.NewScanner(r)

	return &In{reader: r, scanner: scanner}
}

func newReader(uri string) io.Reader {
	if uri == "" {
		panic("argument is empty")
	}

	// first try to read file from local file system
	f, err := os.Open(uri)
	if err == nil {
		if strings.HasSuffix(uri, ".gz") {
			gz, err := gzip.NewReader(f)
			if err != nil {
				panic(err)
			}
			return gz
		} else {
			return f
		}
	} else {
		// URL from web
		resp, err := http.Get(uri)
		if err != nil {
			panic(err)
		}
		return resp.Body
	}

}

func (in *In) ReadString() string {
	in.next()
	return in.scanner.Text()
}

func (in *In) ReadLine() string {
	in.next()
	return in.scanner.Text()
}

func (in *In) ReadInt() int {
	i, _ := strconv.Atoi(in.ReadString())
	return i
}

func (in *In) IsEmpty() bool {
	return !in.HasNext()
}

func (in *In) HasNext() bool {
	if !in.hasScanned {
		in.hasNext = in.scanner.Scan()
		in.hasScanned = true
	}
	return in.hasNext
}

func (in *In) next() bool {
	if in.hasScanned {
		in.hasScanned = false
		return in.hasNext
	}
	return in.scanner.Scan()
}

func (in *In) readAll() string {
	data, err := io.ReadAll(in.reader)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func (in *In) ReadAllStrings() []string {
	str := in.readAll()

	return strings.Fields(str)
}

func (in *In) ReadAllInts() []int {
	strSlice := in.ReadAllStrings()
	n := len(strSlice)
	intSlice := make([]int, n)

	for i := 0; i < n; i++ {
		intSlice[i], _ = strconv.Atoi(strSlice[i])
	}

	return intSlice
}
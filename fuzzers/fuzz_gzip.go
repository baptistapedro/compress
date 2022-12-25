package myfuzz

import (
	"github.com/klauspost/compress/gzip"
	"bytes"
	"time"
)

func Fuzz(data []byte) int {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	
	zw.Name = string(data)
	zw.Comment = string(data)
	zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write(data)
	if err != nil {
		return 1
	}
	if err := zw.Close(); err != nil {
		return 1
	}
	_, err = gzip.NewReader(&buf)
	if err != nil {
		return 1
	}
	return 0
}

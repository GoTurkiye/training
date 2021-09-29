package main

import (
	"bytes"
	"compress/gzip"
)

type GzipWithoutPool struct{}

func NewGzipWithoutPool() *GzipWithoutPool {
	return &GzipWithoutPool{}
}

func (g *GzipWithoutPool) Zip(s string) error {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)

	_, err := w.Write([]byte(s))

	if err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}
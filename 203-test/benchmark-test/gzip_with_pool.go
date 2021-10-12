package main

import (
	"bytes"
	"compress/gzip"
	"sync"
)

type GzipWithPool struct {
	pool sync.Pool
}

func NewGzipWithPool() *GzipWithPool {
	return &GzipWithPool{
		pool: sync.Pool{
			New: func() interface{} {
				var buf bytes.Buffer
				return gzip.NewWriter(&buf)
			},
		},
	}
}

func (g *GzipWithPool) GetGzipWriter() *gzip.Writer {
	return g.pool.Get().(*gzip.Writer)
}

func (g *GzipWithPool) PutGzipWriter(buf *gzip.Writer) {
	buf.Flush()
	g.pool.Put(buf)
}

func (g *GzipWithPool) Zip(s string) error {
	w := g.GetGzipWriter()

	_, err := w.Write([]byte(s))

	g.PutGzipWriter(w)

	if err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil
}
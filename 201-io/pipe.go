package main

import (
	"fmt"
	"io"
	"time"
)

func main() {
	pReader, pWriter := io.Pipe()
	done := make(chan struct{})

	go read(pReader, done)
	go write(pWriter)

	<-done
}

func read(reader io.Reader, done chan struct{}) {
	buff := make([]byte, 1024)
	for {
		readed, err := reader.Read(buff)
		if readed == 0 {
			if err == io.EOF {
				done <- struct{}{}
				break
			}

			if err != nil {
				fmt.Println(err)
				done <- struct{}{}
				break
			}
		} else {
			fmt.Println(buff[:readed])

			if err == io.EOF || err != nil {
				fmt.Println(err)
				done <- struct{}{}
				break
			}
		}
	}
}

func write(writer *io.PipeWriter) {
	i := 0
	for {
		if i == 10 {
			writer.Close()
			break
		}
		writer.Write([]byte(string(i)))
		i++
		time.Sleep(time.Millisecond * 100)
	}
}

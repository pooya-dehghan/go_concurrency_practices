package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

var buf = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, debug string) {
	b := buf.Get().(*bytes.Buffer)

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(debug)
	b.WriteString("\n")

	w.Write(b.Bytes())
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}

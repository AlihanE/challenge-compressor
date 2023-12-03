package writer

import (
	"os"
	"strings"
)

type Writer struct {
	file *os.File
}

func New(name string, ext string, postfix string) *Writer {
	nameWithoutExt := GetFileNameWithoutExt(name)
	f, err := os.Create("./" + nameWithoutExt + postfix + "." + ext)
	if err != nil {
		panic(err)
	}
	return &Writer{file: f}
}

func (w *Writer) Write(b []byte) (int, error) {
	return w.file.Write(b)
}

func (w *Writer) Close() {
	w.file.Close()
}

func GetFileNameWithoutExt(name string) string {
	return name[:strings.Index(name, ".")]
}

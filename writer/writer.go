package writer

import (
	"bufio"
	"os"
	"strings"
)

type Writer struct {
	Writer *bufio.Writer
	file   *os.File
}

func New(name string, ext string, postfix string) *Writer {
	nameWithoutExt := GetFileNameWithoutExt(name)
	f, err := os.Create("./" + nameWithoutExt + postfix + "." + ext)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	return &Writer{Writer: w, file: f}
}

func (w *Writer) Close() {
	w.file.Close()
}

func GetFileNameWithoutExt(name string) string {
	return name[:strings.Index(name, ".")]
}

package output

import (
	"fmt"
	"os"
)

// Writer handles output writing to file or stdout
type Writer struct {
	filePath string
}

// NewWriter creates a new writer
func NewWriter(filePath string) *Writer {
	return &Writer{
		filePath: filePath,
	}
}

// Write outputs data to file or stdout
func (w *Writer) Write(data string) error {
	if w.filePath == "" {
		// Write to stdout
		fmt.Print(data)
		return nil
	}

	// Write to file
	file, err := os.Create(w.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

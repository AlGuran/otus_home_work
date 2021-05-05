package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrNegativeOffset        = errors.New("negative offset")
	ErrNegativeLimit         = errors.New("negative limit")
)

func Copy(fromPath, toPath string, offset, limit int64, test bool) error {
	if offset < 0 {
		return ErrNegativeOffset
	}
	if limit < 0 {
		return ErrNegativeLimit
	}

	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 {
		return ErrUnsupportedFile
	}
	if fileInfo.Size() <= offset {
		return ErrOffsetExceedsFileSize
	}

	fileFrom, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fileFrom.Close()
	fileFrom.Seek(offset, io.SeekStart)

	fileTo, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer fileTo.Close()

	if limit == 0 {
		limit = fileInfo.Size()
	}
	if limit > fileInfo.Size()-offset {
		limit = fileInfo.Size() - offset
	}

	buf := make([]byte, limit)
	count, err := fileFrom.Read(buf)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(buf)
	total := int64(count)

	chunks := getPercent(float64(total), float64(100))

	for writeSize := int64(0); writeSize < total; {
		written, err := io.CopyN(fileTo, reader, chunks)

		writeSize += written
		if !test {
			progressBar(getPercent(float64(writeSize*100), float64(total)))
		}

		// for view run the progress bar uncomment
		// time.Sleep(time.Millisecond * 100)

		if errors.Is(err, io.EOF) {
			err = nil
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func progressBar(percent int64) {
	fmt.Print("\r" + strconv.Itoa(int(percent)) + strings.Repeat("=", int(percent)) + ">")
}

func getPercent(divisor, dividend float64) int64 {
	return int64(math.Ceil(divisor / dividend))
}

package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("copy all", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/test.txt", 0, 0, true)
		require.NoError(t, err)

		fileFrom, _ := os.Open("testdata/input.txt")
		fileFromStat, _ := fileFrom.Stat()
		bufFrom := make([]byte, fileFromStat.Size())
		fileFrom.Read(bufFrom)

		fileTo, _ := os.Open("/tmp/test.txt")
		fileToStat, _ := fileTo.Stat()
		bufTo := make([]byte, fileToStat.Size())
		fileTo.Read(bufTo)

		require.Equal(t, bufFrom, bufTo)
	})

	t.Run("copy 10 bytes", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/test.txt", 0, 10, true)
		require.NoError(t, err)

		fileFrom, _ := os.Open("testdata/input.txt")
		bufFrom := make([]byte, 10)
		fileFrom.Read(bufFrom)

		fileTo, _ := os.Open("/tmp/test.txt")
		fileToStat, _ := fileTo.Stat()
		bufTo := make([]byte, fileToStat.Size())
		fileTo.Read(bufTo)

		require.Equal(t, bufFrom, bufTo)
	})

	t.Run("copy 10 bytes offset 10 bytes", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/test.txt", 10, 10, true)
		require.NoError(t, err)

		fileFrom, _ := os.Open("testdata/input.txt")
		bufFrom := make([]byte, 10)
		fileFrom.Seek(10, io.SeekStart)
		fileFrom.Read(bufFrom)

		fileTo, _ := os.Open("/tmp/test.txt")
		fileToStat, _ := fileTo.Stat()
		bufTo := make([]byte, fileToStat.Size())
		fileTo.Read(bufTo)

		require.Equal(t, bufFrom, bufTo)
	})

	t.Run("limit more than file size", func(t *testing.T) {
		err := Copy("testdata/input.txt", "/tmp/test.txt", 0, 7000, true)
		require.NoError(t, err)

		fileFrom, _ := os.Open("testdata/input.txt")
		fileFromStat, _ := fileFrom.Stat()
		bufFrom := make([]byte, fileFromStat.Size())
		fileFrom.Read(bufFrom)

		fileTo, _ := os.Open("/tmp/test.txt")
		fileToStat, _ := fileTo.Stat()
		bufTo := make([]byte, fileToStat.Size())
		fileTo.Read(bufTo)

		require.Equal(t, bufFrom, bufTo)
	})
}

func TestCopyInvalidData(t *testing.T) {
	t.Run("Invalid data", func(t *testing.T) {
		err := Copy("/dev/urandom", "/tmp/test.txt", 0, 0, true)
		require.Equal(t, err, ErrUnsupportedFile)

		err = Copy("testdata/input.txt", "/tmp/test.txt", -1, 0, true)
		require.Equal(t, err, ErrNegativeOffset)

		err = Copy("testdata/input.txt", "/tmp/test.txt", 0, -1, true)
		require.Equal(t, err, ErrNegativeLimit)

		err = Copy("testdata/input.txt", "/tmp/test.txt", 10000, 0, true)
		require.Equal(t, err, ErrOffsetExceedsFileSize)
	})
}

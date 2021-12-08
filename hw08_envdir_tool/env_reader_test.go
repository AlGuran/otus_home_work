package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	testData := "testdata/env"
	_, err := ioutil.ReadDir(testData)

	t.Run("Test '='", func(t *testing.T) {
		tempFile, _ := ioutil.TempFile(testData, "Test=File")
		defer os.Remove(tempFile.Name())

		env, _ := ReadDir(testData)
		require.Nil(t, err)
		require.Equal(t, "", env["Test=File"].Value)
	})

	t.Run("Test remove ' ' and '\t'", func(t *testing.T) {
		tempFile, _ := ioutil.TempFile(testData, "TestFile")
		defer os.Remove(tempFile.Name())

		ioutil.WriteFile(tempFile.Name(), []byte("test string   \t"), os.ModePerm)
		envMap, err := ReadDir(testData)
		require.Nil(t, err)
		require.Equal(t, "test string", envMap[filepath.Base(tempFile.Name())].Value)
	})

	t.Run("Test replace '\x00'", func(t *testing.T) {
		tempFile, _ := ioutil.TempFile(testData, "TestFile")
		defer os.Remove(tempFile.Name())

		ioutil.WriteFile(tempFile.Name(), []byte("text about\x00 replacement"), os.ModePerm)
		envMap, err := ReadDir(testData)
		require.Nil(t, err)
		require.Equal(t, "text about\n replacement", envMap[filepath.Base(tempFile.Name())].Value)
	})

	t.Run("Test empty file", func(t *testing.T) {
		tempFile, _ := ioutil.TempFile(testData, "TestFile")
		defer os.Remove(tempFile.Name())

		envMap, err := ReadDir(testData)
		require.Nil(t, err)
		require.Equal(t, "", envMap[filepath.Base(tempFile.Name())].Value)
	})
}

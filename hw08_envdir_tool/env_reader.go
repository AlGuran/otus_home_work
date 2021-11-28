package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)
	for _, f := range files {
		if strings.Contains(f.Name(), "=") {
			continue
		}

		file, err := os.Open(dir + "/" + f.Name())
		if err != nil {
			return nil, err
		}

		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}

		str = strings.TrimRight(str, "\n\t ")
		str = strings.ReplaceAll(str, "\x00", "\n")

		env[f.Name()] = EnvValue{Value: str}
	}

	return env, nil
}

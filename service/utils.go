package service

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func getFileExtension(path string) string {
	args := strings.Split(path, ".")
	if len(args) == 0 {
		return ""
	}

	return args[len(args)-1]
}

func changeExtension(path string, newExt string) string {
	args := strings.Split(path, ".")
	if len(args) == 0 {
		panic("pls check extension not empty before pass to this func")
	}

	args[len(args)-1] = newExt
	return strings.Join(args, ".")
}

func createFile(path string, raw []byte) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err, "create file error")
		return err
	}

	defer f.Close()
	_, err = io.Copy(f, bytes.NewReader(raw))
	if err != nil {
		fmt.Println(err, "copy data to file error")
		return err
	}

	fmt.Println("write file success: ", path)
	return nil
}

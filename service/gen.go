package service

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type IGenerator interface {
	Gen(string) error
}

type _service struct {
	// projectRootPath the location of result
	projectRootPath string

	// tmplRootPath the location of template, will map 1-1 with result
	tmplRootPath string

	// module in go.mod
	module string
}

func New(tmplRootPath, projectRootPath, module string) IGenerator {
	return &_service{
		projectRootPath: projectRootPath,
		tmplRootPath:    tmplRootPath,
		module:          module,
	}
}

func (s *_service) Gen(serviceName string) error {
	return filepath.Walk(s.tmplRootPath, s.customTravel)
}

func (s *_service) makeDir(path string) error {
	resultPath := strings.Replace(path, s.tmplRootPath, s.projectRootPath, 1)
	return os.Mkdir(resultPath, 0777)
}

func (s *_service) genWithTmpl(path string, newExt string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Println(err, "read file error")
		return err
	}

	tmpl, err := template.New("x").Parse(string(raw))
	if err != nil {
		fmt.Println(err, "parse template error")
		return err
	}

	rawResult := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(rawResult, GenInput{
		Module: s.module,
	})

	if err != nil {
		log.Println(err, "exe template error")
		return err
	}

	resultPath := strings.Replace(path, s.tmplRootPath, s.projectRootPath, 1)
	resultPath = changeExtension(resultPath, newExt)

	return createFile(resultPath, rawResult.Bytes())
}

func (s *_service) customTravel(path string, info fs.FileInfo, err error) error {
	if path == s.tmplRootPath {
		return nil
	}

	if info.IsDir() {
		return s.makeDir(path)
	}

	ext := getFileExtension(path)
	newExt, shouldExeTemplate := tmplReplacer[ext]
	if shouldExeTemplate {
		return s.genWithTmpl(path, newExt)
	}

	fmt.Println("Travel: ", path)
	raw, err := os.ReadFile(path)
	if err != nil {
		log.Println(err, "read file error")
		return err
	}

	fmt.Println("read file success", path)

	resultPath := strings.Replace(path, s.tmplRootPath, s.projectRootPath, 1)
	return createFile(resultPath, raw)
}

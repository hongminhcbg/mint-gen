package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"

	"github.com/hongminhcbg/mint-gen/service"
)

func getUserChoice() (templateName, projectName, goModule string, err error) {
	dirs, err := os.ReadDir("templates")
	if err != nil {
		return
	}

	tmpls := make([]string, 0)
	for i := 0; i < len(dirs); i++ {
		tmpls = append(tmpls, dirs[i].Name())
	}

	prompt := promptui.Select{
		Label: "Select project template",
		Items: tmpls,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
	templateName = result

	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("don't allow empty")
		}
		return nil
	}

	promptProjectName := promptui.Prompt{
		Label:    "ProjectName",
		Validate: validate,
	}

	projectName, err = promptProjectName.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	promptGomod := promptui.Prompt{
		Label:    "Go module",
		Validate: validate,
	}

	goModule, err = promptGomod.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	return
}

func gen(tmplName, project, gomod string) {
	err := os.Mkdir(project, 0777)
	if err != nil {
		panic(err)
	}

	gen := service.New("templates/"+tmplName, project, gomod)
	err = gen.Gen(project)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Only for unix")
	tmplName, project, goModule, err := getUserChoice()
	if err != nil {
		panic(err)
	}

	fmt.Printf("starting generate project: '%s' with template '%s' and go module '%s' ", project, tmplName, goModule)
	gen(tmplName, project, goModule)
}

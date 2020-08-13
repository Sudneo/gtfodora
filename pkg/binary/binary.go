package binary

import (
	"fmt"
	"strings"
)

type Binary struct {
	Name        string
	Type        string
	Description string
	Commands    []Command
}

type Command struct {
	Function string
	Details  []FunctionSpec
}

type FunctionSpec struct {
	Description string
	Code        string
}

func (b *Binary) HasFunction(function string) (bool, []FunctionSpec) {
	for _, cmd := range b.Commands {
		if cmd.Function == function {
			return true, cmd.Details
		}
	}
	return false, []FunctionSpec{}
}

func (b *Binary) Print() {
	fmt.Printf("Information about: %v\n", b.Name)
	if b.Description != "" {
		fmt.Printf("Description:\n%v\n", b.Description)
	}
	for _, cmd := range b.Commands {
		fmt.Printf("[+] %v:\n", cmd.Function)
		fmt.Println(strings.Repeat("-", len(cmd.Function)+5))
		for _, f := range cmd.Details {
			f.Print()
		}
	}
}

func (f *FunctionSpec) Print() {
	if len(f.Description) > 0 {
		fmt.Printf("- Description:\n")
		fmt.Printf("%v\n", f.Description)
	}
	if len(f.Code) > 0 {
		fmt.Printf("- Code:\n")
		fmt.Printf("%s\n", f.Code)
	}
	if len(f.Code) > 0 || len(f.Description) > 0 {
		fmt.Printf("\n")
	}

}

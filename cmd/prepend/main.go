package main

import (
	"fmt"
	λ "github.com/4thel00z/lambda/v1"
	"os"
)

func printUsage() {
	fmt.Printf(`usage:				echo "stuff" | %s <filename>
filename:			Name of the file which should get "stuff" prepended
`, os.Args[0])

	os.Exit(64) // EX_USAGE
}

func die(message string) func(error) error {
	return func(err error) error {
		fmt.Printf(`Something went wrong while %s:
%s
`, message, err.Error())
		os.Exit(70) // EX_SOFTWARE
		return nil
	}
}
func main() {
	if len(os.Args) != 2 {
		printUsage()
	}

	toPrepend := λ.Slurp(os.Stdin).UnwrapString()
	content := λ.Open(os.Args[1]).Slurp().UnwrapString()
	λ.Create(os.Args[1]).Catch(die("WriteFromString")).WriteFromString(toPrepend + content)

}

package cli

import (
	"fmt"
	"os"
	clr     "github.com/gookit/color"
)

func Notify(a ...interface{}) {
	clr.Printf("<green>::</> ")
	for i := range a {
		clr.Print(a[i])
		print(" ")
	}
	fmt.Println()
}

func Warning(a ...interface{}) {
	clr.Fprintf(os.Stderr, "<lightYellow>::</> ")
	for i := range a {
		clr.Print(a[i])
		print(" ")
	}
	fmt.Println()
}

func Fatal(a ...interface{}) {
	clr.Fprintf(os.Stderr, "<red>::</> ")
	for i := range a {
		clr.Print(a[i])
		print(" ")
	}
	fmt.Println()
}

package main

import (
	"fmt"
	"io"
	"os"
)

func strPrint(w io.Writer, a []interface{}) {
	for _, e := range a {
		if str, ok := e.(string); ok {
			str = fmt.Sprintf("%s\n", str)
			io.WriteString(os.Stdout, str)
		}
	}
}

func main() {
	f := os.Stdout
	strPrint(f, []interface{}{"Hi", 1, 2, "0", "Bey!", 'd'})
}

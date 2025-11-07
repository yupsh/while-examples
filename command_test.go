package while_test

import (
	"fmt"
	"os"

	yup "github.com/gloo-foo/framework"
	"github.com/yupsh/echo"
	"github.com/yupsh/grep"
	. "github.com/gloo-foo/pipe"
	"github.com/yupsh/while"
)

func ExampleWhile() {
	// echo "1\n2\n3" | while read line; do echo "Line: $line"; done
	if err := yup.Run(Pipeline(
		echo.Echo("1\n2\n3"),
		while.While(grep.Grep(".")),
	)); err != nil {
		fmt.Fprintf(os.Stderr, "while: %v\n", err)
	}
	// Output:
	// 1
	// 2
	// 3
}

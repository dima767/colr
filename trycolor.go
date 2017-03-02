package main

import (
	"github.com/dima767/mypackage/pkg1"
	"github.com/dima767/mypackage/pkg2"
	"github.com/fatih/color"
)

func main() {

	color.Cyan("Prints text in cyan.")

	// a newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// More default foreground colors..
	color.Red(pkg1.Word())
	color.White(pkg2.Word())
	color.Yellow("Yellow color too!")
	color.Magenta("And many others ..")

	// Mix up foreground and background colors, create new mixes!
	/*red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)
	boldRedPrintf := boldRed.FprintFunc()
	boldRedPrintf("This will print text in bold red.")*/

	// Create a custom print function for convenient
	red := color.New(color.Bold, color.FgHiGreen).PrintfFunc()
	red("SUCCESS..\n")

}

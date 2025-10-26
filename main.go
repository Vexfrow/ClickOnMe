package main

import (
	"ClickOnMe/cmd"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewFigure("Click On Me", "larry3d", true)
	myFigure.Print()

	cmd.Execute()
}

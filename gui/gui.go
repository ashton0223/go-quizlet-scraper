package main

import (
	_ "embed"
	"fmt"

	"github.com/webview/webview"
)

var (
	//go:embed index.html
	index string
)

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle("test")
	w.SetSize(800, 600, webview.HintNone)

	w.Navigate(fmt.Sprintf("data:text/html,%s", index))
	w.Run()
}

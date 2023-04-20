// Package bannerx
// Date: 2022/10/2 22:31
// Author: Amu
// Description:
package bannerx

import (
	"fmt"

	"github.com/mattn/go-colorable"

	"github.com/dimiro1/banner"
)

func GenerateBanner(bannerString string) {
	templ := fmt.Sprintf(`{{ .Title "%s" "" 4 }}
	  {{ .AnsiColor.BrightRed }}GoVersion: {{ .GoVersion }}
	  {{ .AnsiColor.BrightRed }}GOOS: {{ .GOOS }}
	  {{ .AnsiColor.BrightGreen }}GOARCH: {{ .GOARCH }}
	  {{ .AnsiColor.BrightGreen }}NumCPU: {{ .NumCPU }}
	  {{ .AnsiColor.BrightGreen }}GOPATH: {{ .GOPATH }}
	  {{ .AnsiColor.BrightGreen }}GOROOT: {{ .GOROOT }}
	  Compiler: {{ .Compiler }}
	  ENV: {{ .Env "GOPATH" }}
	  Now: {{ .Now "Monday, 2 Jan 2006" }}`, bannerString)
	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
}

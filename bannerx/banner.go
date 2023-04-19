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

func GenerateBanner() {
	templ := `{{ .Title "Banner" "" 4 }}`
	temp := "{{ .Title " + "\"Hello\"" + " \"\" 4 }}"
	fmt.Println(templ)
	fmt.Println(temp)
	isEnabled := true
	isColorEnabled := true
	banner.InitString(colorable.NewColorableStdout(), isEnabled, isColorEnabled, temp)
}

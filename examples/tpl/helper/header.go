// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: examples/tpl/helper/header.gohtml

package helper

import (
	"io"
	"strings"
)

func Header() string {
	var _b strings.Builder
	RenderHeader(&_b)
	return _b.String()
}

func RenderHeader(_buffer io.StringWriter) {
	_buffer.WriteString("<div>Welcome</div>")

}

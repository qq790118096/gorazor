// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: docs/hello/tpl/index.gohtml

package tpl

import (
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
	"time"
)

func Index() string {
	var _b strings.Builder
	RenderIndex(&_b)
	return _b.String()
}

func RenderIndex(_buffer io.StringWriter) {
	_buffer.WriteString("\n\n<p>This is Index</p>")

	t := time.Now()
	StrTime := t.Format("2006-01-02 15:04:05")

	_buffer.WriteString("<p>Time now is:  ")
	_buffer.WriteString(gorazor.HTMLEscStr(StrTime))
	_buffer.WriteString(" </p>")

}

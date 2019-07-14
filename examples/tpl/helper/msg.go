// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: examples/tpl/helper/msg.gohtml

package helper

import (
	"github.com/sipin/gorazor/examples/models"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

func Msg(u *models.User) string {
	var _b strings.Builder
	RenderMsg(&_b, u)
	return _b.String()
}

func RenderMsg(_buffer io.StringWriter, u *models.User) {

	username := u.Name
	if u.Email != "" {
		username += "(" + u.Email + ")"
	}

	_buffer.WriteString("\n<div class=\"welcome\">\n<h4>Hello ")
	_buffer.WriteString(gorazor.HTMLEscStr(username))
	_buffer.WriteString("</h4>\n\n<div>")
	_buffer.WriteString((u.Intro))
	_buffer.WriteString("</div>\n</div>")

}

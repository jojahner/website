// This file is automatically generated by qtc from "stylesheet.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/stylesheet.qtpl:1
package templates

//line templates/stylesheet.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/stylesheet.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/stylesheet.qtpl:1
func StreamStylesheet(qw422016 *qt422016.Writer) {
	//line templates/stylesheet.qtpl:1
	qw422016.N().S(`
body {
  background-color: #e3e3e3;

  display: flex;
  justify-content: center;
  align-items: center;
}
`)
//line templates/stylesheet.qtpl:9
}

//line templates/stylesheet.qtpl:9
func WriteStylesheet(qq422016 qtio422016.Writer) {
	//line templates/stylesheet.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/stylesheet.qtpl:9
	StreamStylesheet(qw422016)
	//line templates/stylesheet.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line templates/stylesheet.qtpl:9
}

//line templates/stylesheet.qtpl:9
func Stylesheet() string {
	//line templates/stylesheet.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/stylesheet.qtpl:9
	WriteStylesheet(qb422016)
	//line templates/stylesheet.qtpl:9
	qs422016 := string(qb422016.B)
	//line templates/stylesheet.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/stylesheet.qtpl:9
	return qs422016
//line templates/stylesheet.qtpl:9
}
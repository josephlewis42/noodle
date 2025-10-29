package reporter

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"text/tabwriter"
)

// Markdown reporter
type Markdown struct {
	// Destination for the markdown report.
	Writer io.Writer
}

var _ Reporter = (*Markdown)(nil)

// Heading implements Reporter.Heading
func (m *Markdown) Heading(s string) {
	fmt.Fprintf(m.Writer, "# %s\n\n", s)
}

// Paragraph implements Reporter.Paragraph
func (m *Markdown) Paragraph(s string) {
	fmt.Fprintf(m.Writer, "%s\n\n", s)
}

// Pre implements Reporter.Pre
func (m *Markdown) Pre(b []byte) {
	fmt.Fprintln(m.Writer, "```")
	fmt.Fprintf(m.Writer, string(b))
	fmt.Fprintln(m.Writer)
	fmt.Fprintln(m.Writer, "```")
	fmt.Fprintln(m.Writer)
}

// List implements Reporter.List
func (m *Markdown) List(list interface{}) {
	v := reflect.ValueOf(list)

	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			fmt.Fprintf(m.Writer, "* %v\n", v.Index(i).Interface())
		}

	default:
		fmt.Fprintf(m.Writer, "* %v\n", list)
	}

	fmt.Fprintln(m.Writer)
}

// Table implements Reporter.Table
func (m *Markdown) Table(callback func(w io.Writer)) {
	b := &bytes.Buffer{}
	inner := tabwriter.NewWriter(b, 8, 4, 2, byte(' '), 0)
	callback(inner)
	inner.Flush()
	fmt.Fprintln(b)

	m.Pre(b.Bytes())
}

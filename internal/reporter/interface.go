package reporter

import (
	"io"
)

type Reporter interface {
	Heading(s string)

	Paragraph(s string)

	Pre(content []byte)

	Table(func(w io.Writer))

	List(list interface{})
}

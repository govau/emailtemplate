package emailtemplate

import (
	"bytes"
	"html/template"
	texttemplate "text/template"
)

// Template encapsulates a common email template that is made up of a subject,
// a HTML part and a text part.
// By making the subject a template, dynamic data can be passed to it (e.g.
// recipient's name). This also allows the subject to be themed or changed for
// internationalization.
// Another possible use of subject being a dynamic template is AB testing.
type Template struct {
	Subject *texttemplate.Template
	HTML    *template.Template
	Text    *texttemplate.Template
}

// Execute takes the given data and executes each of the subject, HTML and text
// templates, returning the generated bytes for each.
func (t *Template) Execute(data interface{}) (subject *bytes.Buffer, html *bytes.Buffer, text *bytes.Buffer, err error) {
	subject = &bytes.Buffer{}
	html = &bytes.Buffer{}
	text = &bytes.Buffer{}
	if err := t.Subject.Execute(subject, data); err != nil {
		return subject, html, text, err
	}
	if err := t.HTML.Execute(html, data); err != nil {
		return subject, html, text, err
	}
	// Allow for the text template to be nil.
	if t.Text != nil {
		if err := t.Text.Execute(text, data); err != nil {
			return subject, text, text, err
		}
	}
	return subject, html, text, nil
}

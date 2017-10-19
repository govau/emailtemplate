package emailtemplate

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	texttemplate "text/template"
)

const (
	// SubjectFile is the name of the subject template file.
	SubjectFile = "subject.txt"
	// HTMLFile is the name of the HTML template file.
	HTMLFile = "html.html"
	// TextFile is the name of the text template file.
	TextFile = "text.txt"
)

// Loader is a type which loads email templates located at a root path.
type Loader struct {
	rootPath       string
	subjectFuncMap texttemplate.FuncMap
	htmlFuncMap    template.FuncMap
	textFuncMap    texttemplate.FuncMap
	logger         *log.Logger
}

func newLoader(opts ...LoaderOpt) *Loader {
	l := &Loader{}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

// Load loads email templates and returns a Getter that gets templates by key.
func Load(opts ...LoaderOpt) (*Getter, error) {
	l := newLoader(opts...)

	templatePaths := map[Key]string{}

	if err := filepath.Walk(l.rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == l.rootPath {
			return nil
		}
		if !info.IsDir() {
			return nil
		}
		base := filepath.Base(path)
		templatePaths[Key(base)] = path
		return nil
	}); err != nil {
		return nil, err
	}

	return l.load(templatePaths)
}

// load loads a map of email templates from the given templatePaths.
// See loader.loadParts.
func (l *Loader) load(templatePaths map[Key]string) (*Getter, error) {
	g := newGetter()

	for k, tp := range templatePaths {
		t, err := l.loadParts(tp)
		if err != nil {
			return nil, err
		}
		g.set(k, t)
	}

	return g, nil
}

// loadParts loads a single template from the template file parts found under
// templatePath.
// The template is loaded by combining the base path of the given template and
// the type of template (either subject.txt, html.html or text.txt).
// Both the subject and HTML templates are required however the text template
// does not have to exist.
func (l *Loader) loadParts(templatePath string) (*Template, error) {
	st, err := texttemplate.
		New(SubjectFile).
		Funcs(l.subjectFuncMap).
		ParseFiles(filepath.Join(templatePath, SubjectFile))
	if err != nil {
		return nil, err
	}
	ht, err := template.
		New(HTMLFile).
		Funcs(l.htmlFuncMap).
		ParseFiles(filepath.Join(templatePath, HTMLFile))
	if err != nil {
		return nil, err
	}
	tt, err := texttemplate.
		New(TextFile).
		Funcs(l.textFuncMap).
		ParseFiles(filepath.Join(templatePath, TextFile))
	switch {
	case os.IsNotExist(err):
		if l.logger != nil {
			l.logger.Printf("warning: text email template not found. error: %s", err)
		}
	case err != nil:
		return nil, err
	}
	return &Template{
		Subject: st,
		HTML:    ht,
		Text:    tt,
	}, nil
}

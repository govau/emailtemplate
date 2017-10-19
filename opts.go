package emailtemplate

import (
	"html/template"
	"log"
	texttemplate "text/template"
)

// LoaderOpt is a type which defines loader options.
type LoaderOpt func(l *Loader)

// WithRootPath configures the loader to use a custom root path.
func WithRootPath(rootPath string) LoaderOpt {
	return func(l *Loader) {
		l.rootPath = rootPath
	}
}

// WithSubjectFuncMap configures the loader to use a func map for the subject
// template.
func WithSubjectFuncMap(m texttemplate.FuncMap) LoaderOpt {
	return func(l *Loader) {
		l.subjectFuncMap = m
	}
}

// WithHTMLFuncMap configures the loader to use a func map for the HTML
// template.
func WithHTMLFuncMap(m template.FuncMap) LoaderOpt {
	return func(l *Loader) {
		l.htmlFuncMap = m
	}
}

// WithTextFuncMap configures the loader to use a func map for the text
// template.
func WithTextFuncMap(m texttemplate.FuncMap) LoaderOpt {
	return func(l *Loader) {
		l.textFuncMap = m
	}
}

// WithLogger configures the loader to use a logger.
func WithLogger(logger *log.Logger) LoaderOpt {
	return func(l *Loader) {
		l.logger = logger
	}
}

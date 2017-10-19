package emailtemplate

import (
	"errors"
	"html/template"
	"path/filepath"
	"strings"
	"testing"
	texttemplate "text/template"
)

func TestLoad(t *testing.T) {
	g, err := Load(WithRootPath(filepath.Join("testdata", "load")))
	if err != nil {
		t.Fatal(err)
	}
	if expected := 2; g.len() != expected {
		t.Errorf("Load(): getter.len() = %d, want %d", g.len(), expected)
	}
}

func TestLoaderLoad(t *testing.T) {
	tests := []struct {
		name          string
		rootPath      string
		templatePaths map[Key]string
		found         int
		err           error
	}{
		{
			name:          "no template paths (1)",
			rootPath:      "",
			templatePaths: nil,
		},
		{
			name:          "no template paths (2)",
			rootPath:      "",
			templatePaths: map[Key]string{},
		},
		{
			name:     "subject not exists",
			rootPath: filepath.Join("testdata", "loader-load"),
			templatePaths: map[Key]string{
				Key(1): "template-subject-not-exists",
			},
			err: errors.New("open testdata/loader-load/template-subject-not-exists/subject.txt: no such file or directory"),
		},
		{
			name:     "html not exists",
			rootPath: filepath.Join("testdata", "loader-load"),
			templatePaths: map[Key]string{
				Key(1): "template-html-not-exists",
			},
			err: errors.New("open testdata/loader-load/template-html-not-exists/html.html: no such file or directory"),
		},
		{
			name:     "text not exists",
			rootPath: filepath.Join("testdata", "loader-load"),
			templatePaths: map[Key]string{
				Key(1): "template-text-not-exists",
			},
			found: 1,
		},
		{
			name:     "subject, html and text exist",
			rootPath: filepath.Join("testdata", "loader-load"),
			templatePaths: map[Key]string{
				Key(1): "template-1",
			},
			found: 1,
		},
		{
			name:     "multiple",
			rootPath: filepath.Join("testdata", "loader-load"),
			templatePaths: map[Key]string{
				Key(1): "template-1",
				Key(2): "template-2",
			},
			found: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.templatePaths {
				tt.templatePaths[k] = filepath.Join(tt.rootPath, v)
			}
			l := newLoader()
			g, err := l.load(tt.templatePaths)
			if err != nil {
				if tt.err == nil {
					t.Fatalf("load(): err = %v, want nil", err)
				}
				if err.Error() != tt.err.Error() {
					t.Fatalf("load(): err = %v, want %v", err, tt.err)
				}
			} else {
				if tt.err != nil {
					t.Fatalf("load(): err = nil, want %v", tt.err)
				}
			}
			if g.len() != tt.found {
				t.Errorf("load(): getter.len() = %d, want %d", g.len(), tt.found)
			}
		})
	}
}

func TestFuncMaps(t *testing.T) {
	toUpper := func(v string) string {
		return strings.ToUpper(v)
	}

	rootPath := filepath.Join("testdata", "func-maps")

	tests := []struct {
		name           string
		templateName   string
		subjectFuncMap texttemplate.FuncMap
		htmlFuncMap    template.FuncMap
		textFuncMap    texttemplate.FuncMap
		data           map[string]interface{}
		subject        string
		html           string
		text           string
	}{
		{
			name:         "",
			templateName: "template-func-maps",
			subjectFuncMap: texttemplate.FuncMap{
				"toUpper": toUpper,
			},
			htmlFuncMap: template.FuncMap{
				"toUpper": toUpper,
			},
			textFuncMap: texttemplate.FuncMap{
				"toUpper": toUpper,
			},
			data: map[string]interface{}{
				"name": "Jane",
			},
			subject: "hello JANE",
			html:    "<p>dear JANE</p>",
			text:    "dear JANE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLoader(
				WithRootPath(rootPath),
				WithSubjectFuncMap(tt.subjectFuncMap),
				WithHTMLFuncMap(tt.htmlFuncMap),
				WithTextFuncMap(tt.textFuncMap),
			)
			key := Key(tt.templateName)
			g, err := l.load(map[Key]string{
				key: filepath.Join(rootPath, tt.templateName),
			})
			if err != nil {
				t.Fatal(err)
			}
			te, ok := g.Get(key)
			if !ok {
				t.Fatalf("template with key %v does not exist", key)
			}
			subject, html, text, err := te.Execute(tt.data)
			if err != nil {
				t.Fatal(err)
			}
			if subject.String() != tt.subject {
				t.Errorf("Execute(): subject = %s, want %s", subject.String(), tt.subject)
			}
			if html.String() != tt.html {
				t.Errorf("Execute(): html = %s, want %s", html.String(), tt.html)
			}
			if text.String() != tt.text {
				t.Errorf("Execute(): text = %s, want %s", text.String(), tt.text)
			}
		})
	}
}

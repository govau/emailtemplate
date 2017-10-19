package emailtemplate

import (
	"fmt"
	"html/template"
	"strings"
)

func Example() {
	// Define template keys to be used to look up templates later.
	// The value of the key matches the directory name of the template.

	const (
		emailTemplateResetPassword      Key = "resetPassword"
		emailTemplatePasswordUpdated        = "passwordUpdated"
		emailTemplateProjectMemberAdded     = "projectMemberAdded"
		emailTemplateWeeklyUsage            = "weeklyUsage"
	)

	// Load templates (e.g. in main).
	//
	// The directory structure should look like:
	// .
	// └── templates
	//     ├── resetPassword
	//     │   ├── html.html
	//     │   ├── subject.txt
	//     │   └── text.txt
	//     ├── passwordUpdated
	//     │   ├── html.html
	//     │   ├── subject.txt
	//     │   └── text.txt
	//     │── projectMemberAdded
	//     │   ├── html.html
	//     │   ├── subject.txt
	//     │   └── text.txt
	//     └── weeklyUsage
	//         ├── html.html
	//         ├── subject.txt
	//         └── text.txt

	g, err := Load(WithRootPath("./templates"))
	if err != nil {
		// ...
	}

	// Get and execute a template (e.g. in a HTTP handler).

	t, ok := g.Get(emailTemplateResetPassword)
	if !ok {
		// ...
	}

	subject, html, text, err := t.Execute(map[string]interface{}{
		"url": "https://example.com/reset-password",
	})
	if err != nil {
		// ...
	}

	// Pass the generated content to your mail sender.

	fmt.Println(subject.String())
	fmt.Println(html.String())
	fmt.Println(text.String())
}

func ExampleWithHTMLFuncMap() {
	const emailTemplateNotification Key = "notification"

	// Contents of templates/notification/html.html are:
	// <p>Dear {{trimSpace .name}}</p>

	g, err := Load(
		WithRootPath("./templates"),
		WithHTMLFuncMap(template.FuncMap{
			"trimSpace": func(v string) string {
				return strings.TrimSpace(v)
			},
		}),
	)
	if err != nil {
		// ...
	}

	t, ok := g.Get(emailTemplateNotification)
	if !ok {
		// ...
	}

	_, html, _, err := t.Execute(map[string]interface{}{
		"name": " Jane Doe ",
	})
	if err != nil {
		// ...
	}

	fmt.Println(html.String()) // <p>Dear Jane Doe</p>
}

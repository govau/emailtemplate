# emailtemplate &middot; [![Travis-CI](https://travis-ci.org/govau/emailtemplate.svg)](https://travis-ci.org/govau/emailtemplate) [![GoDoc](https://godoc.org/github.com/govau/emailtemplate?status.svg)](http://godoc.org/github.com/govau/emailtemplate) [![Report card](https://goreportcard.com/badge/github.com/govau/emailtemplate)](https://goreportcard.com/report/github.com/govau/emailtemplate)

`emailtemplate` is a Go package for loading and executing email templates. An email template is a directory which consists of a subject, HTML and text Go template file.

## Rationale

Your backend server is written in Go. It passes recipient data to Go templates which are then executed and used when sending mail.

However, you don't want to hand-code the email template parts. Most notably, you don't want to hand-code the HTML part because the tooling for writing HTML in Go templates is not as suitable as tooling in the frontend space. Further, HTML emails are a beast of their own and so using an external toolkit to do things like CSS classes to inline styles is desirable. So, instead, you generate the parts externally (perhaps using a frontend toolkit like React and friends).

`emailtemplate` lets you structure your email templates into a directory containing 3 parts, each of which are Go templates: `subject.txt`, `html.html` and `text.txt`. The directory name is used as the template name when fetching a particular template.

## Usage and examples

[Read the documentation](https://godoc.org/github.com/govau/emailtemplate).

## Development

```sh
go test -race ./...
```

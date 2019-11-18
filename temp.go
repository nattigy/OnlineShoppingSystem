package main

import (
	"net/http"
	"text/template"
)

type Context struct {
	FirstName string
	Message   string
	URL       string
	Beer      []string
	Title     string
	writer    http.ResponseWriter
}

var templ2 = template.Must(template.ParseFiles("index.html"))

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>First Template</title>
</head>
<body>
    <h1>My name is {{ .FirstName }}</h1>
	<p>{{ .Message }}</p>
	{{ if eq .URL "/nobeer" }}
		Add another beer
		<ul>
			{{ range .Beer }}
			<li>{{ . }}</li>
			{{ end }}
		</ul>
		{{ .URL }}
	{{ else }}
		Thank you!!
	{{ end }}
{{.}}
</body>
</html>
`

func toddFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err == nil {
		context := Context{
			"Todd",
			"more Go, please",
			r.URL.Path,
			[]string{"Giorgis", "Heniken", "Waliya", "Habesha", "Anbesa"},
			"Favorite Beers",
			w,
		}
		tmpl.Execute(w, context)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/nobeer", toddFunc)
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err == nil {
		tmpl.Execute(writer, request.URL.Path)
	}
}

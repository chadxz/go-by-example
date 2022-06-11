package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	t1 = template.Must(t1.Parse("Value is {{.}}\n"))

	_ = t1.Execute(os.Stdout, "some text")
	_ = t1.Execute(os.Stdout, 5)
	_ = t1.Execute(os.Stdout, []string{"Go", "Rust"})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")
	_ = t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	_ = t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	_ = t3.Execute(os.Stdout, "not empty")
	_ = t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	_ = t4.Execute(os.Stdout, []string{"Foo", "Bar", "Baz"})
}

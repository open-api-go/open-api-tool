package tmpl

import (
	"bytes"
	"text/template"
)

var makefileTemplate = `build:
	protoc --go_out $(GOPATH)/src  {{ .Name }}.proto
	protoc --open_api_out $(GOPATH)/src {{ .Name }}.proto
`

func GenMakefile(name string) (string, error) {
	t, err := template.New("makefile").Parse(makefileTemplate)
	if err != nil {
		return "", err
	}
	bs := bytes.NewBuffer(nil)
	m := map[string]interface{}{
		"Name": name,
	}
	err = t.Execute(bs, m)
	if err != nil {
		return "", err
	}
	return bs.String(), nil
}

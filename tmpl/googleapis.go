package tmpl

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed third_party/google/api/annotations.proto
var protoannotations string

//go:embed third_party/google/api/http.proto
var protohttp string

func GenAnnotations(_ string) (string, error) {
	t, err := template.New("annotations").Parse(protoannotations)
	if err != nil {
		return "", err
	}
	bs := bytes.NewBuffer(nil)
	err = t.Execute(bs, nil)
	if err != nil {
		return "", err
	}
	return bs.String(), nil
}

func GenHTTP(_ string) (string, error) {
	t, err := template.New("http").Parse(protohttp)
	if err != nil {
		return "", err
	}
	bs := bytes.NewBuffer(nil)
	err = t.Execute(bs, nil)
	if err != nil {
		return "", err
	}
	return bs.String(), nil
}

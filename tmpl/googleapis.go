package tmpl

import (
	"bytes"
	"text/template"
)

func GenAnnotations(_ string) (string, error) {
	t, err := template.ParseFiles("third_party/google/api/annotations.proto")
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
	t, err := template.ParseFiles("third_party/google/api/http.proto")
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

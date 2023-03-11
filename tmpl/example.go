package tmpl

import (
	"bytes"
	"text/template"
)

var exampleTemplate = `syntax = "proto3";
package {{ .Name }}.api; // 填openapi的域名,不需要带https, 只支持https
import "google/api/annotations.proto";

option go_package = "github.com/open-api-go/{{ .Name }}";

service Greeter {
    rpc Hello(HelloReq) returns (HelloRes) {
		option (google.api.http) = {
			post:"/hello/{userid}?name={name}", 
			body: "body"
		};
	};
}

message HelloReq {
    string user_id = 1;
    string name = 2;
    string body = 3;
}

message HelloRes {
}
`

func GenExample(name string) (string, error) {
	t, err := template.New("example").Parse(exampleTemplate)
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

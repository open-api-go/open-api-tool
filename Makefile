build:
	go fmt ./...

test:
	go install
	opa new -n helloworld -d testdata
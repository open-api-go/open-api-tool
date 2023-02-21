package command

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/open-api-go/openapi-tool/tmpl"
	"github.com/spf13/cobra"
)

var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "Create an open-api template",
	Long:  "Create an open-api template",
	Run:   run,
}

var (
	dir  string
	name string
)

func init() {
	CmdNew.Flags().StringVarP(&dir, "dir", "d", dir, "dir")
	CmdNew.Flags().StringVarP(&name, "name", "n", name, "name")
}

func run(_ *cobra.Command, _ []string) {
	err := os.MkdirAll(path.Join(dir, "google", "api"), 0777)
	if err != nil {
		log.Fatalln("mkdir err:", err)
		return
	}

	// Makefile
	generate(tmpl.GenMakefile, dir, "Makefile")
	// Example proto
	generate(tmpl.GenExample, dir, fmt.Sprintf("%s.proto", name))
	// google/api proto file
	generate(tmpl.GenAnnotations, dir, "google/api", "annotations.proto")
	generate(tmpl.GenHTTP, dir, "google/api", "http.proto")
}

func generate(fn func(name string) (string, error), fileName ...string) {
	fd, err := os.Create(path.Join(fileName...))
	if err != nil {
		log.Fatalln("create err:", err)
		return
	}
	c, err := fn(name)
	if err != nil {
		log.Fatalln(err)
		return
	}
	_, err = fd.WriteString(c)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

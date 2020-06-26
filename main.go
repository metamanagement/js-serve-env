package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/NYTimes/gziphandler"
)

func main() {
	runtimeVars := strings.Split(os.Getenv("RUNTIME_VARS"), ",")
	envVals := map[string]string{}
	for _, v := range runtimeVars {
		envVals[v] = os.Getenv(v)
	}

	tmpl, err := template.New("js").Parse("{{range $k, $v := .}}window.RUNTIME_{{$k}}={{ printf \"%q\" $v }}; {{end}}")
	if err != nil {
		panic(err)
	}
	f, err := os.Create("./dist/runtime_env_vars.js")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, envVals)

	http.Handle("/", gziphandler.GzipHandler(http.FileServer(http.Dir("./dist"))))
	log.Println("SERVING...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

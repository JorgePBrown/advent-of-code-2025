package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

func main() {
	flag.Parse()

	day := flag.Arg(0)
	if day == "" {
		log.Fatal("day not defined. Usage 'template <day>'")
	}
	templ, err := template.ParseFiles("./template/main.templ")
	if err != nil {
		log.Fatal("failed to parse template", err)
	}
	err = os.Mkdir("./day"+day, os.ModePerm)
	if err != nil {
		log.Fatal("failed to create dir", err)
	}

	w, err := os.Create("./day" + day + "/main.go")

	err = templ.Execute(w, map[string]string{
		"Day": day,
	})
	if err != nil {
		log.Fatal("failed to execute template", err)
	}
}

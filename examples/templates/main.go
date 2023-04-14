package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"text/template"
)

func main() {
	//helloWorld()
	complexExample()
}

func helloWorld() {
	parse, err := template.New("Greeting").Parse(`Hello, {{ . }}!`)
	if err != nil {
		log.Println(err)
		return
	}

	err = parse.Execute(os.Stdout, "world")
	if err != nil {
		log.Println(err)
		return
	}
}

func complexExample() {
	funcs := template.FuncMap{
		"mult": func(a, b int) int {
			return a * b
		},
	}

	tmpl, err := template.New("main.tmpl").Funcs(funcs).ParseFiles("templates/main.tmpl", "templates/common.tmpl")
	if err != nil {
		log.Println(err)
		return
	}

	result, err := os.Create("templates/result")
	if err != nil {
		log.Println()
		return
	}
	defer result.Close()

	data := struct {
		FieldS     string
		FieldI     int
		FieldA     []int
		FieldBT    bool
		FieldBF    bool
		FieldM     map[string]int
		FieldFunc  func() string
		Calculator *Calculator
		fieldH     string
	}{
		FieldS:     "example string",
		FieldI:     123,
		FieldA:     []int{4, 5, 6},
		FieldBT:    true,
		FieldBF:    false,
		FieldM:     map[string]int{"a": 1, "b": 2, "c": 3},
		FieldFunc:  func() string { return "string from func" },
		Calculator: &Calculator{},
		fieldH:     "hidden field",
	}

	err = tmpl.Execute(result, data)
	if err != nil {
		log.Println(err)
		return
	}

}

type Calculator struct {
	LastResult int
}

func (c *Calculator) PositiveSum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("args must be positive")
	}

	res := a + b
	c.LastResult = res

	return res, nil
}

func (c *Calculator) LastToString() string {
	return strconv.Itoa(c.LastResult)
}

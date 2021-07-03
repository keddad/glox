package main

import (
	"fmt"
	"os"
	"strings"
)

var structs = map[string]string{
	"Binary":   "Left interface{},Operator Token,Right interface{}",
	"Grouping": "Expr interface{}",
	"Literal":  "Value interface{}",
	"Unary":    "Operator Token,Right interface{}",
}

func generateStruct(name string, fields string) string {
	var builder strings.Builder
	builder.Grow(len(name) + len(fields) + 16) // Because I can

	builder.WriteString("type ")
	builder.WriteString(name)
	builder.WriteString(" struct {\n")

	for _, element := range strings.Split(fields, ",") {
		builder.WriteString("\t")
		builder.WriteString(element)
		builder.WriteString("\n")
	}

	builder.WriteString("}\n\n")

	builder.WriteString(fmt.Sprintf("func (obj *%s) accept(v AstVisitor) {\n\tv.Visit%s(obj)\n}", name, name))
	return builder.String()
}

func buildFile(pkg string, structs ...string) string {
	var builder strings.Builder

	builder.WriteString("package ")
	builder.WriteString(pkg)
	builder.WriteString("\n\n")

	builder.WriteString("import ")
	builder.WriteString(". \"glox/glox/tokens\"")

	builder.WriteString("\n\n")

	for _, element := range structs {
		builder.WriteString(element)
		builder.WriteString("\n\n")
	}

	return builder.String()
}

func generateInterface(structs ...string) string {
	var builder strings.Builder
	builder.WriteString("type AstVisitor interface {\n")

	for _, v := range structs {
		builder.WriteString(fmt.Sprintf("\tVisit%s(x *%s)\n", v, v))
	}

	builder.WriteString("}\n\n")
	return builder.String()
}

func main() {
	path := os.Args[1]

	f, _ := os.Create(path)
	defer f.Close()

	sts := make([]string, 0)
	names := make([]string, 0)

	for k, v := range structs {
		sts = append(sts, generateStruct(k, v))
		names = append(names, k)
	}

	f.WriteString(buildFile("ast", append(sts, generateInterface(names...))...))
}

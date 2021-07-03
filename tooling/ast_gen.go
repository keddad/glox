package main

import (
	"os"
	"strings"
)

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

	builder.WriteString("}")
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

func main() {
	path := os.Args[1]

	f, _ := os.Create(path)
	defer f.Close()

	f.WriteString(buildFile("ast",
		generateStruct("Binary", "Left interface{},Operator Token,Right interface{}"),
		generateStruct("Grouping", "Expr interface{}"),
		generateStruct("Literal", "Value interface{}"),
		generateStruct("Unary", "Operator Token,Right interface{}")))

}

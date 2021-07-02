package main

import (
	"bufio"
	"fmt"
	"glox/glox/runtime"
	"glox/glox/tokens"
	"io/ioutil"
	"os"
)

func run(line string, state *runtime.State) {
	tokens := tokens.ScanTokens(line, state)
	if state.HadError {
		return
	}

	fmt.Printf("%+v\n", tokens) // Just print the tokens for now

}

func runFile(path string) error {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	state := runtime.NewState()
	run(string(body), &state)
	if state.HadError {
		os.Exit(65)
	}

	return nil
}

func runInteractive() {
	scanner := bufio.NewScanner(os.Stdin)
	state := runtime.NewState()
	for {
		print("> ")
		scanner.Scan()
		text := scanner.Text()

		if text == "" {
			return
		}

		run(text, &state)
		state.HadError = false
	}
}

func main() {
	argsNum := len(os.Args) - 1

	if argsNum > 1 {
		println("Usage: glox [script]")
	} else if argsNum == 1 {
		err := runFile(os.Args[1])
		if err != nil {
			println(err)
		}
	} else {
		runInteractive()
	}
}

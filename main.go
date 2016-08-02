package main

import (
	"bufio"
	"fmt"
	"os"
)

func runRepl(reader *bufio.Reader) {
	fmt.Print("> ")
	input, _ := reader.ReadString('\n')
	p := newParser(input)
	fmt.Println(p.parseExpression())
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		runRepl(reader)
	}
}

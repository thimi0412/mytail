package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	intOpt = flag.Int("n", 10, "help message for \"n\" option")
)

func main() {
	flag.Parse()
	args := flag.Args()

	lines := readFile(args[0])
	lastIndex := len(lines) - 1

	for i := lastIndex - *intOpt; i < lastIndex; i++ {
		fmt.Println(lines[i])
	}

}

func readFile(file string) []string {
	fp, err := os.Open(file)
	defer fp.Close()

	if err != nil {
		fmt.Println("ファイルが存在しません")
		os.Exit(0)
	}

	lines := []string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

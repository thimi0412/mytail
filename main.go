package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	n = flag.Int("n", 10, "show line")
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("ファイルを指定してください")
		os.Exit(0)
	}

	// コマンドライン引数の数回す
	for _, file := range args {
		fmt.Printf("==> %s <==\n", file)
		lines := readFile(file)

		// 空ファイルは飛ばす
		if lines == nil {
			continue
		}
		
		lastIndex := len(lines) - 1

		var showStart int = lastIndex - *n
		// ファイルの行数がデフォルト値(10)かn(任意の数)より小さければ
		// showRangeを0とし表示できるだけ表示する
		if showStart < 0 {
			showStart = 0
		}
		showLines := lines[showStart:]
		

		for _, line := range showLines {
			fmt.Println(line)
		}
		
		fmt.Println("/////////////////////////////////////")
	}
}


func readFile(file string) []string {
	fp, err := os.Open(file)
	defer fp.Close()

	if err != nil {
		fmt.Println("ファイルが存在しません")
		return nil
	}

	lines := []string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

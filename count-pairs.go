package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	result := make(map[string]([]string))

	//srcfile, errfile := os.Open("/home/yeahnoob/.gvm/pkgsets/go1.3/global/src/github.com/yeahnoob/golang-perf/word-pairs.txt")
	srcfile, errfile := os.Open(os.Args[1])
	if errfile != nil {
		fmt.Println(errfile)
		return
	}

	scanner := bufio.NewScanner(srcfile)
	for scanner.Scan() {
		p := strings.Split(scanner.Text(), " ")
		result[p[0]] = append(result[p[0]], p[1])
	}
	fmt.Println("Count of \"it\" = ", len(result["it"]))
	fmt.Println("Count of \"as\" = ", len(result["as"]))
	fmt.Println("Count of \"of\" = ", len(result["of"]))
	if errscan := scanner.Err(); errscan != nil {
		fmt.Println(errscan)
		return
	}
}

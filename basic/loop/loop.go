package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(1345435),
	)
	printFile("xiangnan/learngo/basic/branch/abc.txt")
	s := `abc"d"
kkkk
123

p`
	printFileContents(strings.NewReader(s))
}

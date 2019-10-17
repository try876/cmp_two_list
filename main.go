// 比较两个文件中的列表，列表以行分隔

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 从utf-8文本文件中取出列表，一行一个
func getArray(path string) []string {
	a := make([]string, 0, 40)

	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("err", err)
		return a
	}
	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)

	str := string(fd)

	lines := strings.Split(str, "\n")

	for _, row := range lines {
		if row == "" {
			continue
		}

		a = append(a, row)
	}

	return a
}

func main() {

	more := getArray("1.txt")
	less := getArray("2.txt")
	for _, aa := range more {
		got := false
		for _, bb := range less {
			if aa == bb {
				got = true
				break
			}
		}

		if got == false {
			fmt.Println(aa)
		}
	}
}

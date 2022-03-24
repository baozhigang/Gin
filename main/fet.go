package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

// RE := regexp.MustCompile("\\w+\\.\\w+$")
var RE = regexp.MustCompile("\\w+\\.\\w+$")

func main() {
	urls := []string{
		"http://www.qq.com",
		"http://www.sina.com",
	}

	start := time.Now()

	for _, url := range urls {
		start := time.Now()

		res, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(res.Body)

		res.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fileName := getFileName(url)

		ioutil.WriteFile(fileName, body, 0644)

		// 消耗的时间
		elapsed := time.Since(start).Seconds()

		fmt.Printf("%.2fs %s\n", elapsed, fileName)
	}

	// 消耗的时间
	elapsed := time.Since(start).Seconds()

	fmt.Printf("%.2fs elapsed\n", elapsed)
}

func getFileName(url string) string {
	return RE.FindString(url) + ".txt"
}

// 恭喜你，通了，下一步，弄清楚里面的语法
// vscode go 语法提示
// debug 调通

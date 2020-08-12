package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/cheggaaa/pb/v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	const baseUrl = "http://tmp.o1o.win/"
	if len(os.Args) < 2 {
		log.Fatal("命令后面需要带上文件路径")
	}
	filepath := os.Args[1]
	sl := strings.Split(filepath, "/")
	filename := sl[len(sl)-1]

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal("文件打开错误")
	}
	url := baseUrl + filename
	fileInfo, _ := os.Stat(filepath)
	filesize := fileInfo.Size()
	bar := pb.Full.Start64(filesize)
	file2 := bar.NewProxyReader(file)
	req, err := http.NewRequest(http.MethodPut, url, file2)
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	bar.Finish()
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err := clipboard.WriteAll(string(body)); err != nil {
		fmt.Println("链接复制失败")
	} else {
		fmt.Println("链接已复制到剪切板：")
	}
	fmt.Println(string(body))
}

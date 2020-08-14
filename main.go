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
	"sync"
)

var wg sync.WaitGroup

const baseUrl = "http://tmp.o1o.win/"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("命令后面需要带上文件路径")
	}
	filenames := os.Args[1:]
	wg.Add(len(filenames))
	for _, value := range filenames {
		go UploadFile(value)
	}
	wg.Wait()
	fmt.Println("全部上传完成")
}

func UploadFile(filepath string) {
	defer wg.Done()
	sl := strings.Split(filepath, "/")
	filename := sl[len(sl)-1]

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("文件打开错误")
	}
	defer file.Close()
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
	if err != nil {
		fmt.Println(filepath + "上传失败")
	}
	bar.Finish()
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if err := clipboard.WriteAll(string(body)); err != nil {
		// fmt.Println("链接复制失败")
	} else {
		fmt.Println("链接已复制到剪切板：")
	}
	fmt.Println(string(body))
}

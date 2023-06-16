package myhttp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Start(port string) {
	http.HandleFunc("/api/file", func(w http.ResponseWriter, r *http.Request) {
		// 设置跨域
		Cors(w)
		// 获取文件目录
		path := r.URL.Query()["path"][0]
		// 读取文件目录 path
		FileInfos, _ := ioutil.ReadDir(path)
		infos := make([]string, 0, len(FileInfos))
		for _, entry := range FileInfos {
			info := entry.Name()
			infos = append(infos, info)
		}
		fmt.Println(infos)
		jsonString, _ := json.Marshal(infos)
		j := string(jsonString)
		// 返回响应
		fmt.Fprintf(w, j)
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

package main

import (
    "io"
    "os"
    "log"
    "fmt"
    "strings"
    "net/http"
)

// 返回文件目录
func getFileInfo(file string) (fileName, filePath string) {
    index := strings.LastIndex(file, "/")
    return file[index + 1:], file[0:index + 1]
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        io.WriteString(w, "<p>Hey <del>gays</del> guys, I'm ready for that, you know.</p>")
        return
    }
    if r.Method == "POST" {
        f, _, err := r.FormFile("file")
        if err != nil {
            http.Error(w, err.Error(), 
                http.StatusInternalServerError)
            return
        }
        defer f.Close()
        // to已经包含完整路径了
        to   := r.FormValue("to")
        // 首先判断一遍目录是否存在
        _, filePath := getFileInfo(to)
        _, err = os.Open(filePath)
        if err != nil && os.IsNotExist(err) {
            // 错误都不处理了
            os.MkdirAll(filePath, 0666)
        }
    
        t, err := os.Create(to)
        if err != nil {            
            http.Error(w, err.Error(),
                http.StatusInternalServerError)
            return
        }
        defer t.Close()
        if _, err := io.Copy(t, f); err != nil {
            http.Error(w, err.Error(), 
                http.StatusInternalServerError)
            return
        }
        fmt.Println(r.FormValue("to"))
        // success
        io.WriteString(w, "0")
    }
}

func main () {
    port := "8527"
    
    if len(os.Args) == 2 {
        port = os.Args[1]
    }
    http.HandleFunc("/", uploadHandler)

    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
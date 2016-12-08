package main

import (
    "io"
    "os"
    "log"
    "fmt"
    "syscall"
    "strings"
    "net/http"
)

// 返回文件目录
func getFileInfo(file string) (fileName, filePath string) {
    index := strings.LastIndex(file, "/")
    return file[index + 1:], file[0:index + 1]
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // 不支持get请求
    if r.Method == "POST" {
        token := r.FormValue("token")
        if (token != secret) {
            io.WriteString(w, "error")
            return
        }
        f, _, err := r.FormFile("file")
        if err != nil {
            http.Error(w, err.Error(),
                http.StatusInternalServerError)
            return
        }
        defer f.Close()
        // to已经包含完整路径了
        to    := r.FormValue("to")

        // 首先判断一遍目录是否存在
        _, filePath := getFileInfo(to)
        _, err = os.Open(filePath)
        if err != nil && os.IsNotExist(err) {
            // 错误都不处理了
            oldMask := syscall.Umask(0)
            os.MkdirAll(filePath, os.ModePerm)
            syscall.Umask(oldMask)
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

var secret string

func main () {
    port := "8527"
    length := len(os.Args)

    if length == 1 {
        fmt.Println("please add token");
        return
    }
    if length == 2 {
        secret = os.Args[1]
    } else if length == 3 {
        secret = os.Args[1]
        port   = os.Args[2]
    }

    http.HandleFunc("/", uploadHandler)

    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}

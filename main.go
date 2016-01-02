package main

import (
    "io"
    "os"
    "log"
    "fmt"
    "strings"
    "net/http"
)

const (
    UPLOAD_DIR = "./temp"
)

// 返回文件目录
func getFileInfo(file string) (fileName, filePath string) {
    var index int = strings.LastIndex(file, "/")
    return file[index + 1:], file[0:index + 1]
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        io.WriteString(w, "<p>Hey <del>gays</del> guys, I'm ready for that, you know.</p>")
        return
    }
    if r.Method == "POST" {
        // to := r.FormValue("to")
        f, h, err := r.FormFile("file")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        file := h.Filename
        defer f.Close()

        _, filePath := getFileInfo(file)

        _, err = os.Open(UPLOAD_DIR + filePath)
        if err != nil {
            if os.IsNotExist(err) {
                // 路径不存在，需要创建
                err := os.MkdirAll(UPLOAD_DIR + filePath, 0666)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return    
                }
            } else {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        }
        t, err := os.Create(UPLOAD_DIR + file)
        if err != nil {            
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer t.Close()
        if _, err := io.Copy(t, f); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Println(r.FormValue("to"))
        // success
        io.WriteString(w, "0")
    }
}

func main () {
    http.HandleFunc("/", uploadHandler)

    err := http.ListenAndServe(":8527", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
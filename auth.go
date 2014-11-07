package main

import (
    "fmt"
    "io"
    "net/http"
    "log"
    "encoding/base64"
    "strings"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    auth := req.Header.Get("Authorization")
    if auth == "" {
        w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
    fmt.Println(auth)

    auths := strings.SplitN(auth, " ", 2)
    if len(auths) != 2 {
        fmt.Println("error")
        return
    }

    authMethod := auths[0]
    authB64 := auths[1]

    switch authMethod {
    case "Basic":
        authstr, err := base64.StdEncoding.DecodeString(authB64)
        if err != nil {
            fmt.Println(err)
            io.WriteString(w, "Unauthorized!\n")
            return
        }
        fmt.Println(string(authstr))

        userPwd := strings.SplitN(string(authstr), ":", 2)
        if len(userPwd) != 2 {
            fmt.Println("error")
            return
        }

        username := userPwd[0]
        password := userPwd[1]

        fmt.Println("Username:", username)
        fmt.Println("Password:", password)
        fmt.Println()

    default:
        fmt.Println("error")
        return
    }
    

    io.WriteString(w, "hello, world!\n")
}

func main() {
    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

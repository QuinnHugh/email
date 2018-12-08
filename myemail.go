package main

import (
    "fmt"
    "./m"
)

func main() {
    msg := email.Message{
        To: "764976884@qq.com", // do not add < > or name in quotes
        From: "quinnhahah@haha.fff", // do not add < > or name in quotes
        Subject: "测试测试",
        Body: "域名是haha.fff",
    }

    err := msg.Send()
    if err != nil {
        fmt.Println(err)
        fmt.Println("1")
    }
    fmt.Println("2")
}
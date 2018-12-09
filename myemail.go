package main

import (
    "fmt"
    "./m"
)

func main() {
    msg := email.Message{
        To: "764976884@qq.com", // do not add < > or name in quotes
        From: "auto@yumn.tk", // do not add < > or name in quotes
        Subject: "测试测试",
        Body: "测试yumn.tk",
    }

    err := msg.Send()
    if err != nil {
        fmt.Println(err)
        fmt.Println("1")
    }
    fmt.Println("2")
}
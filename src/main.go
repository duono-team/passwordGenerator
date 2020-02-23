package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  passwd := ""
  for i:=0;;i++{
    n := rand.Intn(94) + 33
    if len(passwd) >= 16 { //调整左侧数字以确定密码长度
      break
    }
    // 设置字符范围（默认字母加数字）
    if n >= 48 && n <= 57 {
      passwd += string(n)
    } else if n >= 65 && n <= 90 {
      passwd += string(n)
    } else if n >= 97 && n <= 122 {
      passwd += string(n)
    }
  }
  fmt.Println(passwd)
  for {}
}

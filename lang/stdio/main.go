package main

import (
  "fmt"
  "os"
  "bufio"
)

func main () {
  stdin := bufio.NewScanner(os.Stdin)
  
  // pipeされていた入力の取得
  piped := stdin.Text()
  fmt.Println(piped)
  
  // 入力を待つ
  if stdin.Scan() {
    if err := stdin.Err(); err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
    fmt.Println(stdin.Text())
  }
}
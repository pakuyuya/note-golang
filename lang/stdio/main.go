package main

import (
  "fmt"
  "os"
  "bufio"
)

func main () {
  stdin := bufio.NewScanner(os.Stdin)
  
  // pipe����Ă������͂̎擾
  piped := stdin.Text()
  fmt.Println(piped)
  
  // ���͂�҂�
  if stdin.Scan() {
    if err := stdin.Err(); err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
    fmt.Println(stdin.Text())
  }
}
package main

import (
  "fmt"
  "os"
  "bufio"
)

func main () {
  stdin := bufio.NewScanner(os.Stdin)
  
  // pipe‚³‚ê‚Ä‚¢‚½“ü—Í‚ÌŽæ“¾
  piped := stdin.Text()
  fmt.Println(piped)
  
  // “ü—Í‚ð‘Ò‚Â
  if stdin.Scan() {
    if err := stdin.Err(); err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
    fmt.Println(stdin.Text())
  }
}
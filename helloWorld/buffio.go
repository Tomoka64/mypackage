package main

import (
  "fmt"
  "bufio"
  "strings"
)

func main() {
  s := "Hey world"
  scannar := bufio.NewScanner(strings.NewReader(s))
  scannar.Split(bufio.ScanRunes)

  for scannar.Scan() {
    line := scannar.Text()
    fmt.Println(line)
  }
}

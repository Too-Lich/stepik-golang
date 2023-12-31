//На вход подается строка. Нужно определить, является ли она правильной или нет. 
//Правильная строка начинается с заглавной буквы и заканчивается точкой. 
//Если строка правильная - вывести Right иначе - вывести Wrong

package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "unicode"
)
func main() {
  text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
  text = strings.TrimSpace(text)
  s := []rune(text)
  
  if unicode.IsUpper(s[0]) && strings.HasSuffix(text, ".") {
    fmt.Print("Right")
  } else {
    fmt.Print("Wrong")
  }
}

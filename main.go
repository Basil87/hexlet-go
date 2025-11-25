// Файл ./main.go
package main

import (
  "./greeting" // Относительный импорт
  "fmt"
)

 func main() {
       fmt.Println(greeting.Hello())
   }
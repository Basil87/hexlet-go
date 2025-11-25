   package greeting

   import "github.com/fatih/color"

   var greeting string = "Golang for Brave!"

   func Hello() string {
      return color.RedString(greeting)
   }
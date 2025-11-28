   package greeting

   import "github.com/fatih/color"

   var greeting string = "Hexlet for Brave!"

   func Hello() string {
      return color.RedString(greeting)
   }
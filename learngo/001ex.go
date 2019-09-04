package learngo

/*
This is a multi-line comment.
import keyword makes another package available
  for this .go "file".
import "fmt" lets you access fmt package's functionality
  here in this file.
*/
import "fmt"

// Run some docs
func Ex001() {
	// after: import "fmt"
	// Println function of "fmt" package becomes available

	// Look at what it looks like by typing in the console:
	//   go doc -src fmt Println

	// Println is just an exported function from
	//   "fmt" package

	// Exported = First Letter is uppercase
	fmt.Println("Hello Gopher!")

	// Go cannot call Println function by itself.
	// That's why you need to call it here.
	// It only calls `func main` automatically.

	// Go supports Unicode characters in string literals
	// And also in source-code: KÖSTEBEK!
	// Because: Literal ~= Source Code
	fmt.Println("Merhaba Köstebek means Hello Gopher in Turkish")
	fmt.Println("हेलो गोफर means Hello Gopher in Hindi")
}

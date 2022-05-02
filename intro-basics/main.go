package main

// import ("fmt"; "strings")
import (
  "fmt"
  // "strings"
  "sort"
)

// ===== STANDARD LIBRARY
func main() {
  // = "strings" package
  // greeting := "Hello there friends!"

  // fmt.Println(strings.Contains(greeting, "Hello"))
  // fmt.Println(strings.ReplaceAll(greeting, "Hello", "Ni hao"))
  // fmt.Println(strings.ToUpper(greeting))
  // fmt.Println(strings.Index(greeting, "friends")) // 12
  // fmt.Println(strings.Split(greeting, " ")) // [Hello there friends!]

  // = "sort" package
  ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
  
  sort.Ints(ages) // alters original var!
  fmt.Println(ages)

  index := sort.SearchInts(ages, 30)
  fmt.Printf("index of %v inside %v is: %v \n", 30, ages, index)
  fmt.Println(index)

  names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
  sort.Strings(names) // mutates the slice
  fmt.Println(names)

  fmt.Println(sort.SearchStrings(names, "bowser")) // 0
}



// ===== ARRAYS & SLICES
// func main() {
//   // === Arrays (fixed length but can change values)
//   // var ages [3]int = [3]int{7, 8, 10}
//   var ages = [3]int{7, 8, 10}
//   names := [4]string{"yoshi", "mario", "luigi", "peach"}
//   names[1] = "toad"

//   fmt.Println(ages, len(ages))
//   fmt.Println(names, len(names))

//   // === Slices (length is mutable but use arrays under the hood)
//   var scores = []int{100, 50, 60}
//   scores[2] = 25
//   scores = append(scores, 85) // returns new slice! Need to reassign if we want to change

//   fmt.Println(scores, len(scores))

//   // === Slice Ranges
//   rangeOne := names[1:3]
//   rangeTwo := names[2:]
//   rangeThree := names[:3]

//   fmt.Println(rangeOne)
//   fmt.Println(rangeTwo)
//   fmt.Println(rangeThree)
//   fmt.Printf("range slices are type: %T \n", rangeOne) // []string

//   rangeOne = append(rangeOne, "Koopa") // reassignment
//   fmt.Println(rangeOne)
// }





// ===== STRINGS, VARS, PRINT
// func main() {
//   // === strings
//   // var nameOne string = "Mario"
//   // var nameTwo = "Luigi"
//   // var nameThree string
//   // nameFour := "Yoshi"

//   // fmt.Println(nameOne, nameTwo, nameThree)

//   // nameOne = "Peach"
//   // nameThree = "Bowser"

//   // fmt.Println(nameOne, nameTwo, nameThree, nameFour)

//   // === ints
//   // var ageOne int = 20
//   // var ageTwo = 30
//   // ageThree := 40
//   // fmt.Println(ageOne, ageTwo, ageThree)

//   // === bits & memory
//   // var numOne int8 = 25
//   // var numTwo int8 = -128
//   // var numThree uint16 = 256 // unsized is positive num only

//   // === floats
//   // var scoreOne float32 = 25.98
//   // var scoreTwo float64 = 34134.32
//   // scoreThree := 1.5 // default is type float64

//   // === fmt print options
//   age := 35
//   name := "gaylon"
//   // = Print
//   // fmt.Print(scoreOne) // no newline
//   // fmt.Print(scoreTwo)
//   // fmt.Print(scoreThree)

//   // = Println
//   // fmt.Println("my name is", age, "and my name is", name)

//   // = Printf (formatted strings) %_ = format specifier (lots to choose from)
//   fmt.Printf("my age is %v and my name is %v. \n", age, name) // %v is default for vars
//   fmt.Printf("my age is %q and my name is %q. \n", age, name) // %q quotes for strings
//   fmt.Printf("age is of type %T \n", age) // %T for type e.g. int
//   fmt.Printf("you scored %f points! \n", 225.55) // 225.550000
//   fmt.Printf("you scored %0.1f points! \n", 225.55) // 225.6

//   // = Sprintf (save formatted strings)
//   mySavedStr := fmt.Sprintf("my age is %v and my name is %v. \n", age, name)
//   fmt.Println("The saved string is:", mySavedStr)
// }

package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello Anshul currently the time is", time.Now())
	fmt.Println("Random No : ", rand.Intn(2))
	fmt.Println("Square root : ", math.Sqrt(4))
	fmt.Println("Square root : ", math.Pi)
	fmt.Println("Sum of two : ", add(10, 15))

	greet, name, question := greetings()
	fmt.Println(greet, name, question)
	fmt.Println("Square :", square(12))
	declareDemo()
	initilizeDemo()
	initilizeDemo2()
	datatypeDemo()
	typeCasting()
	constDemo()
	forloopDemo()
	conditionDemo(10, 100)
	deferDemo1()
	deferDemo2()
	pointerDemo()
	structDemo()
	arrayDemo()
	slicesDemo()
	rangeDemo()
	mapDemo()
	functionDemo()
	methodsDemo()
	interfaceDemo()
	stringerDemo()
	errorDemo()
	readerTesting()
	typeParameterDemo()
	routineDemo()
	channelDemo()
	selectChannelDemo()
	mutexDemo()
}

// functions
func add(a int, b int) int {
	fmt.Println("\n-----------------------funtions-----------------------\n")
	return a + b
}

// we can same data type to all the vars in one write
func add2(a, b int) int {
	return a + b
}

// function can return any number of values: Not like other languages
func greetings() (string, string, string) {
	greet := "Hello"
	name := "Anshul!."
	question := "How are you ??"
	return greet, name, question
}

// If we write return without telling what to return it will return named variables
// It has a name calle "naked return" and only be used with small functions as it might affect readability.
// Note: for this to work I need to add the naked variable with return type: (sq int)
func square(num int) (sq int) {
	sq = num * num
	return
}

// We first need to write "var" then variable "name" then the "type"
// Can be used on package and funtion level.
// no undefined concept int become 0 bool become false
var a, b, c bool

func declareDemo() {
	fmt.Println("\n-----------------------variable-----------------------\n")
	var i int
	fmt.Println(i, a, b, c)
}

// We can initilize value of many variables in one line
// If we don't give the type and just initilize it will take the type of the value it is initilized
func initilizeDemo() {
	var a, b, c bool = true, false, false
	d, e, f := 0, 5.3, "no!"
	fmt.Println(a, b, c, d, e, f)
}

// We can replace "var" and "datatype" with ":" colon before "=".
// It only works inside function ouside the function we must use "var"
func initilizeDemo2() {
	a, b, c := true, false, false
	d, e, f := 0, 5.3, "no!"
	fmt.Println(a, b, c, d, e, f)
}

// There are 7 types fo datatype in Go:
// bool
// string
// int	int8	int16	int32	int64
// uint	uint8	uint16	uint32	uint64	uintptr
// byte	:alias for uint8
// rune	:alias for int32	:represent a "Unicode" code print
// float32	float64
// complex64 complex128

// Another way declare and initialize variables
// "1<<64" means shifting 1 bit to 64 left
// This "1<<64 - 1" equation gives a very big integer value that can not be stored in "int" type
// If we remove "uint64" type of variable "MaxInt" It will show error.
// because by default it takes variable as int so if assigned value can not be contained inside variable it will throw error
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// %T for printing "type" for the variable
// %v for printing "value" for the variable
func datatypeDemo() {
	fmt.Println("\n-----------------------datatypes-----------------------\n")
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// Type Casting
// To change the value of the var we can just do:
// Type(value)	: It will change "value" to the type "Type"

func typeCasting() {
	fmt.Println("\n-----------------------typecasting-----------------------\n")
	var i int = 32
	var f float32 = float32(i)
	var u uint = uint(f)
	fmt.Printf("Type of i is %T and value is %v \n", i, i)
	fmt.Printf("Type of f is %T and value is %v \n", f, f)
	fmt.Printf("Type of u is %T and value is %v \n", u, u)
}

// Constants:
// Constants are declared using "const" keyword
// It can be of anytype but can not be declared like this ":=" as it means "var"

// If a constant is untyped it cna change it's type according to contex.
const (
	big   = 1 << 100
	small = big >> 99 // so we end up at big<<1 which is equal to 2
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func constDemo() {
	fmt.Println("\n-----------------------const-----------------------\n")
	fmt.Println("Changed to int : ", needInt(small))
	fmt.Println("Changed to float64 : ", needFloat(small))
	fmt.Println("Changed to float64 : ", needFloat(big))
}

// Loops:
// There is no "while" loop in Go
// "for" does not need "()" to start the loop we can write directly
// init statement and post statement are optional
func forloopDemo() {
	fmt.Println("\n-----------------------loops-----------------------\n")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// "for; sum < 100 ;" can also be written like this.
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	// To make a loop run forever we can write:
	// for {
	// }
}

// Conditional statements:
// Does not require "()" same as loop just needs "{}"
// Switch statement same as if and it will automatically break do not require break statement.

func conditionDemo(num, i int) {
	fmt.Println("\n-----------------------conditions-----------------------\n")
	if num < 10 {
		fmt.Println(num, "is smaller than 10.")
	} else if num > 10 {
		fmt.Println(num, "is greater than 10.")
	} else {
		fmt.Println(num, "is equal to 10.")
	}

	// We can also do some statement execution befor conditional checking like for loop
	// We can do it with every if statement
	if sqt := math.Sqrt(float64(i)); sqt == float64(num) {
		fmt.Println("Square root of ", i, "is equal to", num)
	} else if sqt := math.Sqrt(float64(i)); sqt < float64(num) {
		fmt.Println("Square root of ", i, "is smaller to", num)
	} else {
		fmt.Println("Square root of ", i, "is greater to", num)
	}

	// We can not access the variable declared in init statement of if statement after the endind of consition
	// Or outside the space of "if" statement
	// fmt.Println(sqt) //will show error

	fmt.Printf("Go is running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	fmt.Printf("When's Saturday? ")
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switch without a condition is the same as switch true.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Defer:
// It might be evaluated immediately but it is only called when the function is finished it is inside in.

func deferDemo1() {
	fmt.Println("\n-----------------------defer-----------------------\n")
	defer fmt.Println("But called at the end.")
	fmt.Println("Defer is evaluated first.")
}

// defer funtion calls are pushed on to a stack and the order they came out of stack is last-in-fist-out (LIFO)
func deferDemo2() {
	fmt.Println("Loop Starting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("loop done")
}

// Pointers:
// Go also have pointer concept like "C++" using "&" and "*"
// "&" is used to point at the variable and "*" is used to access pointer value

func pointerDemo() {
	fmt.Println("\n-----------------------pointer-----------------------\n")
	v := 10
	p := &v        // creating pointer to v using &v
	*p = *p * 10   // accessing pointer value using *p
	fmt.Println(v) // Can see the change of value inside v
}

// Struct:
// It is a collection of fields.
// Fields can be accessed using dots

type Rectangle struct {
	L int
	B int
}

func structDemo() {
	fmt.Println("\n-----------------------struct-----------------------\n")
	rec := Rectangle{9, 18}
	// rec := &Rectangle{9, 18} // fields can still be accessed using "."
	// rec := Rectangle{L:9,B:18}
	// rec := Rectangle{B: 9, L: 18}
	// rec := Rectangle{}	// B : 0 && L : 0 implicitly
	fmt.Println(rec.L)

	recPointer := &rec // Can access the values of rec without "*"
	fmt.Println(recPointer.B)
}

// Arrays:
// Work same as other languages the syntax changes:
// The type [n]T is an array of n values of type T.

func arrayDemo() {
	fmt.Println("\n-----------------------arrays-----------------------\n")
	var arr [5]int // 1st way to declare
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	// arr[4] = 5 	//It will "0" by default
	fmt.Println(arr)

	arr1 := [2]string{"Hello", "world!"} // 2nd way to declare
	fmt.Println(arr1)
}

// Slices:
// Slices is a dinamic in size. We can also give lower and upperboud while slicing an array shown in example below :

func slicesDemo() {
	fmt.Println("\n-----------------------slices-----------------------\n")

	prime := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(prime)

	// slice := prime[1 : 4]; 	// slice prime array from index 1 to 4. Does not include index 4.
	// var slice = prime[1 : 4];
	var slice []int = prime[1:4]
	fmt.Println(slice)

	// Slice does not store any data. It just describe the array it silced from.
	// Changing anything in slice will change the values in the underlying array also.
	slice2 := prime[1:5]
	slice2[2] = 111 // Changin the slice means changin prime array.
	fmt.Println(slice2, prime)

	// Slice literal is like an array literal but without the size
	// Under the hood it will create and array and then store the slice of that array inside slice
	nameSlice := []string{"Anshul", "Rahul", "Mayank", "Rohan"}
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(nameSlice)
	fmt.Println(s)

	// Not givinf the lower and upper boud means the full array.
	// lower bound means 0 by default and upperbound by default is the length of the array
	// sliceNames := nameSlice[0:]
	// sliceNames := nameSlice[:4]
	// sliceNames := nameSlice[0:4]
	sliceNames := nameSlice[:]
	fmt.Println(sliceNames)

	// A slice has a length of how many items it has but it has a capacity if its underlying array.
	fmt.Println(len(slice), cap(slice)) // Length = 3 : as it has 3 elements but capacity is 5 as its underlying array *prime* has the capacity of 5

	// The zero value of the slice is nil.
	// This type of slices are called "nil slices"
	// It has a capacity of 0 elements and no underlying array
	var nilSlice []int
	fmt.Printf("Slice : %v, Length : %v, Capacity : %v \n", nilSlice, len(nilSlice), cap(nilSlice))

	// Slices can be created using builtin "make()"" function.
	// It is used to create dynamic sized array
	// It created a zeroed sized array and return a slice that point to that array

	makeArray2Arg := make([]int, 5) // integer array with length and capacity of 5 by deafult all values are zero
	fmt.Println(makeArray2Arg, len(makeArray2Arg), cap(makeArray2Arg))
	makeArray3Arg := make([]int, 0, 10) // Integer array length 0 and capacity of 10, because length is 0 so the slice is empty
	fmt.Println(makeArray3Arg, len(makeArray3Arg), cap(makeArray3Arg))

	// A slice can contain any type including other slices.
	// board is a 2 dimentional slice containg the slice of string datatype
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[0][2] = "X"
	board[1][0] = "O"
	board[0][1] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// if we want to append to slices we can use "append()" function
	// func append(s []T, vs ...T) []T
	// The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
	appendSlice := []int{1, 2, 5, 3, 6}
	fmt.Println(appendSlice, len(appendSlice), cap(appendSlice))

	appendSlice = append(appendSlice, 10)                        // You can see it only had capacity of 5 and it was full
	fmt.Println(appendSlice, len(appendSlice), cap(appendSlice)) // But its capacity is increased automaticaly to append new element

	appendSlice = append(appendSlice, 15, 25, 35) // We can also append more then 1 element at a time
	fmt.Println(appendSlice, len(appendSlice), cap(appendSlice))
}

// Range in for loop:
// The "range" form of "for" loop can be used on a slice or map.

func rangeDemo() {
	fmt.Println("\n-----------------------slices-----------------------\n")
	p := []int{1, 3, 5, 9, 2, 7, 8, 2, 5}

	// i is index and v is the element
	for i, v := range p {
		fmt.Print(i, v, " ")
	}
	fmt.Println()

	// You can skip the index or value by assigning to _.
	// for _, value := range pow
	for _, v := range p {
		fmt.Print(v, " ")
	}

	fmt.Println()
	// If you only want the index, you can omit the second variable.
	for i := range p {
		fmt.Print(i, " ")
	}
}

// Map
// Can store key and value. (Map the values)
// can be created using "make()" function
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
func mapDemo() {
	fmt.Println("\n-----------------------map-----------------------\n")
	type Fullname struct {
		firstName, lastName string
	}

	name := make(map[int]Fullname)
	name[1] = Fullname{
		"Anshul", "Saini",
	}

	fmt.Printf("Full name : %v %v \n", name[1].firstName, name[1].lastName)

	// Map literals are like struct literal but it require values
	type Location struct {
		longitude, latitude float32
	}
	locations := map[string]Location{
		// We can write Location here to the value
		"ladwa": {
			1.26445468546, 2.65659526559,
		},
		// Or we can ignore it if the type is just type name
		"radaur": {
			1.46844658435, 2.35465496646,
		},
	}
	fmt.Println("Locations : ", locations["ladwa"], locations["radaur"])

	// Mutating Maps
	m := make(map[string]int)
	// Inserting and updating value
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	// Retriving value
	elem := m["Answer"]
	fmt.Println("The value:", elem)
	// Deleting key and value
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	// To check if map has the key or not
	// If it has that key then "ok" will be "true" otherwise "false"
	// If key is not in the map, then elem is the zero value for the map's element type.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// Function are like just values a variable can store function like it can store a value
func functionDemo() {
	fmt.Println("\n-----------------------function again-----------------------\n")
	functionAsValues()
	clouserDemo()
	fibonacciExercise()
}

// Function as values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionAsValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// Function closures
// Defination: A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
// A function may be a clsure.
// Adder function return a function everytime adder function is called it will return a clouser
// every cluser is bound to different "sum"
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func clouserDemo() {
	for i := 0; i < 10; i++ {
		fmt.Print(adder()(i), " ")
	}
	fmt.Println()
}

// Fibonacci Exercise
func fibonacci() func() int {
	current, next := 0, 1

	return func() int {
		ret := current

		current, next = next, current+next
		return ret
	}
}

func fibonacciExercise() {
	f := fibonacci() // We have instance of that function so acny variable insode that is not going to rest.
	fmt.Print("Fibonacci Series: ")
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
	fmt.Println()
}

// Methods:
// Go does not have classes. However, you can define methods on types.
// Rectangle type is already defined in the file
// A method is a function with a special receiver argument.
// Receiver is written between "func" and "function name"
// Rectangle act as a reciever which will recieve "area()" then we can call it using reciever.
// A method is just a function with a receiver argument.

// Method
func (r Rectangle) area() {
	fmt.Println("(Method) Area of the rectangle is : ", r.B*r.L)
}

// Function
func Area(r Rectangle) {
	fmt.Println("(Function) Area of the rectangle is : ", r.B*r.L)
}

// Method can be declare on non struct type too. (Rectangle is struct)
// It will only work on the types that are declared in the same package.
type Dog string

func (d Dog) speak() {
	fmt.Println(d, " : bhaww bhaww bhaww")
}

// Pointer receivers:
// Same as receiver but the type "T" is written with "*" in front of it like "*T"
// Remember T itself can be a pointer.
// Methods with pointer recievers can change/modify the value of reciever as scale do.
func (r *Rectangle) scale(s int) {
	r.B = r.B * s
	r.L = r.L * s
	fmt.Printf("L : %v, B : %v \n", r.B, r.L)
}

// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

func methodsDemo() {
	fmt.Println("\n-----------------------Methods-----------------------\n")
	r := Rectangle{12, 6}
	r.area()
	Area(r)

	d := Dog("Sheru")
	d.speak()

	r.scale(10)
	// Values are changed by scale method
	r.area()
}

// Interfaces:
// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
type ShapeArea interface {
	area()
}
type Circle struct {
	radius int
}

// We can create methods with same name but there receivers should be different
func (c *Circle) area() {
	if c == nil {
		fmt.Println("(Method) Rectangle is nil, cannot calculate area.")
		return
	}
	area := math.Pi * (float64(c.radius * c.radius))
	fmt.Println("Area of the circle : ", area)
}

// Interfaces are implemented implicitly
// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

func interfaceDemo() {
	fmt.Println("\n-----------------------Interface-----------------------\n")
	var a ShapeArea
	r := Rectangle{10, 12}
	c := &Circle{10}
	a = r // if Rectangle does not implement Shapearea it will throw error
	a.area()
	a = c // if Circle does not implement Shapearea it will throw error
	a.area()

	// Better way to for implementation check
	var rect ShapeArea = Rectangle{10, 12}
	rect.area()

	// 	Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
	// (value, type)
	fmt.Printf("%v , %T \n", rect, rect) // {10 12} , main.Rectangle
	// Interface values with nil underlying values
	// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
	// In some languages this would trigger a null pointer exception,
	// but in Go it is common to write methods that gracefully handle being called with a nil receiver
	var nilCircle *Circle
	a = nilCircle
	fmt.Printf("%v , %T \n", a, a)
	a.area()

	// Nil interface values
	// A nil interface value holds neither value nor concrete type.
	// Calling a method on a nil interface will give a run-time error as there is nothing to run that on.
	// There is no tuple that implemented that method.
	// var s ShapeArea;
	// s.area() //It will cause run time error

	// 	The interface type that specifies zero methods is known as the empty interface:
	// interface{}
	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	fmt.Println("Empty Interface : ")
	var i interface{}
	fmt.Printf("%v , %T \n", i, i)
	i = 42
	fmt.Printf("%v , %T \n", i, i)
	i = "hello"
	fmt.Printf("%v , %T \n", i, i)

	// Type assertions
	fmt.Println("Type assertion check: ")
	check1 := i.(string)
	fmt.Println(check1)
	// 	This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
	// If i does not hold a T, the statement will trigger a panic.
	// We can check it using "ok" as it returns 2 values: 1. values and 2. status (True/false)
	// In this case "i" has concrete type "string" so the value is "hello" and status is "true"
	check2, ok := i.(string)
	fmt.Println(check2, ok)
	// In this case "i" does not has concrete type "int" so the value is "0" and status is "false"
	check3, ok := i.(int)
	fmt.Println(check3, ok)

	// 	Type switches
	// A type switch is a construct that permits several type assertions in series.
	// A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.
	/* switch v := i.(type) {
	case T:
		// here v has type T
	case S:
		// here v has type S
	default:
		// no match; here v has the same type as i
		} */
	fmt.Println("Type switches check: ")
	typeSwitchCheck(21)
	typeSwitchCheck("Hello")
	typeSwitchCheck(true)
}

func typeSwitchCheck(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

//Stringers
// It comes with fmt package
//One of the most universal interfaces is Stringer defined by the fmt package.
/*
	type Stringer interface {
		String() string
		}
*/
// A Stringer is a type that can describe itself as a string.
// The fmt package (and many others) look for this interface to print values.
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringerDemo() {
	fmt.Println("\n-----------------------Stringer-----------------------\n")
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	// Note that other struct describe themselve like {"__", 215} ,
	// but it is describing itself as string "Arthur Dent (42 years)"
	fmt.Println(a, z)
}

/*
Errors
Go programs express error state with error values.
The error type is a built-in interface similar to fmt.Stringer:

type error interface {
    Error() string
	}

	(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)
	Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
		}
		fmt.Println("Converted integer:", i)

		A nil error denotes success; a non-nil error denotes failure.
*/
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func errorDemo() {
	fmt.Println("\n-----------------------Error-----------------------\n")
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// Readers:
// The io package specifies the io.Reader interface,
// which represents the read end of a stream of data.
// The io.Reader interface has a Read method:
// func (T) Read(b []byte) (n int, err error)
// It return "io.EOF" error when stream ends.

func readerTesting() {
	fmt.Println("\n-----------------------Readers-----------------------\n")
	str := strings.NewReader("This is the reader text.")

	buffer := make([]byte, 8)
	for {
		n, err := str.Read(buffer)
		if err != nil {
			fmt.Println("Stream ended")
			break
		}
		fmt.Printf("Stream text: %q \n", buffer[:n])
	}
}

// Images
// Package image defines the Image interface:
// package image
/*
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
	}
*/
// Note: the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.

// Exercise: Images
// Inside : "exercise(Image).go"

// Type parameters
// Function in go can be written to work on multiple type inputs.
// The "type parameter" function is added between "function name" and "function agrument" inside "[] brackets"
// func Index[T comparable](s []T, x T) int
// "compareable" is a constraint that let us use "==" and "!=" compareable operater.
// We can change constraint to "string", "int" or multiple like "string | int"
// We can also add multiple "type parameters" like "[T string | int, S string | int]"
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// We can add multiple "type parameters" same as "func type parameters"
type List[T any] struct {
	next *List[T]
	val  T
}

func listImplement() {
	head := List[int]{nil, 10}
	l1 := List[int]{nil, 20}
	head.next = &l1
	l2 := List[int]{nil, 30}
	l1.next = &l2
	l3 := List[int]{nil, 40}
	l2.next = &l3

	fmt.Print("List Implementation using type parameter in struct : ")
	for i := &head; i != nil; i = i.next {
		fmt.Print(i.val, " ")
	}
	fmt.Println()
}

func typeParameterDemo() {
	fmt.Println("\n-----------------------Type Parameters-----------------------\n")
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println("On 'int' type with index function:", Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println("On 'string' type with index function:", Index(ss, "hello"))

	listImplement()
}

// Go routines :
/*
A goroutine is a lightweight thread managed by the Go runtime.

go f(x, y, z)
starts a new goroutine running

f(x, y, z)
The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func routineDemo() {
	fmt.Println("\n-----------------------Go Routines-----------------------\n")
	go say("New go routine thread")
	say("Current go routine thread")
}

// Channels :
// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
// ch <- v    // Send v to channel ch.
// v := <-ch  // Receive from ch, and
//            // assign value to v.
// (The data flows in the direction of the arrow.)
// Like maps and slices, channels must be created before use:
// ch := make(chan int)
// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

// Remember we need to write "datatype" whlie passing "channel" as "chan" keywoed only tells it is a channel not its datatype
func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}

// Buffered Channels
// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
// ch := make(chan int, 100)
// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

func bufferChannelDemo() {
	// 1. Create a channel with a buffer capacity of 2
	ch := make(chan int, 2)

	// 2. Fill the buffer completely
	ch <- 1
	ch <- 2

	// 3. Try to send a 3rd item (Overfilling!)
	// The program stops here and waits for someone to read.
	// Since no one is reading, it crashes.
	ch <- 3

	// This code will never be reached
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Range and Close
// A sender can close a channel to indicate that no more values will be sent.
// Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
// v, ok := <-ch
// ok is false if there are no more values to receive and the channel is closed.
// Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
// If we want to pass a read only channel or write only channel to a function we can use this while declaring function parameters:
// Read-only
// ch <-chan string
// Write-only
// ch chan<- string

func fibonacciClose(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Another note: Channels aren't like files; you don't usually need to close them.
	// Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	close(c)
}

func closeDemo() {
	c := make(chan int, 10)
	go fibonacciClose(cap(c), c)
	// The loop for i := range c receives values from the channel repeatedly until it is closed.
	for i := range c {
		fmt.Println(i)
	}
	a, ok := <-c
	fmt.Println(a, ok)
}

func channelDemo() {
	fmt.Println("\n-----------------------Channel Demo-----------------------\n")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	ch := make(chan int)
	// go fmt.Println(sum(arr[:len(arr)/2], ch))
	// go fmt.Println(sum(arr[len(arr)/2:], ch))

	// With channel we must use go routines
	go sum(arr[:len(arr)/2], ch)
	go sum(arr[len(arr)/2:], ch)
	fh, sh := <-ch
	// fmt.Printf("%T", ch)
	fmt.Println(fh)
	fmt.Println(sh)

	// It is going to cause error, thats why it is commented
	// bufferChannelDemo()

	// Channel Close
	closeDemo()
}

// Select :
// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// sending value through channel to the side routine
		// it waits here untill the side routine recieves the value
		case c <- x:
			x, y = y, x+y
			// sending value through channel to the side routine
			// When the c stops receiving the value in side routine it receives the value and quit.
		case <-quit:
			fmt.Print("quit")
			fmt.Println()
			return
		}
	}
}

func selectChannelDemo() {
	fmt.Println("\n-----------------------Select Channel Demo-----------------------\n")
	c := make(chan int) // Unbuffered channel 1 time 1 value and forced sync
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// Recieving and printing value through channel from the main go routine
			// After the loop ends (10th time) it will stop receive the value
			fmt.Print(<-c, " ")
		}
		// Sending value through channel from the main go routine
		// It will send the value to the main routine to stop waiting and quit.
		quit <- 0
	}()
	fibonacciSelect(c, quit)

	selectDefaultDemo()
}

func selectDefaultDemo() {
	start := time.Now()
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// sync.Mutex :
// if we don't need communication if and we just want to make sure that,
// only one goroutine can access a variable at a time to avoid conflicts?
// This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.
// Go's standard library provides mutual exclusion with sync.Mutex and its two methods:
// Lock
// Unlock
// We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) NotLockedInc(key string) {
	// not locked so alot of goroutines can access the map c.v. at the same time
	c.v[key]++
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}


func mutexDemo() {
	fmt.Println("\n-----------------------sync.Mutex-----------------------\n")
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
		// go c.NotLockedInc("somekey") // Will throw error because many routines will try to write at the same time.
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

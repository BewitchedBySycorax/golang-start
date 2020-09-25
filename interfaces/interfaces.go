package main

import (
	"fmt"
	"reflect"
)

type animal interface {
	makeSound()
}

type cat struct{}
type dog struct{}
type eagle struct{}

// Methods
func (c *cat) makeSound() {
	fmt.Println("meow!")
}

func (d *dog) makeSound() {
	fmt.Println("woof!")
}

//=============================================================

type greeter interface {
	greet(string) string
}

type russian struct{}
type american struct{}

// Methods
func (r *russian) greet(name string) string {
	return fmt.Sprintf("Привет, %s", name)
}

func (a *american) greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

// Function
// Passing an interface as an argument we can omit that thing what exactly object is incoming - nevermind
func sayHello(g greeter, name string) {
	fmt.Println(g.greet(name))
}

//=============================================================
// Composite interfaces

type compositeAnimal interface {
	walker
	runner
}

type compositeBird interface {
	walker
	flier
}

type walker interface {
	walk()
}

type flier interface {
	fly()
}

type runner interface {
	run()
}

// Methods
func (c *cat) walk() {
	fmt.Println("cat is walking!")
}

func (c *cat) run() {
	fmt.Println("cat is running!")
}

func (e *eagle) walk() {
	fmt.Println("eagle is walking!")
}

func (e *eagle) fly() {
	fmt.Println("eagle is flying!")
}

// Function
// Passing the 'walker' interface
func walk(w walker) {
	// Either bird or animal
	w.walk() // 'walk' method is called (from 'walker' interface)
}

func main() {
	// Two vars assigning type 'cat' and type 'dog'
	var c, d animal = &cat{}, &dog{}
	var cc compositeAnimal = &cat{}
	var e compositeBird = &eagle{}
	c.makeSound()
	d.makeSound()
	cc.walk()
	cc.run()
	e.walk()
	e.fly()
	walk(cc)
	walk(e)
	sayHello(&russian{}, "Алексей")
	sayHello(&american{}, "Aleksei")
	m := map[string]interface{}{} // Map (associative array? No! It's map (карта)) which has 'string' key and empty 'interface' value
	m["one"] = 1                  // Let's add any key to out map
	m["two"] = 2.0                // Let's add any key to out map
	m["three"] = true             // Let's add some keys to out map

	for k, v := range m {
		switch v.(type) {
		case int:
			fmt.Printf("%s is an integer\n", k)
		case float64:
			fmt.Printf("%s is a float64\n", k)
		default:
			fmt.Printf("%s is %v\n", k, reflect.TypeOf(v))
		}
	}
}

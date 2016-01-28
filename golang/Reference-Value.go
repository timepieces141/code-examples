package main

import (
	"fmt"
)

type Foo interface {
	Bar()
}

type FooImpl struct {
	Foobar int
}

func (fi *FooImpl) Bar() {
	fi.Foobar++
}

func main() {
	foo := FooImpl{5}
	doSomething(&foo)
	fmt.Println(foo.Foobar)
}

func doSomething(f Foo) {
	f.Bar()
	doSomethingFurther(f)
}

func doSomethingFurther(f Foo) {
	f.Bar()
}
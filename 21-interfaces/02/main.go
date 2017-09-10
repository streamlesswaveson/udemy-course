package main

import "fmt"

type Foo interface {
	DoFoo()
}

type Bar interface {
	DoBar()
}

type FooBar interface {
	Foo
	Bar
}

type MyFooBar struct {
}

func (m MyFooBar) DoFoo()  {
	fmt.Println("doing foo")
}

func (m MyFooBar) DoBar()  {
	fmt.Println("doing bar")
}

func runFoo(f Foo)  {
	f.DoFoo()
}

func runBar(b Bar)  {
	b.DoBar()
}

func runFooBar(fb FooBar)  {
	fb.DoFoo()
	fb.DoBar()

}

func main()  {

	m := MyFooBar{}

	runFoo(m)
	runBar(m)
	runFooBar(m)
	
}

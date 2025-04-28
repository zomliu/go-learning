package main

type Character interface {
	Show()
}

type ConcreteCharacter struct {
	name string
}

func (c *ConcreteCharacter) SetName(name string) {
	c.name = name
}

func (c *ConcreteCharacter) Show() {
	println(c.name)
}

type FlyweightFactory struct {
	pool map[string]Character
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		pool: make(map[string]Character),
	}
}

func (f *FlyweightFactory) GetCharacter(name string) Character {
	if _, ok := f.pool[name]; !ok {
		f.pool[name] = &ConcreteCharacter{name: name}
	}
	return f.pool[name]
}

func main() {
	factory := NewFlyweightFactory()
	c1 := factory.GetCharacter("A")
	c2 := factory.GetCharacter("B")
	c3 := factory.GetCharacter("C")

	c1.Show()
	c2.Show()
	c3.Show()
}
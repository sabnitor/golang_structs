package zoo

type Animal struct {
	name string
}

func (a Animal) beenAnimal() {
}

type Animaler interface {
	beenAnimal()
}

type Elephant struct {
	Animal
}

type Giraffe struct {
	Animal
}

type Camel struct {
	Animal
}

type Tiger struct {
	Animal
}

type Lion struct {
	Animal
}

type Monkey struct {
	Animal
}

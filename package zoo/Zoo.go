package zoo

type zoo struct {
	Zookeepers []Zookeeper
	Jails      []Jail
	Animals    []Animal
}

type Jail struct {
	close   bool
	whoose  string
	consist Animal
}

type Zookeeper struct {
	name string
}

func (z Zookeeper) CheckJail(jail Jail) {
}

func (z Zookeeper) LookForAnimal() {

}

func (z Zookeeper) CatchAnimal() {

}

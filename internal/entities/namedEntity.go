package entities

type NamedEntity struct {
	name string
}

func (n *NamedEntity) GetName() string {
	return n.name
}

func (n *NamedEntity) SetName(name string) {
	n.name = name
}

package flyweight

import "fmt"

type Soldier struct {
	name string
}

func newSoldier(name string) *Soldier {
	return &Soldier{
		name: name,
	}
}

func (s *Soldier) Promote(context *Context) {
	fmt.Printf("%s %s promoted %d\n", s.name, context.id, context.star)
}

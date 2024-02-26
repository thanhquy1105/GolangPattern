package composite

import "fmt"

type Folder struct {
	Components []Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("searching recursively for keyword %s in folder %s\n", keyword, f.Name)

	for _, i := range f.Components {
		i.Search(keyword)
	}
}

func (f *Folder) Add(component Component) {
	f.Components = append(f.Components, component)
}

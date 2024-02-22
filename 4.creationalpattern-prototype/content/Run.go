package prototype

import "fmt"

func Run() {
	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	file3 := &File{Name: "File3"}

	folder1 := &Folder{
		Childrens: []INode{file1},
		Name:      "Folder 1",
	}
	folder2 := &Folder{
		Childrens: []INode{folder1, file2, file3},
		Name:      "Folder 2",
	}

	fmt.Println("\n Printing for Folder 2")
	folder2.Print("   ")

	cloneFolder := folder2.Clone()
	fmt.Println("\n Printing for clone Folder 2")
	cloneFolder.Print("   ")

}

package composite

func Run() {
	file1 := &File{"file1"}
	file2 := &File{"file2"}
	file3 := &File{"file3"}

	folder1 := &Folder{Name: "folder1"}
	folder1.Add(file1)

	folder2 := &Folder{Name: "folder2"}
	folder2.Add(file2)
	folder2.Add(file3)

	folder2.Add(folder1)

	folder2.Search("value")
}

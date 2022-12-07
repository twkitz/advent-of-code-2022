package helper

type Item interface {
	GetType() string
	GetName() string
	GetSize() int
}

type File struct {
	name string
	size int
}

type Dir struct {
	parent       *Dir
	name         string
	fileChildren []*File
	dirChildren  []*Dir
}

func (file File) GetType() string {
	return "File"
}

func (file File) GetName() string {
	return file.name
}

func (file File) GetSize() int {
	return file.size
}

func (dir Dir) GetType() string {
	return "Dir"
}

func (dir Dir) GetName() string {
	return dir.name
}

func (dir Dir) GetSize() int {
	sum := 0
	for _, item := range dir.fileChildren {
		sum += (*item).GetSize()
	}
	for _, item := range dir.dirChildren {
		sum += (*item).GetSize()
	}
	return sum
}

func (dir *Dir) AddFile(name string, size int) {
	file := File{name, size}
	dir.fileChildren = append(dir.fileChildren, &file)
}

func (dir *Dir) AddDir(name string) *Dir {
	childDir := Dir{parent: dir, name: name, fileChildren: []*File{}, dirChildren: []*Dir{}}
	dir.dirChildren = append(dir.dirChildren, &childDir)
	return &childDir
}

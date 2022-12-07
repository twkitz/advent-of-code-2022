package helper

import "testing"

func TestFileGetSize(t *testing.T) {
	file := File{name: "fileName.txt", size: 100}
	got := file.GetSize()
	if got != file.size {
		t.Error("want:", file.size, "got:", got)
	}
}

func TestDirGetSize(t *testing.T) {
	dir := Dir{
		name: "a",
		fileChildren: []*File{
			&File{name: "a.txt", size: 100},
			&File{name: "b.txt", size: 50},
		},
	}
	got := dir.GetSize()
	if got != 150 {
		t.Error("want:", 150, "got:", got)
	}
}

func TestDirGetSize_FilesAndDirs(t *testing.T) {
	dir := Dir{
		name: "a",
		fileChildren: []*File{
			&File{name: "a.txt", size: 100},
			&File{name: "b.txt", size: 50},
		},
		dirChildren: []*Dir{
			&Dir{
				name: "c",
				fileChildren: []*File{
					&File{name: "d.txt", size: 100},
					&File{name: "e.txt", size: 50},
				},
				dirChildren: []*Dir{},
			},
		},
	}
	got := dir.GetSize()
	if got != 300 {
		t.Error("want:", 300, "got:", got)
	}
}

func TestDirAddFile(t *testing.T) {
	dir := Dir{name: "a", fileChildren: []*File{}}
	dir.AddFile("b.txt", 100)
	if len(dir.fileChildren) != 1 {
		t.Error("add file error")
	}
}

func TestDirAddDir(t *testing.T) {
	dir := Dir{name: "a", dirChildren: []*Dir{}}
	dir.AddDir("b")
	if len(dir.dirChildren) != 1 {
		t.Error("add file error")
	}
}

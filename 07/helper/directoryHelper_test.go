package helper

import "testing"

func TestChangeDir(t *testing.T) {
	root := Dir{nil, "/", []*File{}, []*Dir{}}
	root.AddDir("a")
	system := System{workingDir: &root}

	system.ChangeDir("a")

	child := root.dirChildren[0]
	if system.workingDir.GetName() != child.GetName() {
		t.Error("change dir error")
	}
}

func TestChangeDir_Parent(t *testing.T) {
	root := Dir{nil, "/", []*File{}, []*Dir{}}
	root.AddDir("a")
	child := root.dirChildren[0]
	system := System{workingDir: child}

	system.ChangeDir("..")

	if system.workingDir.GetName() != root.GetName() {
		t.Error("change dir error")
	}
}

func TestGetAnswerPart1(t *testing.T) {
	want := 95437
	got := GetAnswerPart1("../example.txt")
	if got != want {
		t.Error("want: ", want, "got: ", got)
	}
}

func TestGetAnswerPart2(t *testing.T) {
	want := 24933642
	got := GetAnswerPart2("../example.txt")
	if got != want {
		t.Error("want: ", want, "got: ", got)
	}
}

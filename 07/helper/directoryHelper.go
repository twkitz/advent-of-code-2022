package helper

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type System struct {
	workingDir *Dir
}

func (system *System) ChangeDir(target string) {
	if target == ".." {
		system.workingDir = system.workingDir.parent
		return
	}

	for _, item := range system.workingDir.dirChildren {
		if item.GetName() == target {
			system.workingDir = item
			break
		}
	}
}

func GetDirList(filePath string) (*[]*Dir, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rootDir := Dir{nil, "/", []*File{}, []*Dir{}}
	system := System{&rootDir}
	dirList := []*Dir{&rootDir}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		if line == "$ cd /" {

		} else if line == "$ ls" {
			scanner.Scan()
			ln := scanner.Text()

			for ln[0] != '$' {
				if ln[0:3] == "dir" {
					childAddr := system.workingDir.AddDir(ln[4:])
					dirList = append(dirList, childAddr)
				} else {
					fileInfo := strings.Split(ln, " ")
					fileSize, _ := strconv.Atoi(fileInfo[0])
					system.workingDir.AddFile(fileInfo[1], fileSize)
				}
				scanner.Scan()
				ln = scanner.Text()
				if ln == "" {
					break
				}
				if ln[0:4] == "$ cd" {
					system.ChangeDir(ln[5:])
					break
				}
			}
		} else if line[0:4] == "$ cd" {
			system.ChangeDir(line[5:])
		}
	}

	return &dirList, nil
}

func GetAnswerPart1(filePath string) int {
	dirList, err := GetDirList(filePath)

	if err != nil {
		return -1
	}

	sum := 0
	for _, dir := range *dirList {
		dirSize := dir.GetSize()
		if dirSize <= 100000 {
			sum += dirSize
		}
	}

	return sum
}

func GetAnswerPart2(filePath string) int {
	dirListAddr, err := GetDirList(filePath)

	if err != nil {
		return -1
	}

	dirList := *dirListAddr
	totalSpace := 70000000
	usingSpace := dirList[0].GetSize()
	remainSpace := totalSpace - usingSpace
	requireAtLeast := 30000000 - remainSpace
	answer := math.MaxInt
	for _, dir := range dirList {
		dirSize := dir.GetSize()
		if dirSize >= requireAtLeast && dirSize < answer {
			answer = dirSize
		}
	}

	return answer
}

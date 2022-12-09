package helper

import "fmt"

type Rope struct {
	headX int
	headY int
	tailX int
	tailY int
}

func (rope Rope) GetTailPosition() string {
	return fmt.Sprintf("%d,%d", rope.tailX, rope.tailY)
}

func (rope *Rope) MoveUp() {
	rope.headY -= 1
	rope.UpdateTailPosition()
}

func (rope *Rope) MoveDown() {
	rope.headY += 1
	rope.UpdateTailPosition()
}

func (rope *Rope) MoveLeft() {
	rope.headX -= 1
	rope.UpdateTailPosition()
}

func (rope *Rope) MoveRight() {
	rope.headX += 1
	rope.UpdateTailPosition()
}

func (rope *Rope) UpdateTailPosition() {
	if rope.headX == rope.tailX-2 {
		rope.tailX -= 1
		rope.tailY = rope.headY
	} else if rope.headX == rope.tailX+2 {
		rope.tailX += 1
		rope.tailY = rope.headY
	} else if rope.headY == rope.tailY-2 {
		rope.tailY -= 1
		rope.tailX = rope.headX
	} else if rope.headY == rope.tailY+2 {
		rope.tailY += 1
		rope.tailX = rope.headX
	}
}

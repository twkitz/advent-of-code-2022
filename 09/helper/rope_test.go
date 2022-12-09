package helper

import (
	"fmt"
	"testing"
)

func TestGetTailPosition(t *testing.T) {
	type test struct {
		x    int
		y    int
		want string
	}

	tests := []test{
		{0, 0, "0,0"},
		{1, 0, "1,0"},
		{0, 1, "0,1"},
		{-1, 0, "-1,0"},
		{0, -1, "0,-1"},
	}

	for _, tc := range tests {
		rope := Rope{tailX: tc.x, tailY: tc.y}
		got := rope.GetTailPosition()
		if got != tc.want {
			t.Error(fmt.Sprintf(`want: "%s", got: "%s"`, tc.want, got))
		}
	}
}

func TestMoveUp(t *testing.T) {
	rope := Rope{headX: 0, headY: 0}
	want := -1

	rope.MoveUp()

	if rope.headX != 0 || rope.headY != want {
		t.Error(fmt.Sprintf(`want: 0,%d -> got: %d,%d`, want, rope.headX, rope.headY))
	}
}

func TestMoveDown(t *testing.T) {
	rope := Rope{headX: 0, headY: 0}
	want := 1

	rope.MoveDown()

	if rope.headX != 0 || rope.headY != want {
		t.Error(fmt.Sprintf(`want: 0,%d -> got: %d,%d`, want, rope.headX, rope.headY))
	}
}

func TestMoveLeft(t *testing.T) {
	rope := Rope{headX: 0, headY: 0}
	want := -1

	rope.MoveLeft()

	if rope.headX != want || rope.headY != 0 {
		t.Error(fmt.Sprintf(`want: %d,0 -> got: %d,%d`, want, rope.headX, rope.headY))
	}
}

func TestMoveRight(t *testing.T) {
	rope := Rope{headX: 0, headY: 0}
	want := 1

	rope.MoveRight()

	if rope.headX != want || rope.headY != 0 {
		t.Error(fmt.Sprintf(`want: %d,0 -> got: %d,%d`, want, rope.headX, rope.headY))
	}
}

func TestUpdateTailPosition(t *testing.T) {
	type test struct {
		headX int
		headY int
		wantX int
		wantY int
	}

	tests := []test{
		{0, 0, 0, 0},
		{-1, 0, 0, 0},
		{1, 0, 0, 0},
		{0, -1, 0, 0},
		{0, 1, 0, 0},
		{-2, 0, -1, 0},
		{2, 0, 1, 0},
		{0, -2, 0, -1},
		{0, 2, 0, 1},
		{-2, -1, -1, -1},
		{-2, 1, -1, 1},
		{2, -1, 1, -1},
		{2, 1, 1, 1},
		{-1, -2, -1, -1},
		{1, -2, 1, -1},
		{-1, 2, -1, 1},
		{1, 2, 1, 1},
	}

	for _, tc := range tests {
		rope := Rope{tc.headX, tc.headY, 0, 0}
		rope.UpdateTailPosition()
		if rope.tailX != tc.wantX || rope.tailY != tc.wantY {
			t.Error(fmt.Sprintf(`want: %d,%d -> got: %d,%d`, tc.wantX, tc.wantY, rope.tailX, rope.tailY))
		}
	}
}

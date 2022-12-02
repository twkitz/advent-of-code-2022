package rockPaperScissorPart2

import "testing"

func TestGetShape(t *testing.T) {
    type test struct {
        input   string
        want    string
    }

    tests := []test{
        {input: "A", want: "Rock"},
        {input: "B", want: "Paper"},
        {input: "C", want: "Scissor"},
    }

    for _, tc := range tests {
        got, err := GetShape(tc.input)
        if err != "" {
            t.Error("Error: " + err)
        } else if got != tc.want {
            t.Error("input:", tc.input, ", want:", tc.want, ", got:", got)
        }
    }
}

func TestGetOurShape(t *testing.T) {
    type test struct {
        opponent    string
        us          string
        want        string
    }

    tests := []test{
        {opponent: "Rock", us: "X", want: "Scissor"},
        {opponent: "Rock", us: "Y", want: "Rock"},
        {opponent: "Rock", us: "Z", want: "Paper"},
        {opponent: "Paper", us: "X", want: "Rock"},
        {opponent: "Paper", us: "Y", want: "Paper"},
        {opponent: "Paper", us: "Z", want: "Scissor"},
        {opponent: "Scissor", us: "X", want: "Paper"},
        {opponent: "Scissor", us: "Y", want: "Scissor"},
        {opponent: "Scissor", us: "Z", want: "Rock"},
    }

    for _, tc := range tests {
        got, _ := GetOurShape(tc.opponent, tc.us)
        if got != tc.want {
            t.Error("opponent:", tc.opponent, ", us:", tc.us, ", want:", tc.want, ", got:", got)
        }
    }
}

func TestGetShapePoint(t *testing.T) {
    type test struct {
        input   string
        want    int
    }

    tests := []test{
        {input: "Rock", want: 1},
        {input: "Paper", want: 2},
        {input: "Scissor", want: 3},
        {input: "A", want: 0},
    }

    for _, tc := range tests {
        got := GetShapePoint(tc.input)
        if got != tc.want {
            t.Error("input:", tc.input, ", want:", tc.want, ", got:", got)
        }
    }
}

func TestGetChallengePoint(t *testing.T) {
    type test struct {
        opponent    string
        us          string
        want        int
    }

    tests := []test{
        {opponent: "Rock", us: "Rock", want: 3},
        {opponent: "Rock", us: "Scissor", want: 0},
        {opponent: "Rock", us: "Paper", want: 6},
        {opponent: "Scissor", us: "Rock", want: 6},
        {opponent: "Scissor", us: "Scissor", want: 3},
        {opponent: "Scissor", us: "Paper", want: 0},
        {opponent: "Paper", us: "Rock", want: 0},
        {opponent: "Paper", us: "Scissor", want: 6},
        {opponent: "Paper", us: "Paper", want: 3},
    }

    for _, tc := range tests {
        got := GetChallengePoint(tc.opponent, tc.us)
        if got != tc.want {
            t.Error("opponent:", tc.opponent, ", us:", tc.us, ", want:", tc.want, ", got:", got)
        }
    }
}

func TestGetPoint(t *testing.T) {
    type test struct {
        opponent    string
        us          string
        want        int
    }

    tests := []test{
        {opponent: "A", us: "X", want: 3 + 0},
        {opponent: "A", us: "Y", want: 1 + 3},
        {opponent: "A", us: "Z", want: 2 + 6},
        {opponent: "B", us: "X", want: 1 + 0},
        {opponent: "B", us: "Y", want: 2 + 3},
        {opponent: "B", us: "Z", want: 3 + 6},
        {opponent: "C", us: "X", want: 2 + 0},
        {opponent: "C", us: "Y", want: 3 + 3},
        {opponent: "C", us: "Z", want: 1 + 6},
    }

    for _, tc := range tests {
        got := GetPoint(tc.opponent, tc.us)
        if got != tc.want {
            t.Error("opponent:", tc.opponent, ", us:", tc.us, ", want:", tc.want, ", got:", got)
        }
    }
}
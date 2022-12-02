package rockPaperScissorPart2

import "fmt"

func GetShape(alphabet string) (string, string) {
    if alphabet == "A" { return "Rock", "" }
    if alphabet == "B" { return "Paper", "" }
    if alphabet == "C" { return "Scissor", "" }
    return "", "unsupported alphabet" + alphabet
}

func GetOurShape(opponent, us string) (string, string) {
    if us == "X" {
        switch opponent {
        case "Rock": return "Scissor", ""
        case "Paper": return "Rock", ""
        case "Scissor": return "Paper", ""
        }
    }
    if us == "Y" {
        return opponent, ""
    }
    if us == "Z" {
        switch opponent {
            case "Rock": return "Paper", ""
            case "Paper": return "Scissor", ""
            case "Scissor": return "Rock", ""
            }
    }
    return "", "unsupported opponent: " + opponent + ", us: " + us
}

func GetShapePoint(shape string) int {
    if shape == "Rock" { return 1 }
    if shape == "Paper" { return 2 }
    if shape == "Scissor" { return 3 }
    return 0
}

func GetChallengePoint(opponent, us string) int {
    if opponent == us { return 3 }
    if 
        (opponent == "Rock" && us == "Paper") || 
        (opponent == "Scissor" && us == "Rock") ||
        (opponent == "Paper" && us == "Scissor") { 
            return 6
        }
    return 0
}

func GetPoint(opponent, us string) int {
    opponentShape, opponentShapeErr := GetShape(opponent)
    if opponentShapeErr != "" {
        fmt.Println("opponentShape error: " + opponentShapeErr)
    }

    ourShape, ourShapeErr := GetOurShape(opponentShape, us)
    if ourShapeErr != "" {
        fmt.Println("ourShape error: " + ourShapeErr)
    }

    shapePoint := GetShapePoint(ourShape)
    challengePoint := GetChallengePoint(opponentShape, ourShape)
    return shapePoint + challengePoint
}
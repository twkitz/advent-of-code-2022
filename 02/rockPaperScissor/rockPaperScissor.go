package rockPaperScissor

import "fmt"

func GetShape(alphabet string) (string, string) {
    if alphabet == "A" || alphabet == "X" { return "Rock", "" }
    if alphabet == "B" || alphabet == "Y" { return "Paper", "" }
    if alphabet == "C" || alphabet == "Z" { return "Scissor", "" }
    return "", "unsupported alphabet" + alphabet
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

    ourShape, ourShapeErr := GetShape(us)
    if ourShapeErr != "" {
        fmt.Println("ourShape error: " + ourShapeErr)
    }

    shapePoint := GetShapePoint(ourShape)
    challengePoint := GetChallengePoint(opponentShape, ourShape)
    return shapePoint + challengePoint
}
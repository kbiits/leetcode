func executeInstructions(n int, startPos []int, s string) []int {
    res := make([]int, 0, len(s))
    for startInstruction := 0; startInstruction < len(s); startInstruction++ {

        sTemp := s[startInstruction:]
        countRes := count(startPos, n, sTemp)
        // fmt.Println(sTemp, startPos, countRes)
        // fmt.Println(countRes)
        res = append(res, countRes)
    }

    return res
}

func count(startPos []int, n int, s string) int {
    counter := 0
    posY := startPos[0]
    posX := startPos[1]

    for _, c := range s {
        nextPosX := posX
        nextPosY := posY
        switch c {
            case 'R':
                nextPosX++
            case 'L':
                nextPosX--
            case 'D':
                nextPosY++
            case 'U':
                nextPosY--
        }

        if isOutOfBound(n, nextPosX, nextPosY) {
            return counter
        }

        counter++
        posY = nextPosY
        posX = nextPosX
    }

    return counter
}

func isOutOfBound(n, posX, posY int) bool {
    return posX >= n || posX < 0 || posY >= n || posY < 0
}

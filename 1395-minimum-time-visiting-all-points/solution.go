func minTimeToVisitAllPoints(points [][]int) int {
    n := len(points)

    if n <= 1 {
        return 0
    }

    time := 0
    startPoint := points[0]
    for i := 1; i < n; i++ {
        time += calculateSteps(startPoint, points[i])
        startPoint = points[i]
    }

    return time
}

func calculateSteps(pointA, pointB []int) int {
// ex: point A = (3, 2), point B = (6, -4)
// we just need to compute the different from each X and Y
// and the maximum value of X and Y will be 
// the minimum (optimal) steps to reach the point B from point A
    xA := pointA[0]
    xB := pointB[0]
    
    yA := pointA[1]
    yB := pointB[1]

    xDiff := xA - xB
    yDiff := yA - yB
    if xDiff < 0 {
        xDiff = -xDiff
    }
    if yDiff < 0 {
        yDiff = -yDiff
    }

    if xDiff > yDiff {
        return xDiff
    }

    return yDiff
}


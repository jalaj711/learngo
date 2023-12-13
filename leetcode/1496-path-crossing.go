package main

// import "fmt"

func isPathCrossing(path string) bool {
	isCrossed := false

	visited := make([][2]int, 0, len(path))
	pos := [2]int{0, 0}

	for _, dir := range path {
		visited = append(visited, [2]int{pos[0], pos[1]})
		switch dir {
		case 'N':
			pos[0]++
		case 'S':
			pos[0]--
		case 'E':
			pos[1]++
		case 'W':
			pos[1]--
		}
		// fmt.Printf("pos=%d, dir=%U\n", pos, dir)
		for _, point := range visited {
			if pos[0] == point[0] && pos[1] == point[1] {
				// fmt.Println("match found")
				isCrossed = true
				break
			}
		}
		if isCrossed {
			break
		}
	}

	return isCrossed
}

func main() {
	isPathCrossing("NESWW")
}

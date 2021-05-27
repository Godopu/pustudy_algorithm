package main

import "fmt"

type Point struct {
	row    int
	column int
}

type Problem struct {
	size        int
	problemMap  [30][30]int
	level       int
	exper       int
	location    *Point
	minValue    int
	minLocation int
}

func makeList(len, initValue int) []int {
	list := make([]int, len)

	for i := 0; i < len; i++ {
		list[i] = initValue
	}

	return list
}

func make2DList(len, intValue int) [][]int {
	var list [][]int

	for i := 0; i < len; i++ {
		list = append(list, makeList(len, intValue))
	}

	return list
}

func (prob *Problem) readInput() {
	fmt.Scanf("%d", &prob.size)

	for i := 0; i < prob.size+2; i++ {
		for j := 0; j < prob.size+2; j++ {
			if i%(prob.size+1) == 0 || j%(prob.size+1) == 0 {
				prob.problemMap[i][j] = 100
			} else {
				fmt.Scanf("%d", &prob.problemMap[i][j])
				if prob.problemMap[i][j] == 9 {
					prob.location = &Point{i, j}
				}
			}
		}
	}
}

func print2DArray(array [30][30]int, len int) {
	for i := 0; i < len; i++ {
		fmt.Println(array[i])
	}
}

var weightedMap [30][30]int
var visitedMap [30][30]int

func eatFish(p *Problem) {
	// stack
	var stack []*Point
	_ = stack

	for i := 0; i < p.size+2; i++ {
		for j := 0; j < p.size+2; j++ {
			weightedMap[i][j] = 100
			visitedMap[i][j] = 0
		}
	}

	weightedMap[p.location.row][p.location.column] = 0
	visitedMap[p.location.row][p.location.column] = 1
	stack = append(stack, p.location)

	for len(stack) != 0 {
		coord := stack[0]
		stack = stack[1:]

		weight := weightedMap[coord.row][coord.column]

		if p.problemMap[coord.row][coord.column] != 0 && p.problemMap[coord.row][coord.column] < p.level {
			if weight < p.minValue || weight == p.minValue && (coord.row*30+coord.column) < p.minLocation {
				p.minValue = weight
				p.minLocation = coord.row*30 + coord.column
			}
		}

		stack = pushNode(stack, p, &Point{coord.row + 1, coord.column}, weightedMap, visitedMap, weight+1)
		stack = pushNode(stack, p, &Point{coord.row - 1, coord.column}, weightedMap, visitedMap, weight+1)
		stack = pushNode(stack, p, &Point{coord.row, coord.column + 1}, weightedMap, visitedMap, weight+1)
		stack = pushNode(stack, p, &Point{coord.row, coord.column - 1}, weightedMap, visitedMap, weight+1)
	}
}

func pushNode(stack []*Point, p *Problem, coord *Point, weightedMap, visitedMap [30][30]int, weight int) []*Point {
	if visitedMap[coord.row][coord.column] == 1 ||
		p.problemMap[coord.row][coord.column] > p.level {
		return stack
	}

	visitedMap[coord.row][coord.column] = 1
	weightedMap[coord.row][coord.column] = weight
	return append(stack, coord)
}

func main() {
	var result = 0
	_ = result
	p := &Problem{
		size:        0,
		level:       2,
		exper:       0,
		minValue:    400,
		minLocation: 621,
	}
	p.readInput()

	p.problemMap[p.location.row][p.location.column] = 0
	eatFish(p)

	for p.minValue != 400 {
		p.location = &Point{p.minLocation / 30, p.minLocation % 30}
		p.problemMap[p.location.row][p.location.column] = 0
		result += p.minValue

		p.exper = p.exper + 1
		if p.exper == p.level {
			p.level = p.level + 1
			p.exper = 0
		}

		p.minLocation = 621
		p.minValue = 400
		eatFish(p)
	}

	fmt.Println(result)
}

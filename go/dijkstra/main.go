package main

// import "fmt"

// type Point struct {
// 	row    int
// 	column int
// }

// type Problem struct {
// 	size        int
// 	problemMap  [][]int
// 	level       int
// 	exper       int
// 	location    *Point
// 	minValue    int
// 	minLocation int
// 	result      int
// }

// func makeList(len, initValue int) []int {
// 	list := make([]int, len)

// 	for i := 0; i < len; i++ {
// 		list[i] = initValue
// 	}

// 	return list
// }

// func make2DList(len, intValue int) [][]int {
// 	var list [][]int

// 	for i := 0; i < len; i++ {
// 		list = append(list, makeList(len, intValue))
// 	}

// 	return list
// }

// func (self *Problem) readInput() {
// 	fmt.Scanf("%d", &self.size)

// 	self.problemMap = append(self.problemMap, makeList(self.size+2, 100))
// 	for i := 0; i < self.size; i++ {
// 		line := makeList(self.size+2, 100)
// 		for j := 0; j < self.size; j++ {
// 			fmt.Scanf("%d", &line[j+1])
// 			if line[j+1] == 9 {
// 				self.location = &Point{i + 1, j + 1}
// 			}
// 		}
// 		self.problemMap = append(self.problemMap, line)
// 	}
// 	self.problemMap = append(self.problemMap, makeList(self.size+2, 100))
// }

// func print2DArray(array [][]int, len int) {
// 	for i := 0; i < len; i++ {
// 		fmt.Println(array[i])
// 	}
// }

// func eatFish(p *Problem) {
// 	// stack
// 	var stack []*Point
// 	_ = stack
// 	weightedMap := make2DList(p.size+2, 10)
// 	visitedMap := make2DList(p.size+2, 0)

// 	weightedMap[p.location.row][p.location.column] = 0
// 	visitedMap[p.location.row][p.location.column] = 1
// 	stack = append(stack, p.location)

// 	for len(stack) != 0 {
// 		coord := stack[0]
// 		stack = stack[1:]

// 		weight := weightedMap[coord.row][coord.column]

// 		if p.problemMap[coord.row][coord.column] != 0 && p.problemMap[coord.row][coord.column] < p.level {
// 			if weight < p.minValue || weight == p.minValue && (coord.row*30+coord.column) < p.minLocation {
// 				p.minValue = weight
// 				p.minLocation = coord.row*30 + coord.column
// 			}
// 		}

// 		stack = pushNode(stack, p, &Point{coord.row + 1, coord.column}, weightedMap, visitedMap, weight+1)
// 		stack = pushNode(stack, p, &Point{coord.row - 1, coord.column}, weightedMap, visitedMap, weight+1)
// 		stack = pushNode(stack, p, &Point{coord.row, coord.column + 1}, weightedMap, visitedMap, weight+1)
// 		stack = pushNode(stack, p, &Point{coord.row, coord.column - 1}, weightedMap, visitedMap, weight+1)
// 	}
// }

// func pushNode(stack []*Point, p *Problem, coord *Point, weightedMap, visitedMap [][]int, weight int) []*Point {
// 	if visitedMap[coord.row][coord.column] == 1 ||
// 		p.problemMap[coord.row][coord.column] > p.level {
// 		return stack
// 	}

// 	visitedMap[coord.row][coord.column] = 1
// 	weightedMap[coord.row][coord.column] = weight
// 	return append(stack, coord)
// }

// func main() {
// 	var result = 0
// 	p := &Problem{
// 		size:        0,
// 		level:       2,
// 		exper:       0,
// 		minValue:    400,
// 		minLocation: 621,
// 	}
// 	p.readInput()

// 	// print2DArray(p.problemMap, p.size+2)
// 	p.problemMap[p.location.row][p.location.column] = 0
// 	eatFish(p)

// 	for p.minValue != 400 {
// 		p.location = &Point{p.minLocation / 30, p.minLocation % 30}
// 		p.problemMap[p.location.row][p.location.column] = 0
// 		result += p.minValue

// 		p.exper = p.exper + 1
// 		if p.exper == p.level {
// 			p.level = p.level + 1
// 			p.exper = 0
// 		}

// 		p.minLocation = 621
// 		p.minValue = 400
// 		eatFish(p)
// 	}

// 	fmt.Println(result)
// }

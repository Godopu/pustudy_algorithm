class Problem:
    size = 0
    map = []
    level = 2
    exper = 0
    location = (0, 0)
    minValue = 400
    minLocation = 621
    result = 0

    def readInput(self):
        self.size = int(input())

        self.map.append([100]*(self.size + 2))
        for i in range(self.size):
            line = [int(item) for item in input().split(' ')]
            self.map.append([100] + line + [100])
            for j in range(self.size):
                if line[j] == 9:
                    self.location = (i + 1, j + 1)
        self.map.append([100]*(self.size + 2))


def eatFish(p):
    stack = []
    weightedMap = [[10]*(p.size + 2) for i in range(p.size+2)]
    visitedMap = [[False]*(p.size + 2) for i in range(p.size+2)]

    weightedMap[p.location[0]][p.location[1]] = 0
    visitedMap[p.location[0]][p.location[1]] = True
    stack.append(p.location)

    while len(stack) != 0:
        coord = stack.pop(0)
        weight = weightedMap[coord[0]][coord[1]]

        if p.map[coord[0]][coord[1]] != 0 and p.map[coord[0]][coord[1]] < p.level:
            if (weight < p.minValue) or (weight == p.minValue and (coord[0]*30 + coord[1]) < p.minLocation):
                p.minValue = weight
                p.minLocation = coord[0]*30 + coord[1]

        pushNode(stack, p, (coord[0]+1, coord[1]),
                 weightedMap, visitedMap, weight + 1)
        pushNode(stack, p, (coord[0]-1, coord[1]),
                 weightedMap, visitedMap, weight + 1)
        pushNode(stack, p, (coord[0], coord[1]+1),
                 weightedMap, visitedMap, weight + 1)
        pushNode(stack, p, (coord[0], coord[1]-1),
                 weightedMap, visitedMap, weight + 1)

    return


def pushNode(stack, p, coord, weightedMap, visitedMap, weight):
    if visitedMap[coord[0]][coord[1]] == True:
        return

    if p.map[coord[0]][coord[1]] > p.level:
        return

    visitedMap[coord[0]][coord[1]] = True
    weightedMap[coord[0]][coord[1]] = weight
    stack.append(coord)


result = 0
p = Problem()
p.readInput()
# p.printMap()
p.map[p.location[0]][p.location[1]] = 0
eatFish(p)


while p.minValue != 400:
    p.location = (int(p.minLocation / 30), p.minLocation % 30)
    p.map[p.location[0]][p.location[1]] = 0
    result += p.minValue
    # print(p.minValue)
    p.exper = p.exper + 1
    if p.exper == p.level:
        p.level = p.level + 1
        p.exper = 0

    p.minLocation = 621
    p.minValue = 400
    eatFish(p)

# p.printMap()
print(result)

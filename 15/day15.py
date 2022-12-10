import sys
import heapq
from collections import defaultdict
from math import inf, floor

import aocd



if len(sys.argv) > 2 and sys.argv[2] == "submit":
    SUBMIT = True
else:
    SUBMIT = False


def submit(val, part, day, year):
    if SUBMIT:
        aocd.submit(val, part=part, day=day, year=year)
    else:
        print(f"Not submiting {val} for 12/{day}/{year} part {part}")


with open(sys.argv[1]) as f:
    data = f.readlines()


cavern = [[int (i) for i in l.strip()] for l in data]

def get_adjacent(cur, imax, jmax):
    adjacent = []
    i = cur[0]
    j = cur[1]
    for di, dj in [[1,0], [-1,0], [0,1], [0,-1]]:
        newi = i + di
        newj = j + dj
        if 0 <= newi <= imax and 0 <= newj <= jmax: 
            adjacent.append((newi, newj))
    return adjacent



def solve1(cavern):
    squares_to_visit = []

    imax = len(cavern) -1
    jmax = len(cavern[0])-1

    goal = (imax, jmax)
    start = (0,0)

    square_risk_values = defaultdict(lambda: inf)
    square_risk_values[start] = 0 

    #keep track of the nodes we NEED to visit
    unvisited_squares = set()
    for i in range(len(cavern)):
        for j in range(len(cavern[0])):
            unvisited_squares.add((i,j))

    heapq.heappush(squares_to_visit, (0,start))

    steps = 0
    while len(unvisited_squares) > 0 and len(squares_to_visit) > 0:
        visit = heapq.heappop(squares_to_visit)
        risk, cur = visit
        steps +=1 
        adjacents = get_adjacent(cur, imax, jmax)

        if cur not in unvisited_squares:
            continue

        for square in adjacents:
            if square in unvisited_squares:
                #lower the risk
                square_risk = min(
                    square_risk_values[square],
                    square_risk_values[cur] + cavern[square[0]][square[1]]
                )
                square_risk_values[square] = square_risk
                heapq.heappush(squares_to_visit, (square_risk, square))

        if cur in unvisited_squares:
            unvisited_squares.remove(cur)
    return square_risk_values[goal]
    

print(solve1(cavern))


def get_val(x,y, max_x, max_y):
    xsub = floor(x/max_x)
    ysub= floor(y/max_y)
    risk = cavern[x%max_x][y%max_y]
    new_risk = (risk + xsub + ysub -1) %9 + 1
    return new_risk

fulli = len(cavern) * 5
fullj = len(cavern[0]) *5
basei = len(cavern)
basej = len(cavern[0])


bigcavern = [[] for i in range(fulli)]
for i in range(fulli):
    for j in range(fullj):
        bigcavern[i].append(get_val(i,j, basei, basej))

print(bigcavern)
print(solve1(bigcavern))
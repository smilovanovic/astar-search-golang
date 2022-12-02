package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"
)

type GameStateNode struct {
	n      []string
	g      int
	parent *GameStateNode
	key    string
	f      int
}

func runningtime() time.Time {
	log.Println("Start")
	return time.Now()
}

func track(startTime time.Time) {
	endTime := time.Now()
	log.Println("End, took", endTime.Sub(startTime))
}

func getSolution() []string {
	return strings.Split("rrrrggggbbbb", "")
}

func getKey(n *[]string) string {
	return strings.Join(*n, "")
}

func SliceIndex(n *[]string, el string, start int) int {
	for i := start; i < len(*n); i++ {
		if (*n)[i] == el {
			return i
		}
	}
	return -1
}

func getH(n *[]string) int {
	result := 0
	for i, c := range getSolution() {
		if (*n)[i] == c {
			result -= int(math.Pow(10, float64(i+1)))
			continue
		}
		result += SliceIndex(n, c, i)
		return result
	}
	return result
}

func getF(n *[]string, g int) int {
	return g + getH(n)
}

func getChildren(node *GameStateNode) []GameStateNode {
	moves := []GameStateNode{}
	for i := 0; i < len((*node).n)-1; i++ {
		if (*node).n[i] != (*node).n[i+1] {
			n := make([]string, 0, len((*node).n))
			n = append(n, (*node).n[0:i]...)
			n = append(n, (*node).n[i+1], (*node).n[i])
			n = append(n, (*node).n[i+2:]...)
			g := node.g + 1
			moves = append(moves, GameStateNode{n: n, g: g, parent: node, key: getKey(&n), f: getF(&n, g)})
		}
	}
	return moves
}

type Result struct {
	moves        []string
	closedStates int
}

func astar(n *[]string) Result {
	solution := getSolution()
	solutionKey := getKey(&solution)
	start := GameStateNode{n: *n, g: 0, key: getKey(n), f: getH(n)}
	openList := map[string]GameStateNode{}
	closedList := map[string]bool{}
	openList[start.key] = start

	for len(openList) > 0 {
		var currentNode GameStateNode
		for _, node := range openList {
			if currentNode.key == "" || node.f < currentNode.f {
				currentNode = node
			}
		}

		if currentNode.key == solutionKey {
			curr := currentNode
			ret := []string{}
			for curr.parent != nil {
				ret = append(ret, curr.key)
				curr = *curr.parent
			}
			ret = append(ret, curr.key)
			return Result{
				moves:        ret,
				closedStates: len(closedList),
			}
		}

		delete(openList, currentNode.key)
		closedList[currentNode.key] = true

		children := getChildren(&currentNode)
		for _, child := range children {
			_, ok := closedList[child.key]
			if ok {
				continue
			}

			existingNode, ok := openList[child.key]
			if !ok {
				openList[child.key] = child
			} else if existingNode.g > child.g {
				existingNode.g = child.g
				existingNode.parent = child.parent
				existingNode.f = getF(&(child.n), child.g)
			}
		}
	}

	return Result{
		moves:        []string{},
		closedStates: 0,
	}
}

func main() {
	defer track(runningtime())
	niz := strings.Split("gbrgbbrggbrr", "")
	fmt.Println(astar(&niz))
}

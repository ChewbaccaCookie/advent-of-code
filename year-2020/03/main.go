package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func loadInput() (treeMap []string){
	f, err := os.Open("input_3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		treeMap = append(treeMap, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return treeMap
}

func countTrees(treeMap []string, rightStep, downStep int) (trees int) {
	right := 0
	top := 0
	for top < len(treeMap) {
		if treeMap[top][right % 31] == '#' {
			trees++
		}
		right += rightStep
		top += downStep
	}
	return trees
}


func main() {
	treeMap := loadInput()
	countA := countTrees(treeMap, 3,1)
	countB := countTrees(treeMap, 1,1) * countTrees(treeMap, 3,1) * countTrees(treeMap, 5,1)* countTrees(treeMap, 7,1) *countTrees(treeMap, 1,2)
	fmt.Println(countA)

	fmt.Println(countB)
}

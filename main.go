package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const EXPERT = "Customer Engineer"

func main() {
	salesStage := getSalesStage()
	if EXPERT == "Customer Engineer" && salesStage <= 4 || salesStage >= 0 {
		invokeCE(salesStage)
	}
}

func invokeCE(stage int32) {
	switch {
	case stage > 0 && stage <= 4:
		fmt.Printf("A %s will help you with that. go/new-er\n", EXPERT)
	case stage < 0 || stage >= 5:
		fmt.Println("Not  valid Sales Stage")
		fallthrough
	default:
		fmt.Println("Talk to the customer and determine their needs!")
	}
}

func getSalesStage() int32 {
	fmt.Print("Enter your sales stage (1-4), or any character if unknown: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	stage, _ := strconv.Atoi(strings.TrimSuffix(input, "\n"))
	return int32(stage)
}

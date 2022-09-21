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
	var salesStage int32
	fmt.Print("Enter your sales stage: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	stage, _ := strconv.Atoi(strings.TrimSuffix(input, "\n"))
	salesStage = int32(stage)
	if EXPERT == "Customer Engineer" && salesStage < 4 && salesStage > 1 {
		invokeCE(salesStage)
	}
}

func invokeCE(stage int32) {
	switch stage {
	case 1, 2, 3, 4:
		fmt.Println("go/new-er")
	default:
		fmt.Println("Call an FSR")
	}
}

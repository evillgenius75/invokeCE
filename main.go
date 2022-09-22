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
	if EXPERT == "Customer Engineer" && salesStage <= 4 && salesStage >= 0 {
		invokeCE(salesStage)
	} else {
		fmt.Printf("%d is not a valid sales stage\n", salesStage)
	}
}

func invokeCE(stage int32) {
	switch stage {
	case 1, 2, 3, 4:
		fmt.Printf("A %s will help you there. go/new-er\n", EXPERT)
	default:
		fmt.Println("Call an FSR")
	}
}

func getSalesStage() int32 {
	fmt.Print("Enter your sales stage: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	stage, _ := strconv.Atoi(strings.TrimSuffix(input, "\n"))
	return int32(stage)
}

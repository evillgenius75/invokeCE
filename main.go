package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const expert = "Customer Engineer"

func main() {
	salesStage := getSalesStage()
	msg, err := invokeCE(salesStage)
	if err != nil {
		fmt.Printf("ERROR: %s, engage a %s and let's get this done!", err, expert)
	}
	fmt.Println(msg)
}

func invokeCE(stage int32) (string, error) {
	switch stage {
	case 0:
		return "", errors.New("need help qualifying")
	case 1, 2, 3, 4:
		return fmt.Sprintf("A %s will help you with that. go/new-er", expert), nil
	default:
		return "", fmt.Errorf("%d is not a valid Sales Stage", stage)
	}
}

func getSalesStage() int32 {
	fmt.Print("Enter your sales stage (1-4), or any character if unknown: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	stage, _ := strconv.Atoi(strings.TrimSuffix(input, "\n"))
	return int32(stage)
}

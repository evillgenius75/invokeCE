package main

import (
	"errors"
	"flag"
	"fmt"
)

const expert = "Customer Engineer"

var salesStage int

func main() {
	flag.IntVar(&salesStage, "sales-stage", 0, "Current Sales stage you are in (0-4)")
	flag.Parse()
	msg, err := invokeCE(salesStage)
	if err != nil {
		fmt.Printf("ERROR: %s, engage a %s and let's get this done!", err, expert)
	}
	fmt.Println(msg)
}

func invokeCE(stage int) (string, error) {
	switch stage {
	case 0:
		return "", errors.New("need help qualifying")
	case 1, 2, 3, 4:
		return fmt.Sprintf("A %s will help you with that. go/new-er", expert), nil
	default:
		return "", fmt.Errorf("%d is not a valid Sales Stage", stage)
	}
}

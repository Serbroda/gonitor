package main

import (
	"fmt"
	"gonitor/common"
)

func main() {
	args := common.GetArgs()
	fmt.Printf("Args: %v\n", args)
}

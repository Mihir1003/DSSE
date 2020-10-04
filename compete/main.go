package main

import (
	"fmt"
	"github.com/mihir1003/DSSE/api/data"
)

func compete(ops *data.Operations) {
	for _, element := range *ops {
		switch element.Name {
		case "Add":
			fmt.Println(element.Name)

		}

	}
}

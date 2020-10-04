package compete

import (
	"api/data"
	"fmt"
)

func compete(ops *data.Operations) {
	for _, element := range *ops {
		switch element.Name {
		case "Add":
			fmt.Println(element.Name)

		}

	}
}

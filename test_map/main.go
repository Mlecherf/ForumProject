package main

import (
	"fmt"
)

func main() {
	theme := make(map[string]int)
	theme["FastFood"] = 5
	theme["Pizza"] = 5
	theme["Burger"] = 3
	theme["Dessert"] = 4
	theme["American"] = 4
	theme["Italia"] = 0
	theme["Mexican"] = 0
	theme["Indian"] = 0
	theme["Japan"] = 0
	theme["French"] = 0
	theme["African"] = 0
	theme["BBQ"] = 0
	theme["Korea"] = 0
	theme["Vegan"] = 0
	fmt.Println(theme)
	pre_order := [5]int{}
	order_theme := [5]string{}
	fmt.Println(order_theme)

	for i := 0; i < 5; i++ {
		last_add := ""
		for name, v := range theme {
			if i == 0 {
				if v >= pre_order[i] {
					pre_order[i] = v
					order_theme[i] = name
					last_add = name
				}
			} else {
				if v <= pre_order[i-1] && v >= pre_order[i] {
					pre_order[i] = v
					order_theme[i] = name
					last_add = name
				}
			}

		}
		fmt.Println(last_add)
		theme[last_add] = -1

	}
	fmt.Println(order_theme, "--", pre_order)

}

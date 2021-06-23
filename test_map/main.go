package main

import "fmt"

func main() {
	theme := make(map[string]int)
	theme["FastFood"] = 5
	theme["Pizza"] = 4
	theme["Burger"] = 44
	theme["Dessert"] = 32
	theme["American"] = 8
	theme["Italia"] = 9
	theme["Mexican"] = 55
	theme["Indian"] = 2
	theme["Japan"] = 5
	theme["French"] = 8
	theme["African"] = 26
	theme["BBQ"] = 1
	theme["Korea"] = 2
	theme["Vegan"] = 0
	fmt.Println(theme)
	pre_order := [5]int{}
	order_theme := [5]string{}
	fmt.Println(order_theme)

	for i := 0; i < 5; i++ {
		// fmt.Println(i)
		for name, v := range theme {
			fmt.Println(name, pre_order[i], v, i)
			if i == 0 {
				if v >= pre_order[i] {
					if order_theme[i] != name {
						pre_order[i] = v
						order_theme[i] = name
					}
				}
			} else {
				if v < pre_order[i-1] && v >= pre_order[i] {
					if order_theme[i] != name {
						pre_order[i] = v
						order_theme[i] = name
					}
				}
			}
		}
	}
	fmt.Println(order_theme, "--", pre_order)

}

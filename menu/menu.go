package menu

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type menuItem struct {
	name  string
	price map[string]float64
}

var in = bufio.NewReader(os.Stdin)

func MainPrompt() {
	fmt.Println("Please make a selection:")
	fmt.Println("[1] Print menu")
	fmt.Println("[2] Add an item")
	fmt.Println("[3] Remove an item")
	fmt.Println("[Q] Quit the program")
	input := readInput()

	chooseOption(input)
}

func readInput() string {
	input, err := in.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)

	return input
}

func chooseOption(input string) {
	switch input {
	case "1":
		for _, item := range Menu {
			fmt.Println(item.name)
			fmt.Println(strings.Repeat("-", 22))
			for size, price := range item.price {
				fmt.Printf("%10s: %10.2f\n", size, price)
			}
		}
		MainPrompt()
	case "2":
		fmt.Println("Please enter the name of the item you would like to add:")
		input := readInput()
		Menu = append(Menu, menuItem{name: input, price: make(map[string]float64)})
		fmt.Println("Please enter the price of the item for a size small:")
		input = readInput()
		numInput, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("That is not a valid number.")
			MainPrompt()
		} else {
			Menu[len(Menu)-1].price["small"] = numInput
			fmt.Println("Please enter the price of the item for a size medium:")
			input = readInput()
			numInput, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("That is not a valid number.")
				MainPrompt()
			} else {
				Menu[len(Menu)-1].price["medium"] = numInput
				fmt.Println("Please enter the price of the item for a size large:")
				input = readInput()
				numInput, err := strconv.ParseFloat(input, 64)
				if err != nil {
					fmt.Println("That is not a valid number.")
					MainPrompt()
				} else {
					Menu[len(Menu)-1].price["large"] = numInput
				}
			}
		}
		MainPrompt()
	case "3":
		fmt.Println("Please enter the name of the item you would like to remove:")
		input := readInput()

		for i, v := range Menu {
			if v.name == input {
				slices.Delete(Menu, i, i+1)
			}
		}
		MainPrompt()
	case "Q", "q":
		fmt.Println("Exiting the program...")
	}
}

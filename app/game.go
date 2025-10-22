package main

import (
	"bufio"
	"fmt"
	"os"
	"poker-go/app/types"
	"poker-go/app/utils"
	"strings"
)

func pokerGame(pokerGameInput types.PokerGameInput) (string, error) {
	hand1 := pokerGameInput.Hand1
	hand2 := pokerGameInput.Hand2

	return utils.CompareHands(hand1, hand2)
}

func RunGame() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter two poker hands (space separated), type end for ending the game\n")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		if parts[0] == "end" {
			fmt.Println("Game ended.")
			break
		}

		if len(parts) < 2 {
			fmt.Println("You have provided only one, Please enter two inputs.")
			continue
		}

		input1 := parts[0]
		input2 := parts[1]

		fmt.Println("You entered:", input1, "and", input2)

		pokerGameInput := types.PokerGameInput{
			Hand1: input1,
			Hand2: input2,
		}

		result, err := pokerGame(pokerGameInput)

		if err != nil {
			fmt.Printf("There was a problem while playing the poker game %v\n", err)
			break
		}

		fmt.Printf("Your result of the poker game is %s\n", result)

		fmt.Print("Do you still want to play more? (yes/no) ")
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "no" {
			fmt.Print("Have a good day !!!")
			break
		}
	}
}

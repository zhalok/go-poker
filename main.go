package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type POKER_TYPE string

const FOUR_OF_A_KIND POKER_TYPE = "four_of_a_kind"
const FULL_HOUSE POKER_TYPE = "full_house"
const TRIPLE POKER_TYPE = "triple"
const TWO_PAIR POKER_TYPE = "two_pair"
const HIGH_CARD POKER_TYPE = "high_card"

var PokerTypeOrder = [5]POKER_TYPE{FOUR_OF_A_KIND, FULL_HOUSE, TRIPLE, TWO_PAIR, HIGH_CARD}

const POKER_SYMBOL_ORDER = "23456789TJQKA"

type PokerGameInput struct {
	Hand1 string
	Hand2 string
}

func BuildCounterFreqArray(inputString string) []int {
	counterFreqMap := make(map[rune]int, 0)
	counterFreqArr := make([]int, 0)

	for _, ch := range inputString {
		_, ok := counterFreqMap[ch]
		if !ok {
			counterFreqMap[ch] = 0
		}

		counterFreqMap[ch] += 1
	}

	for _, count := range counterFreqMap {
		counterFreqArr = append(counterFreqArr, count)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counterFreqArr)))

	return counterFreqArr
}

func BuildPokerTypeOrderMap() map[POKER_TYPE]int {
	pokerOrderMap := make(map[POKER_TYPE]int, 0)
	totalPokers := len(PokerTypeOrder)

	for idx, value := range PokerTypeOrder {
		pokerOrderMap[value] = (totalPokers - idx)
	}

	return pokerOrderMap
}

func ComparePokers(poker1 POKER_TYPE, poker2 POKER_TYPE) int {
	pokerOrderMap := BuildPokerTypeOrderMap()
	poker1Order := pokerOrderMap[poker1]
	poker2Order := pokerOrderMap[poker2]

	if poker1Order > poker2Order {
		return 1
	} else if poker1Order < poker2Order {
		return 2
	} else {
		return 0
	}
}

func DeterminePokerTypeFromCounterFreq(counterFreq []int) (POKER_TYPE, error) {
	if len(counterFreq) < 2 {
		return "", fmt.Errorf("not enough values in counter freq")
	}

	if counterFreq[0] == 4 {
		return FOUR_OF_A_KIND, nil
	} else if counterFreq[0] == 3 && counterFreq[1] == 2 {
		return FULL_HOUSE, nil
	} else if counterFreq[0] == 3 && counterFreq[1] == 1 {
		return TRIPLE, nil
	} else if counterFreq[0] == 2 && counterFreq[1] == 2 {
		return TWO_PAIR, nil
	} else if counterFreq[0] == 2 && counterFreq[1] == 1 {
		return TWO_PAIR, nil
	} else {
		return HIGH_CARD, nil
	}

}

func ConvertHandToPokerType(hand string) (POKER_TYPE, error) {
	counterFreqArr := BuildCounterFreqArray(hand)
	pokerType, err := DeterminePokerTypeFromCounterFreq(counterFreqArr)
	if err != nil {
		return "", fmt.Errorf("there was a problem determining the poker type %w", err)
	}

	return pokerType, nil
}

func BuilPokerSymbolOrderMap() map[rune]int {
	pokerSymbolOrderMap := make(map[rune]int, 0)
	totalPokerSymbols := len(POKER_SYMBOL_ORDER)
	for idx, ch := range POKER_SYMBOL_ORDER {
		pokerSymbolOrderMap[ch] = (totalPokerSymbols - idx)
	}

	return pokerSymbolOrderMap
}

func BreakTie(hand1 string, hand2 string) string {
	pokerSymbolOrderMap := BuilPokerSymbolOrderMap()
	hand1PokerValues := make([]int, 0)
	hand2PokerValues := make([]int, 0)

	for _, ch := range hand1{
		hand1PokerValues = append(hand1PokerValues, pokerSymbolOrderMap[ch])
	}

	for _, ch := range hand2{
		hand2PokerValues = append(hand2PokerValues, pokerSymbolOrderMap[ch])
	}	
	
	sort.Ints(hand1PokerValues)
	sort.Ints(hand2PokerValues)

	for i:=0;i<5;i++{
		if hand1PokerValues[i] > hand2PokerValues[i]{
			return hand1
		}else if hand1PokerValues[i] < hand2PokerValues[i]{
			return hand2
		}
	}

	return "tie"
}

func CompareHands(hand1 string, hand2 string) (string, error) {
	hand1PokerType, err := ConvertHandToPokerType(hand1)

	if err != nil {
		return "", fmt.Errorf("there was an error while trying to convert hand1 to poker type %w", err)
	}

	hand2PokerType, err := ConvertHandToPokerType(hand2)

	if err != nil {
		return "", fmt.Errorf("there was an error while trying to convert hand2 to poker type %w", err)
	}

	comparisonResult := ComparePokers(hand1PokerType, hand2PokerType)

	if comparisonResult == 1 {
		return hand1, nil
	} else if comparisonResult == 2 {
		return hand2, nil
	}else{
		return BreakTie(hand1,hand2), nil
	}

}

func PokerGame(pokerGameInput PokerGameInput) (string, error) {
	hand1 := pokerGameInput.Hand1
	hand2 := pokerGameInput.Hand2

	return CompareHands(hand1,hand2)
}

func main(){
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

		pokerGameInput := PokerGameInput{
			Hand1: input1,
			Hand2: input2,
		}

		result, err := PokerGame(pokerGameInput)

		if err != nil{
			fmt.Printf("There was a problem while playing the poker game %v\n", err)
			break
		}

		fmt.Printf("Your result of the poker game is %s\n", result)

		fmt.Print("Do you still want to play more? (yes/no) ")
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "no"{
			fmt.Print("Have a good day !!!")
			break
		}
    }
}

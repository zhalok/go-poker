package utils

import (
	"fmt"
	"poker-go/app/types"
	"sort"
)

func comparePokers(poker1 types.POKER_TYPE, poker2 types.POKER_TYPE) int {
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

func determinePokerTypeFromCounterFreq(counterFreq []int) (types.POKER_TYPE, error) {
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

func convertHandToPokerType(hand string) (types.POKER_TYPE, error) {
	counterFreqArr := BuildCounterFreqArray(hand)
	pokerType, err := determinePokerTypeFromCounterFreq(counterFreqArr)
	if err != nil {
		return "", fmt.Errorf("there was a problem determining the poker type %w", err)
	}

	return pokerType, nil
}

func breakTie(hand1 string, hand2 string) string {
	pokerSymbolOrderMap := BuilPokerSymbolOrderMap()
	hand1PokerValues := make([]int, 0)
	hand2PokerValues := make([]int, 0)

	for _, ch := range hand1 {
		hand1PokerValues = append(hand1PokerValues, pokerSymbolOrderMap[ch])
	}

	for _, ch := range hand2 {
		hand2PokerValues = append(hand2PokerValues, pokerSymbolOrderMap[ch])
	}

	sort.Ints(hand1PokerValues)
	sort.Ints(hand2PokerValues)

	for i := 0; i < 5; i++ {
		if hand1PokerValues[i] > hand2PokerValues[i] {
			return hand1
		} else if hand1PokerValues[i] < hand2PokerValues[i] {
			return hand2
		}
	}

	return "tie"
}

func CompareHands(hand1 string, hand2 string) (string, error) {
	hand1PokerType, err := convertHandToPokerType(hand1)

	if err != nil {
		return "", fmt.Errorf("there was an error while trying to convert hand1 to poker type %w", err)
	}

	hand2PokerType, err := convertHandToPokerType(hand2)

	if err != nil {
		return "", fmt.Errorf("there was an error while trying to convert hand2 to poker type %w", err)
	}

	comparisonResult := comparePokers(hand1PokerType, hand2PokerType)

	if comparisonResult == 1 {
		return hand1, nil
	} else if comparisonResult == 2 {
		return hand2, nil
	} else {
		return breakTie(hand1, hand2), nil
	}

}

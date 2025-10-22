package utils

import (
	"poker-go/app/types"
	"sort"
)

func BuildPokerTypeOrderMap() map[types.POKER_TYPE]int {
	pokerOrderMap := make(map[types.POKER_TYPE]int, 0)
	totalPokers := len(PokerTypeOrder)

	for idx, value := range PokerTypeOrder {
		pokerOrderMap[value] = (totalPokers - idx)
	}

	return pokerOrderMap
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

func BuilPokerSymbolOrderMap() map[rune]int {
	pokerSymbolOrderMap := make(map[rune]int, 0)
	totalPokerSymbols := len(POKER_SYMBOL_ORDER)
	for idx, ch := range POKER_SYMBOL_ORDER {
		pokerSymbolOrderMap[ch] = (totalPokerSymbols - idx)
	}

	return pokerSymbolOrderMap
}

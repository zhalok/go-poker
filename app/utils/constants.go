package utils

import "poker-go/app/types"

const FOUR_OF_A_KIND types.POKER_TYPE = "four_of_a_kind"
const FULL_HOUSE types.POKER_TYPE = "full_house"
const TRIPLE types.POKER_TYPE = "triple"
const TWO_PAIR types.POKER_TYPE = "two_pair"
const HIGH_CARD types.POKER_TYPE = "high_card"

var PokerTypeOrder = [5]types.POKER_TYPE{FOUR_OF_A_KIND, FULL_HOUSE, TRIPLE, TWO_PAIR, HIGH_CARD}

const POKER_SYMBOL_ORDER = "23456789TJQKA"

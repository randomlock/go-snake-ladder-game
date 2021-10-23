package models

import "math/rand"

type Dice struct {
    maxNum int
}

func NewDice(maxNum int) Dice {
    return Dice{maxNum: maxNum}
}


func (dice Dice) Roll() int {
    return rand.Intn(6) + 1
}

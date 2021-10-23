package models

type Snake struct {
    From int
    To int
}

func NewSnake(from int, to int) Snake {
    return Snake{From: from, To: to}
}
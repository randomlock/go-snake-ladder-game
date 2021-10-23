package controllers

import (
    "fmt"
    "math/rand"

    "../models"
    "../utils"
)

type Controller struct {
    playerCount int
    snakeCount int
    ladderCount int
    boxes int
    diceMaxNum int
}

func NewController(playerCount int, snakeCount int, ladderCount int, boxes int, diceMaxNum int) *Controller {
    return &Controller{playerCount: playerCount, snakeCount: snakeCount, ladderCount: ladderCount, boxes: boxes, diceMaxNum: diceMaxNum}
}

func (c Controller) CreateGame() (game *models.Game, err error) {

    if c.boxes < 0 {
        return nil, fmt.Errorf("boxes should be greater than 0")
    }
    if c.diceMaxNum < 0 {
        return nil, fmt.Errorf("dice max number should be greater than 0")
    }


    var players []*models.Player
    var snakes []models.Snake
    var ladders []models.Ladder
    var dice models.Dice
    for i := 0; i < c.playerCount; i++ {
        players = append(players, models.NewPlayer(utils.RandStringRunes(10)))
    }
    ladders = c.GenerateLadder()
    snakes = c.GenerateSnake()
    dice = models.NewDice(6)
    return models.NewGame(players, snakes, ladders, dice, c.boxes), nil
}

func (c Controller) GenerateLadder() (ladders []models.Ladder)  {
    for i := 0; i < c.ladderCount; i++ {
        from := rand.Intn(c.boxes)
        to := rand.Intn(c.boxes-from) + from
        ladders = append(ladders, models.NewLadder(from, to))
    }
    return
}

func (c Controller) GenerateSnake() (snakes []models.Snake)  {
    for i := 0; i < c.snakeCount; i++ {
        from := rand.Intn(c.boxes-1)
        to := from - rand.Intn(c.boxes-from)
        snakes = append(snakes, models.NewSnake(from, to))
    }
    return
}

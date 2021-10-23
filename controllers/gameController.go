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
    boxCount    int
    diceMaxNum  int
}

func NewController() *Controller {
    return &Controller{}
}

func (c *Controller) SetPlayerCount(count int) *Controller {
    c.playerCount = count
    return c
}

func (c *Controller) SetSnakeCount(count int) *Controller {
    c.snakeCount = count
    return c
}

func (c *Controller) SetLadderCount(count int) *Controller {
    c.ladderCount = count
    return c
}

func (c *Controller) SetDiceMaxNum(num int) *Controller {
    c.diceMaxNum = num
    return c
}

func (c *Controller) SetBoxCount(count int) *Controller {
    c.boxCount = count
    return c
}

func (c *Controller) CreateGame() (game *models.Game, err error) {

    if c.boxCount < 0 {
        return nil, fmt.Errorf("boxCount should be greater than 0")
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
    return models.NewGame(players, snakes, ladders, dice, c.boxCount), nil
}

func (c Controller) GenerateLadder() (ladders []models.Ladder)  {
    for i := 0; i < c.ladderCount; i++ {
        from := rand.Intn(c.boxCount)
        to := rand.Intn(c.boxCount-from) + from
        ladders = append(ladders, models.NewLadder(from, to))
    }
    return
}

func (c Controller) GenerateSnake() (snakes []models.Snake)  {
    for i := 0; i < c.snakeCount; i++ {
        from := rand.Intn(c.boxCount -1)
        to := from - rand.Intn(c.boxCount-from)
        snakes = append(snakes, models.NewSnake(from, to))
    }
    return
}

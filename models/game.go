package models

import "fmt"

type Game struct {
    players []*Player
    snake map[int]int
    ladder map[int]int
    stats map[Player]int
    dice Dice
    MaxNum int
}

func NewGame(players []*Player, snakes []Snake, ladders []Ladder, dice Dice, maxNum int) *Game {

    snakeMap := map[int]int{}
    ladderMap := map[int]int{}

    for _, snake := range snakes {
        snakeMap[snake.From] = snake.To
    }
    for _, ladder := range ladders {
        ladderMap[ladder.From] = ladder.To
    }

    return &Game{players: players, snake: snakeMap, ladder: ladderMap, dice: dice, MaxNum: maxNum}
}

func (g *Game) GetNextPlayer() *Player  {
    player := g.players[0]
    g.players = append(g.players[1:], player)
    return player
}

func (g Game) RollDice(player *Player)  {
    num := player.GetNum() + g.dice.Roll()

    if num > g.MaxNum {
        return
    }
    if n, exists := g.snake[num]; exists {
        fmt.Printf("player - %s got snake. Moving from %d to %d", player.GetName(), num, n)
        println()
        num = n
    }
    if n, exists := g.ladder[num]; exists {
        fmt.Printf("player - %s got ladder. Moving from %d to %d", player.GetName(), num, n)
        println()
        num = n
    }
    player.SetNum(num)
}

func (g Game) IsFinished(player Player) bool {
    return player.GetNum() ==  g.MaxNum
}

func (g Game) PrintStats() {
    for _, player := range g.players {
        fmt.Printf("player - %s score is - %d", player.playerName, player.GetNum())
        fmt.Println()
    }
}

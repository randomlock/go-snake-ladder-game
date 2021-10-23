package main

import (
    "fmt"

    "../snakeGame/controllers"
)

func main() {
    gameController := controllers.NewController(4, 5, 5, 100, 6)
    game, err := gameController.CreateGame()
    if err != nil {
        println(fmt.Errorf("error in creating game due to - %s", err.Error()))
        return
    }
    for  {
        player := game.GetNextPlayer()
        fmt.Printf("player - %s is rolling dice", player.GetName())
        println()
        game.RollDice(player)
        game.PrintStats()
        if game.IsFinished(*player) {
            fmt.Printf("player - %s won the game", player.GetName())
            println()
            break
        }
    }
}

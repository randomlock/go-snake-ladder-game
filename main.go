package main

import (
    "fmt"

    "../snakeGame/controllers"
)

func main() {
    gameController := controllers.NewController()
    game, err := gameController.
        SetPlayerCount(5).
        SetLadderCount(5).
        SetSnakeCount(5).
        SetBoxCount(100).
        CreateGame()
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

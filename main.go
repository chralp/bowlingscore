package main

import (
	"./bowling"
	"fmt"
)

func main(){

	// Start bowling game
	bowling.Startgame(1)
	
	// Append rolls
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	
	Game := bowling.GetCurrentGame("player1")
	fmt.Println(Game) // 300
}

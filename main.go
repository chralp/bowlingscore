package main

import (
	"./bowling"
)


func main(){

	// Start bowling game
	bowling.Startgame(1)

	// Append rolls
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("9")
	bowling.GetCurrentGame("player1").AddRoll("-")
	bowling.GetCurrentGame("player1").AddRoll("7")
	bowling.GetCurrentGame("player1").AddRoll("/")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("5")
	bowling.GetCurrentGame("player1").AddRoll("-")
	bowling.GetCurrentGame("player1").AddRoll("8")
	bowling.GetCurrentGame("player1").AddRoll("1")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	bowling.GetCurrentGame("player1").AddRoll("x")
	
}

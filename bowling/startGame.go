package bowling

import (
	"strconv"
	"math/rand"
	"time"
)

// Status codes
const (
	statusCompleted = "completed" // Completed
	statusInprogress = "inprogress" // Not completed Frame
	statusLocked = "locked" // Not playable frames yet
	statusStarted = "Started" // Started
	statusNext = "next" // Next frame to play
)

var GAMES []*Game;

// Frame structure
type Frame struct {
	FrameNumber int
	FrameState  string
	FrameScore  int
	Rolls       []string
}


// Game Structure
type Game struct {
	Lane         int
	GameState    string
	Pos          int
	PlayerName   string	
	Score        int
	CurrentFrame int
	Frames       [10]Frame
}

// Create 10 empty frames
func CreateEmptyFrames() (frames [10]Frame) {
	for i := range frames {
		frames[i].FrameNumber = i + 1
		if(i == 0){
			frames[i].FrameState = statusStarted
		} else {
			frames[i].FrameState = statusLocked
		}
		frames[i].FrameScore = 0
	}
	return
}

// return currentGame of player
func GetCurrentGame(player string) *Game{
	var res *Game
	for i := 0; i < len(GAMES); i++ {
		if(GAMES[i].PlayerName == player){
			res = GAMES[i]
		}
	}
	return res;
}


// Randomize number 1-12
func GetLane() int{
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(11-1)
};

// Start new game
func Startgame(playersAmount int){
	var Lane = GetLane();
	for i := 1; i <= playersAmount; i++ {
		s := strconv.Itoa(i)		
		newGame := &Game{
			PlayerName: "player" + s,
			Pos: i,
			Lane: Lane,
			GameState: statusStarted,
			CurrentFrame: 1,
			Frames: CreateEmptyFrames(),
		}
		GAMES = append(GAMES, newGame);
	}
	return
}

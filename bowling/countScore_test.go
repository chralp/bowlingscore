package bowling

import (
    "testing"
)

func TestCountFrameRolls(t *testing.T) {

	// Should return 10
	strike := Frame {
		FrameNumber: 1,
			FrameState: "completed",
			FrameScore: 1337,
			Rolls: []string{"x"},
		}
	if CountFrameRolls(strike) != 10 {
		t.Fail()
	}

	// Should return 10
	spare := Frame {
		FrameNumber: 1,
			FrameState: "completed",
			FrameScore: 1337,
			Rolls: []string{"9","/"},
		}
	if CountFrameRolls(spare) != 10 {
		t.Fail()
	}
	
	// Should return 8
	miss := Frame {
		FrameNumber: 1,
			FrameState: "completed",
			FrameScore: 1337,
			Rolls: []string{"7","1"},
		}
	if CountFrameRolls(miss) != 8 {
		t.Fail()
	}

	// Frame Ten test
	// Should return 8
	missFrameTen := Frame {
		FrameNumber: 10,
			FrameState: "completed",
			FrameScore: 1337,
			Rolls: []string{"7","1"},
		}
	if CountFrameRolls(missFrameTen) != 8 {
		t.Fail()
	}

	// Frame Ten
	
	// Should return 20
	spareFrameTen := Frame {
		FrameNumber: 10,
			FrameState: "completed",
			FrameScore: 1337,
			Rolls: []string{"9","/","x"},
		}
	if CountFrameRolls(spareFrameTen) != 20 {
		t.Fail()
	}

	// Should return 30
	trippleFrameTen := Frame {
		FrameNumber: 10,
			FrameState: "completed",
			FrameScore: 1337,
			Rolls: []string{"x","x","x"},
		}
	if CountFrameRolls(trippleFrameTen) != 30 {
		t.Fail()
	}

	// Should return 29
	doubleAndNineFrameTen := Frame {
		FrameNumber: 10,
			FrameState: "completed",
			FrameScore: 1337,
			Rolls: []string{"x","x","9"},
		}
	if CountFrameRolls(doubleAndNineFrameTen) != 29 {
		t.Fail()
	}
}


func TestCountScore(t *testing.T) {

	// Use AddRolls function to and check data if countScore calculates correct

	// As reference for calculations https://www.bowlinggenius.com/

	
	// Clear games
	GAMES = nil

	// Start new game
	Startgame(1)
	cg := GetCurrentGame("player1")

	// PerfectGame ( 12 rolls ) 
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")

	// Verify that Game.Score is 300
	if(cg.Score != 300) {
		t.Fail()
	}

	// Verify that each frame have score 30 ( First frame ) + 30 ( per each frame )
	var baseScore int = 30;
	
	for i := 0; i < len(cg.Frames); i++ {
		if(cg.Frames[i].FrameScore != baseScore) {
			t.Fail()
		} else {
			baseScore += 30;
		}
	}

	// Clear games
	GAMES = nil

	// Start new game
	Startgame(1)

	cg = GetCurrentGame("player1")

	// Game Total Should be 136
	GetCurrentGame("player1").AddRoll("9")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("6")
	GetCurrentGame("player1").AddRoll("2")
	GetCurrentGame("player1").AddRoll("7")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("1")
	GetCurrentGame("player1").AddRoll("8")
	GetCurrentGame("player1").AddRoll("5")
	GetCurrentGame("player1").AddRoll("2")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("5")
	GetCurrentGame("player1").AddRoll("3")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("9")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("7")

	// Verify that Game.Score is 136
	if(cg.Score != 136) {
		t.Fail()
	}

	// Verify that each frame have a FrameScore
	for i := 0; i < len(cg.Frames); i++ {
		if(cg.Frames[i].FrameScore == 0) {
			t.Fail()
		}
	}

	// Clear games
	GAMES = nil

	// Start new game
	Startgame(1)

	cg = GetCurrentGame("player1")

	// Game Total Should be 189
	GetCurrentGame("player1").AddRoll("9")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("9")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("6")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("8")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("6")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("9")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("9")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("9")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("x")
	
	// Verify that Game.Score is 189
	if(cg.Score != 189) {
		t.Fail()
	}

	// Verify that each frame have a FrameScore
	for i := 0; i < len(cg.Frames); i++ {
		if(cg.Frames[i].FrameScore == 0) {
			t.Fail()
		}
	}


	GAMES = nil

	// Start new game
	Startgame(1)

	cg = GetCurrentGame("player1")

	GetCurrentGame("player1").AddRoll("9")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("6")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("9")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("4")
	GetCurrentGame("player1").AddRoll("-")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("x")
	GetCurrentGame("player1").AddRoll("6")
	GetCurrentGame("player1").AddRoll("/")
	GetCurrentGame("player1").AddRoll("5")

	// Verify that Game.Score is 194
	if(cg.Score != 194) {
		t.Fail()
	}

	// Verify that each frame have a FrameScore
	for i := 0; i < len(cg.Frames); i++ {
		if(cg.Frames[i].FrameScore == 0) {
			t.Fail()
		}
	}

}

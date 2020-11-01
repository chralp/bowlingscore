package bowling

// Count framePoints according to amount of rolls
func CountFrameRolls(data Frame) int {
	
	if (len(data.Rolls) == 1) {
		
		// Possible frameRoll = [( - or 1-9 or x)]
		return ValueGet(data.Rolls[0]);
	}
	
	if (len(data.Rolls) == 2) {
		if (data.Rolls[1] == "/") {
			
			// Possible frameRolls = [(- or 1-9), /]
			return ValueGet(data.Rolls[1]);
		} else {

			// Possible frameRolls = [(- or 1-9), (- or 1-9)]
			return ValueGet(data.Rolls[0]) + ValueGet(data.Rolls[1])
		}
	}

	// Only for frame ten
	if (len(data.Rolls) == 3) {
		if (data.Rolls[1] == "/") {
			
			// Only Possible frameRolls = [ 9, /, (x or - or 1-9)]
			// We only want the secnd roll and the last
			return ValueGet(data.Rolls[1]) + ValueGet(data.Rolls[2])
		} else {
			
			// Only possible frameRolls = [ x, x, x]
			return ValueGet(data.Rolls[0]) + ValueGet(data.Rolls[1]) + ValueGet(data.Rolls[2])
		}
	}
	return 0;
}

// Checks if its several strikes in a row
func (cg* Game) isStrikeRow(prevFrameID int) bool {
	if(cg.Frames[prevFrameID].FrameScore == 0) {
		if(len(cg.Frames[prevFrameID].Rolls) == 1){
			return true
		}
	}
	return false;
}

// Calculate frameScores depending on input in current frame
// i = indexId ( frameNumber -1 )
func (cg* Game) countScore(i int) {

	// First frame we only need to calculate Framescore if secnd rolls is a miss
	if(cg.Frames[i].FrameNumber == 1) {	
		if (len(cg.Frames[i].Rolls) == 2 && cg.Frames[i].Rolls[1] != "/") {
			cg.calculateScore(i,0)
		}
	} else {

		// FrameNumber 2-10

		// Check if currentFrame have one roll
		if(len(cg.Frames[i].Rolls) == 1) {

			// We only want to calculate score if prev frameScore is 0
			// Otherwise we do this at sencd roll
			if (cg.Frames[i - 1].FrameScore == 0) {
				
				switch (cg.Frames[i].Rolls[0]) {
					
				case "x" :
					
					// Check if prev roll is a strike ( Doubble ) 
					if (cg.isStrikeRow(i - 1)) {
			
						// We want to check if its a tripple, but we can only
						// do it if we have played atleast 3 frames
						if (cg.Frames[i].FrameNumber >= 3) {
							if (cg.isStrikeRow(i - 2)) {

								// Its a tripple, so we only want to update frameScore i - 2.
								cg.calculateScore(i,2)
								break;
							}
						}
					} else {
						
						// Prev frame can only be [- or 1-9, /]

						// Update prevFrameScore  i - 1
						cg.calculateScore(i,1)
					}
					break;
					
				default:
					
					// Can only be [ - or 1-9 ]
					// [ x  , Rolls[0] ]
					// [ 9/ , Rolls[0] ]

					// Check if frameNumber is three or higher.
					// if not we cannot check 2 prev frames

					if (cg.Frames[i].FrameNumber >= 3) {

						// Check if last two frames is scoreLess
						if(cg.Frames[i - 2].FrameScore == 0) {

							// Scoreless = [x, x, i]
							// two strikes before
							cg.calculateScore(i,2)
						} else {

							// currentFrame - 2 have FrameScore
							// [ 5,1 = 6, x  , i]
							// [ 5,1 = 6, 9/ , i]
							
							// We only want to set frameScore - 1 if prev roll was a spare
							if(len(cg.Frames[i - 1].Rolls) != 1) {
								cg.calculateScore(i,1)
							}
						}						
					} else {

						// Only update prev frameScore - 1 if prevRoll is a spare
						if(len(cg.Frames[i - 1].Rolls) != 1) {
							cg.calculateScore(i,1)
						}
					}
					break;
				}
			}
		} else {

			// SecndRoll or third ( per frame )
			
			
			// Special case for Frame 10
			// We only want to calc frameScore - 1 if sencd roll is a spare
			// Otherwise update currentFrame.
			// Its a strike row, thats gets handled in calculateScore.
			if(cg.Frames[i].FrameNumber == 10) {
				if(cg.Frames[i].Rolls[1] != "/") {
					cg.calculateScore(i,1)
				}else{
					cg.calculateScore(i,0)
				}
			} else {

				// we want diffrent calculations depending on prev framescore is 0 or !=0
				if (cg.Frames[i - 1].FrameScore == 0) {

					if(cg.Frames[i].Rolls[1] == "/") {
						cg.calculateScore(i,1)
					} else {

						if (cg.Frames[i].FrameNumber >= 3) {
							if  (cg.Frames[i - 2].FrameScore != 0) {
								cg.calculateScore(i,1)
								cg.calculateScore(i,0)
							}
						} else {
								cg.calculateScore(i,1)
								cg.calculateScore(i,0)
						}
					}
				} else {
					
					// Prev frameScore exists.
					// [1,1 = 2 , Rolls[1]]
					// Only calculate current Framescore if Rolls[1] != "/"
					if(cg.Frames[i].Rolls[1] != "/") {
						cg.calculateScore(i,0)
					}					
				}	
			}
		}
	
	}
	return	
}

// Helper function to update FrameScore
func (cg* Game) calculateScoreSwitch (indexId int, prev int) {
	switch prev {
	case 0:
		cg.Frames[indexId].FrameScore = CountFrameRolls(cg.Frames[indexId]) + cg.Score;
		cg.Score = cg.Frames[indexId].FrameScore
		break;
	case 1:
		cg.Frames[indexId - 1].FrameScore =  CountFrameRolls(cg.Frames[indexId - 1])  + CountFrameRolls(cg.Frames[indexId]) + cg.Score;
		cg.Score = cg.Frames[indexId - 1].FrameScore
		break;
	case 2:
		cg.Frames[indexId - 2].FrameScore =  CountFrameRolls(cg.Frames[indexId - 2])  + CountFrameRolls(cg.Frames[indexId - 1])  + CountFrameRolls(cg.Frames[indexId]) + cg.Score;
		cg.Score = cg.Frames[indexId - 2].FrameScore
		break;
	}
}

func (cg* Game) calculateScore (indexId int, prev int) {

	// Make life easier
	frameNumber := indexId + 1;

	// Special case i frameNumber is 10
	if (frameNumber == 10) {
		if(cg.Frames[indexId - 1].FrameScore == 0) {
			if(cg.Frames[indexId - 2].FrameScore == 0) {
				cg.calculateScoreSwitch(indexId, prev)
			} else {
				cg.calculateScoreSwitch(indexId, 1)
				if(len(cg.Frames[indexId].Rolls) == 2) {
					if(cg.Frames[indexId].Rolls[1] != "/" && cg.Frames[indexId].Rolls[1] != "x"){
						cg.calculateScoreSwitch(indexId, 0)
					}
				}
			}
		} else {
			if(len(cg.Frames[indexId].Rolls) > 1) {
				switch(cg.Frames[indexId].Rolls[1]) {
				case "x":
					if(len(cg.Frames[indexId].Rolls) > 2) {
						cg.calculateScoreSwitch(indexId, prev -1)
					}
					break;
				case "/":
					if(len(cg.Frames[indexId].Rolls) > 2) {
						cg.calculateScoreSwitch(indexId, prev)
					}
					break;
				default:
					if(len(cg.Frames[indexId].Rolls) > 1) {
						cg.calculateScoreSwitch(indexId, 0)
					}
					break;
				}
			}
		}
	} else {
		cg.calculateScoreSwitch(indexId, prev)
	}
	return
}

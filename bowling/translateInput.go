package bowling

import 	"strconv"

// Translate input to int value
func ValueGet (input string) int {
	var value int;
	switch (input){
		case "1","2","3","4","5","6","7","8","9":
		s, _ := strconv.Atoi(input)
		value = s;
		break;

		case "x", "/":
		value = 10;
		break;
		
		case "-":
		value = 0;
		break;

		default:
		value = 0;
		break;
	}
	return value;
}

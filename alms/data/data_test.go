package data

import (
	"fmt"
	"testing"
)

func Test_data(t *testing.T) {
	gameLevels := make([]*GameLevel, 0)
	gl1 := &GameLevel{
		GameID:   1,
		LevelID:  2,
		LowSorce: 11111111,
		IsOpen:   0,
	}
	gl2 := &GameLevel{
		GameID:   2,
		LevelID:  3,
		LowSorce: 2355643,
		IsOpen:   1,
	}
	gl3 := &GameLevel{
		GameID:   4,
		LevelID:  5,
		LowSorce: 324,
		IsOpen:   0,
	}
	gameLevels = append(gameLevels, gl1, gl2, gl3)
	str, err := GameLevelsToJSON(gameLevels)
	fmt.Println(str)
	fmt.Println(err)

	newg, err2 := JSONToGameLevels(str)
	fmt.Println(err2)
	for _, n := range newg {
		fmt.Println("-----------------")
		fmt.Println(n.GameID)
		fmt.Println(n.LevelID)
		fmt.Println(n.LowSorce)
		fmt.Println(n.IsOpen)
	}
}

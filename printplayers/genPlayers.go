package generator

import (
	"fmt"
	"os"
)

// Player represents a football player.
type Player struct {
	Name string
}

// Pair represents a pairing of a striker and a defender.
type Pair struct {
	Striker  *Player
	Defender *Player
}

// WritePairsToFile writes the pairs to a .txt file.
func WritePairsToFile(pairs []*Pair, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, pair := range pairs {
		_, err := fmt.Fprintf(file, "%s - %s\n", pair.Striker.Name, pair.Defender.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

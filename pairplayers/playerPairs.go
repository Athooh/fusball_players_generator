package pairs

import (
	"fmt"
	"math/rand"
	"time"
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

// PairPlayers pairs strikers with defenders.
func PairPlayers(strikers, defenders []*Player) ([]*Pair, error) {
	if len(strikers) != len(defenders) {
		return nil, fmt.Errorf("number of strikers and defenders must be equal")
	}

	pairs := make([]*Pair, 0)
	usedStrikers := make(map[int]bool)
	usedDefenders := make(map[int]bool)

	for len(pairs) < len(strikers) {
		rand.Seed(time.Now().UnixNano())
		sIndex := rand.Intn(len(strikers))
		dIndex := rand.Intn(len(defenders))

		if !usedStrikers[sIndex] && !usedDefenders[dIndex] {
			pairs = append(pairs, &Pair{
				Striker:  strikers[sIndex],
				Defender: defenders[dIndex],
			})
			usedStrikers[sIndex] = true
			usedDefenders[dIndex] = true
		}
	}

	return pairs, nil
}

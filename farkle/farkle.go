package farkle

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

const (
	FOUROFAKIND = 1000
	FIVEOFAKIND = 2000
	SIXOFAKIND  = 3000
	STRAIGHT    = 1500
)

type Die struct {
	Value int
}

type Dice []Die

type Game struct {
	Score int
	Table Dice
	Hand  Dice
}

// Rolls a single die
func (d *Die) Roll() {
	d.Value = rand.Intn(6) + 1
}

// Roll all dice in a []Die
func (d Dice) Roll() {
	for i := range d {
		d[i].Roll()
	}
}

// Formats a Die for output
func (d *Die) String() string {
	return fmt.Sprint(d.Value)
}

// Formats a []Die for output
func (dice Dice) String() string {
	var result strings.Builder
	for d := range dice {
		r := fmt.Sprintf("Die #%d is a %d\n", d+1, dice[d].Value)
		result.WriteString(r)
	}
	return result.String()
}

// Returns a pointer to a Game, with initial values set
func NewGame() *Game {
	fmt.Printf("Starting a new game of Farkle.\nRolling Dice...\n\n")

	set := make([]Die, 6)
	for d := range set {
		set[d].Roll()
	}

	return &Game{
		Score: 0,
		Table: set,
		Hand:  nil,
	}
}

func (dice Dice) Score() int {
	counts := make(map[Die]int)
	for _, d := range dice {
		counts[d]++
	}

	// check for a straight
	if len(counts) == 6 {
		return 1500
	}

	var total int
	for d, v := range counts {
		if v < 3 {
			total += (v * d.Score())
		} else {
			switch v {
			case 3:
				if d.Value == 1 {
					total += (v - 2) * 1000
				} else {
					total += d.Value * 100
				}
			case 4:
				total += FOUROFAKIND
			case 5:
				total += FIVEOFAKIND
			case 6:
				total += SIXOFAKIND
			}
		}

	}
	return total
}

func (d Die) Score() int {
	switch d.Value {
	case 1:
		return 100
	case 5:
		return 50
	default:
		return 0
	}
}

func (g *Game) Hold(nums []int) Dice {
	var set Dice
	for i := len(nums) - 1; i >= 0; i-- {
		set = append(set, g.Table[nums[i]-1])
		g.Hand = append(g.Hand, g.Table[nums[i]-1])
		g.Table = append(g.Table[:(nums[i]-1)], g.Table[nums[i]:]...)
	}
	return set
}

func (g *Game) Play() {
	var gameover bool
	for !gameover {
		fmt.Printf("Score:\n%v\n", g.Score)
		fmt.Printf("Hand:\n%v\n", g.Hand)
		fmt.Printf("Table:\n%v\n", g.Table)
		fmt.Printf("Enter die # to keep, (r)oll, (q)uit: ")

		if g.Table.Score() == 0 {
			gameover = true
		}

		var cmd string
		fmt.Scanf("%s", &cmd)

		switch cmd {
		case "q":
			gameover = true
		default:
			held := numbers(cmd)

			reserved := g.Hold(held)
			g.Score += reserved.Score()

			if len(g.Table) == 0 {
				g.Hand = Dice{}
				g.Table = make(Dice, 6)
			}

			g.Table.Roll()

			if g.Table.Score() == 0 {
				fmt.Println(g.Table)
				fmt.Println("Farkle!")
				gameover = true
			}
		}

	}
}

func numbers(s string) []int {
	fields := strings.Split(s, ",")

	var nums []int
	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			fmt.Println("skipping invalid die")
			continue
		}
		nums = append(nums, n)
	}
	return nums
}

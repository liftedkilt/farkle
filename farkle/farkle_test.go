package farkle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoll(t *testing.T) {
	d := Die{}
	assert.Equal(t, 0, d.Value, "initial value should be 0")

	d.Roll()
	assert.NotEqual(t, 0, d.Value, "die cannot be 0 after rolling")
}

func TestRollDice(t *testing.T) {
	dice := Dice{
		Die{}, Die{}, Die{},
		Die{}, Die{}, Die{},
	}

	dice.Roll()

	for _, d := range dice {
		assert.NotEqual(t, 0, d.Value, "die cannot be 0 after rolling")
	}
}

func TestString(t *testing.T) {
	d := Die{}
	assert.Equal(t, "0", d.String(), "should print 0")

	d.Value = 6
	assert.Equal(t, "6", d.String(), "should print 6")
}

func TestStringDice(t *testing.T) {
	set := Dice{
		Die{1},
		Die{2},
		Die{3},
	}

	expected := "Die #1 is a 1\nDie #2 is a 2\nDie #3 is a 3\n"

	assert.Equal(t, expected, set.String(), "should print the correct text")
}

func TestNewGame(t *testing.T) {
	assert := assert.New(t)

	g := NewGame()

	assert.Equal(0, g.Score, "initial score should be 0")
	assert.Equal(6, len(g.Table), "there should be 6 dice on the table")
	assert.Equal(0, len(g.Hand), "there should be 0 dice in the hand")

	for _, d := range g.Table {
		assert.NotEqual(0, d.Value, "dice should be rolled")
	}
}

func TestScore(t *testing.T) {
	tt := []struct {
		input    Die
		expected int
	}{
		{
			input:    Die{0},
			expected: 0,
		},
		{
			input:    Die{1},
			expected: 100,
		},
		{
			input:    Die{2},
			expected: 0,
		},
		{
			input:    Die{3},
			expected: 0,
		},
		{
			input:    Die{4},
			expected: 0,
		},
		{
			input:    Die{5},
			expected: 50,
		},
		{
			input:    Die{6},
			expected: 0,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, tc.input.Score(), fmt.Sprintf("expected: %v, got %v", tc.expected, tc.input.Score()))
	}
}

func TestScoreDice(t *testing.T) {
	tt := []struct {
		reason   string
		input    Dice
		expected int
	}{
		{
			reason: "uninitialized dice must have no score",
			input: Dice{
				Die{0}, Die{0},
			},
			expected: 0,
		},
		{
			reason: "two ones scores 200",
			input: Dice{
				Die{1}, Die{1},
			},
			expected: 200,
		},
		{
			reason: "one 1 and one 5 score 150, 6 is non-scoring",
			input: Dice{
				Die{1}, Die{5}, Die{6},
			},
			expected: 150,
		},
		{
			reason: "three of a kind scores n * 100",
			input: Dice{
				Die{2}, Die{2}, Die{2},
			},
			expected: 200,
		},
		{
			reason: "three 1s scores 1000",
			input: Dice{
				Die{1}, Die{1}, Die{1},
			},
			expected: 1000,
		},
		{
			reason: "four of a kind scores 1000",
			input: Dice{
				Die{4}, Die{4}, Die{4}, Die{4},
			},
			expected: FOUROFAKIND,
		},
		{
			reason: "five of a kind scores 2000",
			input: Dice{
				Die{5}, Die{5}, Die{5}, Die{5}, Die{5},
			},
			expected: FIVEOFAKIND,
		},
		{
			reason: "six of a kind scores 3000",
			input: Dice{
				Die{6}, Die{6}, Die{6}, Die{6}, Die{6}, Die{6},
			},
			expected: SIXOFAKIND,
		},
		{
			reason: "three 1s and a 5 scores 1050",
			input: Dice{
				Die{1}, Die{1}, Die{1},
				Die{5},
			},
			expected: 1050,
		},
		{
			reason: "a straight scores 1500",
			input: Dice{
				Die{1}, Die{2}, Die{3}, Die{4}, Die{5}, Die{6},
			},
			expected: 1500,
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, tc.input.Score(), tc.reason)
	}
}

func TestHold(t *testing.T) {
	game := Game{
		Score: 0,
		Table: Dice{
			Die{1}, Die{2},
			Die{3}, Die{4},
			Die{5}, Die{5},
		},
		Hand: nil,
	}

	dice := game.Hold([]int{1, 5})
	assert.Equal(t, Dice{Die{5}, Die{1}}, dice, "dice should hold the reserved dice")
	assert.Equal(t, Dice{Die{5}, Die{1}}, game.Hand, "hand should hold the proper dice")
	assert.Equal(t, Dice{Die{2}, Die{3}, Die{4}, Die{5}}, game.Table, "table should hold the remaining dice")

	dice = game.Hold([]int{4})
	assert.Equal(t, Dice{Die{5}}, dice, "should hold the proper dice")
	assert.Equal(t, Dice{Die{5}, Die{1}, Die{5}}, game.Hand, "hand should hold the proper dice")
	assert.Equal(t, Dice{Die{2}, Die{3}, Die{4}}, game.Table, "table should hold the remaining dice")
}

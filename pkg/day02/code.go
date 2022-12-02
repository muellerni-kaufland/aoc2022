package day02

import (
	"encoding/csv"
	"errors"
	"io"
)

type Move int

const (
	Rock Move = iota + 1
	Paper
	Scissors
)

type Round struct {
	Attack   Move
	Response Move
}

func (r Round) Score() int {
	switch {
	// you win
	case (r.Attack == Rock && r.Response == Paper) || (r.Attack == Paper && r.Response == Scissors) || (r.Attack == Scissors && r.Response == Rock):
		return int(r.Response) + 6
	// you lose
	case (r.Attack == Paper && r.Response == Rock) || (r.Attack == Scissors && r.Response == Paper) || (r.Attack == Rock && r.Response == Scissors):
		return int(r.Response)
	// draw
	default:
		return int(r.Response) + 3
	}
}

// HackedScore is my lazy approach of not rewriting any logic. So `Response` in this case means a different thing:
// * Rock means "you need to lose"
// * Paper means "you need to draw"
// * Scissors means "you need to win"
func (r Round) HackedScore() int {
	// you need to lose
	if r.Response == Rock {
		switch r.Attack {
		case Rock:
			return int(Scissors)
		case Paper:
			return int(Rock)
		case Scissors:
			return int(Paper)
		}
	}
	// you need to draw
	if r.Response == Paper {
		return int(r.Attack) + 3
	}
	// you need to win
	if r.Response == Scissors {
		switch r.Attack {
		case Rock:
			return int(Paper) + 6
		case Paper:
			return int(Scissors) + 6
		case Scissors:
			return int(Rock) + 6
		}
	}
	return -1
}

func TaskA(r io.Reader) int {
	rounds := setup(r)
	if rounds == nil {
		return -1
	}
	return calculateScores(rounds, false)
}

func TaskB(r io.Reader) int {
	rounds := setup(r)
	if rounds == nil {
		return -1
	}
	return calculateScores(rounds, true)
}

func calculateScores(rounds []Round, hacked bool) int {
	var result int
	for _, round := range rounds {
		if hacked {
			result += round.HackedScore()
		} else {
			result += round.Score()
		}
	}
	return result
}

func decode(row []string) (Round, error) {
	switch row[0] {
	case "A":
		switch row[1] {
		case "X":
			return Round{Rock, Rock}, nil
		case "Y":
			return Round{Rock, Paper}, nil
		case "Z":
			return Round{Rock, Scissors}, nil
		}
	case "B":
		switch row[1] {
		case "X":
			return Round{Paper, Rock}, nil
		case "Y":
			return Round{Paper, Paper}, nil
		case "Z":
			return Round{Paper, Scissors}, nil
		}
	case "C":
		switch row[1] {
		case "X":
			return Round{Scissors, Rock}, nil
		case "Y":
			return Round{Scissors, Paper}, nil
		case "Z":
			return Round{Scissors, Scissors}, nil
		}
	}
	return Round{}, errors.New("invalid combination of attack and response")
}

func setup(r io.Reader) []Round {
	csvReader := csv.NewReader(r)
	csvReader.Comma = ' '
	csvReader.FieldsPerRecord = 2

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil
	}
	rounds := make([]Round, 0, len(records))
	for _, row := range records {
		round, err := decode(row)
		if err != nil {
			return nil
		}
		rounds = append(rounds, round)
	}
	return rounds
}

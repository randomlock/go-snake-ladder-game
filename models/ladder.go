package models

type Ladder struct {
    From  int
    To    int
}

func NewLadder(from int, to int) Ladder {
    return Ladder{
        From: from,
        To:   to,
    }
}



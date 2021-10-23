package models

type Player struct {
    playerName string
    currentNum int
}

func NewPlayer(playerName string) *Player {
    return &Player{playerName: playerName}
}

func (p Player) GetName() string {
    return p.playerName
}

func (p *Player) SetNum(num int)  {
    p.currentNum  = num
}

func (p Player) GetNum() int  {
    return p.currentNum
}



package model

type Player struct {
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	Surname        string   `json:"surname"`
	Age            int      `json:"age"`
	Foot           string   `json:"foot"`
	Club           string   `json:"club"`
	PlayerPosition Position `json:"position"`
	PlayerValue    int      `json:"playerValue"`
}

type Position string

const (
	Goalkeeper        Position = "Goalkeeper"
	LeftBack                   = "LeftBack"
	CentreBack                 = "CentreBack"
	RightBack                  = "RightBack"
	DefenseMidfield            = "DefenseMidfield"
	CentralMidfield            = "CentralMidfield"
	AttackingMidfield          = "AttackingMidfield"
	RightWinger                = "RightWinger"
	LeftWinger                 = "LeftWinger"
	CentreForward              = "CentreForward"
)

// func (p Position) String() string {
// 	return [...]string{
// 		"Goalkeeper", "LeftBack", "CentreBack", "RightBack", "DefenseMidfield",
// 		"CentralMidfield", "CentralMidfield", "AttackingMidfield", "RightWinger", "LeftWinger", "CentreForward"}[p]
// }

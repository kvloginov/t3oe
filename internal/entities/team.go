package entities

const (
	TEAM_UNDEFINED Team = iota
	TEAM_BLUE
	TEAM_RED
)

type Team int

func (t Team) GetTeam() Team {
	return t
}

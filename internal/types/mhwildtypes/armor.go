package mhwildtypes

type Armor struct {
	GameID     int               `json:"game_id"`
	Rarity     int               `json:"rarity"`
	Names      map[string]string `json:"names"`
	Skills     map[string]int    `json:"skills"`
	SetBonus   *SetBonus         `json:"set_bonus,omitempty"`
	GroupBonus *GroupBonus       `json:"group_bonus,omitempty"`
	Pieces     []ArmorPiece      `json:"pieces,omitempty"`
	// Add any other fields as needed
}

type ArmorPiece struct {
	Kind   string            `json:"kind"`
	Names  map[string]string `json:"names"`
	Skills map[string]int    `json:"skills"`
	// Add any other fields as needed
}

type SetBonus struct {
	SkillID int `json:"skill_id"`
	// optional: ranks, etc.
}

type GroupBonus struct {
	SkillID int `json:"skill_id"`
	// optional: ranks, etc.
}

type ArmorMatchResult struct {
	Set            Armor
	SetLevelMatch  bool
	MatchingPieces []ArmorPieceMatch
}

type ArmorPieceMatch struct {
	Piece      ArmorPiece
	SkillLevel int
}

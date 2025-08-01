package types

type Skill struct {
	GameID int               `json:"game_id"`
	Names  map[string]string `json:"names"`
	Ranks  []SkillRank       `json:"ranks"`
}

type SkillRank struct {
	Level        int               `json:"level"`
	Descriptions map[string]string `json:"descriptions"`
}

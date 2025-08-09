package mhwildtypes

type Skill struct {
	GameID int               `json:"game_id"`
	Names  map[string]string `json:"names"`
	Ranks  []SkillRank       `json:"ranks"`
	// other fields omitted for brevity

}

type SkillRank struct {
	Level        int               `json:"level"`
	Descriptions map[string]string `json:"descriptions"`
	// other fields omitted for brevity

}

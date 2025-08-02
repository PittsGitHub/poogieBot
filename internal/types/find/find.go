package findType

type Weapon struct {
	Names  map[string]string `json:"names"`
	Rarity int               `json:"rarity"`
	Skills map[string]int    `json:"skills"` // SkillID string (to match Skill.json) → level
}

type Armor struct {
	Names  map[string]string `json:"names"`
	Rarity int               `json:"rarity"`
	Pieces []ArmorPiece      `json:"pieces"`
}

type ArmorPiece struct {
	Skills map[string]int `json:"skills"` // SkillID string (to match Skill.json) → level
}

type Skill struct {
	GameID int64             `json:"game_id"`
	Names  map[string]string `json:"names"`
}

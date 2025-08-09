package mhwildtypes

type Decoration struct {
	GameID       int64             `json:"game_id"`
	Names        map[string]string `json:"names"`
	Descriptions map[string]string `json:"descriptions"`
	Rarity       int               `json:"rarity"`
	Price        int               `json:"price"`
	Level        int               `json:"level"`
	Skills       map[string]int    `json:"skills"` // <-- required for this version
	AllowedOn    string            `json:"allowed_on"`
	IconColor    string            `json:"icon_color"`
	IconColorID  int               `json:"icon_color_id"`
}

type DecorationSkillMatch struct {
	DecorationName  string
	DecorationLevel int // This is the level of the decoration, not the skill
	AllowedOn       string
	SkillLevel      int
	Rarity          int
}

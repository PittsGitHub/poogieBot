package mhwildtypes

type Talisman struct {
	GameID int64          `json:"game_id"`
	Ranks  []TalismanRank `json:"ranks"`
}

type TalismanRank struct {
	Names        map[string]string `json:"names"`
	Descriptions map[string]string `json:"descriptions"`
	Rarity       int               `json:"rarity"`
	Level        int               `json:"level"`
	Price        int               `json:"price"`
	Skills       map[string]int    `json:"skills"`
	Recipe       struct {
		Inputs map[string]int `json:"inputs"`
	} `json:"recipe"`
}

type TalismanSkillMatch struct {
	TalismanName string
	SkillLevel   int
	Rarity       int
}

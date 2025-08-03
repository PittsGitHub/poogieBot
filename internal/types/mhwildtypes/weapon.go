package mhwildtypes

type Weapon struct {
	Names        map[string]string `json:"names"`
	Descriptions map[string]string `json:"descriptions"`
	Rarity       int               `json:"rarity"`
	AttackRaw    int               `json:"attack_raw"`
	Affinity     int               `json:"affinity"`
	Skills       map[string]int    `json:"skills"`
	// other fields omitted for brevity

}

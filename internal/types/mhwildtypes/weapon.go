package mhwildtypes

type Weapon struct {
	Names            map[string]string `json:"names"`
	Rarity           int               `json:"rarity"`
	AttackRaw        int               `json:"attack_raw"`
	AttackCalculated int
	Affinity         int            `json:"affinity"`
	Skills           map[string]int `json:"skills"`
	// other fields omitted for brevity

}

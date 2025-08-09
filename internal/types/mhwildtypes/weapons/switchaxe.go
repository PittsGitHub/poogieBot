package weapons

type SwitchAxe struct {
	Names     map[string]string `json:"names"`
	Rarity    int               `json:"rarity"`
	AttackRaw int               `json:"attack_raw"`
	Affinity  int               `json:"affinity"`
	Skills    map[string]int    `json:"skills"`
	Phial     Phial             `json:"phial"`
}

type Phial struct {
	Kind string `json:"kind"`
}

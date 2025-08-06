package weapons

type LightBowgun struct {
	Names       map[string]string `json:"names"`
	Rarity      int               `json:"rarity"`
	AttackRaw   int               `json:"attack_raw"`
	Affinity    int               `json:"affinity"`
	Skills      map[string]int    `json:"skills"`
	SpecialAmmo string            `json:"special_ammo"`
	Ammo        []Ammo            `json:"ammo"`
}

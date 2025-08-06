package weapons

type Ammo struct {
	Kind     string `json:"kind"`
	Level    int    `json:"level"`
	Capacity int    `json:"capacity"`
}

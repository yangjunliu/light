package light

type Light struct {
	Name string
}

func NewLight() *Light {
	return &Light{}
}

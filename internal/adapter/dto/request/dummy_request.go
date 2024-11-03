package request

type Dummy struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewDummy() *Dummy {
	return &Dummy{}
}

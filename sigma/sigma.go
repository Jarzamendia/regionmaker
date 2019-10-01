package sigma

//Payload A Node and Edge struct
type Payload struct {
	Edges []Edges `json:"edges"`
	Nodes []Nodes `json:"nodes"`
}

//Edges An edge struct
type Edges struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

//Nodes A node struct
type Nodes struct {
	Color   string `json:"color"`
	ID      string `json:"id"`
	Label   string `json:"label"`
	Size    int    `json:"size"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Custom1 string `json:"custom1"`
	Custom2 string `json:"custom2"`
	Custom3 string `json:"custom3"`
	Custom4 string `json:"custom4"`
}

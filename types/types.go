package types

//NodeConfig Arquivo que representa os nodes.
type NodeConfig struct {
	Generationtime string `json:"generationtime"`
	Nodes          []struct {
		CPU        string `json:"cpu"`
		Mem        string `json:"mem"`
		Region     string `json:"region"`
		Servername string `json:"servername"`
		Status     string `json:"status"`
	} `json:"nodes"`
	Version string `json:"version"`
}

//SigmaStruct asds
type SigmaStruct struct {
	Edges []struct {
		ID     string `json:"id"`
		Source string `json:"source"`
		Target string `json:"target"`
	} `json:"edges"`
	Nodes []struct {
		Color string  `json:"color"`
		ID    string  `json:"id"`
		Label string  `json:"label"`
		Size  float64 `json:"size"`
		X     float64 `json:"x"`
		Y     float64 `json:"y"`
	} `json:"nodes"`
}

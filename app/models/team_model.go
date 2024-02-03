package models

type Team struct {
	BaseModel
	Layer1 int `json:"layer1"`
	Layer2      int    `json:"layer2"`
	Layer3      int    `json:"layer3"`
	LayerValid1 string `json:"layerValid1"`
	LayerValid2 string `json:"layerValid2"`
	LayerValid3 string `json:"layerValid3"`
}

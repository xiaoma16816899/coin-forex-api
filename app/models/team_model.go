package models

type Team struct {
	BaseModel
	Layer1      int    `json:"layer1"`
	Layer2      int    `json:"layer2"`
	Layer3      int    `json:"layer3"`
	LayerValid1 string `json:"layer_valid1"`
	LayerValid2 string `json:"layer_valid2"`
	LayerValid3 string `json:"layer_valid3"`
}

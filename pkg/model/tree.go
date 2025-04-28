package model

type TreeData struct {
	Id       string `json:"id"`
	Label    string `json:"label"`
	Children []struct {
		Id    string `json:"id"`
		Label string `json:"label"`
	} `json:"children,omitempty"`
}

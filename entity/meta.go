package entity

type MetaData struct {
	ID    int64       `json:"id,omitempty"`
	Key   string      `json:"key,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

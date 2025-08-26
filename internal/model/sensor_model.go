package model

type SensorPayload struct {
	ID1         string  `json:"id1"`
	ID2         int     `json:"id2"`
	SensorType  string  `json:"sensor_type"`
	SensorValue float64 `json:"sensor_value"`
	Timestamp   string  `json:"timestamp"`
}

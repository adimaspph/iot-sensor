package model

type SensorPayload struct {
	ID1         string  `json:"id1"`
	ID2         int     `json:"id2"`
	SensorType  string  `json:"sensor_type"`
	SensorValue float64 `json:"sensor_value"`
	Timestamp   string  `json:"timestamp"`
}

type ChangeIntervalRequest struct {
	Interval int `json:"interval" validate:"required,gte=1,lte=100"`
}

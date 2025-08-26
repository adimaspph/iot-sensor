package messaging

import (
	"encoding/json"
	"iot-sensor/internal/model"
	"math"
	"math/rand"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

type SensorPublisher struct {
	Log        *logrus.Logger
	Rnd        *rand.Rand
	MqttClient mqtt.Client
	topic      string
	id1        string
	id2        int
	sensorType string
	interval   time.Duration
	min        float64
	max        float64
}

func NewSensorPublisher(
	logger *logrus.Logger,
	mqttClient mqtt.Client,
	topic string,
	id1 string,
	id2 int,
	sensorType string,
	interval time.Duration,
	min float64,
	max float64,
) *SensorPublisher {
	return &SensorPublisher{
		Log:        logger,
		Rnd:        rand.New(rand.NewSource(time.Now().UnixNano())),
		MqttClient: mqttClient,
		topic:      topic,
		id1:        id1,
		id2:        id2,
		sensorType: sensorType,
		interval:   interval,
		min:        min,
		max:        max,
	}
}

func (s *SensorPublisher) Start() {

	if s.interval < 0 {
		s.interval = time.Second
	}

	ticker := time.NewTicker(s.interval)
	s.Log.WithFields(logrus.Fields{"topic": s.topic, "interval": s.interval}).Info("sensor publisher started")
	defer func() {
		ticker.Stop()
		s.Log.Info("sensor publisher stopped")
	}()

	for {
		select {
		case <-ticker.C:
			payload := s.nextReading()
			b, err := json.Marshal(payload)
			if err != nil {
				s.Log.WithError(err).Error("marshal payload")
				continue
			}

			tok := s.MqttClient.Publish(s.topic, 0, true, b)
			tok.Wait()
			if err := tok.Error(); err != nil {
				s.Log.WithError(err).Error("publish failed")
				continue
			}
			s.Log.WithField("payload", string(b)).Debug("published")
		}
	}
}

func (s *SensorPublisher) nextReading() model.SensorPayload {
	v := s.min + s.Rnd.Float64()*(s.max-s.min)
	v = math.Round(v*10) / 10 // 1 decimal place
	return model.SensorPayload{
		ID1:         s.id1,
		ID2:         s.id2,
		SensorType:  s.sensorType,
		SensorValue: v,
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
	}
}

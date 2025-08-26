package main

import (
	"iot-sensor/internal/config"
	"iot-sensor/internal/delivery/messaging"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	//validate := config.NewValidator(viperConfig)
	mqttClient := config.NewMqtt(viperConfig, log)

	// Connect MQTT
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	log.Info("MQTT connected")

	sensorType := viperConfig.GetString("SENSOR_TYPE")
	interval := viperConfig.GetInt("PUBLISH_INTERVAL")
	pubInterval := time.Duration(interval) * time.Second
	min := viperConfig.GetFloat64("PUBLISH_MIN")
	max := viperConfig.GetFloat64("PUBLISH_MAX")

	for i := 0; i < 100; i++ {
		go func() {
			id := i + 1
			sensorPubliser := messaging.NewSensorPublisher(
				log,
				mqttClient,
				viperConfig.GetString("MQTT_TOPIC"),
				"SENSOR-"+strconv.Itoa(id),
				id,
				sensorType,
				pubInterval,
				min,
				max,
			)
			sensorPubliser.Start()
		}()
	}

	// Wait for shutdown signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	s := <-sigCh
	log.Infof("Received signal: %s. Shutting down...", s.String())

	// Allow in-flight MQTT work to flush
	mqttClient.Disconnect(250)
	log.Info("MQTT disconnected")

	log.Info("Shutdown complete")
}

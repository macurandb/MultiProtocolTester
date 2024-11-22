package kafka

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// KafkaConfig defines the structure for Kafka configuration loaded from YAML.
type KafkaConfig struct {
	BootstrapServers string `yaml:"bootstrap_servers"`
	Topics           struct {
		SendTopic    string `yaml:"send_topic"`
		ReceiveTopic string `yaml:"receive_topic"`
	} `yaml:"topics"`
	GroupID string `yaml:"group_id"`
	Timeout int    `yaml:"timeout"`
}

// LoadKafkaConfig loads and parses the Kafka configuration from a YAML file.
func LoadKafkaConfig(filePath string) (*KafkaConfig, error) {
	configData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config KafkaConfig
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

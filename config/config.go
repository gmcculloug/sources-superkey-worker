package config

import (
	"os"

	"github.com/spf13/viper"
)

// SuperKeyWorkerConfig is the struct for storing runtime configuration
type SuperKeyWorkerConfig struct {
	Hostname           string
	KafkaBrokers       []string
	KafkaGroupID       string
	MetricsPort        int
	LogLevel           string
	LogGroup           string
	AwsRegion          string
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	SourcesHost        string
	SourcesScheme      string
	SourcesPort        int
}

// Get - returns the config parsed from runtime vars
func Get() *SuperKeyWorkerConfig {
	options := viper.New()

	options.SetDefault("KafkaBrokers", []string{"localhost:9092"})
	options.SetDefault("KafkaGroupID", "sources-superkey-worker")
	options.SetDefault("MetricsPort", 9394)
	options.SetDefault("LogLevel", "INFO")
	options.SetDefault("LogGroup", "platform-dev")
	options.SetDefault("AwsRegion", "us-east-1")
	options.SetDefault("AwsAccessKeyId", os.Getenv("CW_AWS_ACCESS_KEY_ID"))
	options.SetDefault("AwsSecretAccessKey", os.Getenv("CW_AWS_SECRET_ACCESS_KEY"))
	options.SetDefault("SourcesHost", "localhost")
	options.SetDefault("SourcesScheme", "http")
	options.SetDefault("SourcesPort", 3002)

	hostname, _ := os.Hostname()
	options.SetDefault("Hostname", hostname)

	options.AutomaticEnv()

	return &SuperKeyWorkerConfig{
		Hostname:           options.GetString("Hostname"),
		KafkaBrokers:       options.GetStringSlice("KafkaBrokers"),
		KafkaGroupID:       options.GetString("KafkaGroupID"),
		MetricsPort:        options.GetInt("MetricsPort"),
		LogLevel:           options.GetString("LogLevel"),
		LogGroup:           options.GetString("LogGroup"),
		AwsRegion:          options.GetString("AwsRegion"),
		AwsAccessKeyId:     options.GetString("AwsAccessKeyID"),
		AwsSecretAccessKey: options.GetString("AwsSecretAcessKey"),
		SourcesHost:        options.GetString("SourcesHost"),
		SourcesScheme:      options.GetString("SourcesScheme"),
		SourcesPort:        options.GetInt("SourcesPort"),
	}
}
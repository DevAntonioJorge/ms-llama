package main

import (
	"os"
	"time"
)

type Config struct {
	Server ServerConfig
	Kafka KafkaConfig
	Ollama OllamaConfig
	Logger LoggerConfig
}


type ServerConfig struct {
	Host string
	Port string
	ReadTimeout, WriteTimeout, IdleTimeout time.Duration
}

type LoggerConfig struct {
	LogLevel string
}
type KafkaConfig struct {
	Brokers []string
	Topics  Topics
}


type Topics struct {
	ChatRequests string
	ChatResponses string
}


type OllamaConfig struct{
	Host, Port string
}


func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getTimeEnv(key string, defaultValue time.Duration) time.Duration {
	value, err := time.ParseDuration(os.Getenv(key))
	if err != nil{
		return defaultValue
	}
	return value
}

func Load() Config {

	return Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnv("SERVER_PORT", ":8080"),
			ReadTimeout: getTimeEnv("SERVER_READ_TIMEOUT", 30*time.Second),
			WriteTimeout: getTimeEnv("SERVER_WRITE_TIMEOUT", 30*time.Second),
			IdleTimeout: getTimeEnv("SERVER_IDLE_TIMEOUT", 30*time.Second),
		},
		Kafka: KafkaConfig{
			Brokers: []string{getEnv("KAFKA_BROKER1", "localhost:9092")},
			Topics: Topics{
				ChatRequests: getEnv("CHAT_REQUEST_TOPIC", "chat_request"),
				ChatResponses: getEnv("CHAT_RESPONSE_TOPIC", "chat_response"),
			},
		},
		Ollama: OllamaConfig{
			Host: getEnv("OLLAMA_HOST", "localhost"),
			Port: getEnv("OLLAMA_PORT", ":11434"),
		},
		Logger: LoggerConfig{
			LogLevel: getEnv("LOG_LEVEL", "debug"),
		},
	}
}
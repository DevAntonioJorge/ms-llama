package main

import (
	"fmt"
	"ms-llama/internal/kafka"
	"ms-llama/internal/ollama"
	"ms-llama/pkg/logger"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)


func main() {
	c := Load()
	logger := logger.NewLogger(c.Logger.LogLevel)
	logger.Info("Config loaded")
}
func run(c Config, logger *zap.Logger) error {
	

	_ = ollama.NewClient(c.Ollama.Host, c.Ollama.Port)
	_, err  := kafka.NewClient(c.Kafka.Brokers) 
	if err != nil{
		return fmt.Errorf("failed to initialize kafka configuration")
	}
	return nil
}
func init() {
	_ = godotenv.Load()
}


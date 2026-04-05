package configuration

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config - корневой набор настроек для HTTP, БД и поведения календаря
type Config struct {
	HTTP HTTPConfig
}

// HTTPConfig - порт и таймауты HTTP-сервера
type HTTPConfig struct {
	Host            string        // хост HTTP-сервера
	Port            string        // порт HTTP-сервера
	ReadTimeout     time.Duration // таймаут чтения тела запроса
	WriteTimeout    time.Duration // таймаут записи ответа
	IdleTimeout     time.Duration // keep-alive
	ShutdownTimeout time.Duration // ожидание завершения при graceful shutdown
}

// Load - читает окружение и подставляет дефолты (ошибка, если Validate не проходит)
func Load() (Config, error) {

	// подставляем дефолты
	cfg := Config{
		HTTP: HTTPConfig{
			Host:            getEnv("HTTP_HOST", "0.0.0.0"),
			Port:            getEnv("HTTP_PORT", "8081"),
			ReadTimeout:     getDuration("HTTP_READ_TIMEOUT", 15*time.Second),
			WriteTimeout:    getDuration("HTTP_WRITE_TIMEOUT", 15*time.Second),
			IdleTimeout:     getDuration("HTTP_IDLE_TIMEOUT", 60*time.Second),
			ShutdownTimeout: getDuration("HTTP_SHUTDOWN_TIMEOUT", 30*time.Second),
		},
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// Validate - проверка инвариантов после Load
func (c *Config) Validate() error {

	if c.HTTP.Host == "" {
		return fmt.Errorf("configuration: не задан хост сервера (HTTP_HOST)")
	}
	if c.HTTP.Port == "" {
		return fmt.Errorf("configuration: не задан порт сервера (HTTP_PORT)")
	}

	return nil
}

// getEnv получает переменную окружения в виде строки
func getEnv(key, def string) string {

	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}

// getInt получает переменную окружения в виде числа
func getInt(key string, def int) int {

	s := os.Getenv(key)
	if s == "" {
		return def
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return def
	}

	return n
}

// getDuration получает переменную окружения в виде продолжительности
func getDuration(key string, def time.Duration) time.Duration {

	s := os.Getenv(key)
	if s == "" {
		return def
	}

	d, err := time.ParseDuration(s)
	if err != nil {
		return def
	}

	return d
}

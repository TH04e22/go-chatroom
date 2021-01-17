package config

var Redis *RedisConfig

func init() {
	Redis = &RedisConfig{
		Address:    "redis:6379",
		Password:   "",
		HubPrefix:  "hub",
		RoomPrefix: "room",
	}
}

type RedisConfig struct {
	Address    string
	Password   string
	HubPrefix  string
	RoomPrefix string
}

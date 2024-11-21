package config

type Config struct {
	App []App `yaml:"app"` // App Client들을 배열로 관리
}

type App struct {
	App struct {
		Port    string `yaml:"port"`
		Version string `yaml:"version"`
	} `yaml:"app"`

	Http     HttpCfg   `yaml:"http"`
	Producer *Producer `yaml:"kafka"`
	// Producer - 1 Producer
	// HTTP     - N Client
}

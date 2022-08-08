package agent

type RegistryConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type Config struct {
	Registry RegistryConfig `json:"registry"`
	Period   int            `json:"period"`
}

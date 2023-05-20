package main

import (
	"log"
	"recon/config"
	"recon/resource"
)

// implements Config interface and holds resource configuration
type AppConfig struct {
	ResourceCfgs map[string]config.ResourceConfig
}

func (c *AppConfig) GetResourceCfgs() map[string]config.ResourceConfig {
	return c.ResourceCfgs
}

func main() {
	log.Print("INFO starting app...")
	// initialize program configuration
	c := &AppConfig{
		ResourceCfgs: make(map[string]config.ResourceConfig),
	}

	log.Println("INFO loading app configs...")
	// load other program config

	log.Println("INFO loading resource configs...")
	// load resource configs
	err := resource.LoadConfigs("./resource/configs", c)
	if err != nil {
		log.Println("ERROR", err)
		return
	}

	for _, r := range c.GetResourceCfgs() {
		log.Println("found resource:", r.ResourceName)
	}
}

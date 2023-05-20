package main

import (
	"log"
	"os"
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
	// init logging
	LOG_FILE := "./logs/recon_log.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	log.Println("...")
	log.Println("...")
	log.Println("...")

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Print("INFO starting app...")

	// initialize program configuration
	c := &AppConfig{
		ResourceCfgs: make(map[string]config.ResourceConfig),
	}

	log.Println("INFO loading app configs...")
	// load other program config

	log.Println("INFO loading resource configs...")
	// load resource configs
	err = resource.LoadConfigs("./resource/configs", c)
	if err != nil {
		log.Println("ERROR", err)
		return
	}

	for _, r := range c.GetResourceCfgs() {
		log.Println("found resource:", r.ResourceName)
	}

	log.Print("INFO exit")
}

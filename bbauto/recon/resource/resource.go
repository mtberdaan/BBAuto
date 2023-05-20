package resource

import (
	"os"
	"path/filepath"
	"recon/config"

	"gopkg.in/yaml.v2"
)

// load all resource configs
func LoadConfigs(folderPath string, cfg config.Config) error {

	// read all files in config folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	// iterate over the files
	for _, file := range files {
		// skip directories TODO
		if file.IsDir() {
			continue
		}

		// read the file contents
		filePath := filepath.Join(folderPath, file.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		// unmarshal the YAML data into struct
		var resourceCfg config.ResourceConfig
		err = yaml.Unmarshal(data, &resourceCfg)
		if err != nil {
			return err
		}

		// save the config in map
		cfg.GetResourceCfgs()[resourceCfg.ResourceName] = resourceCfg
	}

	return nil
}

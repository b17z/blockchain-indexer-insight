package config

import (
	"github.com/gertjaap/blockchain-indexer-insight/logging"
	"github.com/kelseyhightower/envconfig"
)

type ConfigSettings struct {
    BackendBaseUrl       string		`default:"https://tvtc.blkidx.org/"`
}

var _configSettingsRead bool = false
var _configSettings ConfigSettings = ConfigSettings{}

func GetConfiguration() ConfigSettings {
	if !_configSettingsRead {
		err := envconfig.Process("insightwrapper", &_configSettings)
		if err != nil {
			logging.Error.Println("Error reading config settings")
		}
	   _configSettingsRead = true;
	}
	return _configSettings
}
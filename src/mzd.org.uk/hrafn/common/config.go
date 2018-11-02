package common

/*-----------------------------------------------------------------------------
 **
 ** - hrafn -
 **
 ** Copyright (c) 2015-18 by Krogoth of
 ** Ministry of Zombie Defense - http://www.mzd.org.uk/
 **
 ** This program is free software; you can redistribute it and/or modify it
 ** under the terms of the GNU Affero General Public License as published by the
 ** Free Software Foundation, either version 3 of the License, or (at your option)
 ** any later version.
 **
 ** This program is distributed in the hope that it will be useful, but WITHOUT
 ** ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 ** FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
 ** for more details.
 **
 ** You should have received a copy of the GNU Affero General Public License
 ** along with this program. If not, see <http://www.gnu.org/licenses/>.
 **
 **-----------------------------------------------------------------------------
 **
 ** krogoth @ Ministry of Zombie Defense
 **
-----------------------------------------------------------------------------*/

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
)

func InitConfig() {

	// we look in these dirs for the config file
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.hrafn")
	viper.AddConfigPath("/etc/hrafn")

	// the file is expected to be named hrafn.config.json
	viper.SetConfigName("hrafn.config")
	viper.SetConfigType("json")

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {

		// if not found, write a standard config file and quit...
		writeStandardConfig()

		// quit execution
		LogFatal("Error reading config file. New file dumped.", logrus.Fields{"error": err, "filename": "hrafn.config.json"})
	}
}

// Helper function
func GetStringFromConfig(key string) string {

	return viper.GetString(key)
}

func GetIntFromConfig(key string) int {

	return viper.GetInt(key)
}

func GetLogLevel() string {

	loglevel := viper.GetString("log.level")
	if loglevel == "" {

		return "warn"
	} else {

		return loglevel
	}
}

func GetSubConfig(key string) *viper.Viper {

	return viper.Sub(key)
}

//
func writeStandardConfig() error {

	err := ioutil.WriteFile("hrafn.config.json", defaultConfig, 0700)

	return err
}

//
var defaultConfig = []byte(`
{
  "log": {
    "level": "debug"
  },
  "export": {
	"csv_separator": ":"
  },
  "tools": {
    "job_limit": "2",
    "sslyze": "python -m sslyze",
    "nmap": "nmap"
  },
  "files": {
    "domains": "domains.csv",
	"ciphers": "ciphers.csv",
	"scans": "scans",
    "output": "results_current.csv"
  }
}
`)

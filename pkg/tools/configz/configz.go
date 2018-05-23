/*
Copyright [yyyy] [name of copyright owner]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package configz

import (
	"fmt"
	"os"

	"ipaas/pkg/tools/log"

	"github.com/Unknwon/goconfig"
	"github.com/howeyc/fsnotify"
)

var (
	cfg *goconfig.ConfigFile
	err error
)

func init() {
	config := os.Getenv("CONFIG_PATH")
	if config == "" {
		workDir, err := os.Getwd()
		if err != nil {
			panic(`
        -------------------------------------------------------
		you must set the config file path,there two method set:
		1. you should define CONFIG_PATH environment
		2. current dir app/app.conf file muste be exist
		-------------------------------------------------------`,
			)
		}
		config = fmt.Sprintf("%v/%v", workDir, "conf/app.conf")
	}
	cfg, err = goconfig.LoadConfigFile(config)
	if err != nil {
		panic(err)
	}
}

func GetString(section, key, defaults string) string {
	return cfg.MustValue(section, key, defaults)
}

func GetStringArray(section, key, delim string) []string {
	return cfg.MustValueArray(section, key, delim)
}

func MustBool(section, key string, defaultVal bool) bool {
	return cfg.MustBool(section, key, defaultVal)
}

func MustFloat64(section, key string, defaultVal float64) float64 {
	return cfg.MustFloat64(section, key, defaultVal)
}

func MustInt(section, key string, defaultVal int) int {
	return cfg.MustInt(section, key, defaultVal)
}

func MustInt64(section, key string, defaultVal int64) int64 {
	return cfg.MustInt64(section, key, defaultVal)
}

//HeatLoad watcher notify the config file, when the file was changed, reload the file to memory
func HeatLoad() {
	config := os.Getenv("CONFIG_PATH")
	if config == "" {
		workDir, err := os.Getwd()
		if err != nil {
			panic(`
        -------------------------------------------------------
		you must set the config file path,there two method set:
		1. you should define CONFIG_PATH environment
		2. current dir app/app.conf file muste be exist
		-------------------------------------------------------`,
			)
		}
		config = fmt.Sprintf("%v/%v", workDir, "conf/app.conf")
	}

	wacther, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error("create the file watcher err: %v", err)
	}
	defer func() {
		if err = wacther.Close(); err != nil {
			log.Critical("close the file wather err:%v", err)
		}
	}()

	wacther.Watch(config)
	done := make(chan bool)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error("panic is happend: %v", err)
			}
		}()
		n := 0
		for {
			select {
			case event := <-wacther.Event:
				if event.IsModify() && event.IsAttrib() {
					n++
				}
				if n == 2 {
					cfg, err = goconfig.LoadConfigFile(config)
					if err != nil {
						panic(err)
					}
					n = 0
					log.Notice("the config %v file has bee modify, we reloaded the config file success", event.Name)
				}
			case err := <-wacther.Error:
				log.Error("the file watcher err: %v", err)
			}
		}
	}()
	<-done
}

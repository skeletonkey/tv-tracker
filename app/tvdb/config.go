// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package tvdb

import "github.com/skeletonkey/lib-core-go/config"

var cfg *tvdb

func getConfig() *tvdb {
	config.LoadConfig("tvdb", &cfg)
	return cfg
}

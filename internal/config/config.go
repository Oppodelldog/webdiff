package config

import "os"

func DataDir() string {
	return getString("WEBDIFF_DATA_DIR", ".")
}

func getString(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return def
}

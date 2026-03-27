package proxy

import (
	"fmt"
	"os"
	"strings"
)

func readProxies(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("proxy read: %s", err.Error())
	}

	return strings.Split(string(file), "\n"), nil
}

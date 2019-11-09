package plugins

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

// Load attempts to load plugins from the plugins directory
func Load(dir string) Plugins {
	reg := make(Plugins)

	for _, path := range getPluginPaths(dir) {
		// load module
		// 1. open the so file to load the symbols
		plug, err := plugin.Open(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 2. look up a symbol (an exported function or variable)
		// in this case, variable Greeter
		symGreeter, err := plug.Lookup("Greeter")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		reg[strings.Replace(path, ".so", "", 1)] = symGreeter.(Greeter)
	}

	return reg
}

func getPluginPaths(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	paths := make([]string, 0)

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".so" {
			paths = append(paths, file.Name())
		}
	}

	return paths
}

// Registry contains the loaded plugins
var Registry Plugins

func init() {
	fmt.Printf("Loading plugins...")
	Registry = Load(os.Getenv("PLUGINS_DIR"))
	fmt.Printf("ok\n")
}

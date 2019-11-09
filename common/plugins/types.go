package plugins

type Greeter interface {
	Greet()
}

type Plugins map[string]Greeter

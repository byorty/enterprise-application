package application

type Arguments struct {
	ConfigFilename string `short:"c" long:"config" description:"A path to config file" required:"true"`
}

package connection

// Config Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {

	// TODO MOCK THIS!
	c.Server = "localhost"
	c.Database = "hpc-rest-api"

	// if _, err := toml.DecodeFile("config.toml", &c); err != nil {
	// 	log.Fatal(err)
	// }
}

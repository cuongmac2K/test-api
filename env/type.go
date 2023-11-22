package env

type DataYaml struct {
	ConfigMysql ConfigMysql `yaml:"configMysql"`
}

type ConfigMysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

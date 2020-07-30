package config

type ServerConfig struct {
	WorkID     int       `toml:"workId"`
	BusinessID int       `toml:"businessId"`
	Log        Log       `toml:"log"`
	Settings   Settings  `toml:"settings"`
	Broker     Broker    `toml:"broker"`
	Redis      Redis     `toml:"redis"`
	RPCServer  RPCServer `toml:"rpcServer"`
	Mysql      Mysql     `toml:"mysql"`
	Queue      Queue     `toml:"queue"`
}

type Log struct {
	LogPath  string `toml:"logPath"`
	LogLevel string `toml:"logLevel"`
}

type Settings struct {
	Listeners []string `toml:"listeners"`
}

type Broker struct {
	Mode             string `toml:"mode"`
	BindingKeyPrefix string `toml:"bindingKeyPrefix"`
	RoutingKey       string `toml:"routingKey"`
	URL              string `toml:"url"`
}

type Redis struct {
	Password       string        `toml:"password"`
	Index          int           `toml:"index"`
	MasterHost     string        `toml:"masterHost"`
	ReplicateHosts []interface{} `toml:"replicateHosts"`
	ConnTimeout    int           `toml:"connTimeout"`
	ReadTimeout    int           `toml:"readTimeout"`
	WriteTimeout   int           `toml:"writeTimeout"`
	KeepAlive      int           `toml:"keepAlive"`
	AliveTime      int           `toml:"aliveTime"`
}

type RPCServer struct {
	Addr string `toml:"addr"`
}

type Masters struct {
	Name   string `toml:"name"`
	Dsn    string `toml:"dsn"`
	Active int    `toml:"active"`
	Idle   int    `toml:"idle"`
}

type Slaves struct {
	Name   string `toml:"name"`
	Dsn    string `toml:"dsn"`
	Active int    `toml:"active"`
	Idle   int    `toml:"idle"`
}

type Mysql struct {
	Masters []Masters `toml:"masters"`
	Slaves  []Slaves  `toml:"slaves"`
}

type Queue struct {
	Routines int  `toml:"routines"`
	Timeout  int  `toml:"timeout"`
	Cap      int  `toml:"cap"`
	Debug    bool `toml:"debug"`
}

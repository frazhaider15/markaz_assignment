package conf

var (
	Config GbeConfig
)

type GbeConfig struct {
	AWS          AWS              `json:"aws"`
	Mysql        MysqlConfig      `json:"mysqlConfig"`
	RestServer   RestServerConfig `json:"restServer"`
	KafkaBrokers []string         `json:"kafka"`
}

type GrpcConn struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
}

type AWS struct {
	AWSRegion string `json:"awsRegion"`
}

type MysqlConfig struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	DbName     string `json:"dbName"`
	DbUserName string `json:"dbUserName"`
	DbPassword string `json:"dbPassword"`
}

type RestServerConfig struct {
	Addr string `json:"addr"`
}

func SetConfig() {
	Config = GbeConfig{
		AWS: AWS{
			AWSRegion: "",
		},
		Mysql: MysqlConfig{
			Host:       GetValueFromSSMByKey("DB_HOST", false),
			Port:       GetValueFromSSMByKey("DB_PORT", false),
			DbName:     GetValueFromSSMByKey("DB_NAME", true),
			DbUserName: GetValueFromSSMByKey("DB_USER", false),
			DbPassword: GetValueFromSSMByKey("DB_PASSWORD", true),
		},
		RestServer: RestServerConfig{
			Addr: ":8001",
		},
	}
}

package conf

type Conf struct {
	Config *Config `json:"config"`
}

type Config struct {
	Log    Log `json:"log"`
	Server struct {
		Name string `json:"name"`
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
}

type Log struct {
	Level          int    `json:"level"`
	Formatter      string `json:"formatter"`
	CutTime        string `json:"CutTime"`
	CutTimeLimitIn string `json:"CutTimeLimitIn"`
	LogFileSaveNum int    `json:"LogFileSaveNum"`
}

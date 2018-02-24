package utils

type VfConfig struct {
	Addr string `json:"addr"`
}

var (
	vfConf VfConfig
)

func InitVf(conf VfConfig) {
	vfConf = conf
	if vfConf.Addr == "" {
		panic("vfunctions config error")
	}
}

func GetVfAddr() string {
	return vfConf.Addr
}

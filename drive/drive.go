package drive

type GoIpInterface interface {
	GetIp() string
	GetCountryCode() (string, error)
}

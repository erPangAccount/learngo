package config

const (
	//rpc service host
	ItemServiceHost   = ":1234"
	WorkerServiceHost = ":1235"

	//rpcService
	WorkerService = "WorkService.Process"

	//elastic
	ElasticIndex = "test"
	ElasticHost  = "192.168.12.13:9200"

	//ParserName
	CityParser           = "CityParser"
	UserInfoParser       = "UserInfoParser"
	CityListParser       = "CityListParser"
	NilRequestResultFunc = "NilRequestResultFunc"
)

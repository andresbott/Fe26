package fe26


type Settings struct {
	port int
	ip string
	docRoot string
	FeBase string
}

var Config = Settings{
	port: 7070,
	ip: "127.0.0.1",
	docRoot: "",
	FeBase: "fe26",
}

func Fe26()  {
	preStartChecks()
	fe26Router()
}


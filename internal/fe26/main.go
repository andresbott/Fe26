package fe26


type Settings struct {
	port int
	docRoot string
	FeBase string
}

var Config = Settings{
	port: 8080,
	docRoot: "",
	FeBase: "fe26",
}

func Fe26()  {
	preStartChecks()
	fe26Router()
}


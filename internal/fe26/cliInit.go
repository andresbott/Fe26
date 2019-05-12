package fe26

import (
	"os"
	"strconv"
	log "github.com/sirupsen/logrus"
	"flag"
	"github.com/AndresBott/Fe26/pkg/f"
)


func preStartChecks () {

	//flag.Usage = func() {
	//	fmt.Println("help")
	//	os.Exit(0)
	//}

	// Read Cli Parameters
	cliLogLevel := flag.String("loglevel", "info", "set the log level [debug|info|warn|error] (ENV: FE26_LOGLEVEL) ")
	cliPort := flag.Int("port", Config.port, "what port to listen to (ENV: FE26_PORT)")
	flag.Parse()
	cliDocRoot := ""
	if len(flag.Args()) > 0{
		cliDocRoot = flag.Args()[0]
	}

	setDocumentRoot(cliDocRoot)
	setPort(*cliPort)
	setLogLevel(*cliLogLevel)

	log.Info("Starting FE26")
	log.Info("{" +
		"root: \""+Config.docRoot+"\", " +
		"port: "+strconv.Itoa(Config.port)+", " +
		"log-level: \""+log.GetLevel().String()+"\", " +
		"}")

}

func setDocumentRoot(cliDocRoot string){

	envDocRoot := os.Getenv("FE26_ROOT")
	docRoot := ""
	if envDocRoot != ""{
		docRoot = envDocRoot
	}else if cliDocRoot != ""{
		docRoot = cliDocRoot
	}else {
		docRoot = "./"
	}
	docRoot,err:= f.GetAbsPathDir(docRoot)
	if err != nil {
		log.Fatal(err)
	}
	Config.docRoot = docRoot
}

func setPort(cliPort int)  {
	envPort, _ := strconv.Atoi(os.Getenv("FE26_PORT"))
	port := 0

	if envPort != 0{
		port = envPort
	}else if cliPort != 0{
		port = cliPort
	}
	Config.port =  port
}

func setLogLevel(cliLogLevel string){
	envLogLevel := os.Getenv("FE26_LOGLEVEL")
	logLevel := ""
	if envLogLevel != ""{
		logLevel = envLogLevel
	}else if cliLogLevel != ""{
		logLevel = cliLogLevel
	}else {
		logLevel = "info"
	}

	if logLevel == "info"{
		log.SetLevel(log.InfoLevel)
	} else if logLevel == "debug"{
		log.SetLevel(log.DebugLevel)
	} else if logLevel == "warn"{
		log.SetLevel(log.WarnLevel)
	} else if logLevel == "error"{
		log.SetLevel(log.ErrorLevel)
	}
}
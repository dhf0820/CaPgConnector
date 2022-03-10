package main
import (
	"fmt"
	"os"
	src "github.com/dhf0820/caDbConnector/src"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {

		fmt.Println("Start app")
		godotenv.Load()
	
		Formatter := new(log.TextFormatter)
		Formatter.TimestampFormat = "02-01-2006 15:04:05"
		Formatter.FullTimestamp = true
	
		log.SetFormatter(Formatter)
		log.SetLevel(log.DebugLevel)
		log.Infof("Getting MONGODB url")
		mongoURL, err := os.LookupEnv("MONGODB")
		if err != true {
			log.Errorf("Error in MONGODB lookup: %v", err)
			log.Fatal(err)
		}
	
		log.Infof("Calling Initialize all")
		config := src.InitializeAll(mongoURL, "")
}
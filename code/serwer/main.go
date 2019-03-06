package main

import "restapi"

import "github.com/gobuffalo/packr"
import "sync"
import "os"
import "fmt"
import "logger"
import "encoding/json"

var wg sync.WaitGroup

type Configuration struct {
	Port, Adres, Haslo string
	Debug              bool
}

func main() {
	file, _ := os.Open("./config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	logger.Start("./log.txt")
	sigs := make(chan os.Signal, 1)

	box := packr.NewBox("templates")
	srv, _ := restapi.Start(configuration.Adres, configuration.Port, "sqlite3", "./DANE.db", box, false, false, configuration.Haslo, configuration.Debug)
	//	srv2, _ := restapi.Start("0.0.0.0", "443", "sqlite3", "/tmp/test3.db", box, true, false)
	wg.Add(1)

	go func() {
		sig := <-sigs
		srv.Close()
		//		srv2.Close()
		fmt.Println(sig)
		wg.Done()
	}()
	wg.Wait()
}

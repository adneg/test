package restapi

//restapi ustawienia oraz ładowanie metod

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"time"

	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	//"tools"

	"github.com/gobuffalo/packr"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"api"
	"logger"
	//"sessiontoolsu"
)

var (
	DB   *gorm.DB
	REST *httprouter.Router
	Box  packr.Box
)

//func Stop() {
//	SRV.Close()
//}

func Start(ip, port, driver, dbplace string, box packr.Box, ssl, rqcrt bool, pass string, debug bool) (srv *http.Server, db *gorm.DB) {

	// inicjalizacja bazy danych
	db, err := gorm.Open(driver, dbplace)

	if err != nil {
		logger.FatalError.Panic(err.Error())
	}
	Box = box
	db.Exec("PRAGMA foreign_keys = ON")
	db.LogMode(debug)
	//up up up up
	REST := httprouter.New()

	api.Up(REST, db, box, pass)

	return Up(ip, port, ssl, rqcrt, REST), db
}

func Up(ip string, port string, ssl, rqcrt bool, REST *httprouter.Router) (srv *http.Server) {
	var err error
	if ssl {
		cert, err := tls.LoadX509KeyPair("./server.crt", "./server.key")
		if err != nil {
			logger.FatalError.Panic(err.Error())
		}

		caCert, err := ioutil.ReadFile("./CA.crt")
		if err != nil {
			logger.FatalError.Panic(err.Error())
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		config := &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
			Certificates: []tls.Certificate{cert},

			RootCAs:   caCertPool,
			ClientCAs: caCertPool,
		}
		if rqcrt {
			config.ClientAuth = tls.RequireAndVerifyClientCert
		}

		config.Rand = rand.Reader
		srv = &http.Server{
			Addr:              ip + ":" + port,
			TLSConfig:         config,
			Handler:           REST,
			ReadHeaderTimeout: 180 * time.Second,
		}
	} else {

		srv = &http.Server{
			Addr: ip + ":" + port,

			Handler:           REST,
			ReadHeaderTimeout: 180 * time.Second,
		}

	}
	srv.SetKeepAlivesEnabled(true)
	fmt.Println("---- SERWER HTTP START LISTENING ----\n")

	go func() {
		if ssl {

			err = srv.ListenAndServeTLS("server.crt", "server.key")
		} else {
			err = srv.ListenAndServe()
		}

		if err != nil {
			logger.FatalError.Panic(err.Error())
		}
		fmt.Println("---- SERWER HTTP STOP LISTENING ----\n")
	}()
	return srv

}

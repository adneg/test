package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"sync"

	"github.com/gorilla/sessions"

	//	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key           = []byte("super-secret-key")
	store         = sessions.NewCookieStore(key)
	numer         = 0
	mapa_kanalow  = make(map[int]*(chan Mes))
	mutex_kanalow = &sync.Mutex{}
)

func Createcallhandle() {
	store.MaxAge(0)
	//	fmt.Println("zaladowany")
	REST.GET("/", Info)

	REST.GET("/api/loginclient/:did", Login)
	REST.GET("/api/loginclientM/:did", Login2)
	REST.GET("/api/ws", DecoAcces(Wsocket))
	REST.GET("/api/data/last/:did", DecoAcces(GetRecords2))
	REST.GET("/api/data/between/:od/:do", DecoAcces(GetRecords3))
	REST.GET("/api/data", DecoAcces(GetRecords))       //pobranie wszystkich rekorgow
	REST.POST("/api/data/:did", DecoAcces(PutRecord))  //dodanie rekordu
	REST.DELETE("/api/data/:did", DecoAcces(RmRecord)) //usunienice po id

	REST.GET("/api/grupa", DecoAcces(GetGrupy))        //pobranie wszystkich rekorgow
	REST.POST("/api/grupa/:did", DecoAcces(PutGrupa))  //dodanie rekordu
	REST.DELETE("/api/grupa/:did", DecoAcces(RmGrupa)) //usunienice po id

	REST.GET("/api/produkt", DecoAcces(GetProdukty))       //pobranie wszystkich rekorgow
	REST.POST("/api/produkt/:did", DecoAcces(PutProdukt))  //dodanie rekordu
	REST.DELETE("/api/produkt/:did", DecoAcces(RmProdukt)) //usunienice po id
	REST.GET("/api/ww", DecoAcces(GetWw))
	REST.GET("/api/wo", DecoAcces(GetWo))
	REST.POST("/api/ww/:did", DecoAcces(PutWw))
	REST.POST("/api/wo/:did", DecoAcces(PutWo))

	REST.GET("/api/percent", DecoAcces(GetPercent))
	//usunienice po id
	//	REST.GET("/api/data/ptest/:did", DecoAcces(PutRecord))
	//	REST.GET("/api/data/test/:did", DecoAcces(RmRecord)) //

}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session, _ := store.Get(r, "sessiontopgoid")

	did := ps.ByName("did")
	if did == "test" {
		session.Values["authenticated"] = true
		session.Save(r, w)
		w.Write([]byte("ok"))
	}

}

func Login2(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session, _ := store.Get(r, "sessiontopgoid")

	did := ps.ByName("did")
	if did == pass {
		session.Values["authenticated"] = true
		session.Save(r, w)
		w.Write([]byte("ok"))
	}

}

func GetWw(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//	DB.Table("DANE").Select().Where("d")
	//	year, month, day := time.Now().Date()
	//	midnight := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	ww := Wspw{}

	DB.Last(&ww)
	ret, _ := json.Marshal(ww)
	w.Write(ret)
	return
	//	w.Write([]byte("Serwer Aktywizacji Danych get all data"))

}
func GetWo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//	DB.Table("DANE").Select().Where("d")
	//	year, month, day := time.Now().Date()
	//	midnight := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	wo := NormaOsobowa{}

	DB.Last(&wo)
	ret, _ := json.Marshal(wo)
	w.Write(ret)
	return
	//	w.Write([]byte("Serwer Aktywizacji Danych get all data"))

}

func PutWo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//w.Write([]byte("Serwer Aktywizacji Danych put"))
	did := ps.ByName("did")
	fmt.Println(did)
	record := NormaOsobowa{}
	record.Data = time.Now()

	err := json.Unmarshal([]byte(did), &record)
	if err == nil {

		n := DB.Save(&record).RowsAffected
		fmt.Println("id nowego rekordu: ", record.Id_norma_osobowa)
		if n == 1 {
			odp := Odp{}
			odp.Stat = true
			odp.Id = record.Id_norma_osobowa
			ret, _ := json.Marshal(odp)
			w.Write(ret)
			return
		}

	}
	fmt.Println(err.Error())
	odp := Odp{}
	odp.Stat = false

	ret, _ := json.Marshal(odp)

	w.Write(ret)

}
func PutWw(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//w.Write([]byte("Serwer Aktywizacji Danych put"))
	did := ps.ByName("did")
	fmt.Println(did)
	record := Wspw{}
	record.Data = time.Now()

	err := json.Unmarshal([]byte(did), &record)
	if err == nil {

		n := DB.Save(&record).RowsAffected
		fmt.Println("id nowego rekordu: ", record.Id_wspw)
		if n == 1 {
			odp := Odp{}
			odp.Stat = true
			odp.Id = record.Id_wspw
			ret, _ := json.Marshal(odp)
			w.Write(ret)
			return
		}

	}
	fmt.Println(err.Error())
	odp := Odp{}
	odp.Stat = false

	ret, _ := json.Marshal(odp)

	w.Write(ret)

}

func GetProdukty(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//	DB.Table("DANE").Select().Where("d")
	//	year, month, day := time.Now().Date()
	//	midnight := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	danes := []Produkty{}

	DB.Find(&danes)
	ret, _ := json.Marshal(danes)
	w.Write(ret)
	return
	//	w.Write([]byte("Serwer Aktywizacji Danych get all data"))

}

func PutProdukt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//w.Write([]byte("Serwer Aktywizacji Danych put"))
	did := ps.ByName("did")
	fmt.Println(did)
	record := Produkty{}

	err := json.Unmarshal([]byte(did), &record)
	if err == nil {

		n := DB.Save(&record).RowsAffected
		fmt.Println("id nowego rekordu: ", record.Id_produktu)
		if n == 1 {
			odp := Odp{}
			odp.Stat = true
			odp.Id = record.Id_produktu
			ret, _ := json.Marshal(odp)
			w.Write(ret)
			mutex_kanalow.Lock()
			for _, v := range mapa_kanalow {
				test := Mes{}
				test.Instrukcja = "ZmianaProdukty"

				//				//				test.Data, err = json.Marshal(record)
				//				if err != nil {
				//					fmt.Println(err.Error())
				//				}
				//				fmt.Println("KATUALIZACJ PRODUKTY")
				*v <- test

			}
			mutex_kanalow.Unlock()
			return
		}

	}
	fmt.Println(err.Error())
	odp := Odp{}
	odp.Stat = false

	ret, _ := json.Marshal(odp)

	w.Write(ret)

}

func RmProdukt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	did := ps.ByName("did")
	record := Produkty{}
	var err error
	record.Id_produktu, err = strconv.Atoi(did)

	if err == nil {
		n := DB.Delete(&record).RowsAffected
		if n == 1 {
			odp := Odp{}
			odp.Stat = true

			ret, _ := json.Marshal(odp)
			w.Write(ret)
			mutex_kanalow.Lock()
			for _, v := range mapa_kanalow {
				test := Mes{}
				test.Instrukcja = "ZmianaProdukty"

				//				//				test.Data, err = json.Marshal(record)
				//				if err != nil {
				//					fmt.Println(err.Error())
				//				}
				fmt.Println("KATUALIZACJ PRODUKTY")
				*v <- test

			}
			mutex_kanalow.Unlock()
			return
		}
	}
	odp := Odp{}
	odp.Stat = false
	ret, _ := json.Marshal(odp)
	w.Write(ret)
}

func RmGrupa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	did := ps.ByName("did")
	record := Grupy{}
	var err error
	record.Id_grupy, err = strconv.Atoi(did)

	if err == nil {
		n := DB.Delete(&record).RowsAffected
		if n == 1 {
			odp := Odp{}
			odp.Stat = true

			ret, _ := json.Marshal(odp)
			w.Write(ret)

			mutex_kanalow.Lock()
			for _, v := range mapa_kanalow {
				test := Mes{}
				test.Instrukcja = "ZmianaWydzialy"

				//				//				test.Data, err = json.Marshal(record)
				//				if err != nil {
				//					fmt.Println(err.Error())
				//				}
				fmt.Println("KATUALIZACJ WYDZIALOW")
				*v <- test

			}
			mutex_kanalow.Unlock()

			return
		}
	}
	odp := Odp{}
	odp.Stat = false
	ret, _ := json.Marshal(odp)
	w.Write(ret)
}

func GetGrupy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//	DB.Table("DANE").Select().Where("d")
	//	year, month, day := time.Now().Date()
	//	midnight := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	danes := []Grupy{}

	DB.Find(&danes)
	ret, _ := json.Marshal(danes)
	w.Write(ret)
	return
	//	w.Write([]byte("Serwer Aktywizacji Danych get all data"))

}

func PutGrupa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//w.Write([]byte("Serwer Aktywizacji Danych put"))
	did := ps.ByName("did")
	fmt.Println(did)
	record := Grupy{}

	err := json.Unmarshal([]byte(did), &record)
	if err == nil {

		n := DB.Save(&record).RowsAffected
		fmt.Println("id nowego rekordu: ", record.Id_grupy)
		if n == 1 {
			odp := Odp{}
			odp.Stat = true
			odp.Id = record.Id_grupy
			ret, _ := json.Marshal(odp)
			w.Write(ret)
			mutex_kanalow.Lock()
			for _, v := range mapa_kanalow {
				test := Mes{}
				test.Instrukcja = "ZmianaWydzialy"

				//				//				test.Data, err = json.Marshal(record)
				//				if err != nil {
				//					fmt.Println(err.Error())
				//				}
				fmt.Println("KATUALIZACJ WYDZIALOW")
				*v <- test

			}
			mutex_kanalow.Unlock()
			return
		}

	}
	//fmt.Println(err.Error())
	odp := Odp{}
	odp.Stat = false

	ret, _ := json.Marshal(odp)

	w.Write(ret)

}

func GetPercent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	year, month, day := time.Now().Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	danes := []Dane{}

	DB.Order("id desc").Where("Data BETWEEN ? AND ?", midnight, time.Now()).Find(&danes)

	ilosc := len(danes)
	//	var iloscP int
	if ilosc > 0 {
		var mapa map[string]*PercentProdukt
		mapa = make(map[string]*PercentProdukt)
		var percent float32
		for _, element := range danes {
			mapa[element.Nazwa_produktu] = &PercentProdukt{}
			//			percent = percent + element.ZZZ
		}
		for _, element := range danes {
			mapa[element.Nazwa_produktu].Sum = mapa[element.Nazwa_produktu].Sum + element.ZZZ
			mapa[element.Nazwa_produktu].Ilosc = mapa[element.Nazwa_produktu].Ilosc + 1
			//			percent = percent + element.ZZZ
		}
		for _, v := range mapa {
			//			iloscP = iloscP + 1

			percent = percent + v.Sum
		}
		odp := int(math.Round(float64(percent)))
		ret, _ := json.Marshal(odp)
		w.Write(ret)
	} else {
		ret, _ := json.Marshal(-100)
		w.Write(ret)
	}
}

func PutRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//w.Write([]byte("Serwer Aktywizacji Danych put"))
	did := ps.ByName("did")
	//	fmt.Println(did)
	record := Dane{}

	err := json.Unmarshal([]byte(did), &record)
	if err == nil {
		record.Data = time.Now()
		wspw := Wspw{}
		normao := NormaOsobowa{}
		DB.Last(&normao)
		DB.Last(&wspw)

		record.WspW = wspw.WspW
		record.Norma_osobowa = normao.Norma_osobowa
		record.ZZZ = ((((float32(record.Ilosc) * record.Wsp_produktu) / (record.Stan_osobowy * record.Norma_osobowa)) * record.WspW) * 100.00)
		//		fmt.Println(math.Round(float64(record.ZZZ)))
		//		fmt.Println(int(math.Round(float64(record.ZZZ))))
		//((ilosc_sztuk * NORMA_PRODUKTOWA)/(STAN_OSOBOWY * NormaOsobowa)) * W
		n := DB.Save(&record).RowsAffected
		fmt.Println("id nowego rekordu: ", record.Id)
		if n == 1 {
			odp := Odp{}
			odp.Stat = true
			odp.Id = record.Id
			ret, _ := json.Marshal(odp)
			w.Write(ret)
			mutex_kanalow.Lock()
			for _, v := range mapa_kanalow {
				test := Mes{}
				test.Instrukcja = "DodanieElementu"

				test.Data, err = json.Marshal(record)
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println("WYSYLAM! NOWY RECORD")
				*v <- test

			}
			mutex_kanalow.Unlock()

			return
		}

	}
	fmt.Println(err.Error())
	odp := Odp{}
	odp.Stat = false

	ret, _ := json.Marshal(odp)

	w.Write(ret)

}

func GetRecords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//	DB.Table("DANE").Select().Where("d")
	year, month, day := time.Now().Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	danes := []Dane{}

	DB.Order("id desc").Where("Data BETWEEN ? AND ?", midnight, time.Now()).Find(&danes)
	ret, _ := json.Marshal(danes)
	w.Write(ret)
	return
	//	w.Write([]byte("Serwer Aktywizacji Danych get all data"))

}
func GetRecords2(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//	DB.Table("DANE").Select().Where("d")
	did := ps.ByName("did")
	danes := []Dane{}
	n, e := strconv.Atoi(did)
	if e == nil {
		DB.Order("id desc").Limit(n).Find(&danes)
	}

	ret, _ := json.Marshal(danes)
	w.Write(ret)
	return
	//	w.Write([]byte("Serwer Aktywizacji Danych get all data"))

}
func GetRecords3(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	od := ps.ByName("od")
	do := ps.ByName("do")
	fmt.Println(od, do)
	//	DB.Table("DANE").Select().Where("d")
	//	year, month, day := time.Now().Date()
	//	midnight := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	danes := []Dane{}

	DB.Order("id desc").Where("Data BETWEEN ? AND ?", od, do).Find(&danes)
	ret, _ := json.Marshal(danes)
	w.Write(ret)
	return
	//	w.Write([]byte("Serwer Aktywizacji Danych get all data"))

}
func RmRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	did := ps.ByName("did")
	record := Dane{}
	err := json.Unmarshal([]byte(did), &record)

	if err == nil {
		n := DB.Delete(&record).RowsAffected
		if n == 1 {
			odp := Odp{}
			odp.Stat = true
			ret, _ := json.Marshal(odp)
			w.Write(ret)

			mutex_kanalow.Lock()
			for _, v := range mapa_kanalow {
				test := Mes{}
				test.Instrukcja = "UsuniencieElementu"
				test.Data = []byte(did)

				*v <- test
			}
			mutex_kanalow.Unlock()

			return
			return
		}
	}
	odp := Odp{}
	odp.Stat = false
	ret, _ := json.Marshal(odp)
	w.Write(ret)
}

func Info(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Serwer Aktywizacji Danych"))
}

func DecoAcces(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		session, _ := store.Get(r, "sessiontopgoid")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		} else {
			fn(w, r, p)

		}
	}
}

func Wsocket(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notify := w.(http.CloseNotifier).CloseNotify()
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//	fmt.Println("Conn")
	mutex_kanalow.Lock()

	numer2 := numer + 1
	numer = numer + 1
	kanal := make(chan Mes)
	mapa_kanalow[numer2] = &kanal
	mutex_kanalow.Unlock()

	//	fmt.Println("true")
GO_TO_IN:
	for true {
		//		time.Sleep(1 * time.Second)
		//		fmt.Println("start jest")
		select {
		case <-notify:
			fmt.Println("HTTP connection just closed.")
			goto GO_TO_OUT

		case msg := <-*mapa_kanalow[numer2]:

			dat, err := json.Marshal(msg)
			if err != nil {
				fmt.Println(err.Error())
				goto GO_TO_IN
				//				return

			}
			_, err = w.Write(dat)
			if err != nil {
				goto GO_TO_OUT
				//				return

			}
			_, err = w.Write([]byte("\u0004"))
			if err != nil {
				goto GO_TO_OUT
				//		return
			}
			//			fmt.Println("juz kniec?")
			flusher.Flush()

		}
		//		fmt.Println("co jest")
	}
GO_TO_OUT:
	//	fmt.Println("usuwamy kanal")
	mutex_kanalow.Lock()
	delete(mapa_kanalow, numer2)
	mutex_kanalow.Unlock()
	return

}

func Licze() {

}

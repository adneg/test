package rclient

import (
	"bufio"
	"encoding/json"
	"fmt"
	"mstr"
	"net/http"
)

func Message(ms []byte) {
	odp := mstr.Mes{}
	err := json.Unmarshal(ms, &odp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(odp.Instrukcja)
	if odp.Instrukcja == "DodanieElementu" {
		dane := mstr.Dane{}
		json.Unmarshal(odp.Data, &dane)
		fmt.Println(dane.Nazwa_grupy)
	}
	if odp.Instrukcja == "UsuniencieElementu" {
		dane := mstr.Dane{}
		json.Unmarshal(odp.Data, &dane)
		fmt.Println(dane.Id, dane.Index)
	}

}

func ConWs() {

	adr := "http://" + host + ":" + port + "/api/ws"
	req, err := http.NewRequest("GET", adr, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.AddCookie(&coocke)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	reader := bufio.NewReader(resp.Body)
	fmt.Println("co jest")
	for {
		line, err := reader.ReadBytes('\u0004')

		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("koniec")
				//				close(Closechan)
				return
				//				return "Pobrano: " + filepath
			} else {
				fmt.Println("blad", err.Error())
				//				close(Closechan)
				return
				//				return "Nie Pobrano: " + err.Error()
			}
		} else {
			go Message(line[0 : len(line)-1])
			//			Datachan <- line[0 : len(line)-1]
		}
	}
	//	responseReader := bytes.NewReader(body)

}

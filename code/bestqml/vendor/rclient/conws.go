package rclient

import (
	"bufio"

	"encoding/json"
	"fmt"
	"mstr"
	"net/http"
)

func ConWs(ch *chan []byte) {

	adr := "http://" + host + ":" + port + "/api/ws"
	req, err := http.NewRequest("GET", adr, nil)
	if err != nil {
		koniec := mstr.Mes{}
		koniec.Instrukcja = "EOF"
		kk, _ := json.Marshal(koniec)
		*ch <- kk
		fmt.Println(err.Error())
		return
	}

	req.AddCookie(&coocke)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		koniec := mstr.Mes{}
		koniec.Instrukcja = "EOF"
		kk, _ := json.Marshal(koniec)
		*ch <- kk
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

				//				return "Pobrano: " + filepath
			} else {
				fmt.Println("blad", err.Error())
				//				close(Closechan)

				//				return "Nie Pobrano: " + err.Error()
			}
			koniec := mstr.Mes{}
			koniec.Instrukcja = "EOF"
			kk, _ := json.Marshal(koniec)
			*ch <- kk
			return
		} else {
			fmt.Println("wysylam")
			*ch <- line[0 : len(line)-1]
			//			Datachan <- line[0 : len(line)-1]
		}
	}
	//	responseReader := bytes.NewReader(body)

}

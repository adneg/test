package rclient

import (
	"fmt"
	//"time"
	//	"bytes"
	"encoding/json"
	//	"io/ioutil"
	//	"bufio"
	"mstr"
	"net/http"
	"strconv"

	"gopkg.in/resty.v1"
)

var (
	port   = "8080"
	host   = "localhost"
	coocke http.Cookie
)

func Logon(s string, p int, pass string) (stat bool, info string) {
	host = s
	port = strconv.Itoa(p)
	adr := "http://" + host + ":" + port + "/api/loginclientM/" + pass
	resp, err := resty.R().Get(adr)

	if err != nil {
		fmt.Println(err.Error())
		return false, err.Error()
	}

	if "ok" == string(resp.Body()) {
		fmt.Println("ZALOGOWANY")
		cos := resp.Cookies()
		if len(cos) > 0 {
			coocke = *cos[0]
			//			for _, co := range cos {
			//			}

		}
		return true, ""
	}
	return false, ""
}

func PutProdukt(n mstr.Produkty) (odp mstr.Odp) {
	nb, _ := json.Marshal(n)
	adr := "http://" + host + ":" + port + "/api/produkt/" + string(nb)
	resp, err := resty.R().Post(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}
	err = json.Unmarshal(resp.Body(), &odp)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
	//	if (err == nil) && len(co) > 0  {
	//		fmt.Println([0])

	//	}

}

func RmProdukt(id int) (odp mstr.Odp) {
	ids := strconv.Itoa(id)
	adr := "http://" + host + ":" + port + "/api/produkt/" + ids
	resp, err := resty.R().Delete(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}
	err = json.Unmarshal(resp.Body(), &odp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return

}

func PutWw(n mstr.Wspw) (odp mstr.Odp) {
	nb, _ := json.Marshal(n)
	adr := "http://" + host + ":" + port + "/api/ww/" + string(nb)
	resp, err := resty.R().Post(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}
	err = json.Unmarshal(resp.Body(), &odp)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
	//	if (err == nil) && len(co) > 0  {
	//		fmt.Println([0])

	//	}

}

func PutWo(n mstr.NormaOsobowa) (odp mstr.Odp) {
	nb, _ := json.Marshal(n)
	adr := "http://" + host + ":" + port + "/api/wo/" + string(nb)
	resp, err := resty.R().Post(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}
	err = json.Unmarshal(resp.Body(), &odp)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
	//	if (err == nil) && len(co) > 0  {
	//		fmt.Println([0])

	//	}

}
func GetWo() (stat bool, ww mstr.NormaOsobowa) {
	adr := "http://" + host + ":" + port + "/api/wo"
	resp, err := resty.R().Get(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}

	err = json.Unmarshal(resp.Body(), &ww)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	stat = true
	return
}

func GetWw() (stat bool, ww mstr.Wspw) {
	adr := "http://" + host + ":" + port + "/api/ww"
	resp, err := resty.R().Get(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}

	err = json.Unmarshal(resp.Body(), &ww)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	stat = true
	return

}

func GetProdukty() (stat bool, danes []mstr.Produkty) {
	adr := "http://" + host + ":" + port + "/api/produkt"
	resp, err := resty.R().Get(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}

	err = json.Unmarshal(resp.Body(), &danes)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	stat = true
	return

}

func PutGrupa(n mstr.Grupy) (odp mstr.Odp) {
	nb, _ := json.Marshal(n)
	adr := "http://" + host + ":" + port + "/api/grupa/" + string(nb)
	resp, err := resty.R().Post(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}
	err = json.Unmarshal(resp.Body(), &odp)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
	//	if (err == nil) && len(co) > 0  {
	//		fmt.Println([0])

	//	}

}

func RmGrupa(id int) (odp mstr.Odp) {
	ids := strconv.Itoa(id)
	adr := "http://" + host + ":" + port + "/api/grupa/" + ids
	resp, err := resty.R().Delete(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}
	err = json.Unmarshal(resp.Body(), &odp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return

}

func GetGrupy() (stat bool, danes []mstr.Grupy) {
	adr := "http://" + host + ":" + port + "/api/grupa"
	resp, err := resty.R().Get(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}

	err = json.Unmarshal(resp.Body(), &danes)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	stat = true
	return

}

func PutData(n mstr.Dane) (odp mstr.Odp) {
	nb, _ := json.Marshal(n)
	adr := "http://" + host + ":" + port + "/api/data/" + string(nb)
	resp, err := resty.R().Post(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}
	err = json.Unmarshal(resp.Body(), &odp)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
	//	if (err == nil) && len(co) > 0  {
	//		fmt.Println([0])

	//	}

}

func RmData(n mstr.Dane) (odp mstr.Odp) {
	nb, _ := json.Marshal(n)
	adr := "http://" + host + ":" + port + "/api/data/" + string(nb)
	resp, err := resty.R().Delete(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}
	err = json.Unmarshal(resp.Body(), &odp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return

}

func GetData() (stat bool, danes []mstr.Dane) {
	adr := "http://" + host + ":" + port + "/api/data"
	resp, err := resty.R().Get(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}

	err = json.Unmarshal(resp.Body(), &danes)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	stat = true
	return

}

func GetData2(od, do string) (stat bool, danes []mstr.Dane) {
	// /api/data/between/:od/:do
	adr := "http://" + host + ":" + port + "/api/data/between/" + od + "/" + do
	resp, err := resty.R().Get(adr)

	if err != nil {
		fmt.Println(err.Error())
		return

	}

	err = json.Unmarshal(resp.Body(), &danes)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	stat = true
	return

}

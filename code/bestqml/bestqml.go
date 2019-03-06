package main

import (
	"fmt"

	"encoding/json"
	"os"

	//	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"

	"github.com/therecipe/qt/qml"

	//"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"

	"models"
	"mstr"
	"rclient"
	"time"
)

var (
	cSizeini, lSizeini, eSizeini, bSizeini, iSizeini int
	dpiScaling                                       bool
	styl                                             string

	hostname string
	port     int
)

func StartBridge() {

	BridgeQML_QmlRegisterType2("CustomQmlTypes", 1, 0, "BridgeQML")
	//BridgeLoader_QmlRegisterType2("CustomQmlTypes", 1, 0, "BridgeLoader")

}

type BridgeQML struct {
	core.QObject
	_ int            `property:"cSize"`
	_ int            `property:"lSize"`
	_ int            `property:"eSize"`
	_ int            `property:"bSize"`
	_ int            `property:"iSize"`
	_ func(stat int) `signal:"returnStatusOn"`
	//	_ func(string, string, int) `signal:"sendCheckStatusOn,auto"`
	//	_ func(string, string)      `signal:"sendWakeOnLan,auto"`
	// nazwagrupy, idg, stanosobowy,produktName,produktw,produktid,ilosc
	_ func(string, int, float32, string, float32, int, int) `signal:"sendZatwierdz,auto"`
	_ func(int, int)                                        `signal:"sendRemoveE,auto"`
	//	_ func(n int)               `signal:"afterWake_afterHalt,auto"`
	_ func() `constructor:"init"`
}

func (t *BridgeQML) init() {

	t.SetCSize(cSizeini)
	t.SetLSize(lSizeini)
	t.SetESize(eSizeini)
	t.SetBSize(bSizeini)
	t.SetISize(iSizeini)

	fmt.Println("startuje template,")
	go func() {
		ch := make(chan []byte, 20)
		go rclient.ConWs(&ch)
		for true {
			msg := <-ch
			fmt.Println("jest cos")
			t.Message(msg)
		}

	}()

	go func() {
		time.Sleep(time.Second * 5)
		stat, percent := rclient.GetPercent()
		if stat && percent > 0 {
			fmt.Println(percent)
			t.ReturnStatusOn(percent + 1000)

		}
	}()
}

func (t *BridgeQML) Message(ms []byte) {
	odp := mstr.Mes{}
	err := json.Unmarshal(ms, &odp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(odp.Instrukcja)
	if odp.Instrukcja == "ZmianaWydzialy" {
		//		number := len(models.GM.Grupaitem())
		//		fmt.Println(number)
		models.GM.RemoveAll()
		models.LoadGrupy()

	}
	if odp.Instrukcja == "ZmianaProdukty" {
		//		number := len(models.PM.Produktitem())
		models.PM.RemoveAll()
		models.LoadProdukty()
	}
	if odp.Instrukcja == "DodanieElementu" {
		dane := mstr.Dane{}
		json.Unmarshal(odp.Data, &dane)
		fmt.Println(dane.Nazwa_grupy)
		var p = models.NewElement(nil)
		p.SetElementGName(dane.Nazwa_grupy)
		p.SetElementStan(dane.Stan_osobowy)
		p.SetElementPName(dane.Nazwa_produktu)
		p.SetElementIlosc(dane.Ilosc)
		p.SetElementId(dane.Id)

		p.SetElementData(dane.Data.Format("2006-01-02 15:04"))
		//		fmt.Println("co jest")
		models.EM.AddElement(p)
		t.ReturnStatusOn(1)
		go func() {

			stat, percent := rclient.GetPercent()
			if stat && percent > 0 {
				fmt.Println(percent)
				t.ReturnStatusOn(percent + 1000)

			}
		}()
	}
	if odp.Instrukcja == "UsuniencieElementu" {
		dane := mstr.Dane{}
		json.Unmarshal(odp.Data, &dane)
		fmt.Println(dane.Id, dane.Index)
		t.ReturnStatusOn(-2)
		models.EM.RemoveElementId(dane.Id)
		//		models.EM.RemoveElement(dane.Index)
		go func() {

			stat, percent := rclient.GetPercent()
			if stat && percent > 0 {
				fmt.Println(percent)
				t.ReturnStatusOn(percent + 1000)

			}
			if percent == -100 {
				t.ReturnStatusOn(percent)

			}

		}()
	}
	if odp.Instrukcja == "EOF" {
		t.ReturnStatusOn(-1000)

	}

}

func (t *BridgeQML) sendRemoveE(index, eid int) {

	fmt.Println(index, eid)
	n := mstr.Dane{}
	n.Index = index
	n.Id = eid
	odp := rclient.RmData(n)
	if odp.Stat {
		//		models.EM.RemoveElement(index)
	} else {

		t.ReturnStatusOn(-1)

	}
}
func (t *BridgeQML) sendZatwierdz(gupName string, grupId int, presonLen float32, produktName string,
	produktW float32, produktId int, ilosc int) {
	fmt.Println("ZATWIERDZONO:", gupName, grupId, presonLen, produktName,
		produktW, produktId, ilosc)

	n := mstr.Dane{}
	n.Id_grupy = grupId
	n.Id_produktu = produktId

	n.Nazwa_grupy = gupName
	n.Stan_osobowy = presonLen
	n.Nazwa_produktu = produktName
	n.Wsp_produktu = produktW
	n.Ilosc = ilosc

	odp := rclient.PutData(n)
	if odp.Stat {

		//		t.ReturnStatusOn(1)
	} else {
		t.ReturnStatusOn(0)
	}

}

func main() {
	core.QCoreApplication_SetApplicationName("Akwizycja Danych")
	core.QCoreApplication_SetOrganizationName("kamil.sala@gmail.com")

	if _, err := os.Stat("./settings.ini"); err == nil {

		setting_ini := core.NewQSettings4("./settings.ini", core.QSettings__IniFormat, nil)
		var test bool

		cSizeini = setting_ini.Value("cSize", core.NewQVariant()).ToInt(&test)
		lSizeini = setting_ini.Value("lSize", core.NewQVariant()).ToInt(&test)
		eSizeini = setting_ini.Value("eSize", core.NewQVariant()).ToInt(&test)
		bSizeini = setting_ini.Value("bSize", core.NewQVariant()).ToInt(&test)
		iSizeini = setting_ini.Value("iSize", core.NewQVariant()).ToInt(&test)
		dpiScaling = setting_ini.Value("dpiScaling", core.NewQVariant()).ToBool()
		styl = setting_ini.Value("styl", core.NewQVariant()).ToString()
		hostname = setting_ini.Value("hostname", core.NewQVariant()).ToString()
		port = setting_ini.Value("port", core.NewQVariant()).ToInt(&test)
		statlogin, _ := rclient.Logon(hostname, port, "test")

		if statlogin {

			core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, dpiScaling)
			NewQGuiApplication := gui.NewQGuiApplication(len(os.Args), os.Args)
			quickcontrols2.QQuickStyle_SetStyle(styl)

			StartBridge()

			var view = qml.NewQQmlApplicationEngine(nil)
			models.LoadGrupy()
			models.LoadProdukty()
			models.LoadElementy()

			view.RootContext().SetContextProperty("GrupyModel", models.GM)
			view.RootContext().SetContextProperty("ProduktyModel", models.PM)
			view.RootContext().SetContextProperty("ElementyModel", models.EM)

			view.Load(core.NewQUrl3("qrc:/qml/start.qml", 0))

			//		var view = qml.NewQQmlApplicationEngine(nil)
			//		view.Load(core.NewQUrl3("qrc:/qml/error.qml", 0))

			NewQGuiApplication.Exec()
		} else {
			core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, dpiScaling)
			NewQGuiApplication := gui.NewQGuiApplication(len(os.Args), os.Args)
			var view = qml.NewQQmlApplicationEngine(nil)
			view.Load(core.NewQUrl3("qrc:/qml/error_no_connect.qml", 0))
			NewQGuiApplication.Exec()
		}

	} else {
		core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, dpiScaling)
		NewQGuiApplication := gui.NewQGuiApplication(len(os.Args), os.Args)
		var view = qml.NewQQmlApplicationEngine(nil)
		view.Load(core.NewQUrl3("qrc:/qml/error_no_ini.qml", 0))

		//		var view = qml.NewQQmlApplicationEngine(nil)
		//		view.Load(core.NewQUrl3("qrc:/qml/error.qml", 0))

		NewQGuiApplication.Exec()

	}

	//	// Default, Fusion, Imagine, Universal
	//	var (
	//		settings = core.NewQSettings5(nil)
	//		style    = quickcontrols2.QQuickStyle_Name()
	//	)
	//	if style != "" {
	//		settings.SetValue("style", core.NewQVariant14(style))
	//	} else {
	//		quickcontrols2.QQuickStyle_SetStyle(settings.Value("style", core.NewQVariant14("")).ToString())
	//	}

}

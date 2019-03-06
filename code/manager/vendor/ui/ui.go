package ui

import (
	//	"mstr"

	"encoding/csv"
	"fmt"
	"os"

	//	"strconv"
	"rclient"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var (
	hostname string
	port     int
	fontsize int
)

type Pilot struct {
	_ func(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9 string) `slot:"addTreeItem"`
	//slots
	//chan
	// do polaczenia jak zachowac sie podczas utraty polaczenia
	//varibles
	widgets.QMainWindow
	//	LogoutBool bool
	//	CloseBool  bool
	//	UserAccess mstru.UserLoginInfo
	//	//widgets
	treeRecord      *widgets.QTreeWidget
	menuApplication *widgets.QMenu
	menuManagment   *widgets.QMenu
	toolbar         *widgets.QToolBar
	//	menuModules     *widgets.QMenu
	//	menuEdit        *widgets.QMenu
	//	Area            *widgets.QMdiArea

	//	statusBar   *widgets.QStatusBar
	//	systray     *widgets.QSystemTrayIcon
	//	systrayMenu *widgets.QMenu

	//windows
	okno_logowania    *widgets.QDialog
	daty_przegladania *widgets.QDialog
	//sql

}

func IntPilot() *Pilot {
	return NewPilot(nil, 0)

}

func (p *Pilot) initOknoLogowania() {
	p.okno_logowania = widgets.NewQDialog(nil, 0)
	p.okno_logowania.SetAttribute(core.Qt__WA_DeleteOnClose, true)
	//	var haslob int
	//	var status bool
	//	login_textline := widgets.NewQLineEdit(p.okno_logowania)
	//	login_textline.SetText("kamil.sala@gmail.com")
	haslo_textline := widgets.NewQLineEdit(p.okno_logowania)
	haslo_textline.SetText("")
	grid := widgets.NewQGridLayout(nil)
	haslo_textline.SetEchoMode(2)
	p.okno_logowania.SetLayout(grid)
	//	login_label := widgets.NewQLabel(p.okno_logowania, 0)
	//	login_label.SetText("Email: ")
	haslo_label := widgets.NewQLabel(p.okno_logowania, 0)
	haslo_label.SetText("Hasło: ")

	//	grid.AddWidget(login_label, 0, 1, 0)
	//	grid.AddWidget(login_textline, 0, 2, 0)
	grid.AddWidget(haslo_label, 1, 1, 0)
	grid.AddWidget(haslo_textline, 1, 2, 0)
	haslo_textline.ConnectReturnPressed(func() {
		stat, info := rclient.Logon(hostname, port, haslo_textline.Text())
		if stat {
			p.okno_logowania.Close()
			p.loadUi()

		} else {
			if info != "" {
				widgets.QMessageBox_Warning(nil, "Uwaga!!", "Błąd połączenia z serwerem!",
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

			}
		}

	})
	//	login_textline.ConnectReturnPressed(haslo_textline.SetFocus2)
	//buttonZaloguj := widgets.NewQPushButton2("Zaloguj", p.okno_logowania)
	//buttonZaloguj.ConnectClicked(func(checked bool) {
	//		haslo_textline.ReturnPressed()
	//	})

	//	buttonAnuluj := widgets.NewQPushButton2("Anuluj", p.okno_logowania)
	//	buttonAnuluj.ConnectClicked(func(checked bool) {
	//		p.okno_logowania.Close()
	//	})
	//grid.AddWidget(buttonAnuluj, 2, 1, 0)
	//grid.AddWidget(buttonZaloguj, 2, 2, 2)
	//layout.AddWidget(buttonL, 0, core.Qt__AlignCenter)
	p.okno_logowania.SetWindowTitle("Logowanie")
	p.okno_logowania.SetWindowIcon(gui.NewQIcon5(":/qml/password.png"))
	p.okno_logowania.SetMinimumSize2(280, 90)

	p.ConnectAddTreeItem(p.addTreeItem)
	p.okno_logowania.Show()
}

func (p *Pilot) Start() {
	if _, err := os.Stat("./settings.ini"); err == nil {
		var test bool
		setting := core.NewQSettings4("./settings.ini", core.QSettings__IniFormat, nil)
		hostname = setting.Value("hostname", core.NewQVariant()).ToString()
		port = setting.Value("port", core.NewQVariant()).ToInt(&test)
		fontsize = setting.Value("fontsize", core.NewQVariant()).ToInt(&test)
		fmt.Println("wynik czcionki: ", fontsize)
		//		logindb := setting.Value("logindb", core.NewQVariant()).ToString()
		//		haslodb := setting.Value("haslodb", core.NewQVariant()).ToString()
		// firebird : /Dane/topconector.db?auth_plugin_name=Legacy_Auth&wire_auth=false"

		if len(hostname) == 0 {
			widgets.QMessageBox_Critical(nil, "Uwaga!!", "Parametr 'hostname' jest pusty w settings.ini! \nPrzerywam!!",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			p.Close()
			os.Exit(1)
		}
		if port <= 0 {
			widgets.QMessageBox_Critical(nil, "Uwaga!!", "Parametr 'port' jest pusty w settings.ini! \nPrzerywam!!",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			p.Close()
			os.Exit(1)
		}

		//		if len(logindb) == 0 {
		//			widgets.QMessageBox_Information(nil, "Uwaga!!", "Parametr 'logindb' jest pusty w settings.ini!\nUstawiam wartość domyślną: sysdba !",
		//				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

		//			logindb = "sysdba"
		//		}
		//		if len(haslodb) == 0 {
		//			widgets.QMessageBox_Information(nil, "Uwaga!!", "Parametr 'haslodb' jest pusty w settings.ini!\nUstawiam wartość domyślną: masterkey !",
		//				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		//			haslodb = "masterkey"
		//		}
		//		p.Db = dbtools.IntSqlPoll(logindb, haslodb, host, baza_danych)

		//fmt.Println(logindb, haslodb, host, baza_danych)
		//		if p.Db.ErrorCon != nil {
		//			widgets.QMessageBox_Critical(nil, "Uwaga!!", "Nie można się połączyć do wskazanej Bazy Danych! \nPrzerywam!!",
		//				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		//			p.Close()
		//			os.Exit(1)

		//		} else {
		p.initOknoLogowania()

		//		}

	} else {

		widgets.QMessageBox_Critical(nil, "Uwaga!!", "Brak Pliku settings.ini !",
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		p.Close()
		os.Exit(1)
	}

}

func (p *Pilot) loadUi() {

	p.SetWindowTitle("Akwizytor Danych")
	p.SetWindowIcon(gui.NewQIcon5(":/qml/logo.png"))
	p.loadMenu()
	p.loadToolbar()
	p.loadTable()
	p.SetCentralWidget(p.treeRecord)
	//	p.loadMenu()
	//	p.Area = widgets.NewQMdiArea(p)
	//	p.SetCentralWidget(p.Area)
	//	p.initStatusBar()
	//	p.intSystray()
	//	p.SetMinimumSize2(400, 400)
	//	p.systray.Show()

	//	p.ConnectCloseEvent(func(e *gui.QCloseEvent) {
	//		if p.CloseBool {
	//			e.Accept()
	//		} else {
	//			e.Ignore()
	//			p.Hide()
	//		}

	//	})

	if fontsize > 0 {
		FontNew := gui.NewQFont2("verdana", fontsize, 1, false)
		FontNew.SetBold(false)
		p.SetFont(FontNew)
		p.treeRecord.SetFont(FontNew)
	}

	p.ShowMaximized()
}
func (p *Pilot) loadTable() {
	p.treeRecord = widgets.NewQTreeWidget(p)
	p.treeRecord.SetSortingEnabled(true)
	//SetSortingEnabled(true)
	p.treeRecord.SetRootIsDecorated(false)
	p.treeRecord.SetHeaderLabels([]string{"Id", "Wydział", "Stan Osobowy", "Produkt", "Ilość",
		"Norma Osobowa", "Norma Produktowa [JPS]", "Współczynnik W", "Data", "Godzina"})
	p.treeRecord.Header().SetSectionResizeMode(3)

}

func (p *Pilot) loadToolbar() {
	p.toolbar = widgets.NewQToolBar2(p)
	datyAction := p.toolbar.AddAction("Wyznacz Zakres Przeglądania")
	datyAction.SetIcon(gui.NewQIcon5(":/qml/calender.png"))
	datyAction.ConnectTriggered(func(checked bool) {
		p.datyOkno()
	})

	exportAction := p.toolbar.AddAction("Wyeksportuj Rekordy")
	exportAction.SetIcon(gui.NewQIcon5(":/qml/export.png"))
	exportAction.ConnectTriggered(func(checked bool) {
		dialogF := widgets.NewQFileDialog(p, core.Qt__Dialog)

		place := dialogF.GetSaveFileName(p, "Zapisz", "./", "Pliki CSV (*.csv);;Wszystkie Pliki (*)", ".csv", dialogF.Options())
		if place == "" {
			return
		}
		//		fmt.Println(" to jes miejsce jak nic nie wybierzesz:", place)

		p.Exportuj(p.treeRecord.InvisibleRootItem(), place)

	})
	p.AddToolBar(core.Qt__TopToolBarArea, p.toolbar)

}

func (p *Pilot) Exportuj(root *widgets.QTreeWidgetItem, place string) {
	var file *os.File
	var err error
	//	if _, err = os.Stat(place); err == nil {
	// path/to/whatever exists

	//	} else if os.IsNotExist(err) {
	if _, err = os.Stat(place); err == nil {
		err = os.Remove(place)
		if err != nil {
			widgets.QMessageBox_Warning(nil, "Uwaga!!", "Istniejącego pliku nie da się zmodyfikować !",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			fmt.Println(err.Error())
			return
		}
	} else {
		if place[len(place)-4:] != ".csv" {
			place = place + ".csv"
		}

	}

	file, err = os.Create(place)

	if err != nil {
		widgets.QMessageBox_Warning(nil, "Uwaga!!", "Nie da się utworzyć pliku w wskzanej lokalizacji !",
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		fmt.Println(err.Error())
		return
	}

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer file.Close()
	defer writer.Flush()
	records := [][]string{}
	for i := 0; i < root.ChildCount(); i++ {
		item := root.Child(i)
		//		fmt.Println(item.Text(0), item.Text(1), item.Text(2), item.Text(3),
		//			item.Text(4), item.Text(5), item.Text(6), item.Text(7), item.Text(8), item.Text(9))
		value := []string{item.Text(1), item.Text(2), item.Text(3),
			item.Text(4), "NO", item.Text(5), "NP", item.Text(6), "W", item.Text(7), item.Text(8), item.Text(9)}

		records = append(records, value)
	}
	err = writer.WriteAll(records)
	if err != nil {
		fmt.Println(err.Error())
	}
	//
}

func (p *Pilot) loadMenu() {
	p.menuApplication = p.MenuBar().AddMenu2("Aplikacja")

	closeAction := p.menuApplication.AddAction("Wyłącz")
	//closeAction.SetShortcut(gui.NewQKeySequence2("Ctrl+Q", gui.QKeySequence__NativeText))
	closeAction.SetIcon(gui.NewQIcon5(":/qml/exit.png"))
	closeAction.ConnectTriggered(
		func(checked bool) {
			p.Close()
		})
	p.menuApplication.AddSeparator()
	infoAction := p.menuApplication.AddAction("O Programie")
	infoAction.SetIcon(gui.NewQIcon5(":/qml/logo.png"))
	infoAction.ConnectTriggered(
		func(checked bool) {
			p.infoProgram()
			//			p.Close()
		})
	p.menuManagment = p.MenuBar().AddMenu2("Ustawienia ")
	wskaznikiAction := p.menuManagment.AddAction("Wskaźniki")
	//	wskaznikiAction.SetIcon(gui.NewQIcon5(":/qml/users.png"))
	wskaznikiAction.ConnectTriggered(
		func(checked bool) {
			p.normyProgram()
			//			p.infoProgram()
			//			//			p.Close()
		})
	wydzialyAction := p.menuManagment.AddAction("Wydziały")
	//	wydzialyAction.SetIcon(gui.NewQIcon5(":/qml/settings.png"))
	wydzialyAction.ConnectTriggered(
		func(checked bool) {
			p.wydzialyProgram()
			//			p.infoProgram()
			//			//			p.Close()
		})
	produktyAction := p.menuManagment.AddAction("Produkty")
	//	wydzialyAction.SetIcon(gui.NewQIcon5(":/qml/settings.png"))
	produktyAction.ConnectTriggered(
		func(checked bool) {
			p.prooduktyProgram()
			//			p.infoProgram()
			//			//			p.Close()
		})

}

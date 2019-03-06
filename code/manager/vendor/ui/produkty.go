package ui

import (
	//	"fmt"
	"mstr"
	"rclient"

	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	//"gitlab.top/ksala/pamietnik_hasel/restclient"
)

type ProduktyOkno struct {
	//slots
	//chan
	// do polaczenia jak zachowac sie podczas utraty polaczenia
	//varibles
	widgets.QMainWindow
	tree *widgets.QTreeWidget
}

func (o *ProduktyOkno) Okno_mod_exec() {
	id, _ := strconv.Atoi(o.tree.CurrentItem().Text(0))
	okno_dodaj := widgets.NewQDialog(o, 0)
	okno_dodaj.SetWindowTitle("Modyfikuj Produkt")
	layoutV := widgets.NewQVBoxLayout()
	okno_dodaj.SetLayout(layoutV)
	label := widgets.NewQLabel(okno_dodaj, 0)
	labelW := widgets.NewQLabel(okno_dodaj, 0)
	labelW.SetText("Współczynnik: ")
	nazwa := widgets.NewQLineEdit(okno_dodaj)
	nazwa.SetText(o.tree.CurrentItem().Text(1))
	label.SetText("Nazwa Produktu: ")
	waga := widgets.NewQLineEdit(okno_dodaj)
	regex := core.NewQRegExp()
	regex.SetPattern("([0-9]|[1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9]|[1-9][0-9][0-9][0-9][0-9]||[1-9][0-9][0-9][0-9][0-9][0-9])([,][0-9][0-9][0-9][0-9])$")
	validator := gui.NewQRegExpValidator2(regex, waga)
	waga.SetValidator(validator)

	waga.SetText(o.tree.CurrentItem().Text(2))

	//			self.cena = QtGui.QLineEdit(self)
	//		#regex=QtCore.QRegExp("(6[0-5][0-5][0-3][0-5]|[1-5][0-9][0-9][0-9][0-9]|[7-9][0-9][0-9][0-9])$")
	//		regex=QtCore.QRegExp("([1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9]|[1-9][0-9][0-9][0-9][0-9]||[1-9][0-9][0-9][0-9][0-9][0-9])([.][0-9][0-9])$")
	//		validator=QtGui.QRegExpValidator(regex, self.cena)
	//		self.cena.setValidator(validator)

	buttondodaj := widgets.NewQPushButton2("Dodaj", okno_dodaj)
	buttondodaj.ConnectClicked(func(checked bool) {
		if len(nazwa.Text()) > 0 {
			newG := mstr.Produkty{}
			newG.Id_produktu = id
			newG.Nazwa_produktu = nazwa.Text()
			valfloat64, _ := strconv.ParseFloat(ConvercommaTodot(waga.Text()), 32)
			newG.Wsp_produktu = float32(valfloat64)
			odp := rclient.PutProdukt(newG)
			if odp.Stat && (odp.Id != 0) {
				o.LoadData()
				widgets.QMessageBox_Information(okno_dodaj, "Uwaga!!", "Rekord został ZMODYFIKOWANY ID: "+strconv.Itoa(odp.Id),
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				okno_dodaj.Close()
				//					okno.LoadData()

			} else {
				widgets.QMessageBox_Warning(okno_dodaj, "Uwaga!!", "Rekord nie został ZMODYFIKOWANY..\nNazwa ISTNIEJE lub serwer nie odpowiada !",
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				//					okno.LoadData()

			}

		} else {
			widgets.QMessageBox_Warning(okno_dodaj, "Uwaga!!", "Nazwa nie może być pusta !",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		}

	})
	layoutV.AddWidget(label, 0, core.Qt__AlignCenter)
	layoutV.AddWidget(nazwa, 0, core.Qt__AlignCenter)
	layoutV.AddWidget(labelW, 0, core.Qt__AlignCenter)
	layoutV.AddWidget(waga, 0, core.Qt__AlignCenter)
	layoutV.AddWidget(buttondodaj, 0, core.Qt__AlignCenter)
	if fontsize > 0 {
		FontNew := gui.NewQFont2("verdana", fontsize, 1, false)
		FontNew.SetBold(true)
		okno_dodaj.SetFont(FontNew)
	}
	okno_dodaj.Exec()

}

func (o *ProduktyOkno) Okno_dodaj_exec() {

	okno_dodaj := widgets.NewQDialog(o, 0)
	okno_dodaj.SetWindowTitle("Dodaj Produkt")
	layoutV := widgets.NewQVBoxLayout()
	okno_dodaj.SetLayout(layoutV)
	label := widgets.NewQLabel(okno_dodaj, 0)
	labelW := widgets.NewQLabel(okno_dodaj, 0)
	labelW.SetText("Norma Produktowa [JPS]: ")
	nazwa := widgets.NewQLineEdit(okno_dodaj)
	label.SetText("Nazwa Produktu: ")
	waga := widgets.NewQLineEdit(okno_dodaj)
	regex := core.NewQRegExp()
	regex.SetPattern("([0-9]|[1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9]|[1-9][0-9][0-9][0-9][0-9]||[1-9][0-9][0-9][0-9][0-9][0-9])([,][0-9][0-9][0-9][0-9])$")
	validator := gui.NewQRegExpValidator2(regex, waga)
	waga.SetValidator(validator)

	//			self.cena = QtGui.QLineEdit(self)
	//		#regex=QtCore.QRegExp("(6[0-5][0-5][0-3][0-5]|[1-5][0-9][0-9][0-9][0-9]|[7-9][0-9][0-9][0-9])$")
	//		regex=QtCore.QRegExp("([1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9]|[1-9][0-9][0-9][0-9][0-9]||[1-9][0-9][0-9][0-9][0-9][0-9])([.][0-9][0-9])$")
	//		validator=QtGui.QRegExpValidator(regex, self.cena)
	//		self.cena.setValidator(validator)

	buttondodaj := widgets.NewQPushButton2("Dodaj", okno_dodaj)
	buttondodaj.ConnectClicked(func(checked bool) {
		if len(nazwa.Text()) > 0 {
			newG := mstr.Produkty{}
			newG.Nazwa_produktu = nazwa.Text()
			valfloat64, _ := strconv.ParseFloat(ConvercommaTodot(waga.Text()), 32)
			newG.Wsp_produktu = float32(valfloat64)
			odp := rclient.PutProdukt(newG)
			if odp.Stat && (odp.Id != 0) {
				o.LoadData()
				widgets.QMessageBox_Information(okno_dodaj, "Uwaga!!", "Rekord został dodany ID: "+strconv.Itoa(odp.Id),
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				okno_dodaj.Close()
				//					okno.LoadData()

			} else {
				widgets.QMessageBox_Warning(okno_dodaj, "Uwaga!!", "Rekord nie został dodany..\nNazwa ISTNIEJE lub serwer nie odpowiada !",
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				//					okno.LoadData()

			}

		} else {
			widgets.QMessageBox_Warning(okno_dodaj, "Uwaga!!", "Nazwa Wydziału nie może być pusta !",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		}

	})
	layoutV.AddWidget(label, 0, core.Qt__AlignCenter)
	layoutV.AddWidget(nazwa, 0, core.Qt__AlignCenter)
	layoutV.AddWidget(labelW, 0, core.Qt__AlignCenter)
	layoutV.AddWidget(waga, 0, core.Qt__AlignCenter)
	layoutV.AddWidget(buttondodaj, 0, core.Qt__AlignCenter)
	if fontsize > 0 {
		FontNew := gui.NewQFont2("verdana", fontsize, 1, false)
		FontNew.SetBold(true)
		okno_dodaj.SetFont(FontNew)
	}
	okno_dodaj.Exec()

}

func (o *ProduktyOkno) LoadData() {
	o.tree.Clear()

	stat, dane := rclient.GetProdukty()
	if stat && (len(dane)) > 0 {
		for _, v := range dane {

			o.AddTreeItem(v)

		}
	}

}

func (o *ProduktyOkno) AddTreeItem(v mstr.Produkty) {
	itemP := widgets.NewQTreeWidgetItem3(o.tree, 1)
	itemP.SetText(0, strconv.Itoa(v.Id_produktu))
	itemP.SetText(1, v.Nazwa_produktu)
	itemP.SetText(2, ConverdotTocomma(v.Wsp_produktu))
	o.tree.SetCurrentItem(itemP)

}

func (o *ProduktyOkno) Okno_usun_exec() {
	okno_usun := widgets.NewQDialog(o, 0)
	okno_usun.SetWindowTitle("Usuń Produkt")
	layoutV := widgets.NewQVBoxLayout()
	layoutH := widgets.NewQHBoxLayout()
	nwidget := widgets.NewQWidget(okno_usun, 0)
	okno_usun.SetLayout(layoutV)
	nwidget.SetLayout(layoutH)
	label := widgets.NewQLabel(okno_usun, 0)
	label.SetText("Czy usunąć ten Produkt? ")
	buttonusun := widgets.NewQPushButton2("Usuń", okno_usun)
	buttonusun.ConnectClicked(func(checked bool) {
		i := o.tree.CurrentItem()
		eid, _ := strconv.Atoi(i.Text(0))

		odp := rclient.RmProdukt(eid)

		if odp.Stat {

			z := o.tree.InvisibleRootItem()
			z.RemoveChild(o.tree.CurrentItem())
			widgets.QMessageBox_Information(okno_usun, "Uwaga!!", "Rekord został usunięty: ",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			okno_usun.Close()

		} else {
			widgets.QMessageBox_Warning(okno_usun, "Uwaga!!", "Rekord nie został usuniety..\nRekord NIE ISTNIEJE lub serwer nie odpowiada !",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

		}

	})

	buttonanuluj := widgets.NewQPushButton2("Anuluj", okno_usun)
	buttonanuluj.ConnectClicked(func(checked bool) {
		okno_usun.Close()

	})

	layoutH.AddWidget(buttonusun, 0, core.Qt__AlignLeft)
	layoutH.AddWidget(buttonanuluj, 0, core.Qt__AlignRight)
	layoutV.AddWidget(label, 0, core.Qt__AlignCenter)
	layoutV.AddWidget(nwidget, 0, core.Qt__AlignCenter)
	if fontsize > 0 {
		FontNew := gui.NewQFont2("verdana", fontsize, 1, false)
		FontNew.SetBold(false)
		okno_usun.SetFont(FontNew)

	}

	okno_usun.Exec()

}

func (p *Pilot) prooduktyProgram() {

	p.SetVisible(false)
	okno := NewProduktyOkno(p, 0)
	p.okno_logowania.SetAttribute(core.Qt__WA_DeleteOnClose, true)
	okno.SetWindowTitle("Produkty")

	toolbar := widgets.NewQToolBar2(okno)

	addAction := toolbar.AddAction("Dadaj Produkt")
	addAction.SetIcon(gui.NewQIcon5(":/qml/dodaj.png"))
	addAction.ConnectTriggered(func(checked bool) {
		okno.Okno_dodaj_exec()

	})
	edytujAction := toolbar.AddAction("Modyfikuj Produkt")
	edytujAction.SetIcon(gui.NewQIcon5(":/qml/modyfikuj.png"))
	edytujAction.ConnectTriggered(func(checked bool) {
		okno.Okno_mod_exec()
	})
	usunAction := toolbar.AddAction("Usun Produkt")
	usunAction.SetIcon(gui.NewQIcon5(":/qml/usun.png"))
	usunAction.ConnectTriggered(func(checked bool) {
		if (okno.tree.CurrentItem().IsHidden() == false) && (okno.tree.CurrentItem().Type() > 0) {
			okno.Okno_usun_exec()
		}

	})
	okno.tree = widgets.NewQTreeWidget(okno)
	okno.SetCentralWidget(okno.tree)

	//		p.treeRecord.SetSortingEnabled(true)
	//SetSortingEnabled(true)
	okno.tree.SetHeaderLabels([]string{"Id", "Nazwa Produktu", "Norma Produktowa [JPS]"})
	okno.tree.Header().SetSectionResizeMode(3)

	okno.tree.SetRootIsDecorated(false)
	okno.AddToolBar(core.Qt__TopToolBarArea, toolbar)

	okno.SetAttribute(core.Qt__WA_DeleteOnClose, true)

	okno.SetMinimumSize2(650, 400)

	okno.ConnectCloseEvent(func(e *gui.QCloseEvent) {
		p.Show()

	})
	okno.LoadData()
	if fontsize > 0 {
		FontNew := gui.NewQFont2("verdana", fontsize, 1, false)
		FontNew.SetBold(false)
		okno.SetFont(FontNew)
		okno.tree.SetFont(FontNew)
	}
	okno.ShowMaximized()
}

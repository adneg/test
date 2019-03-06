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

type WydzialyOkno struct {
	widgets.QMainWindow
	//slots
	//chan
	// do polaczenia jak zachowac sie podczas utraty polaczenia
	//varibles
	tree *widgets.QTreeWidget
}

func (o *WydzialyOkno) LoadData() {
	o.tree.Clear()

	stat, dane := rclient.GetGrupy()
	if stat && (len(dane)) > 0 {
		for _, v := range dane {

			o.AddTreeItem(v)

		}
	}

}

func (o *WydzialyOkno) AddTreeItem(v mstr.Grupy) {
	itemP := widgets.NewQTreeWidgetItem3(o.tree, 1)
	itemP.SetText(0, strconv.Itoa(v.Id_grupy))
	itemP.SetText(1, v.Nazwa_grupy)
	o.tree.SetCurrentItem(itemP)

}

func (o *WydzialyOkno) Okno_usun_exec() {
	okno_usun := widgets.NewQDialog(o, 0)
	okno_usun.SetWindowTitle("Usuń Wydział")
	layoutV := widgets.NewQVBoxLayout()
	layoutH := widgets.NewQHBoxLayout()
	nwidget := widgets.NewQWidget(okno_usun, 0)
	okno_usun.SetLayout(layoutV)
	nwidget.SetLayout(layoutH)
	label := widgets.NewQLabel(okno_usun, 0)
	label.SetText("Czy usunąć ten wydział? ")
	buttonusun := widgets.NewQPushButton2("Usuń", okno_usun)
	buttonusun.ConnectClicked(func(checked bool) {
		i := o.tree.CurrentItem()
		eid, _ := strconv.Atoi(i.Text(0))

		odp := rclient.RmGrupa(eid)
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
		FontNew.SetBold(true)
		okno_usun.SetFont(FontNew)
	}

	okno_usun.Exec()

}
func (o *WydzialyOkno) Okno_dodaj_exec() {

	okno_dodaj := widgets.NewQDialog(o, 0)
	okno_dodaj.SetWindowTitle("Dodaj Wydział")
	layoutV := widgets.NewQVBoxLayout()
	okno_dodaj.SetLayout(layoutV)
	label := widgets.NewQLabel(okno_dodaj, 0)
	label.SetText("Nazwa Wydziału: ")
	nazwa := widgets.NewQLineEdit(okno_dodaj)
	buttondodaj := widgets.NewQPushButton2("Dodaj", okno_dodaj)
	buttondodaj.ConnectClicked(func(checked bool) {
		if len(nazwa.Text()) > 0 {
			newG := mstr.Grupy{}
			newG.Nazwa_grupy = nazwa.Text()
			odp := rclient.PutGrupa(newG)
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
	layoutV.AddWidget(buttondodaj, 0, core.Qt__AlignCenter)
	if fontsize > 0 {
		FontNew := gui.NewQFont2("verdana", fontsize, 1, false)
		FontNew.SetBold(true)
		okno_dodaj.SetFont(FontNew)
	}
	okno_dodaj.Exec()

}

func (o *WydzialyOkno) Okno_mod_exec() {
	id, _ := strconv.Atoi(o.tree.CurrentItem().Text(0))
	okno_dodaj := widgets.NewQDialog(o, 0)
	okno_dodaj.SetWindowTitle("Modyfikuj Wydział")
	layoutV := widgets.NewQVBoxLayout()
	okno_dodaj.SetLayout(layoutV)
	label := widgets.NewQLabel(okno_dodaj, 0)
	label.SetText("Nazwa Wydziału: ")
	nazwa := widgets.NewQLineEdit(okno_dodaj)
	nazwa.SetText(o.tree.CurrentItem().Text(1))
	buttondodaj := widgets.NewQPushButton2("Dodaj", okno_dodaj)
	buttondodaj.ConnectClicked(func(checked bool) {
		if len(nazwa.Text()) > 0 {
			newG := mstr.Grupy{}
			newG.Id_grupy = id
			newG.Nazwa_grupy = nazwa.Text()
			odp := rclient.PutGrupa(newG)
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
			widgets.QMessageBox_Warning(okno_dodaj, "Uwaga!!", "Nazwa  nie może być pusta !",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

		}

	})
	layoutV.AddWidget(label, 0, core.Qt__AlignCenter)
	layoutV.AddWidget(nazwa, 0, core.Qt__AlignCenter)
	layoutV.AddWidget(buttondodaj, 0, core.Qt__AlignCenter)
	if fontsize > 0 {
		FontNew := gui.NewQFont2("verdana", fontsize, 1, false)
		FontNew.SetBold(true)
		okno_dodaj.SetFont(FontNew)
	}
	okno_dodaj.Exec()

}

func (p *Pilot) wydzialyProgram() {

	p.SetVisible(false)
	okno := NewWydzialyOkno(p, 0)
	p.okno_logowania.SetAttribute(core.Qt__WA_DeleteOnClose, true)
	okno.SetWindowTitle("Wydziały")
	okno.tree = widgets.NewQTreeWidget(okno)
	okno.tree.SetSortingEnabled(true)
	//SetSortingEnabled(true)
	okno.tree.SetRootIsDecorated(false)

	okno.tree.SetHeaderLabels([]string{"Id", "Nazwa Wydziału"})
	okno.tree.Header().SetSectionResizeMode(3)
	okno.LoadData()
	okno.SetCentralWidget(okno.tree)
	toolbar := widgets.NewQToolBar2(okno)

	addAction := toolbar.AddAction("Dadaj Wydział")
	addAction.SetIcon(gui.NewQIcon5(":/qml/dodaj.png"))
	addAction.ConnectTriggered(func(checked bool) {
		okno.Okno_dodaj_exec()

	})
	edytujAction := toolbar.AddAction("Modyfikuj Wydział")
	edytujAction.SetIcon(gui.NewQIcon5(":/qml/modyfikuj.png"))
	edytujAction.ConnectTriggered(func(checked bool) {
		okno.Okno_mod_exec()
	})
	usunAction := toolbar.AddAction("Usun Wydział")
	usunAction.SetIcon(gui.NewQIcon5(":/qml/usun.png"))
	usunAction.ConnectTriggered(func(checked bool) {
		if (okno.tree.CurrentItem().IsHidden() == false) && (okno.tree.CurrentItem().Type() > 0) {
			okno.Okno_usun_exec()

		}

	})

	okno.AddToolBar(core.Qt__TopToolBarArea, toolbar)

	okno.SetAttribute(core.Qt__WA_DeleteOnClose, true)

	okno.SetMinimumSize2(650, 400)

	okno.ConnectCloseEvent(func(e *gui.QCloseEvent) {
		p.Show()

	})

	if fontsize > 0 {
		FontNew := gui.NewQFont2("verdana", fontsize, 1, false)
		FontNew.SetBold(false)
		okno.SetFont(FontNew)
		okno.tree.SetFont(FontNew)
	}
	okno.ShowMaximized()
}

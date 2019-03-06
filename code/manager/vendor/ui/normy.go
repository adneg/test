package ui

import (
	//"fmt"

	"mstr"
	"rclient"
	"strconv"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	//"gitlab.top/ksala/pamietnik_hasel/restclient"
)

type NormyOkno struct {
	//slots
	//chan
	// do polaczenia jak zachowac sie podczas utraty polaczenia
	//varibles
	widgets.QMainWindow
}

func (p *Pilot) normyProgram() {

	//	p.Hide()
	stat, no := rclient.GetWo()

	if stat == false {
		widgets.QMessageBox_Warning(p, "Uwaga!!", "Nie udało się połączyć z serwerem",
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

		return

	}
	stat, nw := rclient.GetWw()
	if stat == false {

		widgets.QMessageBox_Warning(p, "Uwaga!!", "Nie udało się połączyć z serwerem",
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		return

	}
	okno := widgets.NewQDialog(p, 0)
	widgets.NewQMainWindow(p, 0)
	okno.SetWindowTitle("Normy i Wskaźniki")

	okno.SetAttribute(core.Qt__WA_DeleteOnClose, true)

	//	okno.SetMinimumSize2(650, 400)

	//	okno.ConnectCloseEvent(func(e *gui.QCloseEvent) {
	//		p.Show()

	//	})
	grid := widgets.NewQGridLayout(nil)
	//	okno.SetLayout(grid)
	ww_label := widgets.NewQLabel(okno, 0)
	ww_label.SetText("Współczynnik W: ")
	ww_textline := widgets.NewQLineEdit(okno)

	regex := core.NewQRegExp()
	regex.SetPattern("([0-9]|[1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9]|[1-9][0-9][0-9][0-9][0-9]||[1-9][0-9][0-9][0-9][0-9][0-9])([,][0-9][0-9][0-9][0-9])$")
	validator := gui.NewQRegExpValidator2(regex, ww_textline)
	ww_textline.SetValidator(validator)
	//	ww_textline.SetText("")
	ww_textline.SetText(ConverdotTocomma(nw.WspW))

	wo_label := widgets.NewQLabel(okno, 0)
	wo_label.SetText("Norma Osobowa: ")
	wo_textline := widgets.NewQLineEdit(okno)

	regex2 := core.NewQRegExp()
	regex2.SetPattern("([0-9]|[1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9]|[1-9][0-9][0-9][0-9][0-9]||[1-9][0-9][0-9][0-9][0-9][0-9])([,][0-9][0-9][0-9][0-9])$")
	//	regex2.SetPattern("([1-9]|[1-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9]|[1-9][0-9][0-9][0-9][0-9]||[1-9][0-9][0-9][0-9][0-9][0-9])$")
	validator2 := gui.NewQRegExpValidator2(regex2, wo_textline)
	wo_textline.SetValidator(validator2)

	wo_textline.SetText(ConverdotTocomma(no.Norma_osobowa))

	grid.AddWidget(wo_label, 0, 1, 0)
	grid.AddWidget(wo_textline, 0, 2, 0)
	grid.AddWidget(ww_label, 1, 1, 0)
	grid.AddWidget(ww_textline, 1, 2, 0)
	button := widgets.NewQPushButton2("Aktualizuj", okno)
	button.ConnectClicked(func(checked bool) {
		new_ww := mstr.Wspw{}
		valfloat64, _ := strconv.ParseFloat(ConvercommaTodot(ww_textline.Text()), 32)
		new_ww.WspW = float32(valfloat64)

		odp := rclient.PutWw(new_ww)
		if odp.Stat == false {
			widgets.QMessageBox_Warning(okno, "Uwaga!!", "Nie udało się połączyć z serwerem",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			okno.Close()
			return

		}

		new_wo := mstr.NormaOsobowa{}
		valno, _ := strconv.ParseFloat(ConvercommaTodot(wo_textline.Text()), 32)

		new_wo.Norma_osobowa = float32(valno)
		odp = rclient.PutWo(new_wo)
		if odp.Stat == false {
			widgets.QMessageBox_Warning(okno, "Uwaga!!", "Nie udało się połączyć z serwerem",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

			okno.Close()
			return
		}
		widgets.QMessageBox_Information(okno, "Uwaga!!", "Wartości zostały zmienione",
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

		okno.Close()

	})
	layoutV := widgets.NewQVBoxLayout()
	okno.SetLayout(layoutV)
	layoutV.AddLayout(grid, 0)
	layoutV.AddWidget(button, 0, 0)
	if fontsize > 0 {
		FontNew := gui.NewQFont2("verdana", fontsize, 1, false)
		FontNew.SetBold(false)
		okno.SetFont(FontNew)

	}
	okno.Exec()
}

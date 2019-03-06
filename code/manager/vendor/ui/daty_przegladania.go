package ui

import (
	"fmt"
	//	"mstr"
	"rclient"
	"strconv"
	"strings"

	//	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	//"gitlab.top/ksala/pamietnik_hasel/restclient"
)

func (p *Pilot) datyOkno() {
	okno := widgets.NewQDialog(p, 0)
	okno.SetAttribute(core.Qt__WA_DeleteOnClose, true)
	okno.SetWindowTitle("Wyznacz Zakres")
	grid := widgets.NewQGridLayout(nil)
	layoutH := widgets.NewQHBoxLayout()
	layoutV := widgets.NewQVBoxLayout()
	okno.SetLayout(layoutV)
	labelho := widgets.NewQLabel(okno, 0)
	labelho.SetText("Od: ")
	labelhn := widgets.NewQLabel(okno, 0)
	data_od := widgets.NewQDateEdit(p)
	//	core.NewQDate().CurrentDate()
	data_do := widgets.NewQDateEdit(p)
	data_od.SetDate(core.NewQDate().CurrentDate())
	//	data_od.SetEchoMode(2)
	data_od.SetMaximumDate(core.NewQDate().CurrentDate())
	data_od.ConnectDateChanged(func(c *core.QDate) {
		data_do.SetMinimumDate(data_od.Date())
	})

	data_do.SetDate(core.NewQDate().CurrentDate())
	data_do.ConnectDateChanged(func(c *core.QDate) {
		data_od.SetMaximumDate(data_do.Date())
		data_od.SetMinimumDate(data_do.Date().AddDays(-31))
	})
	data_do.SetMaximumDate(core.NewQDate().CurrentDate())
	data_od.SetMinimumDate(data_do.Date().AddDays(-31))
	data_do.SetCalendarPopup(true)
	data_od.SetCalendarPopup(true)
	//	data_do.SetEchoMode(2)
	labelhn.SetText("Do: ")
	grid.AddWidget(labelho, 0, 0, 0)
	grid.AddWidget(data_od, 0, 1, 0)
	grid.AddWidget(labelhn, 1, 0, 0)
	grid.AddWidget(data_do, 1, 1, 0)

	buttonWyznacz := widgets.NewQPushButton2("Wyznacz", okno)
	buttonWyznacz.ConnectClicked(func(checked bool) {
		p.treeRecord.Clear()
		od := data_od.Date().ToString("yyyy-MM-dd") + " 00:00:00"
		//00:00:00
		do := data_do.Date().ToString("yyyy-MM-dd") + " 23:59:59"
		stat, daneSJ := rclient.GetData2(od, do)

		if stat {
			go func() {
				for _, v := range daneSJ {
					//					time.Sleep()

					//	itemP.SetText(0, strconv.Itoa(v.Id))
					//	itemP.SetText(1, v.Nazwa_grupy)
					//	itemP.SetText(2, ConverdotTocomma(v.Stan_osobowy))
					//	itemP.SetText(3, v.Nazwa_produktu)
					//	itemP.SetText(4, strconv.Itoa(v.Ilosc))
					//	itemP.SetText(5, strconv.Itoa(v.Norma_osobowa))
					//	itemP.SetText(6, ConverdotTocomma(v.Wsp_produktu))
					//	itemP.SetText(7, ConverdotTocomma(v.WspW))
					//	itemP.SetText(8, v.Data.Format("2006-01-02"))
					//	itemP.SetText(9, v.Data.Format("15:04"))

					//					time.Sleep(80 * time.Millisecond)
					go p.AddTreeItem(
						strconv.Itoa(v.Id),
						v.Nazwa_grupy,
						ConverdotTocomma(v.Stan_osobowy),
						v.Nazwa_produktu,
						strconv.Itoa(v.Ilosc),
						ConverdotTocomma(v.Norma_osobowa),
						ConverdotTocomma(v.Wsp_produktu),
						ConverdotTocomma(v.WspW),
						v.Data.Format("2006-01-02"),
						v.Data.Format("15:04"))
					//					fmt.Println("jeden")

				}
			}()

		}
		okno.Close()
		//		daneSJ, _ = restclient.PobierzRekordy("/ph/getall")
		//		for _, v := range daneSJ {
		//			p.AddTreeItem(v)

	})
	buttonAnuluj := widgets.NewQPushButton2("Anuluj", okno)
	buttonAnuluj.ConnectClicked(func(checked bool) {
		okno.Close()
	})
	layoutH.AddWidget(buttonWyznacz, 0, 0)
	layoutH.AddWidget(buttonAnuluj, 0, 0)
	layoutV.AddLayout(grid, 0)
	layoutV.AddLayout(layoutH, 0)
	okno.Exec()
}

func (p *Pilot) addTreeItem(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9 string) {
	itemP := widgets.NewQTreeWidgetItem3(p.treeRecord, 1)

	itemP.SetText(0, p0)
	itemP.SetText(1, p1)
	itemP.SetText(2, p2)
	itemP.SetText(3, p3)
	itemP.SetText(4, p4)
	itemP.SetText(5, p5)
	itemP.SetText(6, p6)
	itemP.SetText(7, p7)
	itemP.SetText(8, p8)
	itemP.SetText(9, p9)

	//	p.treeRecord.SetCurrentItem(itemP)

}

func ConverdotTocomma(v float32) string {
	//fmt.Sprintf("%.2f"
	value := fmt.Sprintf("%.4f", v)
	return strings.Replace(value, ".", ",", -1)

}

func Converfloattostring(v float32) string {
	//fmt.Sprintf("%.2f"
	value := fmt.Sprintf("%.4f", v)
	return value

}

func ConvercommaTodot(v string) string {
	//fmt.Sprintf("%.2f"
	//	value := fmt.Sprintf("%.2f", v)
	return strings.Replace(v, ",", ".", -1)

}

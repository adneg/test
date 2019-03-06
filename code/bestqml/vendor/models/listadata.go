package models

import "github.com/therecipe/qt/core"
import "rclient"

//import "fmt"

var (
	EM *ElementyModel
)

const (
	ElementGName = int(core.Qt__UserRole) + 1
	ElementStan  = int(core.Qt__UserRole) + 2
	ElementPName = int(core.Qt__UserRole) + 3
	ElementIlosc = int(core.Qt__UserRole) + 4
	ElementData  = int(core.Qt__UserRole) + 5
	ElementId    = int(core.Qt__UserRole) + 6
)

type ElementyModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Element               `property:"elementitem"`

	_ func(*Element) `slot:"addElement"`
	_ func(row int, elementGName string,
		elementStan float32, elementPName string,
		elementIlosc int, elementData string, elementId int) `slot:"editElement"`
	_ func(row int)  `slot:"removeElement"`
	_ func(idin int) `slot:"removeElementId"`
}

type Element struct {
	core.QObject

	_ string  `property:"elementGName"`
	_ float32 `property:"elementStan"`
	_ string  `property:"elementPName"`
	_ int     `property:"elementIlosc"`
	_ string  `property:"elementData"`
	_ int     `property:"elementId"`
}

func init() {
	Element_QRegisterMetaType()
	EM = NewElementyModel(nil)
}

func (m *ElementyModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		ElementGName: core.NewQByteArray2("elementGName", len("elementGName")),
		ElementStan:  core.NewQByteArray2("elementStan", len("elementStan")),
		ElementPName: core.NewQByteArray2("elementPName", len("elementPName")),
		ElementIlosc: core.NewQByteArray2("elementIlosc", len("elementIlosc")),
		ElementData:  core.NewQByteArray2("elementData", len("elementData")),
		ElementId:    core.NewQByteArray2("elementId", len("elementId")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddElement(m.addElement)
	m.ConnectEditElement(m.editElement)
	m.ConnectRemoveElement(m.removeElement)
	m.ConnectRemoveElementId(m.removeElementId)
}

func (m *ElementyModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Elementitem()) {
		return core.NewQVariant()
	}

	var p = m.Elementitem()[index.Row()]

	switch role {
	case ElementGName:
		{
			return core.NewQVariant14(p.ElementGName())
		}

	case ElementStan:
		{
			return core.NewQVariant13(p.ElementStan())

		}
	case ElementPName:
		{
			return core.NewQVariant14(p.ElementPName())
		}

	case ElementIlosc:
		{
			return core.NewQVariant7(p.ElementIlosc())
		}
	case ElementData:
		{
			return core.NewQVariant14(p.ElementData())
		}
	case ElementId:
		{
			return core.NewQVariant7(p.ElementId())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *ElementyModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Elementitem())
}

func (m *ElementyModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *ElementyModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *ElementyModel) addElement(p *Element) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Elementitem()), len(m.Elementitem()))

	m.SetElementitem(append(m.Elementitem(), p))
	//	m.SetElementitem(append([...]*Element{p}, m.Elementitem()...))
	m.EndInsertRows()

}

func (m *ElementyModel) editElement(row int, elementGName string,
	elementStan float32, elementPName string, elementIlosc int,
	elementData string, elementId int) {
	var p = m.Elementitem()[row]

	if elementGName != "" {
		p.SetElementGName(elementGName)
	}

	if elementStan != 0 {
		p.SetElementStan(elementStan)
	}
	if elementPName != "" {
		p.SetElementPName(elementPName)
	}
	if elementIlosc != 0 {
		p.SetElementIlosc(elementIlosc)
	}
	if elementData != "" {
		p.SetElementGName(elementData)
	}
	if elementId != 0 {
		p.SetElementId(elementId)
	}

	var pIndex = m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(pIndex, pIndex, []int{ElementGName, ElementStan, ElementPName, ElementIlosc, ElementData, ElementId})
}

func (m *ElementyModel) removeElement(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetElementitem(append(m.Elementitem()[:row], m.Elementitem()[row+1:]...))
	m.EndRemoveRows()
	m.Elementitem()
}

func (m *ElementyModel) removeElementId(idin int) {
	var row int
	row = -1
	for n, ele := range m.Elementitem() {
		if ele.ElementId() == idin {
			row = n

		}
	}
	if row == -1 {
		return
	}

	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetElementitem(append(m.Elementitem()[:row], m.Elementitem()[row+1:]...))
	m.EndRemoveRows()
	m.Elementitem()
}

func LoadElementy() {

	stat, data := rclient.GetData()
	if stat {

		for _, elem := range data {
			//			func() {
			p := NewElement(nil)
			p.SetElementId(elem.Id)
			p.SetElementGName(elem.Nazwa_grupy)
			p.SetElementPName(elem.Nazwa_produktu)
			p.SetElementStan(elem.Stan_osobowy)
			p.SetElementData(elem.Data.Format("2006-01-02 15:04"))
			p.SetElementIlosc(elem.Ilosc)
			EM.AddElement(p)
			//			}()

		}
	}
	//
	//

	//	p.SetElementIlosc(300)
	//	p.SetElementId(2222)
	//	p.SetElementData("2012-01-02 10:10")
	//	EM.AddElement(p)
	//	EM.rowCount()

}

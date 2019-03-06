package models

import "github.com/therecipe/qt/core"
import "rclient"

var (
	PM *ProduktyModel
)

const (
	ProduktName = int(core.Qt__UserRole) + 1
	ProduktW    = int(core.Qt__UserRole) + 2
	ProduktId   = int(core.Qt__UserRole) + 3
)

type ProduktyModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Produkt               `property:"produktitem"`

	_ func(*Produkt)                                                     `slot:"addProdukt"`
	_ func(row int, produktName string, produktW float32, produktId int) `slot:"editProdukt"`
	_ func(row int)                                                      `slot:"removeProdukt"`
}

type Produkt struct {
	core.QObject

	_ string  `property:"produktName"`
	_ float32 `property:"produktW"`
	_ int     `property:"produktId"`
}

func init() {
	Produkt_QRegisterMetaType()
	PM = NewProduktyModel(nil)
}

func (m *ProduktyModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		ProduktName: core.NewQByteArray2("produktName", len("produktName")),
		ProduktW:    core.NewQByteArray2("produktW", len("produktW")),
		ProduktId:   core.NewQByteArray2("produktId", len("produktId")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddProdukt(m.addProdukt)
	m.ConnectEditProdukt(m.editProdukt)
	m.ConnectRemoveProdukt(m.removeProdukt)
}

func (m *ProduktyModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Produktitem()) {
		return core.NewQVariant()
	}

	var p = m.Produktitem()[index.Row()]

	switch role {
	case ProduktName:
		{
			return core.NewQVariant14(p.ProduktName())
		}

	case ProduktW:
		{
			return core.NewQVariant13(p.ProduktW())

		}
	case ProduktId:
		{
			return core.NewQVariant7(p.ProduktId())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *ProduktyModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Produktitem())
}

func (m *ProduktyModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *ProduktyModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *ProduktyModel) addProdukt(p *Produkt) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Produktitem()), len(m.Produktitem()))
	m.SetProduktitem(append(m.Produktitem(), p))
	m.EndInsertRows()
}

func (m *ProduktyModel) editProdukt(row int, produktName string, produktW float32, produktId int) {
	var p = m.Produktitem()[row]

	if produktName != "" {
		p.SetProduktName(produktName)
	}

	if produktW != 0 {
		p.SetProduktW(produktW)
	}
	if produktId != 0 {
		p.SetProduktId(produktId)
	}
	var pIndex = m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(pIndex, pIndex, []int{ProduktName, ProduktW, ProduktId})
}

func (m *ProduktyModel) removeProdukt(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetProduktitem(append(m.Produktitem()[:row], m.Produktitem()[row+1:]...))
	m.EndRemoveRows()
	m.Produktitem()
}

func (m *ProduktyModel) RemoveAll() {
	m.BeginRemoveRows(core.NewQModelIndex(), 0, len(m.Produktitem()))
	m.SetProduktitem(m.Produktitem()[:0])
	m.EndRemoveRows()

}

func LoadProdukty() {

	//rclient.GetProdukty()
	stat, produkty := rclient.GetProdukty()
	if stat {
		for _, elem := range produkty {
			//			go func() {
			p := NewProdukt(nil)

			p.SetProduktId(elem.Id_produktu)
			p.SetProduktName(elem.Nazwa_produktu)
			p.SetProduktW(elem.Wsp_produktu)
			PM.addProdukt(p)

			//			}()

		}
	}

	//	p.SetProduktName("Spodnie")
	//	p.SetProduktW(3.22)
	//	p.SetProduktId(1)
	//	PM.addProdukt(p)
	//	p = NewProdukt(nil)
	//	p.SetProduktName("Koszula")
	//	p.SetProduktW(2.1)
	//	p.SetProduktId(2)
	//	PM.addProdukt(p)

}

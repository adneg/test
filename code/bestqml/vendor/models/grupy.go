package models

import "github.com/therecipe/qt/core"
import "rclient"

var (
	GM *GrupyModel
)

const (
	GrupaName = int(core.Qt__UserRole) + 1
	GrupaId   = int(core.Qt__UserRole) + 2
)

type GrupyModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Grupa                 `property:"grupaitem"`

	_ func(*Grupa)                                 `slot:"addGrupa"`
	_ func(row int, grupaName string, grupaId int) `slot:"editGrupa"`
	_ func(row int)                                `slot:"removeGrupa"`
}

type Grupa struct {
	core.QObject

	_ string `property:"grupaName"`
	_ int    `property:"grupaId"`
}

func init() {
	Grupa_QRegisterMetaType()
	GM = NewGrupyModel(nil)
}

func (m *GrupyModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		GrupaName: core.NewQByteArray2("grupaName", len("grupaName")),
		GrupaId:   core.NewQByteArray2("grupaId", len("grupaId")),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddGrupa(m.addGrupa)
	m.ConnectEditGrupa(m.editGrupa)
	m.ConnectRemoveGrupa(m.removeGrupa)

}

func (m *GrupyModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Grupaitem()) {
		return core.NewQVariant()
	}

	var p = m.Grupaitem()[index.Row()]

	switch role {
	case GrupaName:
		{
			return core.NewQVariant14(p.GrupaName())
		}

	case GrupaId:
		{
			return core.NewQVariant7(p.GrupaId())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *GrupyModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Grupaitem())
}

func (m *GrupyModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *GrupyModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *GrupyModel) addGrupa(p *Grupa) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Grupaitem()), len(m.Grupaitem()))
	m.SetGrupaitem(append(m.Grupaitem(), p))
	m.EndInsertRows()
}

func (m *GrupyModel) editGrupa(row int, grupaName string, grupaId int) {
	var p = m.Grupaitem()[row]

	if grupaName != "" {
		p.SetGrupaName(grupaName)
	}

	if grupaId != 0 {
		p.SetGrupaId(grupaId)
	}
	var pIndex = m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(pIndex, pIndex, []int{GrupaName, GrupaId})
}

func (m *GrupyModel) removeGrupa(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetGrupaitem(append(m.Grupaitem()[:row], m.Grupaitem()[row+1:]...))
	m.EndRemoveRows()
	//	m.Grupaitem()
}
func (m *GrupyModel) RemoveAll() {
	m.BeginRemoveRows(core.NewQModelIndex(), 0, len(m.Grupaitem()))
	m.SetGrupaitem(m.Grupaitem()[:0])
	m.EndRemoveRows()

}
func LoadGrupy() {

	stat, gr := rclient.GetGrupy()
	if stat {
		for _, elem := range gr {
			//			func() {
			p := NewGrupa(nil)
			p.SetGrupaId(elem.Id_grupy)
			p.SetGrupaName(elem.Nazwa_grupy)
			GM.addGrupa(p)

			//			}()

		}
	}

	//	var p = NewGrupa(nil)

}

package models

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type GrupyModel_ITF interface {
	std_core.QAbstractListModel_ITF
	GrupyModel_PTR() *GrupyModel
}

func (ptr *GrupyModel) GrupyModel_PTR() *GrupyModel {
	return ptr
}

func (ptr *GrupyModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *GrupyModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromGrupyModel(ptr GrupyModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.GrupyModel_PTR().Pointer()
	}
	return nil
}

func NewGrupyModelFromPointer(ptr unsafe.Pointer) (n *GrupyModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(GrupyModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *GrupyModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &GrupyModel{QAbstractListModel: *deduced}

		default:
			n = new(GrupyModel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackGrupyModeld7c2e5_Constructor
func callbackGrupyModeld7c2e5_Constructor(ptr unsafe.Pointer) {
	this := NewGrupyModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackGrupyModeld7c2e5_AddGrupa
func callbackGrupyModeld7c2e5_AddGrupa(ptr unsafe.Pointer, v0 unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addGrupa"); signal != nil {
		signal.(func(*Grupa))(NewGrupaFromPointer(v0))
	}

}

func (ptr *GrupyModel) ConnectAddGrupa(f func(v0 *Grupa)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addGrupa"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addGrupa", func(v0 *Grupa) {
				signal.(func(*Grupa))(v0)
				f(v0)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addGrupa", f)
		}
	}
}

func (ptr *GrupyModel) DisconnectAddGrupa() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addGrupa")
	}
}

func (ptr *GrupyModel) AddGrupa(v0 Grupa_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_AddGrupa(ptr.Pointer(), PointerFromGrupa(v0))
	}
}

//export callbackGrupyModeld7c2e5_EditGrupa
func callbackGrupyModeld7c2e5_EditGrupa(ptr unsafe.Pointer, row C.int, grupaName C.struct_Moc_PackedString, grupaId C.int) {
	if signal := qt.GetSignal(ptr, "editGrupa"); signal != nil {
		signal.(func(int, string, int))(int(int32(row)), cGoUnpackString(grupaName), int(int32(grupaId)))
	}

}

func (ptr *GrupyModel) ConnectEditGrupa(f func(row int, grupaName string, grupaId int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "editGrupa"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "editGrupa", func(row int, grupaName string, grupaId int) {
				signal.(func(int, string, int))(row, grupaName, grupaId)
				f(row, grupaName, grupaId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "editGrupa", f)
		}
	}
}

func (ptr *GrupyModel) DisconnectEditGrupa() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "editGrupa")
	}
}

func (ptr *GrupyModel) EditGrupa(row int, grupaName string, grupaId int) {
	if ptr.Pointer() != nil {
		var grupaNameC *C.char
		if grupaName != "" {
			grupaNameC = C.CString(grupaName)
			defer C.free(unsafe.Pointer(grupaNameC))
		}
		C.GrupyModeld7c2e5_EditGrupa(ptr.Pointer(), C.int(int32(row)), C.struct_Moc_PackedString{data: grupaNameC, len: C.longlong(len(grupaName))}, C.int(int32(grupaId)))
	}
}

//export callbackGrupyModeld7c2e5_RemoveGrupa
func callbackGrupyModeld7c2e5_RemoveGrupa(ptr unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "removeGrupa"); signal != nil {
		signal.(func(int))(int(int32(row)))
	}

}

func (ptr *GrupyModel) ConnectRemoveGrupa(f func(row int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeGrupa"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "removeGrupa", func(row int) {
				signal.(func(int))(row)
				f(row)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeGrupa", f)
		}
	}
}

func (ptr *GrupyModel) DisconnectRemoveGrupa() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeGrupa")
	}
}

func (ptr *GrupyModel) RemoveGrupa(row int) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_RemoveGrupa(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackGrupyModeld7c2e5_Roles
func callbackGrupyModeld7c2e5_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__roles_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__roles_newList())
		for k, v := range NewGrupyModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *GrupyModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "roles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "roles", func() map[int]*std_core.QByteArray {
				signal.(func() map[int]*std_core.QByteArray)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "roles", f)
		}
	}
}

func (ptr *GrupyModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *GrupyModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.GrupyModeld7c2e5_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *GrupyModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.GrupyModeld7c2e5_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackGrupyModeld7c2e5_SetRoles
func callbackGrupyModeld7c2e5_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewGrupyModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *GrupyModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRoles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", f)
		}
	}
}

func (ptr *GrupyModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *GrupyModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *GrupyModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackGrupyModeld7c2e5_RolesChanged
func callbackGrupyModeld7c2e5_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *GrupyModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.GrupyModeld7c2e5_ConnectRolesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rolesChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", f)
		}
	}
}

func (ptr *GrupyModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *GrupyModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackGrupyModeld7c2e5_Grupaitem
func callbackGrupyModeld7c2e5_Grupaitem(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "grupaitem"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__grupaitem_newList())
			for _, v := range signal.(func() []*Grupa)() {
				tmpList.__grupaitem_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__grupaitem_newList())
		for _, v := range NewGrupyModelFromPointer(ptr).GrupaitemDefault() {
			tmpList.__grupaitem_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *GrupyModel) ConnectGrupaitem(f func() []*Grupa) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "grupaitem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "grupaitem", func() []*Grupa {
				signal.(func() []*Grupa)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "grupaitem", f)
		}
	}
}

func (ptr *GrupyModel) DisconnectGrupaitem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "grupaitem")
	}
}

func (ptr *GrupyModel) Grupaitem() []*Grupa {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Grupa {
			out := make([]*Grupa, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__grupaitem_atList(i)
			}
			return out
		}(C.GrupyModeld7c2e5_Grupaitem(ptr.Pointer()))
	}
	return make([]*Grupa, 0)
}

func (ptr *GrupyModel) GrupaitemDefault() []*Grupa {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Grupa {
			out := make([]*Grupa, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__grupaitem_atList(i)
			}
			return out
		}(C.GrupyModeld7c2e5_GrupaitemDefault(ptr.Pointer()))
	}
	return make([]*Grupa, 0)
}

//export callbackGrupyModeld7c2e5_SetGrupaitem
func callbackGrupyModeld7c2e5_SetGrupaitem(ptr unsafe.Pointer, grupaitem C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setGrupaitem"); signal != nil {
		signal.(func([]*Grupa))(func(l C.struct_Moc_PackedList) []*Grupa {
			out := make([]*Grupa, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setGrupaitem_grupaitem_atList(i)
			}
			return out
		}(grupaitem))
	} else {
		NewGrupyModelFromPointer(ptr).SetGrupaitemDefault(func(l C.struct_Moc_PackedList) []*Grupa {
			out := make([]*Grupa, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setGrupaitem_grupaitem_atList(i)
			}
			return out
		}(grupaitem))
	}
}

func (ptr *GrupyModel) ConnectSetGrupaitem(f func(grupaitem []*Grupa)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGrupaitem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setGrupaitem", func(grupaitem []*Grupa) {
				signal.(func([]*Grupa))(grupaitem)
				f(grupaitem)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGrupaitem", f)
		}
	}
}

func (ptr *GrupyModel) DisconnectSetGrupaitem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGrupaitem")
	}
}

func (ptr *GrupyModel) SetGrupaitem(grupaitem []*Grupa) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_SetGrupaitem(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__setGrupaitem_grupaitem_newList())
			for _, v := range grupaitem {
				tmpList.__setGrupaitem_grupaitem_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *GrupyModel) SetGrupaitemDefault(grupaitem []*Grupa) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_SetGrupaitemDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__setGrupaitem_grupaitem_newList())
			for _, v := range grupaitem {
				tmpList.__setGrupaitem_grupaitem_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackGrupyModeld7c2e5_GrupaitemChanged
func callbackGrupyModeld7c2e5_GrupaitemChanged(ptr unsafe.Pointer, grupaitem C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "grupaitemChanged"); signal != nil {
		signal.(func([]*Grupa))(func(l C.struct_Moc_PackedList) []*Grupa {
			out := make([]*Grupa, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__grupaitemChanged_grupaitem_atList(i)
			}
			return out
		}(grupaitem))
	}

}

func (ptr *GrupyModel) ConnectGrupaitemChanged(f func(grupaitem []*Grupa)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "grupaitemChanged") {
			C.GrupyModeld7c2e5_ConnectGrupaitemChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "grupaitemChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "grupaitemChanged", func(grupaitem []*Grupa) {
				signal.(func([]*Grupa))(grupaitem)
				f(grupaitem)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "grupaitemChanged", f)
		}
	}
}

func (ptr *GrupyModel) DisconnectGrupaitemChanged() {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_DisconnectGrupaitemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "grupaitemChanged")
	}
}

func (ptr *GrupyModel) GrupaitemChanged(grupaitem []*Grupa) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_GrupaitemChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__grupaitemChanged_grupaitem_newList())
			for _, v := range grupaitem {
				tmpList.__grupaitemChanged_grupaitem_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func GrupyModel_QRegisterMetaType() int {
	return int(int32(C.GrupyModeld7c2e5_GrupyModeld7c2e5_QRegisterMetaType()))
}

func (ptr *GrupyModel) QRegisterMetaType() int {
	return int(int32(C.GrupyModeld7c2e5_GrupyModeld7c2e5_QRegisterMetaType()))
}

func GrupyModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.GrupyModeld7c2e5_GrupyModeld7c2e5_QRegisterMetaType2(typeNameC)))
}

func (ptr *GrupyModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.GrupyModeld7c2e5_GrupyModeld7c2e5_QRegisterMetaType2(typeNameC)))
}

func GrupyModel_QmlRegisterType() int {
	return int(int32(C.GrupyModeld7c2e5_GrupyModeld7c2e5_QmlRegisterType()))
}

func (ptr *GrupyModel) QmlRegisterType() int {
	return int(int32(C.GrupyModeld7c2e5_GrupyModeld7c2e5_QmlRegisterType()))
}

func GrupyModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.GrupyModeld7c2e5_GrupyModeld7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *GrupyModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.GrupyModeld7c2e5_GrupyModeld7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *GrupyModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.GrupyModeld7c2e5_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *GrupyModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *GrupyModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *GrupyModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.GrupyModeld7c2e5_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *GrupyModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *GrupyModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *GrupyModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.GrupyModeld7c2e5_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *GrupyModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *GrupyModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.GrupyModeld7c2e5___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *GrupyModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.GrupyModeld7c2e5___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *GrupyModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.GrupyModeld7c2e5___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *GrupyModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.GrupyModeld7c2e5___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *GrupyModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.GrupyModeld7c2e5___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *GrupyModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *GrupyModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.GrupyModeld7c2e5___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *GrupyModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.GrupyModeld7c2e5___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *GrupyModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.GrupyModeld7c2e5___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *GrupyModel) __roleNames_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___roleNames_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.GrupyModeld7c2e5___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *GrupyModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.GrupyModeld7c2e5___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *GrupyModel) __itemData_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___itemData_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.GrupyModeld7c2e5___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *GrupyModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.GrupyModeld7c2e5___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *GrupyModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.GrupyModeld7c2e5___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *GrupyModel) __match_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___match_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.GrupyModeld7c2e5___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *GrupyModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *GrupyModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.GrupyModeld7c2e5_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *GrupyModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *GrupyModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *GrupyModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.GrupyModeld7c2e5_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *GrupyModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *GrupyModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.GrupyModeld7c2e5___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *GrupyModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.GrupyModeld7c2e5___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *GrupyModel) __findChildren_newList2() unsafe.Pointer {
	return C.GrupyModeld7c2e5___findChildren_newList2(ptr.Pointer())
}

func (ptr *GrupyModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.GrupyModeld7c2e5___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *GrupyModel) __findChildren_newList3() unsafe.Pointer {
	return C.GrupyModeld7c2e5___findChildren_newList3(ptr.Pointer())
}

func (ptr *GrupyModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.GrupyModeld7c2e5___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *GrupyModel) __findChildren_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___findChildren_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.GrupyModeld7c2e5___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *GrupyModel) __children_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___children_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.GrupyModeld7c2e5___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *GrupyModel) __roles_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___roles_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.GrupyModeld7c2e5___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *GrupyModel) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.GrupyModeld7c2e5___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *GrupyModel) __setRoles_roles_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.GrupyModeld7c2e5___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *GrupyModel) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.GrupyModeld7c2e5___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *GrupyModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.GrupyModeld7c2e5___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *GrupyModel) __grupaitem_atList(i int) *Grupa {
	if ptr.Pointer() != nil {
		tmpValue := NewGrupaFromPointer(C.GrupyModeld7c2e5___grupaitem_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __grupaitem_setList(i Grupa_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___grupaitem_setList(ptr.Pointer(), PointerFromGrupa(i))
	}
}

func (ptr *GrupyModel) __grupaitem_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___grupaitem_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __setGrupaitem_grupaitem_atList(i int) *Grupa {
	if ptr.Pointer() != nil {
		tmpValue := NewGrupaFromPointer(C.GrupyModeld7c2e5___setGrupaitem_grupaitem_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __setGrupaitem_grupaitem_setList(i Grupa_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___setGrupaitem_grupaitem_setList(ptr.Pointer(), PointerFromGrupa(i))
	}
}

func (ptr *GrupyModel) __setGrupaitem_grupaitem_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___setGrupaitem_grupaitem_newList(ptr.Pointer())
}

func (ptr *GrupyModel) __grupaitemChanged_grupaitem_atList(i int) *Grupa {
	if ptr.Pointer() != nil {
		tmpValue := NewGrupaFromPointer(C.GrupyModeld7c2e5___grupaitemChanged_grupaitem_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *GrupyModel) __grupaitemChanged_grupaitem_setList(i Grupa_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5___grupaitemChanged_grupaitem_setList(ptr.Pointer(), PointerFromGrupa(i))
	}
}

func (ptr *GrupyModel) __grupaitemChanged_grupaitem_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5___grupaitemChanged_grupaitem_newList(ptr.Pointer())
}

func (ptr *GrupyModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.GrupyModeld7c2e5_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *GrupyModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *GrupyModel) ____roles_keyList_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *GrupyModel) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.GrupyModeld7c2e5_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *GrupyModel) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *GrupyModel) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *GrupyModel) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.GrupyModeld7c2e5_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *GrupyModel) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *GrupyModel) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.GrupyModeld7c2e5_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewGrupyModel(parent std_core.QObject_ITF) *GrupyModel {
	tmpValue := NewGrupyModelFromPointer(C.GrupyModeld7c2e5_NewGrupyModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackGrupyModeld7c2e5_DestroyGrupyModel
func callbackGrupyModeld7c2e5_DestroyGrupyModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~GrupyModel"); signal != nil {
		signal.(func())()
	} else {
		NewGrupyModelFromPointer(ptr).DestroyGrupyModelDefault()
	}
}

func (ptr *GrupyModel) ConnectDestroyGrupyModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~GrupyModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~GrupyModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~GrupyModel", f)
		}
	}
}

func (ptr *GrupyModel) DisconnectDestroyGrupyModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~GrupyModel")
	}
}

func (ptr *GrupyModel) DestroyGrupyModel() {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_DestroyGrupyModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *GrupyModel) DestroyGrupyModelDefault() {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_DestroyGrupyModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackGrupyModeld7c2e5_DropMimeData
func callbackGrupyModeld7c2e5_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *GrupyModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_Index
func callbackGrupyModeld7c2e5_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewGrupyModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *GrupyModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.GrupyModeld7c2e5_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackGrupyModeld7c2e5_Sibling
func callbackGrupyModeld7c2e5_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewGrupyModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *GrupyModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.GrupyModeld7c2e5_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackGrupyModeld7c2e5_Flags
func callbackGrupyModeld7c2e5_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewGrupyModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *GrupyModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.GrupyModeld7c2e5_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackGrupyModeld7c2e5_InsertColumns
func callbackGrupyModeld7c2e5_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *GrupyModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_InsertRows
func callbackGrupyModeld7c2e5_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *GrupyModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_MoveColumns
func callbackGrupyModeld7c2e5_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *GrupyModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_MoveRows
func callbackGrupyModeld7c2e5_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *GrupyModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_RemoveColumns
func callbackGrupyModeld7c2e5_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *GrupyModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_RemoveRows
func callbackGrupyModeld7c2e5_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *GrupyModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_SetData
func callbackGrupyModeld7c2e5_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *GrupyModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_SetHeaderData
func callbackGrupyModeld7c2e5_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *GrupyModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_SetItemData
func callbackGrupyModeld7c2e5_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewGrupyModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *GrupyModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_Submit
func callbackGrupyModeld7c2e5_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *GrupyModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_ColumnsAboutToBeInserted
func callbackGrupyModeld7c2e5_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackGrupyModeld7c2e5_ColumnsAboutToBeMoved
func callbackGrupyModeld7c2e5_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackGrupyModeld7c2e5_ColumnsAboutToBeRemoved
func callbackGrupyModeld7c2e5_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackGrupyModeld7c2e5_ColumnsInserted
func callbackGrupyModeld7c2e5_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackGrupyModeld7c2e5_ColumnsMoved
func callbackGrupyModeld7c2e5_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackGrupyModeld7c2e5_ColumnsRemoved
func callbackGrupyModeld7c2e5_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackGrupyModeld7c2e5_DataChanged
func callbackGrupyModeld7c2e5_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackGrupyModeld7c2e5_FetchMore
func callbackGrupyModeld7c2e5_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewGrupyModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *GrupyModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackGrupyModeld7c2e5_HeaderDataChanged
func callbackGrupyModeld7c2e5_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackGrupyModeld7c2e5_LayoutAboutToBeChanged
func callbackGrupyModeld7c2e5_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackGrupyModeld7c2e5_LayoutChanged
func callbackGrupyModeld7c2e5_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackGrupyModeld7c2e5_ModelAboutToBeReset
func callbackGrupyModeld7c2e5_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackGrupyModeld7c2e5_ModelReset
func callbackGrupyModeld7c2e5_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackGrupyModeld7c2e5_ResetInternalData
func callbackGrupyModeld7c2e5_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewGrupyModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *GrupyModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackGrupyModeld7c2e5_Revert
func callbackGrupyModeld7c2e5_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewGrupyModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *GrupyModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_RevertDefault(ptr.Pointer())
	}
}

//export callbackGrupyModeld7c2e5_RowsAboutToBeInserted
func callbackGrupyModeld7c2e5_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackGrupyModeld7c2e5_RowsAboutToBeMoved
func callbackGrupyModeld7c2e5_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackGrupyModeld7c2e5_RowsAboutToBeRemoved
func callbackGrupyModeld7c2e5_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackGrupyModeld7c2e5_RowsInserted
func callbackGrupyModeld7c2e5_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackGrupyModeld7c2e5_RowsMoved
func callbackGrupyModeld7c2e5_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackGrupyModeld7c2e5_RowsRemoved
func callbackGrupyModeld7c2e5_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackGrupyModeld7c2e5_Sort
func callbackGrupyModeld7c2e5_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewGrupyModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *GrupyModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackGrupyModeld7c2e5_RoleNames
func callbackGrupyModeld7c2e5_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewGrupyModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *GrupyModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.GrupyModeld7c2e5_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackGrupyModeld7c2e5_ItemData
func callbackGrupyModeld7c2e5_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__itemData_newList())
		for k, v := range NewGrupyModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *GrupyModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.GrupyModeld7c2e5_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackGrupyModeld7c2e5_MimeData
func callbackGrupyModeld7c2e5_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewGrupyModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewGrupyModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *GrupyModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.GrupyModeld7c2e5_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackGrupyModeld7c2e5_Buddy
func callbackGrupyModeld7c2e5_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewGrupyModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *GrupyModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.GrupyModeld7c2e5_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackGrupyModeld7c2e5_Parent
func callbackGrupyModeld7c2e5_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewGrupyModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *GrupyModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.GrupyModeld7c2e5_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackGrupyModeld7c2e5_Match
func callbackGrupyModeld7c2e5_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewGrupyModelFromPointer(NewGrupyModelFromPointer(nil).__match_newList())
		for _, v := range NewGrupyModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *GrupyModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewGrupyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.GrupyModeld7c2e5_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackGrupyModeld7c2e5_Span
func callbackGrupyModeld7c2e5_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewGrupyModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *GrupyModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.GrupyModeld7c2e5_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackGrupyModeld7c2e5_MimeTypes
func callbackGrupyModeld7c2e5_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewGrupyModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *GrupyModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.GrupyModeld7c2e5_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackGrupyModeld7c2e5_Data
func callbackGrupyModeld7c2e5_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewGrupyModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *GrupyModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.GrupyModeld7c2e5_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackGrupyModeld7c2e5_HeaderData
func callbackGrupyModeld7c2e5_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewGrupyModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *GrupyModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.GrupyModeld7c2e5_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackGrupyModeld7c2e5_SupportedDragActions
func callbackGrupyModeld7c2e5_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewGrupyModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *GrupyModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.GrupyModeld7c2e5_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackGrupyModeld7c2e5_SupportedDropActions
func callbackGrupyModeld7c2e5_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewGrupyModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *GrupyModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.GrupyModeld7c2e5_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackGrupyModeld7c2e5_CanDropMimeData
func callbackGrupyModeld7c2e5_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *GrupyModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_CanFetchMore
func callbackGrupyModeld7c2e5_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *GrupyModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_HasChildren
func callbackGrupyModeld7c2e5_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *GrupyModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_ColumnCount
func callbackGrupyModeld7c2e5_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewGrupyModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *GrupyModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.GrupyModeld7c2e5_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackGrupyModeld7c2e5_RowCount
func callbackGrupyModeld7c2e5_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewGrupyModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *GrupyModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.GrupyModeld7c2e5_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackGrupyModeld7c2e5_Event
func callbackGrupyModeld7c2e5_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *GrupyModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_EventFilter
func callbackGrupyModeld7c2e5_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupyModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *GrupyModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.GrupyModeld7c2e5_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackGrupyModeld7c2e5_ChildEvent
func callbackGrupyModeld7c2e5_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewGrupyModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *GrupyModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackGrupyModeld7c2e5_ConnectNotify
func callbackGrupyModeld7c2e5_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewGrupyModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *GrupyModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackGrupyModeld7c2e5_CustomEvent
func callbackGrupyModeld7c2e5_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewGrupyModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *GrupyModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackGrupyModeld7c2e5_DeleteLater
func callbackGrupyModeld7c2e5_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewGrupyModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *GrupyModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackGrupyModeld7c2e5_Destroyed
func callbackGrupyModeld7c2e5_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackGrupyModeld7c2e5_DisconnectNotify
func callbackGrupyModeld7c2e5_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewGrupyModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *GrupyModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackGrupyModeld7c2e5_ObjectNameChanged
func callbackGrupyModeld7c2e5_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackGrupyModeld7c2e5_TimerEvent
func callbackGrupyModeld7c2e5_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewGrupyModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *GrupyModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.GrupyModeld7c2e5_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type Produkt_ITF interface {
	std_core.QObject_ITF
	Produkt_PTR() *Produkt
}

func (ptr *Produkt) Produkt_PTR() *Produkt {
	return ptr
}

func (ptr *Produkt) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Produkt) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromProdukt(ptr Produkt_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Produkt_PTR().Pointer()
	}
	return nil
}

func NewProduktFromPointer(ptr unsafe.Pointer) (n *Produkt) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Produkt)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Produkt:
			n = deduced

		case *std_core.QObject:
			n = &Produkt{QObject: *deduced}

		default:
			n = new(Produkt)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackProduktd7c2e5_Constructor
func callbackProduktd7c2e5_Constructor(ptr unsafe.Pointer) {
	this := NewProduktFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackProduktd7c2e5_ProduktName
func callbackProduktd7c2e5_ProduktName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "produktName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewProduktFromPointer(ptr).ProduktNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Produkt) ConnectProduktName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "produktName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "produktName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "produktName", f)
		}
	}
}

func (ptr *Produkt) DisconnectProduktName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "produktName")
	}
}

func (ptr *Produkt) ProduktName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Produktd7c2e5_ProduktName(ptr.Pointer()))
	}
	return ""
}

func (ptr *Produkt) ProduktNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Produktd7c2e5_ProduktNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackProduktd7c2e5_SetProduktName
func callbackProduktd7c2e5_SetProduktName(ptr unsafe.Pointer, produktName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setProduktName"); signal != nil {
		signal.(func(string))(cGoUnpackString(produktName))
	} else {
		NewProduktFromPointer(ptr).SetProduktNameDefault(cGoUnpackString(produktName))
	}
}

func (ptr *Produkt) ConnectSetProduktName(f func(produktName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setProduktName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setProduktName", func(produktName string) {
				signal.(func(string))(produktName)
				f(produktName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setProduktName", f)
		}
	}
}

func (ptr *Produkt) DisconnectSetProduktName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setProduktName")
	}
}

func (ptr *Produkt) SetProduktName(produktName string) {
	if ptr.Pointer() != nil {
		var produktNameC *C.char
		if produktName != "" {
			produktNameC = C.CString(produktName)
			defer C.free(unsafe.Pointer(produktNameC))
		}
		C.Produktd7c2e5_SetProduktName(ptr.Pointer(), C.struct_Moc_PackedString{data: produktNameC, len: C.longlong(len(produktName))})
	}
}

func (ptr *Produkt) SetProduktNameDefault(produktName string) {
	if ptr.Pointer() != nil {
		var produktNameC *C.char
		if produktName != "" {
			produktNameC = C.CString(produktName)
			defer C.free(unsafe.Pointer(produktNameC))
		}
		C.Produktd7c2e5_SetProduktNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: produktNameC, len: C.longlong(len(produktName))})
	}
}

//export callbackProduktd7c2e5_ProduktNameChanged
func callbackProduktd7c2e5_ProduktNameChanged(ptr unsafe.Pointer, produktName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "produktNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(produktName))
	}

}

func (ptr *Produkt) ConnectProduktNameChanged(f func(produktName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "produktNameChanged") {
			C.Produktd7c2e5_ConnectProduktNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "produktNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "produktNameChanged", func(produktName string) {
				signal.(func(string))(produktName)
				f(produktName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "produktNameChanged", f)
		}
	}
}

func (ptr *Produkt) DisconnectProduktNameChanged() {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_DisconnectProduktNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "produktNameChanged")
	}
}

func (ptr *Produkt) ProduktNameChanged(produktName string) {
	if ptr.Pointer() != nil {
		var produktNameC *C.char
		if produktName != "" {
			produktNameC = C.CString(produktName)
			defer C.free(unsafe.Pointer(produktNameC))
		}
		C.Produktd7c2e5_ProduktNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: produktNameC, len: C.longlong(len(produktName))})
	}
}

//export callbackProduktd7c2e5_ProduktW
func callbackProduktd7c2e5_ProduktW(ptr unsafe.Pointer) C.float {
	if signal := qt.GetSignal(ptr, "produktW"); signal != nil {
		return C.float(signal.(func() float32)())
	}

	return C.float(NewProduktFromPointer(ptr).ProduktWDefault())
}

func (ptr *Produkt) ConnectProduktW(f func() float32) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "produktW"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "produktW", func() float32 {
				signal.(func() float32)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "produktW", f)
		}
	}
}

func (ptr *Produkt) DisconnectProduktW() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "produktW")
	}
}

func (ptr *Produkt) ProduktW() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Produktd7c2e5_ProduktW(ptr.Pointer()))
	}
	return 0
}

func (ptr *Produkt) ProduktWDefault() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Produktd7c2e5_ProduktWDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackProduktd7c2e5_SetProduktW
func callbackProduktd7c2e5_SetProduktW(ptr unsafe.Pointer, produktW C.float) {
	if signal := qt.GetSignal(ptr, "setProduktW"); signal != nil {
		signal.(func(float32))(float32(produktW))
	} else {
		NewProduktFromPointer(ptr).SetProduktWDefault(float32(produktW))
	}
}

func (ptr *Produkt) ConnectSetProduktW(f func(produktW float32)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setProduktW"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setProduktW", func(produktW float32) {
				signal.(func(float32))(produktW)
				f(produktW)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setProduktW", f)
		}
	}
}

func (ptr *Produkt) DisconnectSetProduktW() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setProduktW")
	}
}

func (ptr *Produkt) SetProduktW(produktW float32) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_SetProduktW(ptr.Pointer(), C.float(produktW))
	}
}

func (ptr *Produkt) SetProduktWDefault(produktW float32) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_SetProduktWDefault(ptr.Pointer(), C.float(produktW))
	}
}

//export callbackProduktd7c2e5_ProduktWChanged
func callbackProduktd7c2e5_ProduktWChanged(ptr unsafe.Pointer, produktW C.float) {
	if signal := qt.GetSignal(ptr, "produktWChanged"); signal != nil {
		signal.(func(float32))(float32(produktW))
	}

}

func (ptr *Produkt) ConnectProduktWChanged(f func(produktW float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "produktWChanged") {
			C.Produktd7c2e5_ConnectProduktWChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "produktWChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "produktWChanged", func(produktW float32) {
				signal.(func(float32))(produktW)
				f(produktW)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "produktWChanged", f)
		}
	}
}

func (ptr *Produkt) DisconnectProduktWChanged() {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_DisconnectProduktWChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "produktWChanged")
	}
}

func (ptr *Produkt) ProduktWChanged(produktW float32) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_ProduktWChanged(ptr.Pointer(), C.float(produktW))
	}
}

//export callbackProduktd7c2e5_ProduktId
func callbackProduktd7c2e5_ProduktId(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "produktId"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewProduktFromPointer(ptr).ProduktIdDefault()))
}

func (ptr *Produkt) ConnectProduktId(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "produktId"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "produktId", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "produktId", f)
		}
	}
}

func (ptr *Produkt) DisconnectProduktId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "produktId")
	}
}

func (ptr *Produkt) ProduktId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Produktd7c2e5_ProduktId(ptr.Pointer())))
	}
	return 0
}

func (ptr *Produkt) ProduktIdDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Produktd7c2e5_ProduktIdDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackProduktd7c2e5_SetProduktId
func callbackProduktd7c2e5_SetProduktId(ptr unsafe.Pointer, produktId C.int) {
	if signal := qt.GetSignal(ptr, "setProduktId"); signal != nil {
		signal.(func(int))(int(int32(produktId)))
	} else {
		NewProduktFromPointer(ptr).SetProduktIdDefault(int(int32(produktId)))
	}
}

func (ptr *Produkt) ConnectSetProduktId(f func(produktId int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setProduktId"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setProduktId", func(produktId int) {
				signal.(func(int))(produktId)
				f(produktId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setProduktId", f)
		}
	}
}

func (ptr *Produkt) DisconnectSetProduktId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setProduktId")
	}
}

func (ptr *Produkt) SetProduktId(produktId int) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_SetProduktId(ptr.Pointer(), C.int(int32(produktId)))
	}
}

func (ptr *Produkt) SetProduktIdDefault(produktId int) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_SetProduktIdDefault(ptr.Pointer(), C.int(int32(produktId)))
	}
}

//export callbackProduktd7c2e5_ProduktIdChanged
func callbackProduktd7c2e5_ProduktIdChanged(ptr unsafe.Pointer, produktId C.int) {
	if signal := qt.GetSignal(ptr, "produktIdChanged"); signal != nil {
		signal.(func(int))(int(int32(produktId)))
	}

}

func (ptr *Produkt) ConnectProduktIdChanged(f func(produktId int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "produktIdChanged") {
			C.Produktd7c2e5_ConnectProduktIdChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "produktIdChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "produktIdChanged", func(produktId int) {
				signal.(func(int))(produktId)
				f(produktId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "produktIdChanged", f)
		}
	}
}

func (ptr *Produkt) DisconnectProduktIdChanged() {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_DisconnectProduktIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "produktIdChanged")
	}
}

func (ptr *Produkt) ProduktIdChanged(produktId int) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_ProduktIdChanged(ptr.Pointer(), C.int(int32(produktId)))
	}
}

func Produkt_QRegisterMetaType() int {
	return int(int32(C.Produktd7c2e5_Produktd7c2e5_QRegisterMetaType()))
}

func (ptr *Produkt) QRegisterMetaType() int {
	return int(int32(C.Produktd7c2e5_Produktd7c2e5_QRegisterMetaType()))
}

func Produkt_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Produktd7c2e5_Produktd7c2e5_QRegisterMetaType2(typeNameC)))
}

func (ptr *Produkt) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Produktd7c2e5_Produktd7c2e5_QRegisterMetaType2(typeNameC)))
}

func Produkt_QmlRegisterType() int {
	return int(int32(C.Produktd7c2e5_Produktd7c2e5_QmlRegisterType()))
}

func (ptr *Produkt) QmlRegisterType() int {
	return int(int32(C.Produktd7c2e5_Produktd7c2e5_QmlRegisterType()))
}

func Produkt_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Produktd7c2e5_Produktd7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Produkt) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Produktd7c2e5_Produktd7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Produkt) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Produktd7c2e5___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Produkt) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Produkt) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Produktd7c2e5___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Produkt) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Produktd7c2e5___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Produkt) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Produkt) __findChildren_newList2() unsafe.Pointer {
	return C.Produktd7c2e5___findChildren_newList2(ptr.Pointer())
}

func (ptr *Produkt) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Produktd7c2e5___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Produkt) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Produkt) __findChildren_newList3() unsafe.Pointer {
	return C.Produktd7c2e5___findChildren_newList3(ptr.Pointer())
}

func (ptr *Produkt) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Produktd7c2e5___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Produkt) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Produkt) __findChildren_newList() unsafe.Pointer {
	return C.Produktd7c2e5___findChildren_newList(ptr.Pointer())
}

func (ptr *Produkt) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Produktd7c2e5___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Produkt) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Produkt) __children_newList() unsafe.Pointer {
	return C.Produktd7c2e5___children_newList(ptr.Pointer())
}

func NewProdukt(parent std_core.QObject_ITF) *Produkt {
	tmpValue := NewProduktFromPointer(C.Produktd7c2e5_NewProdukt(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackProduktd7c2e5_DestroyProdukt
func callbackProduktd7c2e5_DestroyProdukt(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Produkt"); signal != nil {
		signal.(func())()
	} else {
		NewProduktFromPointer(ptr).DestroyProduktDefault()
	}
}

func (ptr *Produkt) ConnectDestroyProdukt(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Produkt"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Produkt", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Produkt", f)
		}
	}
}

func (ptr *Produkt) DisconnectDestroyProdukt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Produkt")
	}
}

func (ptr *Produkt) DestroyProdukt() {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_DestroyProdukt(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Produkt) DestroyProduktDefault() {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_DestroyProduktDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackProduktd7c2e5_Event
func callbackProduktd7c2e5_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Produkt) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Produktd7c2e5_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackProduktd7c2e5_EventFilter
func callbackProduktd7c2e5_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Produkt) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Produktd7c2e5_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackProduktd7c2e5_ChildEvent
func callbackProduktd7c2e5_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewProduktFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Produkt) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackProduktd7c2e5_ConnectNotify
func callbackProduktd7c2e5_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewProduktFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Produkt) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackProduktd7c2e5_CustomEvent
func callbackProduktd7c2e5_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewProduktFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Produkt) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackProduktd7c2e5_DeleteLater
func callbackProduktd7c2e5_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewProduktFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Produkt) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackProduktd7c2e5_Destroyed
func callbackProduktd7c2e5_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackProduktd7c2e5_DisconnectNotify
func callbackProduktd7c2e5_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewProduktFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Produkt) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackProduktd7c2e5_ObjectNameChanged
func callbackProduktd7c2e5_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackProduktd7c2e5_TimerEvent
func callbackProduktd7c2e5_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewProduktFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Produkt) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Produktd7c2e5_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type ProduktyModel_ITF interface {
	std_core.QAbstractListModel_ITF
	ProduktyModel_PTR() *ProduktyModel
}

func (ptr *ProduktyModel) ProduktyModel_PTR() *ProduktyModel {
	return ptr
}

func (ptr *ProduktyModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *ProduktyModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromProduktyModel(ptr ProduktyModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ProduktyModel_PTR().Pointer()
	}
	return nil
}

func NewProduktyModelFromPointer(ptr unsafe.Pointer) (n *ProduktyModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ProduktyModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ProduktyModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &ProduktyModel{QAbstractListModel: *deduced}

		default:
			n = new(ProduktyModel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackProduktyModeld7c2e5_Constructor
func callbackProduktyModeld7c2e5_Constructor(ptr unsafe.Pointer) {
	this := NewProduktyModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackProduktyModeld7c2e5_AddProdukt
func callbackProduktyModeld7c2e5_AddProdukt(ptr unsafe.Pointer, v0 unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addProdukt"); signal != nil {
		signal.(func(*Produkt))(NewProduktFromPointer(v0))
	}

}

func (ptr *ProduktyModel) ConnectAddProdukt(f func(v0 *Produkt)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addProdukt"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addProdukt", func(v0 *Produkt) {
				signal.(func(*Produkt))(v0)
				f(v0)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addProdukt", f)
		}
	}
}

func (ptr *ProduktyModel) DisconnectAddProdukt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addProdukt")
	}
}

func (ptr *ProduktyModel) AddProdukt(v0 Produkt_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_AddProdukt(ptr.Pointer(), PointerFromProdukt(v0))
	}
}

//export callbackProduktyModeld7c2e5_EditProdukt
func callbackProduktyModeld7c2e5_EditProdukt(ptr unsafe.Pointer, row C.int, produktName C.struct_Moc_PackedString, produktW C.float, produktId C.int) {
	if signal := qt.GetSignal(ptr, "editProdukt"); signal != nil {
		signal.(func(int, string, float32, int))(int(int32(row)), cGoUnpackString(produktName), float32(produktW), int(int32(produktId)))
	}

}

func (ptr *ProduktyModel) ConnectEditProdukt(f func(row int, produktName string, produktW float32, produktId int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "editProdukt"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "editProdukt", func(row int, produktName string, produktW float32, produktId int) {
				signal.(func(int, string, float32, int))(row, produktName, produktW, produktId)
				f(row, produktName, produktW, produktId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "editProdukt", f)
		}
	}
}

func (ptr *ProduktyModel) DisconnectEditProdukt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "editProdukt")
	}
}

func (ptr *ProduktyModel) EditProdukt(row int, produktName string, produktW float32, produktId int) {
	if ptr.Pointer() != nil {
		var produktNameC *C.char
		if produktName != "" {
			produktNameC = C.CString(produktName)
			defer C.free(unsafe.Pointer(produktNameC))
		}
		C.ProduktyModeld7c2e5_EditProdukt(ptr.Pointer(), C.int(int32(row)), C.struct_Moc_PackedString{data: produktNameC, len: C.longlong(len(produktName))}, C.float(produktW), C.int(int32(produktId)))
	}
}

//export callbackProduktyModeld7c2e5_RemoveProdukt
func callbackProduktyModeld7c2e5_RemoveProdukt(ptr unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "removeProdukt"); signal != nil {
		signal.(func(int))(int(int32(row)))
	}

}

func (ptr *ProduktyModel) ConnectRemoveProdukt(f func(row int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeProdukt"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "removeProdukt", func(row int) {
				signal.(func(int))(row)
				f(row)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeProdukt", f)
		}
	}
}

func (ptr *ProduktyModel) DisconnectRemoveProdukt() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeProdukt")
	}
}

func (ptr *ProduktyModel) RemoveProdukt(row int) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_RemoveProdukt(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackProduktyModeld7c2e5_Roles
func callbackProduktyModeld7c2e5_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__roles_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__roles_newList())
		for k, v := range NewProduktyModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *ProduktyModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "roles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "roles", func() map[int]*std_core.QByteArray {
				signal.(func() map[int]*std_core.QByteArray)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "roles", f)
		}
	}
}

func (ptr *ProduktyModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *ProduktyModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.ProduktyModeld7c2e5_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *ProduktyModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.ProduktyModeld7c2e5_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackProduktyModeld7c2e5_SetRoles
func callbackProduktyModeld7c2e5_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewProduktyModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *ProduktyModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRoles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", f)
		}
	}
}

func (ptr *ProduktyModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *ProduktyModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *ProduktyModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackProduktyModeld7c2e5_RolesChanged
func callbackProduktyModeld7c2e5_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *ProduktyModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.ProduktyModeld7c2e5_ConnectRolesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rolesChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", f)
		}
	}
}

func (ptr *ProduktyModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *ProduktyModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackProduktyModeld7c2e5_Produktitem
func callbackProduktyModeld7c2e5_Produktitem(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "produktitem"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__produktitem_newList())
			for _, v := range signal.(func() []*Produkt)() {
				tmpList.__produktitem_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__produktitem_newList())
		for _, v := range NewProduktyModelFromPointer(ptr).ProduktitemDefault() {
			tmpList.__produktitem_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *ProduktyModel) ConnectProduktitem(f func() []*Produkt) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "produktitem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "produktitem", func() []*Produkt {
				signal.(func() []*Produkt)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "produktitem", f)
		}
	}
}

func (ptr *ProduktyModel) DisconnectProduktitem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "produktitem")
	}
}

func (ptr *ProduktyModel) Produktitem() []*Produkt {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Produkt {
			out := make([]*Produkt, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__produktitem_atList(i)
			}
			return out
		}(C.ProduktyModeld7c2e5_Produktitem(ptr.Pointer()))
	}
	return make([]*Produkt, 0)
}

func (ptr *ProduktyModel) ProduktitemDefault() []*Produkt {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Produkt {
			out := make([]*Produkt, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__produktitem_atList(i)
			}
			return out
		}(C.ProduktyModeld7c2e5_ProduktitemDefault(ptr.Pointer()))
	}
	return make([]*Produkt, 0)
}

//export callbackProduktyModeld7c2e5_SetProduktitem
func callbackProduktyModeld7c2e5_SetProduktitem(ptr unsafe.Pointer, produktitem C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setProduktitem"); signal != nil {
		signal.(func([]*Produkt))(func(l C.struct_Moc_PackedList) []*Produkt {
			out := make([]*Produkt, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setProduktitem_produktitem_atList(i)
			}
			return out
		}(produktitem))
	} else {
		NewProduktyModelFromPointer(ptr).SetProduktitemDefault(func(l C.struct_Moc_PackedList) []*Produkt {
			out := make([]*Produkt, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setProduktitem_produktitem_atList(i)
			}
			return out
		}(produktitem))
	}
}

func (ptr *ProduktyModel) ConnectSetProduktitem(f func(produktitem []*Produkt)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setProduktitem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setProduktitem", func(produktitem []*Produkt) {
				signal.(func([]*Produkt))(produktitem)
				f(produktitem)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setProduktitem", f)
		}
	}
}

func (ptr *ProduktyModel) DisconnectSetProduktitem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setProduktitem")
	}
}

func (ptr *ProduktyModel) SetProduktitem(produktitem []*Produkt) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_SetProduktitem(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__setProduktitem_produktitem_newList())
			for _, v := range produktitem {
				tmpList.__setProduktitem_produktitem_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *ProduktyModel) SetProduktitemDefault(produktitem []*Produkt) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_SetProduktitemDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__setProduktitem_produktitem_newList())
			for _, v := range produktitem {
				tmpList.__setProduktitem_produktitem_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackProduktyModeld7c2e5_ProduktitemChanged
func callbackProduktyModeld7c2e5_ProduktitemChanged(ptr unsafe.Pointer, produktitem C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "produktitemChanged"); signal != nil {
		signal.(func([]*Produkt))(func(l C.struct_Moc_PackedList) []*Produkt {
			out := make([]*Produkt, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__produktitemChanged_produktitem_atList(i)
			}
			return out
		}(produktitem))
	}

}

func (ptr *ProduktyModel) ConnectProduktitemChanged(f func(produktitem []*Produkt)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "produktitemChanged") {
			C.ProduktyModeld7c2e5_ConnectProduktitemChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "produktitemChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "produktitemChanged", func(produktitem []*Produkt) {
				signal.(func([]*Produkt))(produktitem)
				f(produktitem)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "produktitemChanged", f)
		}
	}
}

func (ptr *ProduktyModel) DisconnectProduktitemChanged() {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_DisconnectProduktitemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "produktitemChanged")
	}
}

func (ptr *ProduktyModel) ProduktitemChanged(produktitem []*Produkt) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_ProduktitemChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__produktitemChanged_produktitem_newList())
			for _, v := range produktitem {
				tmpList.__produktitemChanged_produktitem_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func ProduktyModel_QRegisterMetaType() int {
	return int(int32(C.ProduktyModeld7c2e5_ProduktyModeld7c2e5_QRegisterMetaType()))
}

func (ptr *ProduktyModel) QRegisterMetaType() int {
	return int(int32(C.ProduktyModeld7c2e5_ProduktyModeld7c2e5_QRegisterMetaType()))
}

func ProduktyModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ProduktyModeld7c2e5_ProduktyModeld7c2e5_QRegisterMetaType2(typeNameC)))
}

func (ptr *ProduktyModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ProduktyModeld7c2e5_ProduktyModeld7c2e5_QRegisterMetaType2(typeNameC)))
}

func ProduktyModel_QmlRegisterType() int {
	return int(int32(C.ProduktyModeld7c2e5_ProduktyModeld7c2e5_QmlRegisterType()))
}

func (ptr *ProduktyModel) QmlRegisterType() int {
	return int(int32(C.ProduktyModeld7c2e5_ProduktyModeld7c2e5_QmlRegisterType()))
}

func ProduktyModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ProduktyModeld7c2e5_ProduktyModeld7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ProduktyModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ProduktyModeld7c2e5_ProduktyModeld7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ProduktyModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyModeld7c2e5_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ProduktyModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ProduktyModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyModeld7c2e5_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ProduktyModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ProduktyModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyModeld7c2e5_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ProduktyModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ProduktyModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.ProduktyModeld7c2e5___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *ProduktyModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.ProduktyModeld7c2e5___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ProduktyModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ProduktyModeld7c2e5___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *ProduktyModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ProduktyModeld7c2e5___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *ProduktyModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyModeld7c2e5___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ProduktyModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ProduktyModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.ProduktyModeld7c2e5___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *ProduktyModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.ProduktyModeld7c2e5___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *ProduktyModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ProduktyModeld7c2e5___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ProduktyModel) __roleNames_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___roleNames_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.ProduktyModeld7c2e5___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ProduktyModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.ProduktyModeld7c2e5___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *ProduktyModel) __itemData_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___itemData_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.ProduktyModeld7c2e5___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ProduktyModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ProduktyModeld7c2e5___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *ProduktyModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ProduktyModeld7c2e5___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *ProduktyModel) __match_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___match_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ProduktyModeld7c2e5___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *ProduktyModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyModeld7c2e5_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ProduktyModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ProduktyModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyModeld7c2e5_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ProduktyModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ProduktyModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ProduktyModeld7c2e5___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ProduktyModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ProduktyModeld7c2e5___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ProduktyModel) __findChildren_newList2() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___findChildren_newList2(ptr.Pointer())
}

func (ptr *ProduktyModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ProduktyModeld7c2e5___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ProduktyModel) __findChildren_newList3() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___findChildren_newList3(ptr.Pointer())
}

func (ptr *ProduktyModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ProduktyModeld7c2e5___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ProduktyModel) __findChildren_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___findChildren_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ProduktyModeld7c2e5___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ProduktyModel) __children_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___children_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ProduktyModeld7c2e5___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ProduktyModel) __roles_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___roles_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.ProduktyModeld7c2e5___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ProduktyModel) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ProduktyModeld7c2e5___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ProduktyModel) __setRoles_roles_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.ProduktyModeld7c2e5___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ProduktyModel) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ProduktyModeld7c2e5___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ProduktyModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.ProduktyModeld7c2e5___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ProduktyModel) __produktitem_atList(i int) *Produkt {
	if ptr.Pointer() != nil {
		tmpValue := NewProduktFromPointer(C.ProduktyModeld7c2e5___produktitem_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __produktitem_setList(i Produkt_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___produktitem_setList(ptr.Pointer(), PointerFromProdukt(i))
	}
}

func (ptr *ProduktyModel) __produktitem_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___produktitem_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __setProduktitem_produktitem_atList(i int) *Produkt {
	if ptr.Pointer() != nil {
		tmpValue := NewProduktFromPointer(C.ProduktyModeld7c2e5___setProduktitem_produktitem_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __setProduktitem_produktitem_setList(i Produkt_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___setProduktitem_produktitem_setList(ptr.Pointer(), PointerFromProdukt(i))
	}
}

func (ptr *ProduktyModel) __setProduktitem_produktitem_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___setProduktitem_produktitem_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) __produktitemChanged_produktitem_atList(i int) *Produkt {
	if ptr.Pointer() != nil {
		tmpValue := NewProduktFromPointer(C.ProduktyModeld7c2e5___produktitemChanged_produktitem_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyModel) __produktitemChanged_produktitem_setList(i Produkt_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5___produktitemChanged_produktitem_setList(ptr.Pointer(), PointerFromProdukt(i))
	}
}

func (ptr *ProduktyModel) __produktitemChanged_produktitem_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5___produktitemChanged_produktitem_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyModeld7c2e5_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ProduktyModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ProduktyModel) ____roles_keyList_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyModeld7c2e5_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ProduktyModel) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ProduktyModel) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *ProduktyModel) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyModeld7c2e5_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ProduktyModel) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ProduktyModel) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.ProduktyModeld7c2e5_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewProduktyModel(parent std_core.QObject_ITF) *ProduktyModel {
	tmpValue := NewProduktyModelFromPointer(C.ProduktyModeld7c2e5_NewProduktyModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackProduktyModeld7c2e5_DestroyProduktyModel
func callbackProduktyModeld7c2e5_DestroyProduktyModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~ProduktyModel"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyModelFromPointer(ptr).DestroyProduktyModelDefault()
	}
}

func (ptr *ProduktyModel) ConnectDestroyProduktyModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~ProduktyModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~ProduktyModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~ProduktyModel", f)
		}
	}
}

func (ptr *ProduktyModel) DisconnectDestroyProduktyModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~ProduktyModel")
	}
}

func (ptr *ProduktyModel) DestroyProduktyModel() {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_DestroyProduktyModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ProduktyModel) DestroyProduktyModelDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_DestroyProduktyModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackProduktyModeld7c2e5_DropMimeData
func callbackProduktyModeld7c2e5_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ProduktyModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_Index
func callbackProduktyModeld7c2e5_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewProduktyModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *ProduktyModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ProduktyModeld7c2e5_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackProduktyModeld7c2e5_Sibling
func callbackProduktyModeld7c2e5_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewProduktyModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *ProduktyModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ProduktyModeld7c2e5_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackProduktyModeld7c2e5_Flags
func callbackProduktyModeld7c2e5_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewProduktyModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *ProduktyModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.ProduktyModeld7c2e5_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackProduktyModeld7c2e5_InsertColumns
func callbackProduktyModeld7c2e5_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ProduktyModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_InsertRows
func callbackProduktyModeld7c2e5_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ProduktyModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_MoveColumns
func callbackProduktyModeld7c2e5_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *ProduktyModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_MoveRows
func callbackProduktyModeld7c2e5_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *ProduktyModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_RemoveColumns
func callbackProduktyModeld7c2e5_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ProduktyModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_RemoveRows
func callbackProduktyModeld7c2e5_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ProduktyModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_SetData
func callbackProduktyModeld7c2e5_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *ProduktyModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_SetHeaderData
func callbackProduktyModeld7c2e5_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *ProduktyModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_SetItemData
func callbackProduktyModeld7c2e5_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewProduktyModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *ProduktyModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_Submit
func callbackProduktyModeld7c2e5_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *ProduktyModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_ColumnsAboutToBeInserted
func callbackProduktyModeld7c2e5_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackProduktyModeld7c2e5_ColumnsAboutToBeMoved
func callbackProduktyModeld7c2e5_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackProduktyModeld7c2e5_ColumnsAboutToBeRemoved
func callbackProduktyModeld7c2e5_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackProduktyModeld7c2e5_ColumnsInserted
func callbackProduktyModeld7c2e5_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackProduktyModeld7c2e5_ColumnsMoved
func callbackProduktyModeld7c2e5_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackProduktyModeld7c2e5_ColumnsRemoved
func callbackProduktyModeld7c2e5_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackProduktyModeld7c2e5_DataChanged
func callbackProduktyModeld7c2e5_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackProduktyModeld7c2e5_FetchMore
func callbackProduktyModeld7c2e5_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewProduktyModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *ProduktyModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackProduktyModeld7c2e5_HeaderDataChanged
func callbackProduktyModeld7c2e5_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackProduktyModeld7c2e5_LayoutAboutToBeChanged
func callbackProduktyModeld7c2e5_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackProduktyModeld7c2e5_LayoutChanged
func callbackProduktyModeld7c2e5_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackProduktyModeld7c2e5_ModelAboutToBeReset
func callbackProduktyModeld7c2e5_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackProduktyModeld7c2e5_ModelReset
func callbackProduktyModeld7c2e5_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackProduktyModeld7c2e5_ResetInternalData
func callbackProduktyModeld7c2e5_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *ProduktyModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackProduktyModeld7c2e5_Revert
func callbackProduktyModeld7c2e5_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *ProduktyModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_RevertDefault(ptr.Pointer())
	}
}

//export callbackProduktyModeld7c2e5_RowsAboutToBeInserted
func callbackProduktyModeld7c2e5_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackProduktyModeld7c2e5_RowsAboutToBeMoved
func callbackProduktyModeld7c2e5_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackProduktyModeld7c2e5_RowsAboutToBeRemoved
func callbackProduktyModeld7c2e5_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackProduktyModeld7c2e5_RowsInserted
func callbackProduktyModeld7c2e5_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackProduktyModeld7c2e5_RowsMoved
func callbackProduktyModeld7c2e5_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackProduktyModeld7c2e5_RowsRemoved
func callbackProduktyModeld7c2e5_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackProduktyModeld7c2e5_Sort
func callbackProduktyModeld7c2e5_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewProduktyModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *ProduktyModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackProduktyModeld7c2e5_RoleNames
func callbackProduktyModeld7c2e5_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewProduktyModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *ProduktyModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.ProduktyModeld7c2e5_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackProduktyModeld7c2e5_ItemData
func callbackProduktyModeld7c2e5_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__itemData_newList())
		for k, v := range NewProduktyModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *ProduktyModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.ProduktyModeld7c2e5_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackProduktyModeld7c2e5_MimeData
func callbackProduktyModeld7c2e5_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewProduktyModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewProduktyModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *ProduktyModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.ProduktyModeld7c2e5_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackProduktyModeld7c2e5_Buddy
func callbackProduktyModeld7c2e5_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewProduktyModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *ProduktyModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ProduktyModeld7c2e5_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackProduktyModeld7c2e5_Parent
func callbackProduktyModeld7c2e5_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewProduktyModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *ProduktyModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ProduktyModeld7c2e5_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackProduktyModeld7c2e5_Match
func callbackProduktyModeld7c2e5_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewProduktyModelFromPointer(NewProduktyModelFromPointer(nil).__match_newList())
		for _, v := range NewProduktyModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *ProduktyModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewProduktyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.ProduktyModeld7c2e5_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackProduktyModeld7c2e5_Span
func callbackProduktyModeld7c2e5_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewProduktyModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *ProduktyModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.ProduktyModeld7c2e5_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackProduktyModeld7c2e5_MimeTypes
func callbackProduktyModeld7c2e5_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewProduktyModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *ProduktyModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.ProduktyModeld7c2e5_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackProduktyModeld7c2e5_Data
func callbackProduktyModeld7c2e5_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewProduktyModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *ProduktyModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.ProduktyModeld7c2e5_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackProduktyModeld7c2e5_HeaderData
func callbackProduktyModeld7c2e5_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewProduktyModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *ProduktyModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.ProduktyModeld7c2e5_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackProduktyModeld7c2e5_SupportedDragActions
func callbackProduktyModeld7c2e5_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewProduktyModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *ProduktyModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.ProduktyModeld7c2e5_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackProduktyModeld7c2e5_SupportedDropActions
func callbackProduktyModeld7c2e5_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewProduktyModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *ProduktyModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.ProduktyModeld7c2e5_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackProduktyModeld7c2e5_CanDropMimeData
func callbackProduktyModeld7c2e5_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ProduktyModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_CanFetchMore
func callbackProduktyModeld7c2e5_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ProduktyModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_HasChildren
func callbackProduktyModeld7c2e5_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ProduktyModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_ColumnCount
func callbackProduktyModeld7c2e5_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewProduktyModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *ProduktyModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyModeld7c2e5_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackProduktyModeld7c2e5_RowCount
func callbackProduktyModeld7c2e5_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewProduktyModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *ProduktyModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyModeld7c2e5_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackProduktyModeld7c2e5_Event
func callbackProduktyModeld7c2e5_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ProduktyModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_EventFilter
func callbackProduktyModeld7c2e5_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ProduktyModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyModeld7c2e5_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackProduktyModeld7c2e5_ChildEvent
func callbackProduktyModeld7c2e5_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewProduktyModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ProduktyModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackProduktyModeld7c2e5_ConnectNotify
func callbackProduktyModeld7c2e5_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewProduktyModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ProduktyModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackProduktyModeld7c2e5_CustomEvent
func callbackProduktyModeld7c2e5_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewProduktyModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ProduktyModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackProduktyModeld7c2e5_DeleteLater
func callbackProduktyModeld7c2e5_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ProduktyModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackProduktyModeld7c2e5_Destroyed
func callbackProduktyModeld7c2e5_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackProduktyModeld7c2e5_DisconnectNotify
func callbackProduktyModeld7c2e5_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewProduktyModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ProduktyModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackProduktyModeld7c2e5_ObjectNameChanged
func callbackProduktyModeld7c2e5_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackProduktyModeld7c2e5_TimerEvent
func callbackProduktyModeld7c2e5_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewProduktyModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ProduktyModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyModeld7c2e5_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type Element_ITF interface {
	std_core.QObject_ITF
	Element_PTR() *Element
}

func (ptr *Element) Element_PTR() *Element {
	return ptr
}

func (ptr *Element) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Element) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromElement(ptr Element_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Element_PTR().Pointer()
	}
	return nil
}

func NewElementFromPointer(ptr unsafe.Pointer) (n *Element) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Element)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Element:
			n = deduced

		case *std_core.QObject:
			n = &Element{QObject: *deduced}

		default:
			n = new(Element)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackElementd7c2e5_Constructor
func callbackElementd7c2e5_Constructor(ptr unsafe.Pointer) {
	this := NewElementFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackElementd7c2e5_ElementGName
func callbackElementd7c2e5_ElementGName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "elementGName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewElementFromPointer(ptr).ElementGNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Element) ConnectElementGName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "elementGName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementGName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementGName", f)
		}
	}
}

func (ptr *Element) DisconnectElementGName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "elementGName")
	}
}

func (ptr *Element) ElementGName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Elementd7c2e5_ElementGName(ptr.Pointer()))
	}
	return ""
}

func (ptr *Element) ElementGNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Elementd7c2e5_ElementGNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackElementd7c2e5_SetElementGName
func callbackElementd7c2e5_SetElementGName(ptr unsafe.Pointer, elementGName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setElementGName"); signal != nil {
		signal.(func(string))(cGoUnpackString(elementGName))
	} else {
		NewElementFromPointer(ptr).SetElementGNameDefault(cGoUnpackString(elementGName))
	}
}

func (ptr *Element) ConnectSetElementGName(f func(elementGName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setElementGName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setElementGName", func(elementGName string) {
				signal.(func(string))(elementGName)
				f(elementGName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setElementGName", f)
		}
	}
}

func (ptr *Element) DisconnectSetElementGName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setElementGName")
	}
}

func (ptr *Element) SetElementGName(elementGName string) {
	if ptr.Pointer() != nil {
		var elementGNameC *C.char
		if elementGName != "" {
			elementGNameC = C.CString(elementGName)
			defer C.free(unsafe.Pointer(elementGNameC))
		}
		C.Elementd7c2e5_SetElementGName(ptr.Pointer(), C.struct_Moc_PackedString{data: elementGNameC, len: C.longlong(len(elementGName))})
	}
}

func (ptr *Element) SetElementGNameDefault(elementGName string) {
	if ptr.Pointer() != nil {
		var elementGNameC *C.char
		if elementGName != "" {
			elementGNameC = C.CString(elementGName)
			defer C.free(unsafe.Pointer(elementGNameC))
		}
		C.Elementd7c2e5_SetElementGNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: elementGNameC, len: C.longlong(len(elementGName))})
	}
}

//export callbackElementd7c2e5_ElementGNameChanged
func callbackElementd7c2e5_ElementGNameChanged(ptr unsafe.Pointer, elementGName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "elementGNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(elementGName))
	}

}

func (ptr *Element) ConnectElementGNameChanged(f func(elementGName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "elementGNameChanged") {
			C.Elementd7c2e5_ConnectElementGNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "elementGNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementGNameChanged", func(elementGName string) {
				signal.(func(string))(elementGName)
				f(elementGName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementGNameChanged", f)
		}
	}
}

func (ptr *Element) DisconnectElementGNameChanged() {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_DisconnectElementGNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "elementGNameChanged")
	}
}

func (ptr *Element) ElementGNameChanged(elementGName string) {
	if ptr.Pointer() != nil {
		var elementGNameC *C.char
		if elementGName != "" {
			elementGNameC = C.CString(elementGName)
			defer C.free(unsafe.Pointer(elementGNameC))
		}
		C.Elementd7c2e5_ElementGNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: elementGNameC, len: C.longlong(len(elementGName))})
	}
}

//export callbackElementd7c2e5_ElementStan
func callbackElementd7c2e5_ElementStan(ptr unsafe.Pointer) C.float {
	if signal := qt.GetSignal(ptr, "elementStan"); signal != nil {
		return C.float(signal.(func() float32)())
	}

	return C.float(NewElementFromPointer(ptr).ElementStanDefault())
}

func (ptr *Element) ConnectElementStan(f func() float32) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "elementStan"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementStan", func() float32 {
				signal.(func() float32)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementStan", f)
		}
	}
}

func (ptr *Element) DisconnectElementStan() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "elementStan")
	}
}

func (ptr *Element) ElementStan() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Elementd7c2e5_ElementStan(ptr.Pointer()))
	}
	return 0
}

func (ptr *Element) ElementStanDefault() float32 {
	if ptr.Pointer() != nil {
		return float32(C.Elementd7c2e5_ElementStanDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackElementd7c2e5_SetElementStan
func callbackElementd7c2e5_SetElementStan(ptr unsafe.Pointer, elementStan C.float) {
	if signal := qt.GetSignal(ptr, "setElementStan"); signal != nil {
		signal.(func(float32))(float32(elementStan))
	} else {
		NewElementFromPointer(ptr).SetElementStanDefault(float32(elementStan))
	}
}

func (ptr *Element) ConnectSetElementStan(f func(elementStan float32)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setElementStan"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setElementStan", func(elementStan float32) {
				signal.(func(float32))(elementStan)
				f(elementStan)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setElementStan", f)
		}
	}
}

func (ptr *Element) DisconnectSetElementStan() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setElementStan")
	}
}

func (ptr *Element) SetElementStan(elementStan float32) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_SetElementStan(ptr.Pointer(), C.float(elementStan))
	}
}

func (ptr *Element) SetElementStanDefault(elementStan float32) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_SetElementStanDefault(ptr.Pointer(), C.float(elementStan))
	}
}

//export callbackElementd7c2e5_ElementStanChanged
func callbackElementd7c2e5_ElementStanChanged(ptr unsafe.Pointer, elementStan C.float) {
	if signal := qt.GetSignal(ptr, "elementStanChanged"); signal != nil {
		signal.(func(float32))(float32(elementStan))
	}

}

func (ptr *Element) ConnectElementStanChanged(f func(elementStan float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "elementStanChanged") {
			C.Elementd7c2e5_ConnectElementStanChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "elementStanChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementStanChanged", func(elementStan float32) {
				signal.(func(float32))(elementStan)
				f(elementStan)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementStanChanged", f)
		}
	}
}

func (ptr *Element) DisconnectElementStanChanged() {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_DisconnectElementStanChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "elementStanChanged")
	}
}

func (ptr *Element) ElementStanChanged(elementStan float32) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_ElementStanChanged(ptr.Pointer(), C.float(elementStan))
	}
}

//export callbackElementd7c2e5_ElementPName
func callbackElementd7c2e5_ElementPName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "elementPName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewElementFromPointer(ptr).ElementPNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Element) ConnectElementPName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "elementPName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementPName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementPName", f)
		}
	}
}

func (ptr *Element) DisconnectElementPName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "elementPName")
	}
}

func (ptr *Element) ElementPName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Elementd7c2e5_ElementPName(ptr.Pointer()))
	}
	return ""
}

func (ptr *Element) ElementPNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Elementd7c2e5_ElementPNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackElementd7c2e5_SetElementPName
func callbackElementd7c2e5_SetElementPName(ptr unsafe.Pointer, elementPName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setElementPName"); signal != nil {
		signal.(func(string))(cGoUnpackString(elementPName))
	} else {
		NewElementFromPointer(ptr).SetElementPNameDefault(cGoUnpackString(elementPName))
	}
}

func (ptr *Element) ConnectSetElementPName(f func(elementPName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setElementPName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setElementPName", func(elementPName string) {
				signal.(func(string))(elementPName)
				f(elementPName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setElementPName", f)
		}
	}
}

func (ptr *Element) DisconnectSetElementPName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setElementPName")
	}
}

func (ptr *Element) SetElementPName(elementPName string) {
	if ptr.Pointer() != nil {
		var elementPNameC *C.char
		if elementPName != "" {
			elementPNameC = C.CString(elementPName)
			defer C.free(unsafe.Pointer(elementPNameC))
		}
		C.Elementd7c2e5_SetElementPName(ptr.Pointer(), C.struct_Moc_PackedString{data: elementPNameC, len: C.longlong(len(elementPName))})
	}
}

func (ptr *Element) SetElementPNameDefault(elementPName string) {
	if ptr.Pointer() != nil {
		var elementPNameC *C.char
		if elementPName != "" {
			elementPNameC = C.CString(elementPName)
			defer C.free(unsafe.Pointer(elementPNameC))
		}
		C.Elementd7c2e5_SetElementPNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: elementPNameC, len: C.longlong(len(elementPName))})
	}
}

//export callbackElementd7c2e5_ElementPNameChanged
func callbackElementd7c2e5_ElementPNameChanged(ptr unsafe.Pointer, elementPName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "elementPNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(elementPName))
	}

}

func (ptr *Element) ConnectElementPNameChanged(f func(elementPName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "elementPNameChanged") {
			C.Elementd7c2e5_ConnectElementPNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "elementPNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementPNameChanged", func(elementPName string) {
				signal.(func(string))(elementPName)
				f(elementPName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementPNameChanged", f)
		}
	}
}

func (ptr *Element) DisconnectElementPNameChanged() {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_DisconnectElementPNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "elementPNameChanged")
	}
}

func (ptr *Element) ElementPNameChanged(elementPName string) {
	if ptr.Pointer() != nil {
		var elementPNameC *C.char
		if elementPName != "" {
			elementPNameC = C.CString(elementPName)
			defer C.free(unsafe.Pointer(elementPNameC))
		}
		C.Elementd7c2e5_ElementPNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: elementPNameC, len: C.longlong(len(elementPName))})
	}
}

//export callbackElementd7c2e5_ElementIlosc
func callbackElementd7c2e5_ElementIlosc(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "elementIlosc"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewElementFromPointer(ptr).ElementIloscDefault()))
}

func (ptr *Element) ConnectElementIlosc(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "elementIlosc"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementIlosc", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementIlosc", f)
		}
	}
}

func (ptr *Element) DisconnectElementIlosc() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "elementIlosc")
	}
}

func (ptr *Element) ElementIlosc() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Elementd7c2e5_ElementIlosc(ptr.Pointer())))
	}
	return 0
}

func (ptr *Element) ElementIloscDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Elementd7c2e5_ElementIloscDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackElementd7c2e5_SetElementIlosc
func callbackElementd7c2e5_SetElementIlosc(ptr unsafe.Pointer, elementIlosc C.int) {
	if signal := qt.GetSignal(ptr, "setElementIlosc"); signal != nil {
		signal.(func(int))(int(int32(elementIlosc)))
	} else {
		NewElementFromPointer(ptr).SetElementIloscDefault(int(int32(elementIlosc)))
	}
}

func (ptr *Element) ConnectSetElementIlosc(f func(elementIlosc int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setElementIlosc"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setElementIlosc", func(elementIlosc int) {
				signal.(func(int))(elementIlosc)
				f(elementIlosc)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setElementIlosc", f)
		}
	}
}

func (ptr *Element) DisconnectSetElementIlosc() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setElementIlosc")
	}
}

func (ptr *Element) SetElementIlosc(elementIlosc int) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_SetElementIlosc(ptr.Pointer(), C.int(int32(elementIlosc)))
	}
}

func (ptr *Element) SetElementIloscDefault(elementIlosc int) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_SetElementIloscDefault(ptr.Pointer(), C.int(int32(elementIlosc)))
	}
}

//export callbackElementd7c2e5_ElementIloscChanged
func callbackElementd7c2e5_ElementIloscChanged(ptr unsafe.Pointer, elementIlosc C.int) {
	if signal := qt.GetSignal(ptr, "elementIloscChanged"); signal != nil {
		signal.(func(int))(int(int32(elementIlosc)))
	}

}

func (ptr *Element) ConnectElementIloscChanged(f func(elementIlosc int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "elementIloscChanged") {
			C.Elementd7c2e5_ConnectElementIloscChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "elementIloscChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementIloscChanged", func(elementIlosc int) {
				signal.(func(int))(elementIlosc)
				f(elementIlosc)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementIloscChanged", f)
		}
	}
}

func (ptr *Element) DisconnectElementIloscChanged() {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_DisconnectElementIloscChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "elementIloscChanged")
	}
}

func (ptr *Element) ElementIloscChanged(elementIlosc int) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_ElementIloscChanged(ptr.Pointer(), C.int(int32(elementIlosc)))
	}
}

//export callbackElementd7c2e5_ElementData
func callbackElementd7c2e5_ElementData(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "elementData"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewElementFromPointer(ptr).ElementDataDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Element) ConnectElementData(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "elementData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementData", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementData", f)
		}
	}
}

func (ptr *Element) DisconnectElementData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "elementData")
	}
}

func (ptr *Element) ElementData() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Elementd7c2e5_ElementData(ptr.Pointer()))
	}
	return ""
}

func (ptr *Element) ElementDataDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Elementd7c2e5_ElementDataDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackElementd7c2e5_SetElementData
func callbackElementd7c2e5_SetElementData(ptr unsafe.Pointer, elementData C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setElementData"); signal != nil {
		signal.(func(string))(cGoUnpackString(elementData))
	} else {
		NewElementFromPointer(ptr).SetElementDataDefault(cGoUnpackString(elementData))
	}
}

func (ptr *Element) ConnectSetElementData(f func(elementData string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setElementData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setElementData", func(elementData string) {
				signal.(func(string))(elementData)
				f(elementData)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setElementData", f)
		}
	}
}

func (ptr *Element) DisconnectSetElementData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setElementData")
	}
}

func (ptr *Element) SetElementData(elementData string) {
	if ptr.Pointer() != nil {
		var elementDataC *C.char
		if elementData != "" {
			elementDataC = C.CString(elementData)
			defer C.free(unsafe.Pointer(elementDataC))
		}
		C.Elementd7c2e5_SetElementData(ptr.Pointer(), C.struct_Moc_PackedString{data: elementDataC, len: C.longlong(len(elementData))})
	}
}

func (ptr *Element) SetElementDataDefault(elementData string) {
	if ptr.Pointer() != nil {
		var elementDataC *C.char
		if elementData != "" {
			elementDataC = C.CString(elementData)
			defer C.free(unsafe.Pointer(elementDataC))
		}
		C.Elementd7c2e5_SetElementDataDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: elementDataC, len: C.longlong(len(elementData))})
	}
}

//export callbackElementd7c2e5_ElementDataChanged
func callbackElementd7c2e5_ElementDataChanged(ptr unsafe.Pointer, elementData C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "elementDataChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(elementData))
	}

}

func (ptr *Element) ConnectElementDataChanged(f func(elementData string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "elementDataChanged") {
			C.Elementd7c2e5_ConnectElementDataChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "elementDataChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementDataChanged", func(elementData string) {
				signal.(func(string))(elementData)
				f(elementData)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementDataChanged", f)
		}
	}
}

func (ptr *Element) DisconnectElementDataChanged() {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_DisconnectElementDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "elementDataChanged")
	}
}

func (ptr *Element) ElementDataChanged(elementData string) {
	if ptr.Pointer() != nil {
		var elementDataC *C.char
		if elementData != "" {
			elementDataC = C.CString(elementData)
			defer C.free(unsafe.Pointer(elementDataC))
		}
		C.Elementd7c2e5_ElementDataChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: elementDataC, len: C.longlong(len(elementData))})
	}
}

//export callbackElementd7c2e5_ElementId
func callbackElementd7c2e5_ElementId(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "elementId"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewElementFromPointer(ptr).ElementIdDefault()))
}

func (ptr *Element) ConnectElementId(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "elementId"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementId", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementId", f)
		}
	}
}

func (ptr *Element) DisconnectElementId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "elementId")
	}
}

func (ptr *Element) ElementId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Elementd7c2e5_ElementId(ptr.Pointer())))
	}
	return 0
}

func (ptr *Element) ElementIdDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Elementd7c2e5_ElementIdDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackElementd7c2e5_SetElementId
func callbackElementd7c2e5_SetElementId(ptr unsafe.Pointer, elementId C.int) {
	if signal := qt.GetSignal(ptr, "setElementId"); signal != nil {
		signal.(func(int))(int(int32(elementId)))
	} else {
		NewElementFromPointer(ptr).SetElementIdDefault(int(int32(elementId)))
	}
}

func (ptr *Element) ConnectSetElementId(f func(elementId int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setElementId"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setElementId", func(elementId int) {
				signal.(func(int))(elementId)
				f(elementId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setElementId", f)
		}
	}
}

func (ptr *Element) DisconnectSetElementId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setElementId")
	}
}

func (ptr *Element) SetElementId(elementId int) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_SetElementId(ptr.Pointer(), C.int(int32(elementId)))
	}
}

func (ptr *Element) SetElementIdDefault(elementId int) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_SetElementIdDefault(ptr.Pointer(), C.int(int32(elementId)))
	}
}

//export callbackElementd7c2e5_ElementIdChanged
func callbackElementd7c2e5_ElementIdChanged(ptr unsafe.Pointer, elementId C.int) {
	if signal := qt.GetSignal(ptr, "elementIdChanged"); signal != nil {
		signal.(func(int))(int(int32(elementId)))
	}

}

func (ptr *Element) ConnectElementIdChanged(f func(elementId int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "elementIdChanged") {
			C.Elementd7c2e5_ConnectElementIdChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "elementIdChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementIdChanged", func(elementId int) {
				signal.(func(int))(elementId)
				f(elementId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementIdChanged", f)
		}
	}
}

func (ptr *Element) DisconnectElementIdChanged() {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_DisconnectElementIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "elementIdChanged")
	}
}

func (ptr *Element) ElementIdChanged(elementId int) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_ElementIdChanged(ptr.Pointer(), C.int(int32(elementId)))
	}
}

func Element_QRegisterMetaType() int {
	return int(int32(C.Elementd7c2e5_Elementd7c2e5_QRegisterMetaType()))
}

func (ptr *Element) QRegisterMetaType() int {
	return int(int32(C.Elementd7c2e5_Elementd7c2e5_QRegisterMetaType()))
}

func Element_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Elementd7c2e5_Elementd7c2e5_QRegisterMetaType2(typeNameC)))
}

func (ptr *Element) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Elementd7c2e5_Elementd7c2e5_QRegisterMetaType2(typeNameC)))
}

func Element_QmlRegisterType() int {
	return int(int32(C.Elementd7c2e5_Elementd7c2e5_QmlRegisterType()))
}

func (ptr *Element) QmlRegisterType() int {
	return int(int32(C.Elementd7c2e5_Elementd7c2e5_QmlRegisterType()))
}

func Element_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Elementd7c2e5_Elementd7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Element) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Elementd7c2e5_Elementd7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Element) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Elementd7c2e5___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Element) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Element) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Elementd7c2e5___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Element) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Elementd7c2e5___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Element) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Element) __findChildren_newList2() unsafe.Pointer {
	return C.Elementd7c2e5___findChildren_newList2(ptr.Pointer())
}

func (ptr *Element) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Elementd7c2e5___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Element) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Element) __findChildren_newList3() unsafe.Pointer {
	return C.Elementd7c2e5___findChildren_newList3(ptr.Pointer())
}

func (ptr *Element) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Elementd7c2e5___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Element) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Element) __findChildren_newList() unsafe.Pointer {
	return C.Elementd7c2e5___findChildren_newList(ptr.Pointer())
}

func (ptr *Element) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Elementd7c2e5___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Element) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Element) __children_newList() unsafe.Pointer {
	return C.Elementd7c2e5___children_newList(ptr.Pointer())
}

func NewElement(parent std_core.QObject_ITF) *Element {
	tmpValue := NewElementFromPointer(C.Elementd7c2e5_NewElement(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackElementd7c2e5_DestroyElement
func callbackElementd7c2e5_DestroyElement(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Element"); signal != nil {
		signal.(func())()
	} else {
		NewElementFromPointer(ptr).DestroyElementDefault()
	}
}

func (ptr *Element) ConnectDestroyElement(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Element"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Element", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Element", f)
		}
	}
}

func (ptr *Element) DisconnectDestroyElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Element")
	}
}

func (ptr *Element) DestroyElement() {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_DestroyElement(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Element) DestroyElementDefault() {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_DestroyElementDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackElementd7c2e5_Event
func callbackElementd7c2e5_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Element) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Elementd7c2e5_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackElementd7c2e5_EventFilter
func callbackElementd7c2e5_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Element) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Elementd7c2e5_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackElementd7c2e5_ChildEvent
func callbackElementd7c2e5_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewElementFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Element) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackElementd7c2e5_ConnectNotify
func callbackElementd7c2e5_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewElementFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Element) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackElementd7c2e5_CustomEvent
func callbackElementd7c2e5_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewElementFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Element) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackElementd7c2e5_DeleteLater
func callbackElementd7c2e5_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewElementFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Element) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackElementd7c2e5_Destroyed
func callbackElementd7c2e5_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackElementd7c2e5_DisconnectNotify
func callbackElementd7c2e5_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewElementFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Element) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackElementd7c2e5_ObjectNameChanged
func callbackElementd7c2e5_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackElementd7c2e5_TimerEvent
func callbackElementd7c2e5_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewElementFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Element) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Elementd7c2e5_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type ElementyModel_ITF interface {
	std_core.QAbstractListModel_ITF
	ElementyModel_PTR() *ElementyModel
}

func (ptr *ElementyModel) ElementyModel_PTR() *ElementyModel {
	return ptr
}

func (ptr *ElementyModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *ElementyModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromElementyModel(ptr ElementyModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ElementyModel_PTR().Pointer()
	}
	return nil
}

func NewElementyModelFromPointer(ptr unsafe.Pointer) (n *ElementyModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ElementyModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ElementyModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &ElementyModel{QAbstractListModel: *deduced}

		default:
			n = new(ElementyModel)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackElementyModeld7c2e5_Constructor
func callbackElementyModeld7c2e5_Constructor(ptr unsafe.Pointer) {
	this := NewElementyModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackElementyModeld7c2e5_AddElement
func callbackElementyModeld7c2e5_AddElement(ptr unsafe.Pointer, v0 unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addElement"); signal != nil {
		signal.(func(*Element))(NewElementFromPointer(v0))
	}

}

func (ptr *ElementyModel) ConnectAddElement(f func(v0 *Element)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addElement"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addElement", func(v0 *Element) {
				signal.(func(*Element))(v0)
				f(v0)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addElement", f)
		}
	}
}

func (ptr *ElementyModel) DisconnectAddElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addElement")
	}
}

func (ptr *ElementyModel) AddElement(v0 Element_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_AddElement(ptr.Pointer(), PointerFromElement(v0))
	}
}

//export callbackElementyModeld7c2e5_EditElement
func callbackElementyModeld7c2e5_EditElement(ptr unsafe.Pointer, row C.int, elementGName C.struct_Moc_PackedString, elementStan C.float, elementPName C.struct_Moc_PackedString, elementIlosc C.int, elementData C.struct_Moc_PackedString, elementId C.int) {
	if signal := qt.GetSignal(ptr, "editElement"); signal != nil {
		signal.(func(int, string, float32, string, int, string, int))(int(int32(row)), cGoUnpackString(elementGName), float32(elementStan), cGoUnpackString(elementPName), int(int32(elementIlosc)), cGoUnpackString(elementData), int(int32(elementId)))
	}

}

func (ptr *ElementyModel) ConnectEditElement(f func(row int, elementGName string, elementStan float32, elementPName string, elementIlosc int, elementData string, elementId int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "editElement"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "editElement", func(row int, elementGName string, elementStan float32, elementPName string, elementIlosc int, elementData string, elementId int) {
				signal.(func(int, string, float32, string, int, string, int))(row, elementGName, elementStan, elementPName, elementIlosc, elementData, elementId)
				f(row, elementGName, elementStan, elementPName, elementIlosc, elementData, elementId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "editElement", f)
		}
	}
}

func (ptr *ElementyModel) DisconnectEditElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "editElement")
	}
}

func (ptr *ElementyModel) EditElement(row int, elementGName string, elementStan float32, elementPName string, elementIlosc int, elementData string, elementId int) {
	if ptr.Pointer() != nil {
		var elementGNameC *C.char
		if elementGName != "" {
			elementGNameC = C.CString(elementGName)
			defer C.free(unsafe.Pointer(elementGNameC))
		}
		var elementPNameC *C.char
		if elementPName != "" {
			elementPNameC = C.CString(elementPName)
			defer C.free(unsafe.Pointer(elementPNameC))
		}
		var elementDataC *C.char
		if elementData != "" {
			elementDataC = C.CString(elementData)
			defer C.free(unsafe.Pointer(elementDataC))
		}
		C.ElementyModeld7c2e5_EditElement(ptr.Pointer(), C.int(int32(row)), C.struct_Moc_PackedString{data: elementGNameC, len: C.longlong(len(elementGName))}, C.float(elementStan), C.struct_Moc_PackedString{data: elementPNameC, len: C.longlong(len(elementPName))}, C.int(int32(elementIlosc)), C.struct_Moc_PackedString{data: elementDataC, len: C.longlong(len(elementData))}, C.int(int32(elementId)))
	}
}

//export callbackElementyModeld7c2e5_RemoveElement
func callbackElementyModeld7c2e5_RemoveElement(ptr unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "removeElement"); signal != nil {
		signal.(func(int))(int(int32(row)))
	}

}

func (ptr *ElementyModel) ConnectRemoveElement(f func(row int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeElement"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "removeElement", func(row int) {
				signal.(func(int))(row)
				f(row)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeElement", f)
		}
	}
}

func (ptr *ElementyModel) DisconnectRemoveElement() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeElement")
	}
}

func (ptr *ElementyModel) RemoveElement(row int) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_RemoveElement(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackElementyModeld7c2e5_RemoveElementId
func callbackElementyModeld7c2e5_RemoveElementId(ptr unsafe.Pointer, idin C.int) {
	if signal := qt.GetSignal(ptr, "removeElementId"); signal != nil {
		signal.(func(int))(int(int32(idin)))
	}

}

func (ptr *ElementyModel) ConnectRemoveElementId(f func(idin int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeElementId"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "removeElementId", func(idin int) {
				signal.(func(int))(idin)
				f(idin)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeElementId", f)
		}
	}
}

func (ptr *ElementyModel) DisconnectRemoveElementId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeElementId")
	}
}

func (ptr *ElementyModel) RemoveElementId(idin int) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_RemoveElementId(ptr.Pointer(), C.int(int32(idin)))
	}
}

//export callbackElementyModeld7c2e5_Roles
func callbackElementyModeld7c2e5_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__roles_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__roles_newList())
		for k, v := range NewElementyModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *ElementyModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "roles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "roles", func() map[int]*std_core.QByteArray {
				signal.(func() map[int]*std_core.QByteArray)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "roles", f)
		}
	}
}

func (ptr *ElementyModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *ElementyModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.ElementyModeld7c2e5_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *ElementyModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.ElementyModeld7c2e5_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackElementyModeld7c2e5_SetRoles
func callbackElementyModeld7c2e5_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewElementyModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *ElementyModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRoles"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", f)
		}
	}
}

func (ptr *ElementyModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *ElementyModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *ElementyModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackElementyModeld7c2e5_RolesChanged
func callbackElementyModeld7c2e5_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		signal.(func(map[int]*std_core.QByteArray))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *ElementyModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.ElementyModeld7c2e5_ConnectRolesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rolesChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", func(roles map[int]*std_core.QByteArray) {
				signal.(func(map[int]*std_core.QByteArray))(roles)
				f(roles)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", f)
		}
	}
}

func (ptr *ElementyModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *ElementyModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackElementyModeld7c2e5_Elementitem
func callbackElementyModeld7c2e5_Elementitem(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "elementitem"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__elementitem_newList())
			for _, v := range signal.(func() []*Element)() {
				tmpList.__elementitem_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__elementitem_newList())
		for _, v := range NewElementyModelFromPointer(ptr).ElementitemDefault() {
			tmpList.__elementitem_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *ElementyModel) ConnectElementitem(f func() []*Element) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "elementitem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementitem", func() []*Element {
				signal.(func() []*Element)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementitem", f)
		}
	}
}

func (ptr *ElementyModel) DisconnectElementitem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "elementitem")
	}
}

func (ptr *ElementyModel) Elementitem() []*Element {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Element {
			out := make([]*Element, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__elementitem_atList(i)
			}
			return out
		}(C.ElementyModeld7c2e5_Elementitem(ptr.Pointer()))
	}
	return make([]*Element, 0)
}

func (ptr *ElementyModel) ElementitemDefault() []*Element {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*Element {
			out := make([]*Element, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__elementitem_atList(i)
			}
			return out
		}(C.ElementyModeld7c2e5_ElementitemDefault(ptr.Pointer()))
	}
	return make([]*Element, 0)
}

//export callbackElementyModeld7c2e5_SetElementitem
func callbackElementyModeld7c2e5_SetElementitem(ptr unsafe.Pointer, elementitem C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setElementitem"); signal != nil {
		signal.(func([]*Element))(func(l C.struct_Moc_PackedList) []*Element {
			out := make([]*Element, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setElementitem_elementitem_atList(i)
			}
			return out
		}(elementitem))
	} else {
		NewElementyModelFromPointer(ptr).SetElementitemDefault(func(l C.struct_Moc_PackedList) []*Element {
			out := make([]*Element, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setElementitem_elementitem_atList(i)
			}
			return out
		}(elementitem))
	}
}

func (ptr *ElementyModel) ConnectSetElementitem(f func(elementitem []*Element)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setElementitem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setElementitem", func(elementitem []*Element) {
				signal.(func([]*Element))(elementitem)
				f(elementitem)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setElementitem", f)
		}
	}
}

func (ptr *ElementyModel) DisconnectSetElementitem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setElementitem")
	}
}

func (ptr *ElementyModel) SetElementitem(elementitem []*Element) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_SetElementitem(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__setElementitem_elementitem_newList())
			for _, v := range elementitem {
				tmpList.__setElementitem_elementitem_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *ElementyModel) SetElementitemDefault(elementitem []*Element) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_SetElementitemDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__setElementitem_elementitem_newList())
			for _, v := range elementitem {
				tmpList.__setElementitem_elementitem_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackElementyModeld7c2e5_ElementitemChanged
func callbackElementyModeld7c2e5_ElementitemChanged(ptr unsafe.Pointer, elementitem C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "elementitemChanged"); signal != nil {
		signal.(func([]*Element))(func(l C.struct_Moc_PackedList) []*Element {
			out := make([]*Element, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__elementitemChanged_elementitem_atList(i)
			}
			return out
		}(elementitem))
	}

}

func (ptr *ElementyModel) ConnectElementitemChanged(f func(elementitem []*Element)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "elementitemChanged") {
			C.ElementyModeld7c2e5_ConnectElementitemChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "elementitemChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "elementitemChanged", func(elementitem []*Element) {
				signal.(func([]*Element))(elementitem)
				f(elementitem)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "elementitemChanged", f)
		}
	}
}

func (ptr *ElementyModel) DisconnectElementitemChanged() {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_DisconnectElementitemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "elementitemChanged")
	}
}

func (ptr *ElementyModel) ElementitemChanged(elementitem []*Element) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_ElementitemChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__elementitemChanged_elementitem_newList())
			for _, v := range elementitem {
				tmpList.__elementitemChanged_elementitem_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func ElementyModel_QRegisterMetaType() int {
	return int(int32(C.ElementyModeld7c2e5_ElementyModeld7c2e5_QRegisterMetaType()))
}

func (ptr *ElementyModel) QRegisterMetaType() int {
	return int(int32(C.ElementyModeld7c2e5_ElementyModeld7c2e5_QRegisterMetaType()))
}

func ElementyModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ElementyModeld7c2e5_ElementyModeld7c2e5_QRegisterMetaType2(typeNameC)))
}

func (ptr *ElementyModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ElementyModeld7c2e5_ElementyModeld7c2e5_QRegisterMetaType2(typeNameC)))
}

func ElementyModel_QmlRegisterType() int {
	return int(int32(C.ElementyModeld7c2e5_ElementyModeld7c2e5_QmlRegisterType()))
}

func (ptr *ElementyModel) QmlRegisterType() int {
	return int(int32(C.ElementyModeld7c2e5_ElementyModeld7c2e5_QmlRegisterType()))
}

func ElementyModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ElementyModeld7c2e5_ElementyModeld7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ElementyModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.ElementyModeld7c2e5_ElementyModeld7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ElementyModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ElementyModeld7c2e5_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ElementyModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ElementyModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *ElementyModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ElementyModeld7c2e5_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ElementyModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ElementyModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *ElementyModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ElementyModeld7c2e5_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ElementyModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ElementyModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.ElementyModeld7c2e5___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *ElementyModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.ElementyModeld7c2e5___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ElementyModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ElementyModeld7c2e5___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *ElementyModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ElementyModeld7c2e5___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *ElementyModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ElementyModeld7c2e5___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ElementyModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ElementyModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.ElementyModeld7c2e5___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *ElementyModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.ElementyModeld7c2e5___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *ElementyModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ElementyModeld7c2e5___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ElementyModel) __roleNames_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___roleNames_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.ElementyModeld7c2e5___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ElementyModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.ElementyModeld7c2e5___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *ElementyModel) __itemData_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___itemData_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.ElementyModeld7c2e5___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ElementyModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ElementyModeld7c2e5___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *ElementyModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ElementyModeld7c2e5___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *ElementyModel) __match_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___match_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ElementyModeld7c2e5___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *ElementyModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *ElementyModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ElementyModeld7c2e5_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ElementyModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ElementyModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *ElementyModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ElementyModeld7c2e5_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ElementyModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ElementyModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ElementyModeld7c2e5___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ElementyModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ElementyModeld7c2e5___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ElementyModel) __findChildren_newList2() unsafe.Pointer {
	return C.ElementyModeld7c2e5___findChildren_newList2(ptr.Pointer())
}

func (ptr *ElementyModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ElementyModeld7c2e5___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ElementyModel) __findChildren_newList3() unsafe.Pointer {
	return C.ElementyModeld7c2e5___findChildren_newList3(ptr.Pointer())
}

func (ptr *ElementyModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ElementyModeld7c2e5___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ElementyModel) __findChildren_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___findChildren_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ElementyModeld7c2e5___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ElementyModel) __children_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___children_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ElementyModeld7c2e5___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ElementyModel) __roles_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___roles_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.ElementyModeld7c2e5___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ElementyModel) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ElementyModeld7c2e5___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ElementyModel) __setRoles_roles_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.ElementyModeld7c2e5___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ElementyModel) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ElementyModeld7c2e5___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ElementyModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.ElementyModeld7c2e5___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *ElementyModel) __elementitem_atList(i int) *Element {
	if ptr.Pointer() != nil {
		tmpValue := NewElementFromPointer(C.ElementyModeld7c2e5___elementitem_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __elementitem_setList(i Element_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___elementitem_setList(ptr.Pointer(), PointerFromElement(i))
	}
}

func (ptr *ElementyModel) __elementitem_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___elementitem_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __setElementitem_elementitem_atList(i int) *Element {
	if ptr.Pointer() != nil {
		tmpValue := NewElementFromPointer(C.ElementyModeld7c2e5___setElementitem_elementitem_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __setElementitem_elementitem_setList(i Element_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___setElementitem_elementitem_setList(ptr.Pointer(), PointerFromElement(i))
	}
}

func (ptr *ElementyModel) __setElementitem_elementitem_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___setElementitem_elementitem_newList(ptr.Pointer())
}

func (ptr *ElementyModel) __elementitemChanged_elementitem_atList(i int) *Element {
	if ptr.Pointer() != nil {
		tmpValue := NewElementFromPointer(C.ElementyModeld7c2e5___elementitemChanged_elementitem_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ElementyModel) __elementitemChanged_elementitem_setList(i Element_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5___elementitemChanged_elementitem_setList(ptr.Pointer(), PointerFromElement(i))
	}
}

func (ptr *ElementyModel) __elementitemChanged_elementitem_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5___elementitemChanged_elementitem_newList(ptr.Pointer())
}

func (ptr *ElementyModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ElementyModeld7c2e5_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ElementyModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ElementyModel) ____roles_keyList_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *ElementyModel) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ElementyModeld7c2e5_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ElementyModel) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ElementyModel) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *ElementyModel) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ElementyModeld7c2e5_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ElementyModel) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ElementyModel) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.ElementyModeld7c2e5_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewElementyModel(parent std_core.QObject_ITF) *ElementyModel {
	tmpValue := NewElementyModelFromPointer(C.ElementyModeld7c2e5_NewElementyModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackElementyModeld7c2e5_DestroyElementyModel
func callbackElementyModeld7c2e5_DestroyElementyModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~ElementyModel"); signal != nil {
		signal.(func())()
	} else {
		NewElementyModelFromPointer(ptr).DestroyElementyModelDefault()
	}
}

func (ptr *ElementyModel) ConnectDestroyElementyModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~ElementyModel"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~ElementyModel", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~ElementyModel", f)
		}
	}
}

func (ptr *ElementyModel) DisconnectDestroyElementyModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~ElementyModel")
	}
}

func (ptr *ElementyModel) DestroyElementyModel() {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_DestroyElementyModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ElementyModel) DestroyElementyModelDefault() {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_DestroyElementyModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackElementyModeld7c2e5_DropMimeData
func callbackElementyModeld7c2e5_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ElementyModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_Index
func callbackElementyModeld7c2e5_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewElementyModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *ElementyModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ElementyModeld7c2e5_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackElementyModeld7c2e5_Sibling
func callbackElementyModeld7c2e5_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewElementyModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *ElementyModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ElementyModeld7c2e5_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackElementyModeld7c2e5_Flags
func callbackElementyModeld7c2e5_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong(signal.(func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewElementyModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *ElementyModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.ElementyModeld7c2e5_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackElementyModeld7c2e5_InsertColumns
func callbackElementyModeld7c2e5_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ElementyModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_InsertRows
func callbackElementyModeld7c2e5_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ElementyModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_MoveColumns
func callbackElementyModeld7c2e5_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *ElementyModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_MoveRows
func callbackElementyModeld7c2e5_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *ElementyModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_RemoveColumns
func callbackElementyModeld7c2e5_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ElementyModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_RemoveRows
func callbackElementyModeld7c2e5_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *std_core.QModelIndex) bool)(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ElementyModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_SetData
func callbackElementyModeld7c2e5_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *ElementyModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_SetHeaderData
func callbackElementyModeld7c2e5_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *ElementyModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_SetItemData
func callbackElementyModeld7c2e5_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewElementyModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *ElementyModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_Submit
func callbackElementyModeld7c2e5_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *ElementyModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_ColumnsAboutToBeInserted
func callbackElementyModeld7c2e5_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackElementyModeld7c2e5_ColumnsAboutToBeMoved
func callbackElementyModeld7c2e5_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackElementyModeld7c2e5_ColumnsAboutToBeRemoved
func callbackElementyModeld7c2e5_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackElementyModeld7c2e5_ColumnsInserted
func callbackElementyModeld7c2e5_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackElementyModeld7c2e5_ColumnsMoved
func callbackElementyModeld7c2e5_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackElementyModeld7c2e5_ColumnsRemoved
func callbackElementyModeld7c2e5_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackElementyModeld7c2e5_DataChanged
func callbackElementyModeld7c2e5_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackElementyModeld7c2e5_FetchMore
func callbackElementyModeld7c2e5_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		signal.(func(*std_core.QModelIndex))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewElementyModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *ElementyModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackElementyModeld7c2e5_HeaderDataChanged
func callbackElementyModeld7c2e5_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		signal.(func(std_core.Qt__Orientation, int, int))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackElementyModeld7c2e5_LayoutAboutToBeChanged
func callbackElementyModeld7c2e5_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackElementyModeld7c2e5_LayoutChanged
func callbackElementyModeld7c2e5_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		signal.(func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackElementyModeld7c2e5_ModelAboutToBeReset
func callbackElementyModeld7c2e5_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackElementyModeld7c2e5_ModelReset
func callbackElementyModeld7c2e5_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		signal.(func())()
	}

}

//export callbackElementyModeld7c2e5_ResetInternalData
func callbackElementyModeld7c2e5_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewElementyModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *ElementyModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackElementyModeld7c2e5_Revert
func callbackElementyModeld7c2e5_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		signal.(func())()
	} else {
		NewElementyModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *ElementyModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_RevertDefault(ptr.Pointer())
	}
}

//export callbackElementyModeld7c2e5_RowsAboutToBeInserted
func callbackElementyModeld7c2e5_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackElementyModeld7c2e5_RowsAboutToBeMoved
func callbackElementyModeld7c2e5_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackElementyModeld7c2e5_RowsAboutToBeRemoved
func callbackElementyModeld7c2e5_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackElementyModeld7c2e5_RowsInserted
func callbackElementyModeld7c2e5_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackElementyModeld7c2e5_RowsMoved
func callbackElementyModeld7c2e5_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackElementyModeld7c2e5_RowsRemoved
func callbackElementyModeld7c2e5_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		signal.(func(*std_core.QModelIndex, int, int))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackElementyModeld7c2e5_Sort
func callbackElementyModeld7c2e5_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		signal.(func(int, std_core.Qt__SortOrder))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewElementyModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *ElementyModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackElementyModeld7c2e5_RoleNames
func callbackElementyModeld7c2e5_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__roleNames_newList())
			for k, v := range signal.(func() map[int]*std_core.QByteArray)() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewElementyModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *ElementyModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.ElementyModeld7c2e5_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackElementyModeld7c2e5_ItemData
func callbackElementyModeld7c2e5_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__itemData_newList())
			for k, v := range signal.(func(*std_core.QModelIndex) map[int]*std_core.QVariant)(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__itemData_newList())
		for k, v := range NewElementyModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *ElementyModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.ElementyModeld7c2e5_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackElementyModeld7c2e5_MimeData
func callbackElementyModeld7c2e5_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData(signal.(func([]*std_core.QModelIndex) *std_core.QMimeData)(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewElementyModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewElementyModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *ElementyModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.ElementyModeld7c2e5_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackElementyModeld7c2e5_Buddy
func callbackElementyModeld7c2e5_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewElementyModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *ElementyModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ElementyModeld7c2e5_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackElementyModeld7c2e5_Parent
func callbackElementyModeld7c2e5_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex(signal.(func(*std_core.QModelIndex) *std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewElementyModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *ElementyModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.ElementyModeld7c2e5_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackElementyModeld7c2e5_Match
func callbackElementyModeld7c2e5_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__match_newList())
			for _, v := range signal.(func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewElementyModelFromPointer(NewElementyModelFromPointer(nil).__match_newList())
		for _, v := range NewElementyModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *ElementyModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewElementyModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.ElementyModeld7c2e5_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackElementyModeld7c2e5_Span
func callbackElementyModeld7c2e5_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize(signal.(func(*std_core.QModelIndex) *std_core.QSize)(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewElementyModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *ElementyModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.ElementyModeld7c2e5_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackElementyModeld7c2e5_MimeTypes
func callbackElementyModeld7c2e5_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := signal.(func() []string)()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
	}
	tempVal := NewElementyModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "|")), len: C.longlong(len(strings.Join(tempVal, "|")))}
}

func (ptr *ElementyModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.ElementyModeld7c2e5_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackElementyModeld7c2e5_Data
func callbackElementyModeld7c2e5_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(*std_core.QModelIndex, int) *std_core.QVariant)(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewElementyModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *ElementyModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.ElementyModeld7c2e5_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackElementyModeld7c2e5_HeaderData
func callbackElementyModeld7c2e5_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewElementyModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *ElementyModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.ElementyModeld7c2e5_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackElementyModeld7c2e5_SupportedDragActions
func callbackElementyModeld7c2e5_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewElementyModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *ElementyModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.ElementyModeld7c2e5_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackElementyModeld7c2e5_SupportedDropActions
func callbackElementyModeld7c2e5_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__DropAction)())
	}

	return C.longlong(NewElementyModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *ElementyModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.ElementyModeld7c2e5_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackElementyModeld7c2e5_CanDropMimeData
func callbackElementyModeld7c2e5_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ElementyModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_CanFetchMore
func callbackElementyModeld7c2e5_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ElementyModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_HasChildren
func callbackElementyModeld7c2e5_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QModelIndex) bool)(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *ElementyModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_ColumnCount
func callbackElementyModeld7c2e5_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewElementyModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *ElementyModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ElementyModeld7c2e5_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackElementyModeld7c2e5_RowCount
func callbackElementyModeld7c2e5_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32(signal.(func(*std_core.QModelIndex) int)(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewElementyModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *ElementyModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ElementyModeld7c2e5_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackElementyModeld7c2e5_Event
func callbackElementyModeld7c2e5_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *ElementyModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_EventFilter
func callbackElementyModeld7c2e5_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewElementyModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ElementyModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ElementyModeld7c2e5_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackElementyModeld7c2e5_ChildEvent
func callbackElementyModeld7c2e5_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewElementyModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ElementyModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackElementyModeld7c2e5_ConnectNotify
func callbackElementyModeld7c2e5_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewElementyModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ElementyModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackElementyModeld7c2e5_CustomEvent
func callbackElementyModeld7c2e5_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewElementyModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ElementyModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackElementyModeld7c2e5_DeleteLater
func callbackElementyModeld7c2e5_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewElementyModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ElementyModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackElementyModeld7c2e5_Destroyed
func callbackElementyModeld7c2e5_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackElementyModeld7c2e5_DisconnectNotify
func callbackElementyModeld7c2e5_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewElementyModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ElementyModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackElementyModeld7c2e5_ObjectNameChanged
func callbackElementyModeld7c2e5_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackElementyModeld7c2e5_TimerEvent
func callbackElementyModeld7c2e5_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewElementyModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ElementyModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ElementyModeld7c2e5_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type Grupa_ITF interface {
	std_core.QObject_ITF
	Grupa_PTR() *Grupa
}

func (ptr *Grupa) Grupa_PTR() *Grupa {
	return ptr
}

func (ptr *Grupa) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Grupa) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromGrupa(ptr Grupa_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Grupa_PTR().Pointer()
	}
	return nil
}

func NewGrupaFromPointer(ptr unsafe.Pointer) (n *Grupa) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Grupa)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Grupa:
			n = deduced

		case *std_core.QObject:
			n = &Grupa{QObject: *deduced}

		default:
			n = new(Grupa)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackGrupad7c2e5_Constructor
func callbackGrupad7c2e5_Constructor(ptr unsafe.Pointer) {
	this := NewGrupaFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackGrupad7c2e5_GrupaName
func callbackGrupad7c2e5_GrupaName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "grupaName"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewGrupaFromPointer(ptr).GrupaNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Grupa) ConnectGrupaName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "grupaName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "grupaName", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "grupaName", f)
		}
	}
}

func (ptr *Grupa) DisconnectGrupaName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "grupaName")
	}
}

func (ptr *Grupa) GrupaName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Grupad7c2e5_GrupaName(ptr.Pointer()))
	}
	return ""
}

func (ptr *Grupa) GrupaNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Grupad7c2e5_GrupaNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackGrupad7c2e5_SetGrupaName
func callbackGrupad7c2e5_SetGrupaName(ptr unsafe.Pointer, grupaName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setGrupaName"); signal != nil {
		signal.(func(string))(cGoUnpackString(grupaName))
	} else {
		NewGrupaFromPointer(ptr).SetGrupaNameDefault(cGoUnpackString(grupaName))
	}
}

func (ptr *Grupa) ConnectSetGrupaName(f func(grupaName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGrupaName"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setGrupaName", func(grupaName string) {
				signal.(func(string))(grupaName)
				f(grupaName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGrupaName", f)
		}
	}
}

func (ptr *Grupa) DisconnectSetGrupaName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGrupaName")
	}
}

func (ptr *Grupa) SetGrupaName(grupaName string) {
	if ptr.Pointer() != nil {
		var grupaNameC *C.char
		if grupaName != "" {
			grupaNameC = C.CString(grupaName)
			defer C.free(unsafe.Pointer(grupaNameC))
		}
		C.Grupad7c2e5_SetGrupaName(ptr.Pointer(), C.struct_Moc_PackedString{data: grupaNameC, len: C.longlong(len(grupaName))})
	}
}

func (ptr *Grupa) SetGrupaNameDefault(grupaName string) {
	if ptr.Pointer() != nil {
		var grupaNameC *C.char
		if grupaName != "" {
			grupaNameC = C.CString(grupaName)
			defer C.free(unsafe.Pointer(grupaNameC))
		}
		C.Grupad7c2e5_SetGrupaNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: grupaNameC, len: C.longlong(len(grupaName))})
	}
}

//export callbackGrupad7c2e5_GrupaNameChanged
func callbackGrupad7c2e5_GrupaNameChanged(ptr unsafe.Pointer, grupaName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "grupaNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(grupaName))
	}

}

func (ptr *Grupa) ConnectGrupaNameChanged(f func(grupaName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "grupaNameChanged") {
			C.Grupad7c2e5_ConnectGrupaNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "grupaNameChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "grupaNameChanged", func(grupaName string) {
				signal.(func(string))(grupaName)
				f(grupaName)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "grupaNameChanged", f)
		}
	}
}

func (ptr *Grupa) DisconnectGrupaNameChanged() {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_DisconnectGrupaNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "grupaNameChanged")
	}
}

func (ptr *Grupa) GrupaNameChanged(grupaName string) {
	if ptr.Pointer() != nil {
		var grupaNameC *C.char
		if grupaName != "" {
			grupaNameC = C.CString(grupaName)
			defer C.free(unsafe.Pointer(grupaNameC))
		}
		C.Grupad7c2e5_GrupaNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: grupaNameC, len: C.longlong(len(grupaName))})
	}
}

//export callbackGrupad7c2e5_GrupaId
func callbackGrupad7c2e5_GrupaId(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "grupaId"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewGrupaFromPointer(ptr).GrupaIdDefault()))
}

func (ptr *Grupa) ConnectGrupaId(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "grupaId"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "grupaId", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "grupaId", f)
		}
	}
}

func (ptr *Grupa) DisconnectGrupaId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "grupaId")
	}
}

func (ptr *Grupa) GrupaId() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Grupad7c2e5_GrupaId(ptr.Pointer())))
	}
	return 0
}

func (ptr *Grupa) GrupaIdDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Grupad7c2e5_GrupaIdDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackGrupad7c2e5_SetGrupaId
func callbackGrupad7c2e5_SetGrupaId(ptr unsafe.Pointer, grupaId C.int) {
	if signal := qt.GetSignal(ptr, "setGrupaId"); signal != nil {
		signal.(func(int))(int(int32(grupaId)))
	} else {
		NewGrupaFromPointer(ptr).SetGrupaIdDefault(int(int32(grupaId)))
	}
}

func (ptr *Grupa) ConnectSetGrupaId(f func(grupaId int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setGrupaId"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setGrupaId", func(grupaId int) {
				signal.(func(int))(grupaId)
				f(grupaId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setGrupaId", f)
		}
	}
}

func (ptr *Grupa) DisconnectSetGrupaId() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setGrupaId")
	}
}

func (ptr *Grupa) SetGrupaId(grupaId int) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_SetGrupaId(ptr.Pointer(), C.int(int32(grupaId)))
	}
}

func (ptr *Grupa) SetGrupaIdDefault(grupaId int) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_SetGrupaIdDefault(ptr.Pointer(), C.int(int32(grupaId)))
	}
}

//export callbackGrupad7c2e5_GrupaIdChanged
func callbackGrupad7c2e5_GrupaIdChanged(ptr unsafe.Pointer, grupaId C.int) {
	if signal := qt.GetSignal(ptr, "grupaIdChanged"); signal != nil {
		signal.(func(int))(int(int32(grupaId)))
	}

}

func (ptr *Grupa) ConnectGrupaIdChanged(f func(grupaId int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "grupaIdChanged") {
			C.Grupad7c2e5_ConnectGrupaIdChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "grupaIdChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "grupaIdChanged", func(grupaId int) {
				signal.(func(int))(grupaId)
				f(grupaId)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "grupaIdChanged", f)
		}
	}
}

func (ptr *Grupa) DisconnectGrupaIdChanged() {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_DisconnectGrupaIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "grupaIdChanged")
	}
}

func (ptr *Grupa) GrupaIdChanged(grupaId int) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_GrupaIdChanged(ptr.Pointer(), C.int(int32(grupaId)))
	}
}

func Grupa_QRegisterMetaType() int {
	return int(int32(C.Grupad7c2e5_Grupad7c2e5_QRegisterMetaType()))
}

func (ptr *Grupa) QRegisterMetaType() int {
	return int(int32(C.Grupad7c2e5_Grupad7c2e5_QRegisterMetaType()))
}

func Grupa_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Grupad7c2e5_Grupad7c2e5_QRegisterMetaType2(typeNameC)))
}

func (ptr *Grupa) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Grupad7c2e5_Grupad7c2e5_QRegisterMetaType2(typeNameC)))
}

func Grupa_QmlRegisterType() int {
	return int(int32(C.Grupad7c2e5_Grupad7c2e5_QmlRegisterType()))
}

func (ptr *Grupa) QmlRegisterType() int {
	return int(int32(C.Grupad7c2e5_Grupad7c2e5_QmlRegisterType()))
}

func Grupa_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Grupad7c2e5_Grupad7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Grupa) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Grupad7c2e5_Grupad7c2e5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Grupa) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Grupad7c2e5___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Grupa) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Grupa) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Grupad7c2e5___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Grupa) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Grupad7c2e5___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Grupa) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Grupa) __findChildren_newList2() unsafe.Pointer {
	return C.Grupad7c2e5___findChildren_newList2(ptr.Pointer())
}

func (ptr *Grupa) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Grupad7c2e5___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Grupa) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Grupa) __findChildren_newList3() unsafe.Pointer {
	return C.Grupad7c2e5___findChildren_newList3(ptr.Pointer())
}

func (ptr *Grupa) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Grupad7c2e5___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Grupa) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Grupa) __findChildren_newList() unsafe.Pointer {
	return C.Grupad7c2e5___findChildren_newList(ptr.Pointer())
}

func (ptr *Grupa) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Grupad7c2e5___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Grupa) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Grupa) __children_newList() unsafe.Pointer {
	return C.Grupad7c2e5___children_newList(ptr.Pointer())
}

func NewGrupa(parent std_core.QObject_ITF) *Grupa {
	tmpValue := NewGrupaFromPointer(C.Grupad7c2e5_NewGrupa(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackGrupad7c2e5_DestroyGrupa
func callbackGrupad7c2e5_DestroyGrupa(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Grupa"); signal != nil {
		signal.(func())()
	} else {
		NewGrupaFromPointer(ptr).DestroyGrupaDefault()
	}
}

func (ptr *Grupa) ConnectDestroyGrupa(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Grupa"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Grupa", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Grupa", f)
		}
	}
}

func (ptr *Grupa) DisconnectDestroyGrupa() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Grupa")
	}
}

func (ptr *Grupa) DestroyGrupa() {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_DestroyGrupa(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Grupa) DestroyGrupaDefault() {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_DestroyGrupaDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackGrupad7c2e5_Event
func callbackGrupad7c2e5_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupaFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Grupa) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Grupad7c2e5_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackGrupad7c2e5_EventFilter
func callbackGrupad7c2e5_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewGrupaFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Grupa) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Grupad7c2e5_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackGrupad7c2e5_ChildEvent
func callbackGrupad7c2e5_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewGrupaFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Grupa) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackGrupad7c2e5_ConnectNotify
func callbackGrupad7c2e5_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewGrupaFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Grupa) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackGrupad7c2e5_CustomEvent
func callbackGrupad7c2e5_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewGrupaFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Grupa) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackGrupad7c2e5_DeleteLater
func callbackGrupad7c2e5_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewGrupaFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Grupa) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackGrupad7c2e5_Destroyed
func callbackGrupad7c2e5_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackGrupad7c2e5_DisconnectNotify
func callbackGrupad7c2e5_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewGrupaFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Grupa) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackGrupad7c2e5_ObjectNameChanged
func callbackGrupad7c2e5_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackGrupad7c2e5_TimerEvent
func callbackGrupad7c2e5_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewGrupaFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Grupa) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Grupad7c2e5_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

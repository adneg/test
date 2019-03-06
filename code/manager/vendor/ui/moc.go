package ui

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
	std_gui "github.com/therecipe/qt/gui"
	std_widgets "github.com/therecipe/qt/widgets"
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

type NormyOkno_ITF interface {
	std_widgets.QMainWindow_ITF
	NormyOkno_PTR() *NormyOkno
}

func (ptr *NormyOkno) NormyOkno_PTR() *NormyOkno {
	return ptr
}

func (ptr *NormyOkno) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QMainWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *NormyOkno) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QMainWindow_PTR().SetPointer(p)
	}
}

func PointerFromNormyOkno(ptr NormyOkno_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.NormyOkno_PTR().Pointer()
	}
	return nil
}

func NewNormyOknoFromPointer(ptr unsafe.Pointer) (n *NormyOkno) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(NormyOkno)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *NormyOkno:
			n = deduced

		case *std_widgets.QMainWindow:
			n = &NormyOkno{QMainWindow: *deduced}

		default:
			n = new(NormyOkno)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackNormyOkno42a7e9_Constructor
func callbackNormyOkno42a7e9_Constructor(ptr unsafe.Pointer) {
	this := NewNormyOknoFromPointer(ptr)
	qt.Register(ptr, this)
}

func NormyOkno_QRegisterMetaType() int {
	return int(int32(C.NormyOkno42a7e9_NormyOkno42a7e9_QRegisterMetaType()))
}

func (ptr *NormyOkno) QRegisterMetaType() int {
	return int(int32(C.NormyOkno42a7e9_NormyOkno42a7e9_QRegisterMetaType()))
}

func NormyOkno_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.NormyOkno42a7e9_NormyOkno42a7e9_QRegisterMetaType2(typeNameC)))
}

func (ptr *NormyOkno) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.NormyOkno42a7e9_NormyOkno42a7e9_QRegisterMetaType2(typeNameC)))
}

func NormyOkno_QmlRegisterType() int {
	return int(int32(C.NormyOkno42a7e9_NormyOkno42a7e9_QmlRegisterType()))
}

func (ptr *NormyOkno) QmlRegisterType() int {
	return int(int32(C.NormyOkno42a7e9_NormyOkno42a7e9_QmlRegisterType()))
}

func NormyOkno_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.NormyOkno42a7e9_NormyOkno42a7e9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *NormyOkno) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.NormyOkno42a7e9_NormyOkno42a7e9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *NormyOkno) __resizeDocks_docks_atList(i int) *std_widgets.QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQDockWidgetFromPointer(C.NormyOkno42a7e9___resizeDocks_docks_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *NormyOkno) __resizeDocks_docks_setList(i std_widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9___resizeDocks_docks_setList(ptr.Pointer(), std_widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *NormyOkno) __resizeDocks_docks_newList() unsafe.Pointer {
	return C.NormyOkno42a7e9___resizeDocks_docks_newList(ptr.Pointer())
}

func (ptr *NormyOkno) __resizeDocks_sizes_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.NormyOkno42a7e9___resizeDocks_sizes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *NormyOkno) __resizeDocks_sizes_setList(i int) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9___resizeDocks_sizes_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *NormyOkno) __resizeDocks_sizes_newList() unsafe.Pointer {
	return C.NormyOkno42a7e9___resizeDocks_sizes_newList(ptr.Pointer())
}

func (ptr *NormyOkno) __tabifiedDockWidgets_atList(i int) *std_widgets.QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQDockWidgetFromPointer(C.NormyOkno42a7e9___tabifiedDockWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *NormyOkno) __tabifiedDockWidgets_setList(i std_widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9___tabifiedDockWidgets_setList(ptr.Pointer(), std_widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *NormyOkno) __tabifiedDockWidgets_newList() unsafe.Pointer {
	return C.NormyOkno42a7e9___tabifiedDockWidgets_newList(ptr.Pointer())
}

func (ptr *NormyOkno) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.NormyOkno42a7e9___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *NormyOkno) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *NormyOkno) __addActions_actions_newList() unsafe.Pointer {
	return C.NormyOkno42a7e9___addActions_actions_newList(ptr.Pointer())
}

func (ptr *NormyOkno) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.NormyOkno42a7e9___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *NormyOkno) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *NormyOkno) __insertActions_actions_newList() unsafe.Pointer {
	return C.NormyOkno42a7e9___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *NormyOkno) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.NormyOkno42a7e9___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *NormyOkno) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *NormyOkno) __actions_newList() unsafe.Pointer {
	return C.NormyOkno42a7e9___actions_newList(ptr.Pointer())
}

func (ptr *NormyOkno) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.NormyOkno42a7e9___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *NormyOkno) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *NormyOkno) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.NormyOkno42a7e9___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *NormyOkno) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.NormyOkno42a7e9___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *NormyOkno) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *NormyOkno) __findChildren_newList2() unsafe.Pointer {
	return C.NormyOkno42a7e9___findChildren_newList2(ptr.Pointer())
}

func (ptr *NormyOkno) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.NormyOkno42a7e9___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *NormyOkno) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *NormyOkno) __findChildren_newList3() unsafe.Pointer {
	return C.NormyOkno42a7e9___findChildren_newList3(ptr.Pointer())
}

func (ptr *NormyOkno) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.NormyOkno42a7e9___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *NormyOkno) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *NormyOkno) __findChildren_newList() unsafe.Pointer {
	return C.NormyOkno42a7e9___findChildren_newList(ptr.Pointer())
}

func (ptr *NormyOkno) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.NormyOkno42a7e9___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *NormyOkno) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *NormyOkno) __children_newList() unsafe.Pointer {
	return C.NormyOkno42a7e9___children_newList(ptr.Pointer())
}

func NewNormyOkno(parent std_widgets.QWidget_ITF, flags std_core.Qt__WindowType) *NormyOkno {
	tmpValue := NewNormyOknoFromPointer(C.NormyOkno42a7e9_NewNormyOkno(std_widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *NormyOkno) DestroyNormyOkno() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_DestroyNormyOkno(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackNormyOkno42a7e9_CreatePopupMenu
func callbackNormyOkno42a7e9_CreatePopupMenu(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createPopupMenu"); signal != nil {
		return std_widgets.PointerFromQMenu(signal.(func() *std_widgets.QMenu)())
	}

	return std_widgets.PointerFromQMenu(NewNormyOknoFromPointer(ptr).CreatePopupMenuDefault())
}

func (ptr *NormyOkno) CreatePopupMenuDefault() *std_widgets.QMenu {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQMenuFromPointer(C.NormyOkno42a7e9_CreatePopupMenuDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackNormyOkno42a7e9_Event
func callbackNormyOkno42a7e9_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNormyOknoFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *NormyOkno) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.NormyOkno42a7e9_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackNormyOkno42a7e9_ContextMenuEvent
func callbackNormyOkno42a7e9_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*std_gui.QContextMenuEvent))(std_gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *NormyOkno) ContextMenuEventDefault(event std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackNormyOkno42a7e9_IconSizeChanged
func callbackNormyOkno42a7e9_IconSizeChanged(ptr unsafe.Pointer, iconSize unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconSizeChanged"); signal != nil {
		signal.(func(*std_core.QSize))(std_core.NewQSizeFromPointer(iconSize))
	}

}

//export callbackNormyOkno42a7e9_SetAnimated
func callbackNormyOkno42a7e9_SetAnimated(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(ptr, "setAnimated"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewNormyOknoFromPointer(ptr).SetAnimatedDefault(int8(enabled) != 0)
	}
}

func (ptr *NormyOkno) SetAnimatedDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_SetAnimatedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackNormyOkno42a7e9_SetDockNestingEnabled
func callbackNormyOkno42a7e9_SetDockNestingEnabled(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(ptr, "setDockNestingEnabled"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewNormyOknoFromPointer(ptr).SetDockNestingEnabledDefault(int8(enabled) != 0)
	}
}

func (ptr *NormyOkno) SetDockNestingEnabledDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_SetDockNestingEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackNormyOkno42a7e9_SetUnifiedTitleAndToolBarOnMac
func callbackNormyOkno42a7e9_SetUnifiedTitleAndToolBarOnMac(ptr unsafe.Pointer, set C.char) {
	if signal := qt.GetSignal(ptr, "setUnifiedTitleAndToolBarOnMac"); signal != nil {
		signal.(func(bool))(int8(set) != 0)
	} else {
		NewNormyOknoFromPointer(ptr).SetUnifiedTitleAndToolBarOnMacDefault(int8(set) != 0)
	}
}

func (ptr *NormyOkno) SetUnifiedTitleAndToolBarOnMacDefault(set bool) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_SetUnifiedTitleAndToolBarOnMacDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))
	}
}

//export callbackNormyOkno42a7e9_TabifiedDockWidgetActivated
func callbackNormyOkno42a7e9_TabifiedDockWidgetActivated(ptr unsafe.Pointer, dockWidget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabifiedDockWidgetActivated"); signal != nil {
		signal.(func(*std_widgets.QDockWidget))(std_widgets.NewQDockWidgetFromPointer(dockWidget))
	}

}

//export callbackNormyOkno42a7e9_ToolButtonStyleChanged
func callbackNormyOkno42a7e9_ToolButtonStyleChanged(ptr unsafe.Pointer, toolButtonStyle C.longlong) {
	if signal := qt.GetSignal(ptr, "toolButtonStyleChanged"); signal != nil {
		signal.(func(std_core.Qt__ToolButtonStyle))(std_core.Qt__ToolButtonStyle(toolButtonStyle))
	}

}

//export callbackNormyOkno42a7e9_Close
func callbackNormyOkno42a7e9_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewNormyOknoFromPointer(ptr).CloseDefault())))
}

func (ptr *NormyOkno) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.NormyOkno42a7e9_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackNormyOkno42a7e9_FocusNextPrevChild
func callbackNormyOkno42a7e9_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNormyOknoFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *NormyOkno) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.NormyOkno42a7e9_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackNormyOkno42a7e9_ActionEvent
func callbackNormyOkno42a7e9_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*std_gui.QActionEvent))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *NormyOkno) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackNormyOkno42a7e9_ChangeEvent
func callbackNormyOkno42a7e9_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *NormyOkno) ChangeEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackNormyOkno42a7e9_CloseEvent
func callbackNormyOkno42a7e9_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*std_gui.QCloseEvent))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *NormyOkno) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackNormyOkno42a7e9_CustomContextMenuRequested
func callbackNormyOkno42a7e9_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*std_core.QPoint))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackNormyOkno42a7e9_DragEnterEvent
func callbackNormyOkno42a7e9_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*std_gui.QDragEnterEvent))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *NormyOkno) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackNormyOkno42a7e9_DragLeaveEvent
func callbackNormyOkno42a7e9_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*std_gui.QDragLeaveEvent))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *NormyOkno) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackNormyOkno42a7e9_DragMoveEvent
func callbackNormyOkno42a7e9_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*std_gui.QDragMoveEvent))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *NormyOkno) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackNormyOkno42a7e9_DropEvent
func callbackNormyOkno42a7e9_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*std_gui.QDropEvent))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *NormyOkno) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackNormyOkno42a7e9_EnterEvent
func callbackNormyOkno42a7e9_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *NormyOkno) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackNormyOkno42a7e9_FocusInEvent
func callbackNormyOkno42a7e9_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *NormyOkno) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackNormyOkno42a7e9_FocusOutEvent
func callbackNormyOkno42a7e9_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *NormyOkno) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackNormyOkno42a7e9_Hide
func callbackNormyOkno42a7e9_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).HideDefault()
	}
}

func (ptr *NormyOkno) HideDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_HideDefault(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_HideEvent
func callbackNormyOkno42a7e9_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*std_gui.QHideEvent))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *NormyOkno) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackNormyOkno42a7e9_InputMethodEvent
func callbackNormyOkno42a7e9_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*std_gui.QInputMethodEvent))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *NormyOkno) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackNormyOkno42a7e9_KeyPressEvent
func callbackNormyOkno42a7e9_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *NormyOkno) KeyPressEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackNormyOkno42a7e9_KeyReleaseEvent
func callbackNormyOkno42a7e9_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *NormyOkno) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackNormyOkno42a7e9_LeaveEvent
func callbackNormyOkno42a7e9_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *NormyOkno) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackNormyOkno42a7e9_Lower
func callbackNormyOkno42a7e9_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).LowerDefault()
	}
}

func (ptr *NormyOkno) LowerDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_LowerDefault(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_MouseDoubleClickEvent
func callbackNormyOkno42a7e9_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *NormyOkno) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackNormyOkno42a7e9_MouseMoveEvent
func callbackNormyOkno42a7e9_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *NormyOkno) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackNormyOkno42a7e9_MousePressEvent
func callbackNormyOkno42a7e9_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *NormyOkno) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackNormyOkno42a7e9_MouseReleaseEvent
func callbackNormyOkno42a7e9_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *NormyOkno) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackNormyOkno42a7e9_MoveEvent
func callbackNormyOkno42a7e9_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*std_gui.QMoveEvent))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *NormyOkno) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackNormyOkno42a7e9_PaintEvent
func callbackNormyOkno42a7e9_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*std_gui.QPaintEvent))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *NormyOkno) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackNormyOkno42a7e9_Raise
func callbackNormyOkno42a7e9_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *NormyOkno) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_RaiseDefault(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_Repaint
func callbackNormyOkno42a7e9_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *NormyOkno) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_RepaintDefault(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_ResizeEvent
func callbackNormyOkno42a7e9_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*std_gui.QResizeEvent))(std_gui.NewQResizeEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *NormyOkno) ResizeEventDefault(event std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(event))
	}
}

//export callbackNormyOkno42a7e9_SetDisabled
func callbackNormyOkno42a7e9_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewNormyOknoFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *NormyOkno) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackNormyOkno42a7e9_SetEnabled
func callbackNormyOkno42a7e9_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewNormyOknoFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *NormyOkno) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackNormyOkno42a7e9_SetFocus2
func callbackNormyOkno42a7e9_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *NormyOkno) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_SetHidden
func callbackNormyOkno42a7e9_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewNormyOknoFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *NormyOkno) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackNormyOkno42a7e9_SetStyleSheet
func callbackNormyOkno42a7e9_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewNormyOknoFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *NormyOkno) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.NormyOkno42a7e9_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackNormyOkno42a7e9_SetVisible
func callbackNormyOkno42a7e9_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewNormyOknoFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *NormyOkno) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackNormyOkno42a7e9_SetWindowModified
func callbackNormyOkno42a7e9_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewNormyOknoFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *NormyOkno) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackNormyOkno42a7e9_SetWindowTitle
func callbackNormyOkno42a7e9_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewNormyOknoFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *NormyOkno) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.NormyOkno42a7e9_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackNormyOkno42a7e9_Show
func callbackNormyOkno42a7e9_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).ShowDefault()
	}
}

func (ptr *NormyOkno) ShowDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ShowDefault(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_ShowEvent
func callbackNormyOkno42a7e9_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*std_gui.QShowEvent))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *NormyOkno) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackNormyOkno42a7e9_ShowFullScreen
func callbackNormyOkno42a7e9_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *NormyOkno) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_ShowMaximized
func callbackNormyOkno42a7e9_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *NormyOkno) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_ShowMinimized
func callbackNormyOkno42a7e9_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *NormyOkno) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_ShowNormal
func callbackNormyOkno42a7e9_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *NormyOkno) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_TabletEvent
func callbackNormyOkno42a7e9_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*std_gui.QTabletEvent))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *NormyOkno) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackNormyOkno42a7e9_Update
func callbackNormyOkno42a7e9_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *NormyOkno) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_UpdateDefault(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_UpdateMicroFocus
func callbackNormyOkno42a7e9_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *NormyOkno) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackNormyOkno42a7e9_WheelEvent
func callbackNormyOkno42a7e9_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*std_gui.QWheelEvent))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *NormyOkno) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackNormyOkno42a7e9_WindowIconChanged
func callbackNormyOkno42a7e9_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*std_gui.QIcon))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackNormyOkno42a7e9_WindowTitleChanged
func callbackNormyOkno42a7e9_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackNormyOkno42a7e9_PaintEngine
func callbackNormyOkno42a7e9_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine(signal.(func() *std_gui.QPaintEngine)())
	}

	return std_gui.PointerFromQPaintEngine(NewNormyOknoFromPointer(ptr).PaintEngineDefault())
}

func (ptr *NormyOkno) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.NormyOkno42a7e9_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackNormyOkno42a7e9_MinimumSizeHint
func callbackNormyOkno42a7e9_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewNormyOknoFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *NormyOkno) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.NormyOkno42a7e9_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackNormyOkno42a7e9_SizeHint
func callbackNormyOkno42a7e9_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewNormyOknoFromPointer(ptr).SizeHintDefault())
}

func (ptr *NormyOkno) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.NormyOkno42a7e9_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackNormyOkno42a7e9_InputMethodQuery
func callbackNormyOkno42a7e9_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewNormyOknoFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *NormyOkno) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.NormyOkno42a7e9_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackNormyOkno42a7e9_HasHeightForWidth
func callbackNormyOkno42a7e9_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewNormyOknoFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *NormyOkno) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.NormyOkno42a7e9_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackNormyOkno42a7e9_HeightForWidth
func callbackNormyOkno42a7e9_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewNormyOknoFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *NormyOkno) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.NormyOkno42a7e9_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackNormyOkno42a7e9_Metric
func callbackNormyOkno42a7e9_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(std_gui.QPaintDevice__PaintDeviceMetric) int)(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewNormyOknoFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *NormyOkno) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.NormyOkno42a7e9_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackNormyOkno42a7e9_EventFilter
func callbackNormyOkno42a7e9_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNormyOknoFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *NormyOkno) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.NormyOkno42a7e9_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackNormyOkno42a7e9_ChildEvent
func callbackNormyOkno42a7e9_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *NormyOkno) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackNormyOkno42a7e9_ConnectNotify
func callbackNormyOkno42a7e9_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewNormyOknoFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *NormyOkno) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackNormyOkno42a7e9_CustomEvent
func callbackNormyOkno42a7e9_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *NormyOkno) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackNormyOkno42a7e9_DeleteLater
func callbackNormyOkno42a7e9_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewNormyOknoFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *NormyOkno) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackNormyOkno42a7e9_Destroyed
func callbackNormyOkno42a7e9_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackNormyOkno42a7e9_DisconnectNotify
func callbackNormyOkno42a7e9_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewNormyOknoFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *NormyOkno) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackNormyOkno42a7e9_ObjectNameChanged
func callbackNormyOkno42a7e9_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackNormyOkno42a7e9_TimerEvent
func callbackNormyOkno42a7e9_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewNormyOknoFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *NormyOkno) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.NormyOkno42a7e9_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type Pilot_ITF interface {
	std_widgets.QMainWindow_ITF
	Pilot_PTR() *Pilot
}

func (ptr *Pilot) Pilot_PTR() *Pilot {
	return ptr
}

func (ptr *Pilot) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QMainWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *Pilot) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QMainWindow_PTR().SetPointer(p)
	}
}

func PointerFromPilot(ptr Pilot_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Pilot_PTR().Pointer()
	}
	return nil
}

func NewPilotFromPointer(ptr unsafe.Pointer) (n *Pilot) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Pilot)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Pilot:
			n = deduced

		case *std_widgets.QMainWindow:
			n = &Pilot{QMainWindow: *deduced}

		default:
			n = new(Pilot)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackPilot42a7e9_Constructor
func callbackPilot42a7e9_Constructor(ptr unsafe.Pointer) {
	this := NewPilotFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackPilot42a7e9_AddTreeItem
func callbackPilot42a7e9_AddTreeItem(ptr unsafe.Pointer, p0 C.struct_Moc_PackedString, p1 C.struct_Moc_PackedString, p2 C.struct_Moc_PackedString, p3 C.struct_Moc_PackedString, p4 C.struct_Moc_PackedString, p5 C.struct_Moc_PackedString, p6 C.struct_Moc_PackedString, p7 C.struct_Moc_PackedString, p8 C.struct_Moc_PackedString, p9 C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addTreeItem"); signal != nil {
		signal.(func(string, string, string, string, string, string, string, string, string, string))(cGoUnpackString(p0), cGoUnpackString(p1), cGoUnpackString(p2), cGoUnpackString(p3), cGoUnpackString(p4), cGoUnpackString(p5), cGoUnpackString(p6), cGoUnpackString(p7), cGoUnpackString(p8), cGoUnpackString(p9))
	}

}

func (ptr *Pilot) ConnectAddTreeItem(f func(p0 string, p1 string, p2 string, p3 string, p4 string, p5 string, p6 string, p7 string, p8 string, p9 string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addTreeItem"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "addTreeItem", func(p0 string, p1 string, p2 string, p3 string, p4 string, p5 string, p6 string, p7 string, p8 string, p9 string) {
				signal.(func(string, string, string, string, string, string, string, string, string, string))(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
				f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addTreeItem", f)
		}
	}
}

func (ptr *Pilot) DisconnectAddTreeItem() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addTreeItem")
	}
}

func (ptr *Pilot) AddTreeItem(p0 string, p1 string, p2 string, p3 string, p4 string, p5 string, p6 string, p7 string, p8 string, p9 string) {
	if ptr.Pointer() != nil {
		var p0C *C.char
		if p0 != "" {
			p0C = C.CString(p0)
			defer C.free(unsafe.Pointer(p0C))
		}
		var p1C *C.char
		if p1 != "" {
			p1C = C.CString(p1)
			defer C.free(unsafe.Pointer(p1C))
		}
		var p2C *C.char
		if p2 != "" {
			p2C = C.CString(p2)
			defer C.free(unsafe.Pointer(p2C))
		}
		var p3C *C.char
		if p3 != "" {
			p3C = C.CString(p3)
			defer C.free(unsafe.Pointer(p3C))
		}
		var p4C *C.char
		if p4 != "" {
			p4C = C.CString(p4)
			defer C.free(unsafe.Pointer(p4C))
		}
		var p5C *C.char
		if p5 != "" {
			p5C = C.CString(p5)
			defer C.free(unsafe.Pointer(p5C))
		}
		var p6C *C.char
		if p6 != "" {
			p6C = C.CString(p6)
			defer C.free(unsafe.Pointer(p6C))
		}
		var p7C *C.char
		if p7 != "" {
			p7C = C.CString(p7)
			defer C.free(unsafe.Pointer(p7C))
		}
		var p8C *C.char
		if p8 != "" {
			p8C = C.CString(p8)
			defer C.free(unsafe.Pointer(p8C))
		}
		var p9C *C.char
		if p9 != "" {
			p9C = C.CString(p9)
			defer C.free(unsafe.Pointer(p9C))
		}
		C.Pilot42a7e9_AddTreeItem(ptr.Pointer(), C.struct_Moc_PackedString{data: p0C, len: C.longlong(len(p0))}, C.struct_Moc_PackedString{data: p1C, len: C.longlong(len(p1))}, C.struct_Moc_PackedString{data: p2C, len: C.longlong(len(p2))}, C.struct_Moc_PackedString{data: p3C, len: C.longlong(len(p3))}, C.struct_Moc_PackedString{data: p4C, len: C.longlong(len(p4))}, C.struct_Moc_PackedString{data: p5C, len: C.longlong(len(p5))}, C.struct_Moc_PackedString{data: p6C, len: C.longlong(len(p6))}, C.struct_Moc_PackedString{data: p7C, len: C.longlong(len(p7))}, C.struct_Moc_PackedString{data: p8C, len: C.longlong(len(p8))}, C.struct_Moc_PackedString{data: p9C, len: C.longlong(len(p9))})
	}
}

func Pilot_QRegisterMetaType() int {
	return int(int32(C.Pilot42a7e9_Pilot42a7e9_QRegisterMetaType()))
}

func (ptr *Pilot) QRegisterMetaType() int {
	return int(int32(C.Pilot42a7e9_Pilot42a7e9_QRegisterMetaType()))
}

func Pilot_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Pilot42a7e9_Pilot42a7e9_QRegisterMetaType2(typeNameC)))
}

func (ptr *Pilot) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Pilot42a7e9_Pilot42a7e9_QRegisterMetaType2(typeNameC)))
}

func Pilot_QmlRegisterType() int {
	return int(int32(C.Pilot42a7e9_Pilot42a7e9_QmlRegisterType()))
}

func (ptr *Pilot) QmlRegisterType() int {
	return int(int32(C.Pilot42a7e9_Pilot42a7e9_QmlRegisterType()))
}

func Pilot_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Pilot42a7e9_Pilot42a7e9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Pilot) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Pilot42a7e9_Pilot42a7e9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Pilot) __resizeDocks_docks_atList(i int) *std_widgets.QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQDockWidgetFromPointer(C.Pilot42a7e9___resizeDocks_docks_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Pilot) __resizeDocks_docks_setList(i std_widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9___resizeDocks_docks_setList(ptr.Pointer(), std_widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *Pilot) __resizeDocks_docks_newList() unsafe.Pointer {
	return C.Pilot42a7e9___resizeDocks_docks_newList(ptr.Pointer())
}

func (ptr *Pilot) __resizeDocks_sizes_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.Pilot42a7e9___resizeDocks_sizes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *Pilot) __resizeDocks_sizes_setList(i int) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9___resizeDocks_sizes_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *Pilot) __resizeDocks_sizes_newList() unsafe.Pointer {
	return C.Pilot42a7e9___resizeDocks_sizes_newList(ptr.Pointer())
}

func (ptr *Pilot) __tabifiedDockWidgets_atList(i int) *std_widgets.QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQDockWidgetFromPointer(C.Pilot42a7e9___tabifiedDockWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Pilot) __tabifiedDockWidgets_setList(i std_widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9___tabifiedDockWidgets_setList(ptr.Pointer(), std_widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *Pilot) __tabifiedDockWidgets_newList() unsafe.Pointer {
	return C.Pilot42a7e9___tabifiedDockWidgets_newList(ptr.Pointer())
}

func (ptr *Pilot) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.Pilot42a7e9___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Pilot) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *Pilot) __addActions_actions_newList() unsafe.Pointer {
	return C.Pilot42a7e9___addActions_actions_newList(ptr.Pointer())
}

func (ptr *Pilot) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.Pilot42a7e9___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Pilot) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *Pilot) __insertActions_actions_newList() unsafe.Pointer {
	return C.Pilot42a7e9___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *Pilot) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.Pilot42a7e9___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Pilot) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *Pilot) __actions_newList() unsafe.Pointer {
	return C.Pilot42a7e9___actions_newList(ptr.Pointer())
}

func (ptr *Pilot) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.Pilot42a7e9___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Pilot) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Pilot) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Pilot42a7e9___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Pilot) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Pilot42a7e9___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Pilot) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Pilot) __findChildren_newList2() unsafe.Pointer {
	return C.Pilot42a7e9___findChildren_newList2(ptr.Pointer())
}

func (ptr *Pilot) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Pilot42a7e9___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Pilot) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Pilot) __findChildren_newList3() unsafe.Pointer {
	return C.Pilot42a7e9___findChildren_newList3(ptr.Pointer())
}

func (ptr *Pilot) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Pilot42a7e9___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Pilot) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Pilot) __findChildren_newList() unsafe.Pointer {
	return C.Pilot42a7e9___findChildren_newList(ptr.Pointer())
}

func (ptr *Pilot) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.Pilot42a7e9___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Pilot) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Pilot) __children_newList() unsafe.Pointer {
	return C.Pilot42a7e9___children_newList(ptr.Pointer())
}

func NewPilot(parent std_widgets.QWidget_ITF, flags std_core.Qt__WindowType) *Pilot {
	tmpValue := NewPilotFromPointer(C.Pilot42a7e9_NewPilot(std_widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Pilot) DestroyPilot() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_DestroyPilot(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackPilot42a7e9_CreatePopupMenu
func callbackPilot42a7e9_CreatePopupMenu(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createPopupMenu"); signal != nil {
		return std_widgets.PointerFromQMenu(signal.(func() *std_widgets.QMenu)())
	}

	return std_widgets.PointerFromQMenu(NewPilotFromPointer(ptr).CreatePopupMenuDefault())
}

func (ptr *Pilot) CreatePopupMenuDefault() *std_widgets.QMenu {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQMenuFromPointer(C.Pilot42a7e9_CreatePopupMenuDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackPilot42a7e9_Event
func callbackPilot42a7e9_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPilotFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *Pilot) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Pilot42a7e9_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackPilot42a7e9_ContextMenuEvent
func callbackPilot42a7e9_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*std_gui.QContextMenuEvent))(std_gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *Pilot) ContextMenuEventDefault(event std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackPilot42a7e9_IconSizeChanged
func callbackPilot42a7e9_IconSizeChanged(ptr unsafe.Pointer, iconSize unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconSizeChanged"); signal != nil {
		signal.(func(*std_core.QSize))(std_core.NewQSizeFromPointer(iconSize))
	}

}

//export callbackPilot42a7e9_SetAnimated
func callbackPilot42a7e9_SetAnimated(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(ptr, "setAnimated"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewPilotFromPointer(ptr).SetAnimatedDefault(int8(enabled) != 0)
	}
}

func (ptr *Pilot) SetAnimatedDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_SetAnimatedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackPilot42a7e9_SetDockNestingEnabled
func callbackPilot42a7e9_SetDockNestingEnabled(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(ptr, "setDockNestingEnabled"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewPilotFromPointer(ptr).SetDockNestingEnabledDefault(int8(enabled) != 0)
	}
}

func (ptr *Pilot) SetDockNestingEnabledDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_SetDockNestingEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackPilot42a7e9_SetUnifiedTitleAndToolBarOnMac
func callbackPilot42a7e9_SetUnifiedTitleAndToolBarOnMac(ptr unsafe.Pointer, set C.char) {
	if signal := qt.GetSignal(ptr, "setUnifiedTitleAndToolBarOnMac"); signal != nil {
		signal.(func(bool))(int8(set) != 0)
	} else {
		NewPilotFromPointer(ptr).SetUnifiedTitleAndToolBarOnMacDefault(int8(set) != 0)
	}
}

func (ptr *Pilot) SetUnifiedTitleAndToolBarOnMacDefault(set bool) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_SetUnifiedTitleAndToolBarOnMacDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))
	}
}

//export callbackPilot42a7e9_TabifiedDockWidgetActivated
func callbackPilot42a7e9_TabifiedDockWidgetActivated(ptr unsafe.Pointer, dockWidget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabifiedDockWidgetActivated"); signal != nil {
		signal.(func(*std_widgets.QDockWidget))(std_widgets.NewQDockWidgetFromPointer(dockWidget))
	}

}

//export callbackPilot42a7e9_ToolButtonStyleChanged
func callbackPilot42a7e9_ToolButtonStyleChanged(ptr unsafe.Pointer, toolButtonStyle C.longlong) {
	if signal := qt.GetSignal(ptr, "toolButtonStyleChanged"); signal != nil {
		signal.(func(std_core.Qt__ToolButtonStyle))(std_core.Qt__ToolButtonStyle(toolButtonStyle))
	}

}

//export callbackPilot42a7e9_Close
func callbackPilot42a7e9_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewPilotFromPointer(ptr).CloseDefault())))
}

func (ptr *Pilot) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.Pilot42a7e9_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackPilot42a7e9_FocusNextPrevChild
func callbackPilot42a7e9_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPilotFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *Pilot) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.Pilot42a7e9_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackPilot42a7e9_ActionEvent
func callbackPilot42a7e9_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*std_gui.QActionEvent))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *Pilot) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackPilot42a7e9_ChangeEvent
func callbackPilot42a7e9_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Pilot) ChangeEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPilot42a7e9_CloseEvent
func callbackPilot42a7e9_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*std_gui.QCloseEvent))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *Pilot) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackPilot42a7e9_CustomContextMenuRequested
func callbackPilot42a7e9_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*std_core.QPoint))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackPilot42a7e9_DragEnterEvent
func callbackPilot42a7e9_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*std_gui.QDragEnterEvent))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *Pilot) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackPilot42a7e9_DragLeaveEvent
func callbackPilot42a7e9_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*std_gui.QDragLeaveEvent))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *Pilot) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackPilot42a7e9_DragMoveEvent
func callbackPilot42a7e9_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*std_gui.QDragMoveEvent))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *Pilot) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackPilot42a7e9_DropEvent
func callbackPilot42a7e9_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*std_gui.QDropEvent))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *Pilot) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackPilot42a7e9_EnterEvent
func callbackPilot42a7e9_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Pilot) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPilot42a7e9_FocusInEvent
func callbackPilot42a7e9_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *Pilot) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackPilot42a7e9_FocusOutEvent
func callbackPilot42a7e9_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *Pilot) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackPilot42a7e9_Hide
func callbackPilot42a7e9_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).HideDefault()
	}
}

func (ptr *Pilot) HideDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_HideDefault(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_HideEvent
func callbackPilot42a7e9_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*std_gui.QHideEvent))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *Pilot) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackPilot42a7e9_InputMethodEvent
func callbackPilot42a7e9_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*std_gui.QInputMethodEvent))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *Pilot) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackPilot42a7e9_KeyPressEvent
func callbackPilot42a7e9_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *Pilot) KeyPressEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackPilot42a7e9_KeyReleaseEvent
func callbackPilot42a7e9_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *Pilot) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackPilot42a7e9_LeaveEvent
func callbackPilot42a7e9_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Pilot) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPilot42a7e9_Lower
func callbackPilot42a7e9_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).LowerDefault()
	}
}

func (ptr *Pilot) LowerDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_LowerDefault(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_MouseDoubleClickEvent
func callbackPilot42a7e9_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *Pilot) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackPilot42a7e9_MouseMoveEvent
func callbackPilot42a7e9_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *Pilot) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackPilot42a7e9_MousePressEvent
func callbackPilot42a7e9_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *Pilot) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackPilot42a7e9_MouseReleaseEvent
func callbackPilot42a7e9_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *Pilot) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackPilot42a7e9_MoveEvent
func callbackPilot42a7e9_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*std_gui.QMoveEvent))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *Pilot) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackPilot42a7e9_PaintEvent
func callbackPilot42a7e9_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*std_gui.QPaintEvent))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *Pilot) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackPilot42a7e9_Raise
func callbackPilot42a7e9_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *Pilot) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_RaiseDefault(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_Repaint
func callbackPilot42a7e9_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *Pilot) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_RepaintDefault(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_ResizeEvent
func callbackPilot42a7e9_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*std_gui.QResizeEvent))(std_gui.NewQResizeEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *Pilot) ResizeEventDefault(event std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(event))
	}
}

//export callbackPilot42a7e9_SetDisabled
func callbackPilot42a7e9_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewPilotFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *Pilot) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackPilot42a7e9_SetEnabled
func callbackPilot42a7e9_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewPilotFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *Pilot) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackPilot42a7e9_SetFocus2
func callbackPilot42a7e9_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *Pilot) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_SetHidden
func callbackPilot42a7e9_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewPilotFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *Pilot) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackPilot42a7e9_SetStyleSheet
func callbackPilot42a7e9_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewPilotFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *Pilot) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.Pilot42a7e9_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackPilot42a7e9_SetVisible
func callbackPilot42a7e9_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewPilotFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *Pilot) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackPilot42a7e9_SetWindowModified
func callbackPilot42a7e9_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewPilotFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *Pilot) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackPilot42a7e9_SetWindowTitle
func callbackPilot42a7e9_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewPilotFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *Pilot) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.Pilot42a7e9_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackPilot42a7e9_Show
func callbackPilot42a7e9_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).ShowDefault()
	}
}

func (ptr *Pilot) ShowDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ShowDefault(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_ShowEvent
func callbackPilot42a7e9_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*std_gui.QShowEvent))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *Pilot) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackPilot42a7e9_ShowFullScreen
func callbackPilot42a7e9_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *Pilot) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_ShowMaximized
func callbackPilot42a7e9_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *Pilot) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_ShowMinimized
func callbackPilot42a7e9_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *Pilot) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_ShowNormal
func callbackPilot42a7e9_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *Pilot) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_TabletEvent
func callbackPilot42a7e9_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*std_gui.QTabletEvent))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *Pilot) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackPilot42a7e9_Update
func callbackPilot42a7e9_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *Pilot) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_UpdateDefault(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_UpdateMicroFocus
func callbackPilot42a7e9_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *Pilot) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackPilot42a7e9_WheelEvent
func callbackPilot42a7e9_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*std_gui.QWheelEvent))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *Pilot) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackPilot42a7e9_WindowIconChanged
func callbackPilot42a7e9_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*std_gui.QIcon))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackPilot42a7e9_WindowTitleChanged
func callbackPilot42a7e9_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackPilot42a7e9_PaintEngine
func callbackPilot42a7e9_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine(signal.(func() *std_gui.QPaintEngine)())
	}

	return std_gui.PointerFromQPaintEngine(NewPilotFromPointer(ptr).PaintEngineDefault())
}

func (ptr *Pilot) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.Pilot42a7e9_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackPilot42a7e9_MinimumSizeHint
func callbackPilot42a7e9_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewPilotFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *Pilot) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.Pilot42a7e9_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackPilot42a7e9_SizeHint
func callbackPilot42a7e9_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewPilotFromPointer(ptr).SizeHintDefault())
}

func (ptr *Pilot) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.Pilot42a7e9_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackPilot42a7e9_InputMethodQuery
func callbackPilot42a7e9_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewPilotFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *Pilot) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.Pilot42a7e9_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackPilot42a7e9_HasHeightForWidth
func callbackPilot42a7e9_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewPilotFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *Pilot) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.Pilot42a7e9_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackPilot42a7e9_HeightForWidth
func callbackPilot42a7e9_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewPilotFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *Pilot) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.Pilot42a7e9_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackPilot42a7e9_Metric
func callbackPilot42a7e9_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(std_gui.QPaintDevice__PaintDeviceMetric) int)(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewPilotFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *Pilot) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.Pilot42a7e9_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackPilot42a7e9_EventFilter
func callbackPilot42a7e9_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPilotFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Pilot) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.Pilot42a7e9_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackPilot42a7e9_ChildEvent
func callbackPilot42a7e9_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Pilot) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackPilot42a7e9_ConnectNotify
func callbackPilot42a7e9_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPilotFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Pilot) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPilot42a7e9_CustomEvent
func callbackPilot42a7e9_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Pilot) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPilot42a7e9_DeleteLater
func callbackPilot42a7e9_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewPilotFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Pilot) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackPilot42a7e9_Destroyed
func callbackPilot42a7e9_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackPilot42a7e9_DisconnectNotify
func callbackPilot42a7e9_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPilotFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Pilot) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPilot42a7e9_ObjectNameChanged
func callbackPilot42a7e9_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackPilot42a7e9_TimerEvent
func callbackPilot42a7e9_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewPilotFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Pilot) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Pilot42a7e9_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type ProduktyOkno_ITF interface {
	std_widgets.QMainWindow_ITF
	ProduktyOkno_PTR() *ProduktyOkno
}

func (ptr *ProduktyOkno) ProduktyOkno_PTR() *ProduktyOkno {
	return ptr
}

func (ptr *ProduktyOkno) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QMainWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *ProduktyOkno) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QMainWindow_PTR().SetPointer(p)
	}
}

func PointerFromProduktyOkno(ptr ProduktyOkno_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ProduktyOkno_PTR().Pointer()
	}
	return nil
}

func NewProduktyOknoFromPointer(ptr unsafe.Pointer) (n *ProduktyOkno) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ProduktyOkno)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ProduktyOkno:
			n = deduced

		case *std_widgets.QMainWindow:
			n = &ProduktyOkno{QMainWindow: *deduced}

		default:
			n = new(ProduktyOkno)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackProduktyOkno42a7e9_Constructor
func callbackProduktyOkno42a7e9_Constructor(ptr unsafe.Pointer) {
	this := NewProduktyOknoFromPointer(ptr)
	qt.Register(ptr, this)
}

func ProduktyOkno_QRegisterMetaType() int {
	return int(int32(C.ProduktyOkno42a7e9_ProduktyOkno42a7e9_QRegisterMetaType()))
}

func (ptr *ProduktyOkno) QRegisterMetaType() int {
	return int(int32(C.ProduktyOkno42a7e9_ProduktyOkno42a7e9_QRegisterMetaType()))
}

func ProduktyOkno_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ProduktyOkno42a7e9_ProduktyOkno42a7e9_QRegisterMetaType2(typeNameC)))
}

func (ptr *ProduktyOkno) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ProduktyOkno42a7e9_ProduktyOkno42a7e9_QRegisterMetaType2(typeNameC)))
}

func ProduktyOkno_QmlRegisterType() int {
	return int(int32(C.ProduktyOkno42a7e9_ProduktyOkno42a7e9_QmlRegisterType()))
}

func (ptr *ProduktyOkno) QmlRegisterType() int {
	return int(int32(C.ProduktyOkno42a7e9_ProduktyOkno42a7e9_QmlRegisterType()))
}

func ProduktyOkno_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.ProduktyOkno42a7e9_ProduktyOkno42a7e9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ProduktyOkno) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.ProduktyOkno42a7e9_ProduktyOkno42a7e9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ProduktyOkno) __resizeDocks_docks_atList(i int) *std_widgets.QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQDockWidgetFromPointer(C.ProduktyOkno42a7e9___resizeDocks_docks_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyOkno) __resizeDocks_docks_setList(i std_widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9___resizeDocks_docks_setList(ptr.Pointer(), std_widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *ProduktyOkno) __resizeDocks_docks_newList() unsafe.Pointer {
	return C.ProduktyOkno42a7e9___resizeDocks_docks_newList(ptr.Pointer())
}

func (ptr *ProduktyOkno) __resizeDocks_sizes_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyOkno42a7e9___resizeDocks_sizes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ProduktyOkno) __resizeDocks_sizes_setList(i int) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9___resizeDocks_sizes_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ProduktyOkno) __resizeDocks_sizes_newList() unsafe.Pointer {
	return C.ProduktyOkno42a7e9___resizeDocks_sizes_newList(ptr.Pointer())
}

func (ptr *ProduktyOkno) __tabifiedDockWidgets_atList(i int) *std_widgets.QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQDockWidgetFromPointer(C.ProduktyOkno42a7e9___tabifiedDockWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyOkno) __tabifiedDockWidgets_setList(i std_widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9___tabifiedDockWidgets_setList(ptr.Pointer(), std_widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *ProduktyOkno) __tabifiedDockWidgets_newList() unsafe.Pointer {
	return C.ProduktyOkno42a7e9___tabifiedDockWidgets_newList(ptr.Pointer())
}

func (ptr *ProduktyOkno) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.ProduktyOkno42a7e9___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyOkno) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *ProduktyOkno) __addActions_actions_newList() unsafe.Pointer {
	return C.ProduktyOkno42a7e9___addActions_actions_newList(ptr.Pointer())
}

func (ptr *ProduktyOkno) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.ProduktyOkno42a7e9___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyOkno) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *ProduktyOkno) __insertActions_actions_newList() unsafe.Pointer {
	return C.ProduktyOkno42a7e9___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *ProduktyOkno) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.ProduktyOkno42a7e9___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyOkno) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *ProduktyOkno) __actions_newList() unsafe.Pointer {
	return C.ProduktyOkno42a7e9___actions_newList(ptr.Pointer())
}

func (ptr *ProduktyOkno) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ProduktyOkno42a7e9___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyOkno) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ProduktyOkno) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ProduktyOkno42a7e9___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ProduktyOkno) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ProduktyOkno42a7e9___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyOkno) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ProduktyOkno) __findChildren_newList2() unsafe.Pointer {
	return C.ProduktyOkno42a7e9___findChildren_newList2(ptr.Pointer())
}

func (ptr *ProduktyOkno) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ProduktyOkno42a7e9___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyOkno) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ProduktyOkno) __findChildren_newList3() unsafe.Pointer {
	return C.ProduktyOkno42a7e9___findChildren_newList3(ptr.Pointer())
}

func (ptr *ProduktyOkno) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ProduktyOkno42a7e9___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyOkno) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ProduktyOkno) __findChildren_newList() unsafe.Pointer {
	return C.ProduktyOkno42a7e9___findChildren_newList(ptr.Pointer())
}

func (ptr *ProduktyOkno) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ProduktyOkno42a7e9___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ProduktyOkno) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ProduktyOkno) __children_newList() unsafe.Pointer {
	return C.ProduktyOkno42a7e9___children_newList(ptr.Pointer())
}

func NewProduktyOkno(parent std_widgets.QWidget_ITF, flags std_core.Qt__WindowType) *ProduktyOkno {
	tmpValue := NewProduktyOknoFromPointer(C.ProduktyOkno42a7e9_NewProduktyOkno(std_widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *ProduktyOkno) DestroyProduktyOkno() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_DestroyProduktyOkno(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackProduktyOkno42a7e9_CreatePopupMenu
func callbackProduktyOkno42a7e9_CreatePopupMenu(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createPopupMenu"); signal != nil {
		return std_widgets.PointerFromQMenu(signal.(func() *std_widgets.QMenu)())
	}

	return std_widgets.PointerFromQMenu(NewProduktyOknoFromPointer(ptr).CreatePopupMenuDefault())
}

func (ptr *ProduktyOkno) CreatePopupMenuDefault() *std_widgets.QMenu {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQMenuFromPointer(C.ProduktyOkno42a7e9_CreatePopupMenuDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackProduktyOkno42a7e9_Event
func callbackProduktyOkno42a7e9_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyOknoFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *ProduktyOkno) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyOkno42a7e9_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackProduktyOkno42a7e9_ContextMenuEvent
func callbackProduktyOkno42a7e9_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*std_gui.QContextMenuEvent))(std_gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) ContextMenuEventDefault(event std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_IconSizeChanged
func callbackProduktyOkno42a7e9_IconSizeChanged(ptr unsafe.Pointer, iconSize unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconSizeChanged"); signal != nil {
		signal.(func(*std_core.QSize))(std_core.NewQSizeFromPointer(iconSize))
	}

}

//export callbackProduktyOkno42a7e9_SetAnimated
func callbackProduktyOkno42a7e9_SetAnimated(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(ptr, "setAnimated"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewProduktyOknoFromPointer(ptr).SetAnimatedDefault(int8(enabled) != 0)
	}
}

func (ptr *ProduktyOkno) SetAnimatedDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_SetAnimatedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackProduktyOkno42a7e9_SetDockNestingEnabled
func callbackProduktyOkno42a7e9_SetDockNestingEnabled(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(ptr, "setDockNestingEnabled"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewProduktyOknoFromPointer(ptr).SetDockNestingEnabledDefault(int8(enabled) != 0)
	}
}

func (ptr *ProduktyOkno) SetDockNestingEnabledDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_SetDockNestingEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackProduktyOkno42a7e9_SetUnifiedTitleAndToolBarOnMac
func callbackProduktyOkno42a7e9_SetUnifiedTitleAndToolBarOnMac(ptr unsafe.Pointer, set C.char) {
	if signal := qt.GetSignal(ptr, "setUnifiedTitleAndToolBarOnMac"); signal != nil {
		signal.(func(bool))(int8(set) != 0)
	} else {
		NewProduktyOknoFromPointer(ptr).SetUnifiedTitleAndToolBarOnMacDefault(int8(set) != 0)
	}
}

func (ptr *ProduktyOkno) SetUnifiedTitleAndToolBarOnMacDefault(set bool) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_SetUnifiedTitleAndToolBarOnMacDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))
	}
}

//export callbackProduktyOkno42a7e9_TabifiedDockWidgetActivated
func callbackProduktyOkno42a7e9_TabifiedDockWidgetActivated(ptr unsafe.Pointer, dockWidget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabifiedDockWidgetActivated"); signal != nil {
		signal.(func(*std_widgets.QDockWidget))(std_widgets.NewQDockWidgetFromPointer(dockWidget))
	}

}

//export callbackProduktyOkno42a7e9_ToolButtonStyleChanged
func callbackProduktyOkno42a7e9_ToolButtonStyleChanged(ptr unsafe.Pointer, toolButtonStyle C.longlong) {
	if signal := qt.GetSignal(ptr, "toolButtonStyleChanged"); signal != nil {
		signal.(func(std_core.Qt__ToolButtonStyle))(std_core.Qt__ToolButtonStyle(toolButtonStyle))
	}

}

//export callbackProduktyOkno42a7e9_Close
func callbackProduktyOkno42a7e9_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyOknoFromPointer(ptr).CloseDefault())))
}

func (ptr *ProduktyOkno) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyOkno42a7e9_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackProduktyOkno42a7e9_FocusNextPrevChild
func callbackProduktyOkno42a7e9_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyOknoFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *ProduktyOkno) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyOkno42a7e9_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackProduktyOkno42a7e9_ActionEvent
func callbackProduktyOkno42a7e9_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*std_gui.QActionEvent))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_ChangeEvent
func callbackProduktyOkno42a7e9_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) ChangeEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_CloseEvent
func callbackProduktyOkno42a7e9_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*std_gui.QCloseEvent))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_CustomContextMenuRequested
func callbackProduktyOkno42a7e9_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*std_core.QPoint))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackProduktyOkno42a7e9_DragEnterEvent
func callbackProduktyOkno42a7e9_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*std_gui.QDragEnterEvent))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_DragLeaveEvent
func callbackProduktyOkno42a7e9_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*std_gui.QDragLeaveEvent))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_DragMoveEvent
func callbackProduktyOkno42a7e9_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*std_gui.QDragMoveEvent))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_DropEvent
func callbackProduktyOkno42a7e9_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*std_gui.QDropEvent))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_EnterEvent
func callbackProduktyOkno42a7e9_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_FocusInEvent
func callbackProduktyOkno42a7e9_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_FocusOutEvent
func callbackProduktyOkno42a7e9_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_Hide
func callbackProduktyOkno42a7e9_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).HideDefault()
	}
}

func (ptr *ProduktyOkno) HideDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_HideDefault(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_HideEvent
func callbackProduktyOkno42a7e9_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*std_gui.QHideEvent))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_InputMethodEvent
func callbackProduktyOkno42a7e9_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*std_gui.QInputMethodEvent))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_KeyPressEvent
func callbackProduktyOkno42a7e9_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) KeyPressEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_KeyReleaseEvent
func callbackProduktyOkno42a7e9_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_LeaveEvent
func callbackProduktyOkno42a7e9_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_Lower
func callbackProduktyOkno42a7e9_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).LowerDefault()
	}
}

func (ptr *ProduktyOkno) LowerDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_LowerDefault(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_MouseDoubleClickEvent
func callbackProduktyOkno42a7e9_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_MouseMoveEvent
func callbackProduktyOkno42a7e9_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_MousePressEvent
func callbackProduktyOkno42a7e9_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_MouseReleaseEvent
func callbackProduktyOkno42a7e9_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_MoveEvent
func callbackProduktyOkno42a7e9_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*std_gui.QMoveEvent))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_PaintEvent
func callbackProduktyOkno42a7e9_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*std_gui.QPaintEvent))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_Raise
func callbackProduktyOkno42a7e9_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *ProduktyOkno) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_RaiseDefault(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_Repaint
func callbackProduktyOkno42a7e9_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *ProduktyOkno) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_RepaintDefault(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_ResizeEvent
func callbackProduktyOkno42a7e9_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*std_gui.QResizeEvent))(std_gui.NewQResizeEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) ResizeEventDefault(event std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_SetDisabled
func callbackProduktyOkno42a7e9_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewProduktyOknoFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *ProduktyOkno) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackProduktyOkno42a7e9_SetEnabled
func callbackProduktyOkno42a7e9_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewProduktyOknoFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *ProduktyOkno) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackProduktyOkno42a7e9_SetFocus2
func callbackProduktyOkno42a7e9_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *ProduktyOkno) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_SetHidden
func callbackProduktyOkno42a7e9_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewProduktyOknoFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *ProduktyOkno) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackProduktyOkno42a7e9_SetStyleSheet
func callbackProduktyOkno42a7e9_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewProduktyOknoFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *ProduktyOkno) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.ProduktyOkno42a7e9_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackProduktyOkno42a7e9_SetVisible
func callbackProduktyOkno42a7e9_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewProduktyOknoFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *ProduktyOkno) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackProduktyOkno42a7e9_SetWindowModified
func callbackProduktyOkno42a7e9_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewProduktyOknoFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *ProduktyOkno) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackProduktyOkno42a7e9_SetWindowTitle
func callbackProduktyOkno42a7e9_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewProduktyOknoFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *ProduktyOkno) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.ProduktyOkno42a7e9_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackProduktyOkno42a7e9_Show
func callbackProduktyOkno42a7e9_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).ShowDefault()
	}
}

func (ptr *ProduktyOkno) ShowDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ShowDefault(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_ShowEvent
func callbackProduktyOkno42a7e9_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*std_gui.QShowEvent))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_ShowFullScreen
func callbackProduktyOkno42a7e9_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *ProduktyOkno) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_ShowMaximized
func callbackProduktyOkno42a7e9_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *ProduktyOkno) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_ShowMinimized
func callbackProduktyOkno42a7e9_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *ProduktyOkno) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_ShowNormal
func callbackProduktyOkno42a7e9_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *ProduktyOkno) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_TabletEvent
func callbackProduktyOkno42a7e9_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*std_gui.QTabletEvent))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_Update
func callbackProduktyOkno42a7e9_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *ProduktyOkno) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_UpdateDefault(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_UpdateMicroFocus
func callbackProduktyOkno42a7e9_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *ProduktyOkno) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackProduktyOkno42a7e9_WheelEvent
func callbackProduktyOkno42a7e9_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*std_gui.QWheelEvent))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_WindowIconChanged
func callbackProduktyOkno42a7e9_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*std_gui.QIcon))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackProduktyOkno42a7e9_WindowTitleChanged
func callbackProduktyOkno42a7e9_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackProduktyOkno42a7e9_PaintEngine
func callbackProduktyOkno42a7e9_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine(signal.(func() *std_gui.QPaintEngine)())
	}

	return std_gui.PointerFromQPaintEngine(NewProduktyOknoFromPointer(ptr).PaintEngineDefault())
}

func (ptr *ProduktyOkno) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.ProduktyOkno42a7e9_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackProduktyOkno42a7e9_MinimumSizeHint
func callbackProduktyOkno42a7e9_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewProduktyOknoFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *ProduktyOkno) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.ProduktyOkno42a7e9_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackProduktyOkno42a7e9_SizeHint
func callbackProduktyOkno42a7e9_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewProduktyOknoFromPointer(ptr).SizeHintDefault())
}

func (ptr *ProduktyOkno) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.ProduktyOkno42a7e9_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackProduktyOkno42a7e9_InputMethodQuery
func callbackProduktyOkno42a7e9_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewProduktyOknoFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *ProduktyOkno) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.ProduktyOkno42a7e9_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackProduktyOkno42a7e9_HasHeightForWidth
func callbackProduktyOkno42a7e9_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyOknoFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *ProduktyOkno) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyOkno42a7e9_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackProduktyOkno42a7e9_HeightForWidth
func callbackProduktyOkno42a7e9_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewProduktyOknoFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *ProduktyOkno) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyOkno42a7e9_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackProduktyOkno42a7e9_Metric
func callbackProduktyOkno42a7e9_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(std_gui.QPaintDevice__PaintDeviceMetric) int)(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewProduktyOknoFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *ProduktyOkno) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ProduktyOkno42a7e9_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackProduktyOkno42a7e9_EventFilter
func callbackProduktyOkno42a7e9_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewProduktyOknoFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ProduktyOkno) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ProduktyOkno42a7e9_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackProduktyOkno42a7e9_ChildEvent
func callbackProduktyOkno42a7e9_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_ConnectNotify
func callbackProduktyOkno42a7e9_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewProduktyOknoFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ProduktyOkno) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackProduktyOkno42a7e9_CustomEvent
func callbackProduktyOkno42a7e9_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackProduktyOkno42a7e9_DeleteLater
func callbackProduktyOkno42a7e9_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewProduktyOknoFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ProduktyOkno) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackProduktyOkno42a7e9_Destroyed
func callbackProduktyOkno42a7e9_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackProduktyOkno42a7e9_DisconnectNotify
func callbackProduktyOkno42a7e9_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewProduktyOknoFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ProduktyOkno) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackProduktyOkno42a7e9_ObjectNameChanged
func callbackProduktyOkno42a7e9_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackProduktyOkno42a7e9_TimerEvent
func callbackProduktyOkno42a7e9_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewProduktyOknoFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ProduktyOkno) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ProduktyOkno42a7e9_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type WydzialyOkno_ITF interface {
	std_widgets.QMainWindow_ITF
	WydzialyOkno_PTR() *WydzialyOkno
}

func (ptr *WydzialyOkno) WydzialyOkno_PTR() *WydzialyOkno {
	return ptr
}

func (ptr *WydzialyOkno) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QMainWindow_PTR().Pointer()
	}
	return nil
}

func (ptr *WydzialyOkno) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QMainWindow_PTR().SetPointer(p)
	}
}

func PointerFromWydzialyOkno(ptr WydzialyOkno_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WydzialyOkno_PTR().Pointer()
	}
	return nil
}

func NewWydzialyOknoFromPointer(ptr unsafe.Pointer) (n *WydzialyOkno) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(WydzialyOkno)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *WydzialyOkno:
			n = deduced

		case *std_widgets.QMainWindow:
			n = &WydzialyOkno{QMainWindow: *deduced}

		default:
			n = new(WydzialyOkno)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackWydzialyOkno42a7e9_Constructor
func callbackWydzialyOkno42a7e9_Constructor(ptr unsafe.Pointer) {
	this := NewWydzialyOknoFromPointer(ptr)
	qt.Register(ptr, this)
}

func WydzialyOkno_QRegisterMetaType() int {
	return int(int32(C.WydzialyOkno42a7e9_WydzialyOkno42a7e9_QRegisterMetaType()))
}

func (ptr *WydzialyOkno) QRegisterMetaType() int {
	return int(int32(C.WydzialyOkno42a7e9_WydzialyOkno42a7e9_QRegisterMetaType()))
}

func WydzialyOkno_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.WydzialyOkno42a7e9_WydzialyOkno42a7e9_QRegisterMetaType2(typeNameC)))
}

func (ptr *WydzialyOkno) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.WydzialyOkno42a7e9_WydzialyOkno42a7e9_QRegisterMetaType2(typeNameC)))
}

func WydzialyOkno_QmlRegisterType() int {
	return int(int32(C.WydzialyOkno42a7e9_WydzialyOkno42a7e9_QmlRegisterType()))
}

func (ptr *WydzialyOkno) QmlRegisterType() int {
	return int(int32(C.WydzialyOkno42a7e9_WydzialyOkno42a7e9_QmlRegisterType()))
}

func WydzialyOkno_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.WydzialyOkno42a7e9_WydzialyOkno42a7e9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *WydzialyOkno) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.WydzialyOkno42a7e9_WydzialyOkno42a7e9_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *WydzialyOkno) __resizeDocks_docks_atList(i int) *std_widgets.QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQDockWidgetFromPointer(C.WydzialyOkno42a7e9___resizeDocks_docks_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WydzialyOkno) __resizeDocks_docks_setList(i std_widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9___resizeDocks_docks_setList(ptr.Pointer(), std_widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *WydzialyOkno) __resizeDocks_docks_newList() unsafe.Pointer {
	return C.WydzialyOkno42a7e9___resizeDocks_docks_newList(ptr.Pointer())
}

func (ptr *WydzialyOkno) __resizeDocks_sizes_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WydzialyOkno42a7e9___resizeDocks_sizes_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *WydzialyOkno) __resizeDocks_sizes_setList(i int) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9___resizeDocks_sizes_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *WydzialyOkno) __resizeDocks_sizes_newList() unsafe.Pointer {
	return C.WydzialyOkno42a7e9___resizeDocks_sizes_newList(ptr.Pointer())
}

func (ptr *WydzialyOkno) __tabifiedDockWidgets_atList(i int) *std_widgets.QDockWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQDockWidgetFromPointer(C.WydzialyOkno42a7e9___tabifiedDockWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WydzialyOkno) __tabifiedDockWidgets_setList(i std_widgets.QDockWidget_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9___tabifiedDockWidgets_setList(ptr.Pointer(), std_widgets.PointerFromQDockWidget(i))
	}
}

func (ptr *WydzialyOkno) __tabifiedDockWidgets_newList() unsafe.Pointer {
	return C.WydzialyOkno42a7e9___tabifiedDockWidgets_newList(ptr.Pointer())
}

func (ptr *WydzialyOkno) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.WydzialyOkno42a7e9___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WydzialyOkno) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *WydzialyOkno) __addActions_actions_newList() unsafe.Pointer {
	return C.WydzialyOkno42a7e9___addActions_actions_newList(ptr.Pointer())
}

func (ptr *WydzialyOkno) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.WydzialyOkno42a7e9___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WydzialyOkno) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *WydzialyOkno) __insertActions_actions_newList() unsafe.Pointer {
	return C.WydzialyOkno42a7e9___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *WydzialyOkno) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.WydzialyOkno42a7e9___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WydzialyOkno) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *WydzialyOkno) __actions_newList() unsafe.Pointer {
	return C.WydzialyOkno42a7e9___actions_newList(ptr.Pointer())
}

func (ptr *WydzialyOkno) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.WydzialyOkno42a7e9___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *WydzialyOkno) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *WydzialyOkno) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.WydzialyOkno42a7e9___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *WydzialyOkno) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WydzialyOkno42a7e9___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WydzialyOkno) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WydzialyOkno) __findChildren_newList2() unsafe.Pointer {
	return C.WydzialyOkno42a7e9___findChildren_newList2(ptr.Pointer())
}

func (ptr *WydzialyOkno) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WydzialyOkno42a7e9___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WydzialyOkno) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WydzialyOkno) __findChildren_newList3() unsafe.Pointer {
	return C.WydzialyOkno42a7e9___findChildren_newList3(ptr.Pointer())
}

func (ptr *WydzialyOkno) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WydzialyOkno42a7e9___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WydzialyOkno) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WydzialyOkno) __findChildren_newList() unsafe.Pointer {
	return C.WydzialyOkno42a7e9___findChildren_newList(ptr.Pointer())
}

func (ptr *WydzialyOkno) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WydzialyOkno42a7e9___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WydzialyOkno) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WydzialyOkno) __children_newList() unsafe.Pointer {
	return C.WydzialyOkno42a7e9___children_newList(ptr.Pointer())
}

func NewWydzialyOkno(parent std_widgets.QWidget_ITF, flags std_core.Qt__WindowType) *WydzialyOkno {
	tmpValue := NewWydzialyOknoFromPointer(C.WydzialyOkno42a7e9_NewWydzialyOkno(std_widgets.PointerFromQWidget(parent), C.longlong(flags)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *WydzialyOkno) DestroyWydzialyOkno() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_DestroyWydzialyOkno(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackWydzialyOkno42a7e9_CreatePopupMenu
func callbackWydzialyOkno42a7e9_CreatePopupMenu(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createPopupMenu"); signal != nil {
		return std_widgets.PointerFromQMenu(signal.(func() *std_widgets.QMenu)())
	}

	return std_widgets.PointerFromQMenu(NewWydzialyOknoFromPointer(ptr).CreatePopupMenuDefault())
}

func (ptr *WydzialyOkno) CreatePopupMenuDefault() *std_widgets.QMenu {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQMenuFromPointer(C.WydzialyOkno42a7e9_CreatePopupMenuDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackWydzialyOkno42a7e9_Event
func callbackWydzialyOkno42a7e9_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWydzialyOknoFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *WydzialyOkno) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WydzialyOkno42a7e9_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackWydzialyOkno42a7e9_ContextMenuEvent
func callbackWydzialyOkno42a7e9_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		signal.(func(*std_gui.QContextMenuEvent))(std_gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) ContextMenuEventDefault(event std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_IconSizeChanged
func callbackWydzialyOkno42a7e9_IconSizeChanged(ptr unsafe.Pointer, iconSize unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconSizeChanged"); signal != nil {
		signal.(func(*std_core.QSize))(std_core.NewQSizeFromPointer(iconSize))
	}

}

//export callbackWydzialyOkno42a7e9_SetAnimated
func callbackWydzialyOkno42a7e9_SetAnimated(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(ptr, "setAnimated"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewWydzialyOknoFromPointer(ptr).SetAnimatedDefault(int8(enabled) != 0)
	}
}

func (ptr *WydzialyOkno) SetAnimatedDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_SetAnimatedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackWydzialyOkno42a7e9_SetDockNestingEnabled
func callbackWydzialyOkno42a7e9_SetDockNestingEnabled(ptr unsafe.Pointer, enabled C.char) {
	if signal := qt.GetSignal(ptr, "setDockNestingEnabled"); signal != nil {
		signal.(func(bool))(int8(enabled) != 0)
	} else {
		NewWydzialyOknoFromPointer(ptr).SetDockNestingEnabledDefault(int8(enabled) != 0)
	}
}

func (ptr *WydzialyOkno) SetDockNestingEnabledDefault(enabled bool) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_SetDockNestingEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackWydzialyOkno42a7e9_SetUnifiedTitleAndToolBarOnMac
func callbackWydzialyOkno42a7e9_SetUnifiedTitleAndToolBarOnMac(ptr unsafe.Pointer, set C.char) {
	if signal := qt.GetSignal(ptr, "setUnifiedTitleAndToolBarOnMac"); signal != nil {
		signal.(func(bool))(int8(set) != 0)
	} else {
		NewWydzialyOknoFromPointer(ptr).SetUnifiedTitleAndToolBarOnMacDefault(int8(set) != 0)
	}
}

func (ptr *WydzialyOkno) SetUnifiedTitleAndToolBarOnMacDefault(set bool) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_SetUnifiedTitleAndToolBarOnMacDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(set))))
	}
}

//export callbackWydzialyOkno42a7e9_TabifiedDockWidgetActivated
func callbackWydzialyOkno42a7e9_TabifiedDockWidgetActivated(ptr unsafe.Pointer, dockWidget unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabifiedDockWidgetActivated"); signal != nil {
		signal.(func(*std_widgets.QDockWidget))(std_widgets.NewQDockWidgetFromPointer(dockWidget))
	}

}

//export callbackWydzialyOkno42a7e9_ToolButtonStyleChanged
func callbackWydzialyOkno42a7e9_ToolButtonStyleChanged(ptr unsafe.Pointer, toolButtonStyle C.longlong) {
	if signal := qt.GetSignal(ptr, "toolButtonStyleChanged"); signal != nil {
		signal.(func(std_core.Qt__ToolButtonStyle))(std_core.Qt__ToolButtonStyle(toolButtonStyle))
	}

}

//export callbackWydzialyOkno42a7e9_Close
func callbackWydzialyOkno42a7e9_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewWydzialyOknoFromPointer(ptr).CloseDefault())))
}

func (ptr *WydzialyOkno) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.WydzialyOkno42a7e9_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackWydzialyOkno42a7e9_FocusNextPrevChild
func callbackWydzialyOkno42a7e9_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWydzialyOknoFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *WydzialyOkno) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.WydzialyOkno42a7e9_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackWydzialyOkno42a7e9_ActionEvent
func callbackWydzialyOkno42a7e9_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		signal.(func(*std_gui.QActionEvent))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_ChangeEvent
func callbackWydzialyOkno42a7e9_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) ChangeEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_CloseEvent
func callbackWydzialyOkno42a7e9_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		signal.(func(*std_gui.QCloseEvent))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_CustomContextMenuRequested
func callbackWydzialyOkno42a7e9_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		signal.(func(*std_core.QPoint))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackWydzialyOkno42a7e9_DragEnterEvent
func callbackWydzialyOkno42a7e9_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		signal.(func(*std_gui.QDragEnterEvent))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_DragLeaveEvent
func callbackWydzialyOkno42a7e9_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		signal.(func(*std_gui.QDragLeaveEvent))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_DragMoveEvent
func callbackWydzialyOkno42a7e9_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		signal.(func(*std_gui.QDragMoveEvent))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_DropEvent
func callbackWydzialyOkno42a7e9_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		signal.(func(*std_gui.QDropEvent))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_EnterEvent
func callbackWydzialyOkno42a7e9_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_FocusInEvent
func callbackWydzialyOkno42a7e9_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) FocusInEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_FocusOutEvent
func callbackWydzialyOkno42a7e9_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		signal.(func(*std_gui.QFocusEvent))(std_gui.NewQFocusEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) FocusOutEventDefault(event std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_Hide
func callbackWydzialyOkno42a7e9_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).HideDefault()
	}
}

func (ptr *WydzialyOkno) HideDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_HideDefault(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_HideEvent
func callbackWydzialyOkno42a7e9_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		signal.(func(*std_gui.QHideEvent))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_InputMethodEvent
func callbackWydzialyOkno42a7e9_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		signal.(func(*std_gui.QInputMethodEvent))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_KeyPressEvent
func callbackWydzialyOkno42a7e9_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) KeyPressEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_KeyReleaseEvent
func callbackWydzialyOkno42a7e9_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QKeyEvent))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_LeaveEvent
func callbackWydzialyOkno42a7e9_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_Lower
func callbackWydzialyOkno42a7e9_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).LowerDefault()
	}
}

func (ptr *WydzialyOkno) LowerDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_LowerDefault(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_MouseDoubleClickEvent
func callbackWydzialyOkno42a7e9_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_MouseMoveEvent
func callbackWydzialyOkno42a7e9_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_MousePressEvent
func callbackWydzialyOkno42a7e9_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_MouseReleaseEvent
func callbackWydzialyOkno42a7e9_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		signal.(func(*std_gui.QMouseEvent))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_MoveEvent
func callbackWydzialyOkno42a7e9_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		signal.(func(*std_gui.QMoveEvent))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_PaintEvent
func callbackWydzialyOkno42a7e9_PaintEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		signal.(func(*std_gui.QPaintEvent))(std_gui.NewQPaintEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) PaintEventDefault(event std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_Raise
func callbackWydzialyOkno42a7e9_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *WydzialyOkno) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_RaiseDefault(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_Repaint
func callbackWydzialyOkno42a7e9_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *WydzialyOkno) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_RepaintDefault(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_ResizeEvent
func callbackWydzialyOkno42a7e9_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		signal.(func(*std_gui.QResizeEvent))(std_gui.NewQResizeEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) ResizeEventDefault(event std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_SetDisabled
func callbackWydzialyOkno42a7e9_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewWydzialyOknoFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *WydzialyOkno) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackWydzialyOkno42a7e9_SetEnabled
func callbackWydzialyOkno42a7e9_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewWydzialyOknoFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *WydzialyOkno) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackWydzialyOkno42a7e9_SetFocus2
func callbackWydzialyOkno42a7e9_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *WydzialyOkno) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_SetHidden
func callbackWydzialyOkno42a7e9_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewWydzialyOknoFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *WydzialyOkno) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackWydzialyOkno42a7e9_SetStyleSheet
func callbackWydzialyOkno42a7e9_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		signal.(func(string))(cGoUnpackString(styleSheet))
	} else {
		NewWydzialyOknoFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *WydzialyOkno) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.WydzialyOkno42a7e9_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackWydzialyOkno42a7e9_SetVisible
func callbackWydzialyOkno42a7e9_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewWydzialyOknoFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *WydzialyOkno) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackWydzialyOkno42a7e9_SetWindowModified
func callbackWydzialyOkno42a7e9_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewWydzialyOknoFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *WydzialyOkno) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackWydzialyOkno42a7e9_SetWindowTitle
func callbackWydzialyOkno42a7e9_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		signal.(func(string))(cGoUnpackString(vqs))
	} else {
		NewWydzialyOknoFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *WydzialyOkno) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.WydzialyOkno42a7e9_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackWydzialyOkno42a7e9_Show
func callbackWydzialyOkno42a7e9_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).ShowDefault()
	}
}

func (ptr *WydzialyOkno) ShowDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ShowDefault(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_ShowEvent
func callbackWydzialyOkno42a7e9_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		signal.(func(*std_gui.QShowEvent))(std_gui.NewQShowEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) ShowEventDefault(event std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_ShowFullScreen
func callbackWydzialyOkno42a7e9_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *WydzialyOkno) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_ShowMaximized
func callbackWydzialyOkno42a7e9_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *WydzialyOkno) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_ShowMinimized
func callbackWydzialyOkno42a7e9_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *WydzialyOkno) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_ShowNormal
func callbackWydzialyOkno42a7e9_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *WydzialyOkno) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_TabletEvent
func callbackWydzialyOkno42a7e9_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		signal.(func(*std_gui.QTabletEvent))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_Update
func callbackWydzialyOkno42a7e9_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *WydzialyOkno) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_UpdateDefault(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_UpdateMicroFocus
func callbackWydzialyOkno42a7e9_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *WydzialyOkno) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackWydzialyOkno42a7e9_WheelEvent
func callbackWydzialyOkno42a7e9_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		signal.(func(*std_gui.QWheelEvent))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_WindowIconChanged
func callbackWydzialyOkno42a7e9_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		signal.(func(*std_gui.QIcon))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackWydzialyOkno42a7e9_WindowTitleChanged
func callbackWydzialyOkno42a7e9_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(title))
	}

}

//export callbackWydzialyOkno42a7e9_PaintEngine
func callbackWydzialyOkno42a7e9_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine(signal.(func() *std_gui.QPaintEngine)())
	}

	return std_gui.PointerFromQPaintEngine(NewWydzialyOknoFromPointer(ptr).PaintEngineDefault())
}

func (ptr *WydzialyOkno) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.WydzialyOkno42a7e9_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackWydzialyOkno42a7e9_MinimumSizeHint
func callbackWydzialyOkno42a7e9_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewWydzialyOknoFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *WydzialyOkno) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.WydzialyOkno42a7e9_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackWydzialyOkno42a7e9_SizeHint
func callbackWydzialyOkno42a7e9_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewWydzialyOknoFromPointer(ptr).SizeHintDefault())
}

func (ptr *WydzialyOkno) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.WydzialyOkno42a7e9_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackWydzialyOkno42a7e9_InputMethodQuery
func callbackWydzialyOkno42a7e9_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant(signal.(func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewWydzialyOknoFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *WydzialyOkno) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.WydzialyOkno42a7e9_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackWydzialyOkno42a7e9_HasHeightForWidth
func callbackWydzialyOkno42a7e9_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewWydzialyOknoFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *WydzialyOkno) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.WydzialyOkno42a7e9_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackWydzialyOkno42a7e9_HeightForWidth
func callbackWydzialyOkno42a7e9_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewWydzialyOknoFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *WydzialyOkno) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WydzialyOkno42a7e9_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackWydzialyOkno42a7e9_Metric
func callbackWydzialyOkno42a7e9_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32(signal.(func(std_gui.QPaintDevice__PaintDeviceMetric) int)(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewWydzialyOknoFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *WydzialyOkno) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WydzialyOkno42a7e9_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackWydzialyOkno42a7e9_EventFilter
func callbackWydzialyOkno42a7e9_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWydzialyOknoFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *WydzialyOkno) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WydzialyOkno42a7e9_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackWydzialyOkno42a7e9_ChildEvent
func callbackWydzialyOkno42a7e9_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_ConnectNotify
func callbackWydzialyOkno42a7e9_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewWydzialyOknoFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *WydzialyOkno) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackWydzialyOkno42a7e9_CustomEvent
func callbackWydzialyOkno42a7e9_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackWydzialyOkno42a7e9_DeleteLater
func callbackWydzialyOkno42a7e9_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewWydzialyOknoFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *WydzialyOkno) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackWydzialyOkno42a7e9_Destroyed
func callbackWydzialyOkno42a7e9_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackWydzialyOkno42a7e9_DisconnectNotify
func callbackWydzialyOkno42a7e9_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewWydzialyOknoFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *WydzialyOkno) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackWydzialyOkno42a7e9_ObjectNameChanged
func callbackWydzialyOkno42a7e9_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackWydzialyOkno42a7e9_TimerEvent
func callbackWydzialyOkno42a7e9_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewWydzialyOknoFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *WydzialyOkno) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WydzialyOkno42a7e9_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

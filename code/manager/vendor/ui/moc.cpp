

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDockWidget>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QList>
#include <QMainWindow>
#include <QMenu>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPoint>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QTabletEvent>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>


class ProduktyOkno42a7e9: public QMainWindow
{
Q_OBJECT
public:
	ProduktyOkno42a7e9(QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QMainWindow(parent, flags) {qRegisterMetaType<quintptr>("quintptr");ProduktyOkno42a7e9_ProduktyOkno42a7e9_QRegisterMetaType();ProduktyOkno42a7e9_ProduktyOkno42a7e9_QRegisterMetaTypes();callbackProduktyOkno42a7e9_Constructor(this);};
	QMenu * createPopupMenu() { return static_cast<QMenu*>(callbackProduktyOkno42a7e9_CreatePopupMenu(this)); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackProduktyOkno42a7e9_ContextMenuEvent(this, event); };
	void Signal_IconSizeChanged(const QSize & iconSize) { callbackProduktyOkno42a7e9_IconSizeChanged(this, const_cast<QSize*>(&iconSize)); };
	void setAnimated(bool enabled) { callbackProduktyOkno42a7e9_SetAnimated(this, enabled); };
	void setDockNestingEnabled(bool enabled) { callbackProduktyOkno42a7e9_SetDockNestingEnabled(this, enabled); };
	void setUnifiedTitleAndToolBarOnMac(bool set) { callbackProduktyOkno42a7e9_SetUnifiedTitleAndToolBarOnMac(this, set); };
	void Signal_TabifiedDockWidgetActivated(QDockWidget * dockWidget) { callbackProduktyOkno42a7e9_TabifiedDockWidgetActivated(this, dockWidget); };
	void Signal_ToolButtonStyleChanged(Qt::ToolButtonStyle toolButtonStyle) { callbackProduktyOkno42a7e9_ToolButtonStyleChanged(this, toolButtonStyle); };
	bool close() { return callbackProduktyOkno42a7e9_Close(this) != 0; };
	bool focusNextPrevChild(bool next) { return callbackProduktyOkno42a7e9_FocusNextPrevChild(this, next) != 0; };
	void actionEvent(QActionEvent * event) { callbackProduktyOkno42a7e9_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackProduktyOkno42a7e9_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackProduktyOkno42a7e9_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackProduktyOkno42a7e9_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackProduktyOkno42a7e9_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackProduktyOkno42a7e9_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackProduktyOkno42a7e9_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackProduktyOkno42a7e9_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackProduktyOkno42a7e9_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackProduktyOkno42a7e9_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackProduktyOkno42a7e9_FocusOutEvent(this, event); };
	void hide() { callbackProduktyOkno42a7e9_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackProduktyOkno42a7e9_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackProduktyOkno42a7e9_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackProduktyOkno42a7e9_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackProduktyOkno42a7e9_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackProduktyOkno42a7e9_LeaveEvent(this, event); };
	void lower() { callbackProduktyOkno42a7e9_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackProduktyOkno42a7e9_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackProduktyOkno42a7e9_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackProduktyOkno42a7e9_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackProduktyOkno42a7e9_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackProduktyOkno42a7e9_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackProduktyOkno42a7e9_PaintEvent(this, event); };
	void raise() { callbackProduktyOkno42a7e9_Raise(this); };
	void repaint() { callbackProduktyOkno42a7e9_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackProduktyOkno42a7e9_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackProduktyOkno42a7e9_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackProduktyOkno42a7e9_SetEnabled(this, vbo); };
	void setFocus() { callbackProduktyOkno42a7e9_SetFocus2(this); };
	void setHidden(bool hidden) { callbackProduktyOkno42a7e9_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackProduktyOkno42a7e9_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackProduktyOkno42a7e9_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackProduktyOkno42a7e9_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackProduktyOkno42a7e9_SetWindowTitle(this, vqsPacked); };
	void show() { callbackProduktyOkno42a7e9_Show(this); };
	void showEvent(QShowEvent * event) { callbackProduktyOkno42a7e9_ShowEvent(this, event); };
	void showFullScreen() { callbackProduktyOkno42a7e9_ShowFullScreen(this); };
	void showMaximized() { callbackProduktyOkno42a7e9_ShowMaximized(this); };
	void showMinimized() { callbackProduktyOkno42a7e9_ShowMinimized(this); };
	void showNormal() { callbackProduktyOkno42a7e9_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackProduktyOkno42a7e9_TabletEvent(this, event); };
	void update() { callbackProduktyOkno42a7e9_Update(this); };
	void updateMicroFocus() { callbackProduktyOkno42a7e9_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackProduktyOkno42a7e9_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackProduktyOkno42a7e9_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackProduktyOkno42a7e9_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackProduktyOkno42a7e9_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackProduktyOkno42a7e9_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackProduktyOkno42a7e9_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackProduktyOkno42a7e9_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackProduktyOkno42a7e9_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackProduktyOkno42a7e9_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackProduktyOkno42a7e9_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackProduktyOkno42a7e9_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackProduktyOkno42a7e9_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackProduktyOkno42a7e9_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackProduktyOkno42a7e9_CustomEvent(this, event); };
	void deleteLater() { callbackProduktyOkno42a7e9_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackProduktyOkno42a7e9_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackProduktyOkno42a7e9_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackProduktyOkno42a7e9_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackProduktyOkno42a7e9_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(ProduktyOkno42a7e9*)


void ProduktyOkno42a7e9_ProduktyOkno42a7e9_QRegisterMetaTypes() {
}

class WydzialyOkno42a7e9: public QMainWindow
{
Q_OBJECT
public:
	WydzialyOkno42a7e9(QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QMainWindow(parent, flags) {qRegisterMetaType<quintptr>("quintptr");WydzialyOkno42a7e9_WydzialyOkno42a7e9_QRegisterMetaType();WydzialyOkno42a7e9_WydzialyOkno42a7e9_QRegisterMetaTypes();callbackWydzialyOkno42a7e9_Constructor(this);};
	QMenu * createPopupMenu() { return static_cast<QMenu*>(callbackWydzialyOkno42a7e9_CreatePopupMenu(this)); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackWydzialyOkno42a7e9_ContextMenuEvent(this, event); };
	void Signal_IconSizeChanged(const QSize & iconSize) { callbackWydzialyOkno42a7e9_IconSizeChanged(this, const_cast<QSize*>(&iconSize)); };
	void setAnimated(bool enabled) { callbackWydzialyOkno42a7e9_SetAnimated(this, enabled); };
	void setDockNestingEnabled(bool enabled) { callbackWydzialyOkno42a7e9_SetDockNestingEnabled(this, enabled); };
	void setUnifiedTitleAndToolBarOnMac(bool set) { callbackWydzialyOkno42a7e9_SetUnifiedTitleAndToolBarOnMac(this, set); };
	void Signal_TabifiedDockWidgetActivated(QDockWidget * dockWidget) { callbackWydzialyOkno42a7e9_TabifiedDockWidgetActivated(this, dockWidget); };
	void Signal_ToolButtonStyleChanged(Qt::ToolButtonStyle toolButtonStyle) { callbackWydzialyOkno42a7e9_ToolButtonStyleChanged(this, toolButtonStyle); };
	bool close() { return callbackWydzialyOkno42a7e9_Close(this) != 0; };
	bool focusNextPrevChild(bool next) { return callbackWydzialyOkno42a7e9_FocusNextPrevChild(this, next) != 0; };
	void actionEvent(QActionEvent * event) { callbackWydzialyOkno42a7e9_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackWydzialyOkno42a7e9_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackWydzialyOkno42a7e9_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackWydzialyOkno42a7e9_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackWydzialyOkno42a7e9_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackWydzialyOkno42a7e9_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackWydzialyOkno42a7e9_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackWydzialyOkno42a7e9_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackWydzialyOkno42a7e9_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackWydzialyOkno42a7e9_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackWydzialyOkno42a7e9_FocusOutEvent(this, event); };
	void hide() { callbackWydzialyOkno42a7e9_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackWydzialyOkno42a7e9_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackWydzialyOkno42a7e9_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackWydzialyOkno42a7e9_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackWydzialyOkno42a7e9_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackWydzialyOkno42a7e9_LeaveEvent(this, event); };
	void lower() { callbackWydzialyOkno42a7e9_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackWydzialyOkno42a7e9_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackWydzialyOkno42a7e9_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackWydzialyOkno42a7e9_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackWydzialyOkno42a7e9_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackWydzialyOkno42a7e9_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackWydzialyOkno42a7e9_PaintEvent(this, event); };
	void raise() { callbackWydzialyOkno42a7e9_Raise(this); };
	void repaint() { callbackWydzialyOkno42a7e9_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackWydzialyOkno42a7e9_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackWydzialyOkno42a7e9_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackWydzialyOkno42a7e9_SetEnabled(this, vbo); };
	void setFocus() { callbackWydzialyOkno42a7e9_SetFocus2(this); };
	void setHidden(bool hidden) { callbackWydzialyOkno42a7e9_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackWydzialyOkno42a7e9_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackWydzialyOkno42a7e9_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackWydzialyOkno42a7e9_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackWydzialyOkno42a7e9_SetWindowTitle(this, vqsPacked); };
	void show() { callbackWydzialyOkno42a7e9_Show(this); };
	void showEvent(QShowEvent * event) { callbackWydzialyOkno42a7e9_ShowEvent(this, event); };
	void showFullScreen() { callbackWydzialyOkno42a7e9_ShowFullScreen(this); };
	void showMaximized() { callbackWydzialyOkno42a7e9_ShowMaximized(this); };
	void showMinimized() { callbackWydzialyOkno42a7e9_ShowMinimized(this); };
	void showNormal() { callbackWydzialyOkno42a7e9_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackWydzialyOkno42a7e9_TabletEvent(this, event); };
	void update() { callbackWydzialyOkno42a7e9_Update(this); };
	void updateMicroFocus() { callbackWydzialyOkno42a7e9_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackWydzialyOkno42a7e9_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackWydzialyOkno42a7e9_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackWydzialyOkno42a7e9_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackWydzialyOkno42a7e9_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackWydzialyOkno42a7e9_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackWydzialyOkno42a7e9_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackWydzialyOkno42a7e9_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackWydzialyOkno42a7e9_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackWydzialyOkno42a7e9_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackWydzialyOkno42a7e9_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWydzialyOkno42a7e9_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackWydzialyOkno42a7e9_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWydzialyOkno42a7e9_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWydzialyOkno42a7e9_CustomEvent(this, event); };
	void deleteLater() { callbackWydzialyOkno42a7e9_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWydzialyOkno42a7e9_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWydzialyOkno42a7e9_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWydzialyOkno42a7e9_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWydzialyOkno42a7e9_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(WydzialyOkno42a7e9*)


void WydzialyOkno42a7e9_WydzialyOkno42a7e9_QRegisterMetaTypes() {
}

class NormyOkno42a7e9: public QMainWindow
{
Q_OBJECT
public:
	NormyOkno42a7e9(QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QMainWindow(parent, flags) {qRegisterMetaType<quintptr>("quintptr");NormyOkno42a7e9_NormyOkno42a7e9_QRegisterMetaType();NormyOkno42a7e9_NormyOkno42a7e9_QRegisterMetaTypes();callbackNormyOkno42a7e9_Constructor(this);};
	QMenu * createPopupMenu() { return static_cast<QMenu*>(callbackNormyOkno42a7e9_CreatePopupMenu(this)); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackNormyOkno42a7e9_ContextMenuEvent(this, event); };
	void Signal_IconSizeChanged(const QSize & iconSize) { callbackNormyOkno42a7e9_IconSizeChanged(this, const_cast<QSize*>(&iconSize)); };
	void setAnimated(bool enabled) { callbackNormyOkno42a7e9_SetAnimated(this, enabled); };
	void setDockNestingEnabled(bool enabled) { callbackNormyOkno42a7e9_SetDockNestingEnabled(this, enabled); };
	void setUnifiedTitleAndToolBarOnMac(bool set) { callbackNormyOkno42a7e9_SetUnifiedTitleAndToolBarOnMac(this, set); };
	void Signal_TabifiedDockWidgetActivated(QDockWidget * dockWidget) { callbackNormyOkno42a7e9_TabifiedDockWidgetActivated(this, dockWidget); };
	void Signal_ToolButtonStyleChanged(Qt::ToolButtonStyle toolButtonStyle) { callbackNormyOkno42a7e9_ToolButtonStyleChanged(this, toolButtonStyle); };
	bool close() { return callbackNormyOkno42a7e9_Close(this) != 0; };
	bool focusNextPrevChild(bool next) { return callbackNormyOkno42a7e9_FocusNextPrevChild(this, next) != 0; };
	void actionEvent(QActionEvent * event) { callbackNormyOkno42a7e9_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackNormyOkno42a7e9_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackNormyOkno42a7e9_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackNormyOkno42a7e9_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackNormyOkno42a7e9_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackNormyOkno42a7e9_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackNormyOkno42a7e9_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackNormyOkno42a7e9_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackNormyOkno42a7e9_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackNormyOkno42a7e9_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackNormyOkno42a7e9_FocusOutEvent(this, event); };
	void hide() { callbackNormyOkno42a7e9_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackNormyOkno42a7e9_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackNormyOkno42a7e9_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackNormyOkno42a7e9_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackNormyOkno42a7e9_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackNormyOkno42a7e9_LeaveEvent(this, event); };
	void lower() { callbackNormyOkno42a7e9_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackNormyOkno42a7e9_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackNormyOkno42a7e9_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackNormyOkno42a7e9_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackNormyOkno42a7e9_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackNormyOkno42a7e9_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackNormyOkno42a7e9_PaintEvent(this, event); };
	void raise() { callbackNormyOkno42a7e9_Raise(this); };
	void repaint() { callbackNormyOkno42a7e9_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackNormyOkno42a7e9_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackNormyOkno42a7e9_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackNormyOkno42a7e9_SetEnabled(this, vbo); };
	void setFocus() { callbackNormyOkno42a7e9_SetFocus2(this); };
	void setHidden(bool hidden) { callbackNormyOkno42a7e9_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackNormyOkno42a7e9_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackNormyOkno42a7e9_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackNormyOkno42a7e9_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackNormyOkno42a7e9_SetWindowTitle(this, vqsPacked); };
	void show() { callbackNormyOkno42a7e9_Show(this); };
	void showEvent(QShowEvent * event) { callbackNormyOkno42a7e9_ShowEvent(this, event); };
	void showFullScreen() { callbackNormyOkno42a7e9_ShowFullScreen(this); };
	void showMaximized() { callbackNormyOkno42a7e9_ShowMaximized(this); };
	void showMinimized() { callbackNormyOkno42a7e9_ShowMinimized(this); };
	void showNormal() { callbackNormyOkno42a7e9_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackNormyOkno42a7e9_TabletEvent(this, event); };
	void update() { callbackNormyOkno42a7e9_Update(this); };
	void updateMicroFocus() { callbackNormyOkno42a7e9_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackNormyOkno42a7e9_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackNormyOkno42a7e9_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackNormyOkno42a7e9_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackNormyOkno42a7e9_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackNormyOkno42a7e9_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackNormyOkno42a7e9_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackNormyOkno42a7e9_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackNormyOkno42a7e9_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackNormyOkno42a7e9_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackNormyOkno42a7e9_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackNormyOkno42a7e9_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackNormyOkno42a7e9_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackNormyOkno42a7e9_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackNormyOkno42a7e9_CustomEvent(this, event); };
	void deleteLater() { callbackNormyOkno42a7e9_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackNormyOkno42a7e9_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackNormyOkno42a7e9_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackNormyOkno42a7e9_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackNormyOkno42a7e9_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(NormyOkno42a7e9*)


void NormyOkno42a7e9_NormyOkno42a7e9_QRegisterMetaTypes() {
}

class Pilot42a7e9: public QMainWindow
{
Q_OBJECT
public:
	Pilot42a7e9(QWidget *parent = Q_NULLPTR, Qt::WindowFlags flags = Qt::WindowFlags()) : QMainWindow(parent, flags) {qRegisterMetaType<quintptr>("quintptr");Pilot42a7e9_Pilot42a7e9_QRegisterMetaType();Pilot42a7e9_Pilot42a7e9_QRegisterMetaTypes();callbackPilot42a7e9_Constructor(this);};
	QMenu * createPopupMenu() { return static_cast<QMenu*>(callbackPilot42a7e9_CreatePopupMenu(this)); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackPilot42a7e9_ContextMenuEvent(this, event); };
	void Signal_IconSizeChanged(const QSize & iconSize) { callbackPilot42a7e9_IconSizeChanged(this, const_cast<QSize*>(&iconSize)); };
	void setAnimated(bool enabled) { callbackPilot42a7e9_SetAnimated(this, enabled); };
	void setDockNestingEnabled(bool enabled) { callbackPilot42a7e9_SetDockNestingEnabled(this, enabled); };
	void setUnifiedTitleAndToolBarOnMac(bool set) { callbackPilot42a7e9_SetUnifiedTitleAndToolBarOnMac(this, set); };
	void Signal_TabifiedDockWidgetActivated(QDockWidget * dockWidget) { callbackPilot42a7e9_TabifiedDockWidgetActivated(this, dockWidget); };
	void Signal_ToolButtonStyleChanged(Qt::ToolButtonStyle toolButtonStyle) { callbackPilot42a7e9_ToolButtonStyleChanged(this, toolButtonStyle); };
	bool close() { return callbackPilot42a7e9_Close(this) != 0; };
	bool focusNextPrevChild(bool next) { return callbackPilot42a7e9_FocusNextPrevChild(this, next) != 0; };
	void actionEvent(QActionEvent * event) { callbackPilot42a7e9_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackPilot42a7e9_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackPilot42a7e9_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackPilot42a7e9_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackPilot42a7e9_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackPilot42a7e9_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackPilot42a7e9_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackPilot42a7e9_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackPilot42a7e9_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackPilot42a7e9_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackPilot42a7e9_FocusOutEvent(this, event); };
	void hide() { callbackPilot42a7e9_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackPilot42a7e9_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackPilot42a7e9_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackPilot42a7e9_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackPilot42a7e9_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackPilot42a7e9_LeaveEvent(this, event); };
	void lower() { callbackPilot42a7e9_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackPilot42a7e9_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackPilot42a7e9_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackPilot42a7e9_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackPilot42a7e9_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackPilot42a7e9_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackPilot42a7e9_PaintEvent(this, event); };
	void raise() { callbackPilot42a7e9_Raise(this); };
	void repaint() { callbackPilot42a7e9_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackPilot42a7e9_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackPilot42a7e9_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackPilot42a7e9_SetEnabled(this, vbo); };
	void setFocus() { callbackPilot42a7e9_SetFocus2(this); };
	void setHidden(bool hidden) { callbackPilot42a7e9_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackPilot42a7e9_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackPilot42a7e9_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackPilot42a7e9_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackPilot42a7e9_SetWindowTitle(this, vqsPacked); };
	void show() { callbackPilot42a7e9_Show(this); };
	void showEvent(QShowEvent * event) { callbackPilot42a7e9_ShowEvent(this, event); };
	void showFullScreen() { callbackPilot42a7e9_ShowFullScreen(this); };
	void showMaximized() { callbackPilot42a7e9_ShowMaximized(this); };
	void showMinimized() { callbackPilot42a7e9_ShowMinimized(this); };
	void showNormal() { callbackPilot42a7e9_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackPilot42a7e9_TabletEvent(this, event); };
	void update() { callbackPilot42a7e9_Update(this); };
	void updateMicroFocus() { callbackPilot42a7e9_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackPilot42a7e9_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackPilot42a7e9_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackPilot42a7e9_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackPilot42a7e9_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackPilot42a7e9_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackPilot42a7e9_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackPilot42a7e9_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackPilot42a7e9_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackPilot42a7e9_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackPilot42a7e9_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackPilot42a7e9_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackPilot42a7e9_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackPilot42a7e9_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackPilot42a7e9_CustomEvent(this, event); };
	void deleteLater() { callbackPilot42a7e9_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackPilot42a7e9_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackPilot42a7e9_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackPilot42a7e9_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackPilot42a7e9_TimerEvent(this, event); };
signals:
public slots:
	void addTreeItem(QString p0, QString p1, QString p2, QString p3, QString p4, QString p5, QString p6, QString p7, QString p8, QString p9) { QByteArray tf187ce = p0.toUtf8(); Moc_PackedString p0Packed = { const_cast<char*>(tf187ce.prepend("WHITESPACE").constData()+10), tf187ce.size()-10 };QByteArray tb78f57 = p1.toUtf8(); Moc_PackedString p1Packed = { const_cast<char*>(tb78f57.prepend("WHITESPACE").constData()+10), tb78f57.size()-10 };QByteArray tc5fd96 = p2.toUtf8(); Moc_PackedString p2Packed = { const_cast<char*>(tc5fd96.prepend("WHITESPACE").constData()+10), tc5fd96.size()-10 };QByteArray te4fbe6 = p3.toUtf8(); Moc_PackedString p3Packed = { const_cast<char*>(te4fbe6.prepend("WHITESPACE").constData()+10), te4fbe6.size()-10 };QByteArray t1b9645 = p4.toUtf8(); Moc_PackedString p4Packed = { const_cast<char*>(t1b9645.prepend("WHITESPACE").constData()+10), t1b9645.size()-10 };QByteArray t4197f5 = p5.toUtf8(); Moc_PackedString p5Packed = { const_cast<char*>(t4197f5.prepend("WHITESPACE").constData()+10), t4197f5.size()-10 };QByteArray t826f05 = p6.toUtf8(); Moc_PackedString p6Packed = { const_cast<char*>(t826f05.prepend("WHITESPACE").constData()+10), t826f05.size()-10 };QByteArray te88b9e = p7.toUtf8(); Moc_PackedString p7Packed = { const_cast<char*>(te88b9e.prepend("WHITESPACE").constData()+10), te88b9e.size()-10 };QByteArray t0cb9b5 = p8.toUtf8(); Moc_PackedString p8Packed = { const_cast<char*>(t0cb9b5.prepend("WHITESPACE").constData()+10), t0cb9b5.size()-10 };QByteArray te0517f = p9.toUtf8(); Moc_PackedString p9Packed = { const_cast<char*>(te0517f.prepend("WHITESPACE").constData()+10), te0517f.size()-10 };callbackPilot42a7e9_AddTreeItem(this, p0Packed, p1Packed, p2Packed, p3Packed, p4Packed, p5Packed, p6Packed, p7Packed, p8Packed, p9Packed); };
private:
};

Q_DECLARE_METATYPE(Pilot42a7e9*)


void Pilot42a7e9_Pilot42a7e9_QRegisterMetaTypes() {
}

void Pilot42a7e9_AddTreeItem(void* ptr, struct Moc_PackedString p0, struct Moc_PackedString p1, struct Moc_PackedString p2, struct Moc_PackedString p3, struct Moc_PackedString p4, struct Moc_PackedString p5, struct Moc_PackedString p6, struct Moc_PackedString p7, struct Moc_PackedString p8, struct Moc_PackedString p9)
{
	QMetaObject::invokeMethod(static_cast<Pilot42a7e9*>(ptr), "addTreeItem", Q_ARG(QString, QString::fromUtf8(p0.data, p0.len)), Q_ARG(QString, QString::fromUtf8(p1.data, p1.len)), Q_ARG(QString, QString::fromUtf8(p2.data, p2.len)), Q_ARG(QString, QString::fromUtf8(p3.data, p3.len)), Q_ARG(QString, QString::fromUtf8(p4.data, p4.len)), Q_ARG(QString, QString::fromUtf8(p5.data, p5.len)), Q_ARG(QString, QString::fromUtf8(p6.data, p6.len)), Q_ARG(QString, QString::fromUtf8(p7.data, p7.len)), Q_ARG(QString, QString::fromUtf8(p8.data, p8.len)), Q_ARG(QString, QString::fromUtf8(p9.data, p9.len)));
}

int Pilot42a7e9_Pilot42a7e9_QRegisterMetaType()
{
	return qRegisterMetaType<Pilot42a7e9*>();
}

int Pilot42a7e9_Pilot42a7e9_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Pilot42a7e9*>(const_cast<const char*>(typeName));
}

int Pilot42a7e9_Pilot42a7e9_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Pilot42a7e9>();
#else
	return 0;
#endif
}

int Pilot42a7e9_Pilot42a7e9_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Pilot42a7e9>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Pilot42a7e9___resizeDocks_docks_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Pilot42a7e9___resizeDocks_docks_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* Pilot42a7e9___resizeDocks_docks_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

int Pilot42a7e9___resizeDocks_sizes_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Pilot42a7e9___resizeDocks_sizes_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* Pilot42a7e9___resizeDocks_sizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* Pilot42a7e9___tabifiedDockWidgets_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Pilot42a7e9___tabifiedDockWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* Pilot42a7e9___tabifiedDockWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

void* Pilot42a7e9___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Pilot42a7e9___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* Pilot42a7e9___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* Pilot42a7e9___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Pilot42a7e9___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* Pilot42a7e9___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* Pilot42a7e9___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Pilot42a7e9___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* Pilot42a7e9___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* Pilot42a7e9___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Pilot42a7e9___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Pilot42a7e9___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Pilot42a7e9___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Pilot42a7e9___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Pilot42a7e9___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Pilot42a7e9___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Pilot42a7e9___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Pilot42a7e9___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Pilot42a7e9___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Pilot42a7e9___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Pilot42a7e9___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Pilot42a7e9___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Pilot42a7e9___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Pilot42a7e9___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Pilot42a7e9_NewPilot(void* parent, long long flags)
{
		return new Pilot42a7e9(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void Pilot42a7e9_DestroyPilot(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->~Pilot42a7e9();
}

void* Pilot42a7e9_CreatePopupMenuDefault(void* ptr)
{
	return static_cast<Pilot42a7e9*>(ptr)->QMainWindow::createPopupMenu();
}

char Pilot42a7e9_EventDefault(void* ptr, void* event)
{
	return static_cast<Pilot42a7e9*>(ptr)->QMainWindow::event(static_cast<QEvent*>(event));
}

void Pilot42a7e9_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void Pilot42a7e9_SetAnimatedDefault(void* ptr, char enabled)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::setAnimated(enabled != 0);
}

void Pilot42a7e9_SetDockNestingEnabledDefault(void* ptr, char enabled)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::setDockNestingEnabled(enabled != 0);
}

void Pilot42a7e9_SetUnifiedTitleAndToolBarOnMacDefault(void* ptr, char set)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::setUnifiedTitleAndToolBarOnMac(set != 0);
}

char Pilot42a7e9_CloseDefault(void* ptr)
{
	return static_cast<Pilot42a7e9*>(ptr)->QMainWindow::close();
}

char Pilot42a7e9_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<Pilot42a7e9*>(ptr)->QMainWindow::focusNextPrevChild(next != 0);
}

void Pilot42a7e9_ActionEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::actionEvent(static_cast<QActionEvent*>(event));
}

void Pilot42a7e9_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::changeEvent(static_cast<QEvent*>(event));
}

void Pilot42a7e9_CloseEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::closeEvent(static_cast<QCloseEvent*>(event));
}

void Pilot42a7e9_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void Pilot42a7e9_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void Pilot42a7e9_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void Pilot42a7e9_DropEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::dropEvent(static_cast<QDropEvent*>(event));
}

void Pilot42a7e9_EnterEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::enterEvent(static_cast<QEvent*>(event));
}

void Pilot42a7e9_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::focusInEvent(static_cast<QFocusEvent*>(event));
}

void Pilot42a7e9_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void Pilot42a7e9_HideDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::hide();
}

void Pilot42a7e9_HideEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::hideEvent(static_cast<QHideEvent*>(event));
}

void Pilot42a7e9_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void Pilot42a7e9_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void Pilot42a7e9_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void Pilot42a7e9_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::leaveEvent(static_cast<QEvent*>(event));
}

void Pilot42a7e9_LowerDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::lower();
}

void Pilot42a7e9_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void Pilot42a7e9_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void Pilot42a7e9_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void Pilot42a7e9_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void Pilot42a7e9_MoveEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::moveEvent(static_cast<QMoveEvent*>(event));
}

void Pilot42a7e9_PaintEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::paintEvent(static_cast<QPaintEvent*>(event));
}

void Pilot42a7e9_RaiseDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::raise();
}

void Pilot42a7e9_RepaintDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::repaint();
}

void Pilot42a7e9_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::resizeEvent(static_cast<QResizeEvent*>(event));
}

void Pilot42a7e9_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::setDisabled(disable != 0);
}

void Pilot42a7e9_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::setEnabled(vbo != 0);
}

void Pilot42a7e9_SetFocus2Default(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::setFocus();
}

void Pilot42a7e9_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::setHidden(hidden != 0);
}

void Pilot42a7e9_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void Pilot42a7e9_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::setVisible(visible != 0);
}

void Pilot42a7e9_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::setWindowModified(vbo != 0);
}

void Pilot42a7e9_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void Pilot42a7e9_ShowDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::show();
}

void Pilot42a7e9_ShowEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::showEvent(static_cast<QShowEvent*>(event));
}

void Pilot42a7e9_ShowFullScreenDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::showFullScreen();
}

void Pilot42a7e9_ShowMaximizedDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::showMaximized();
}

void Pilot42a7e9_ShowMinimizedDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::showMinimized();
}

void Pilot42a7e9_ShowNormalDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::showNormal();
}

void Pilot42a7e9_TabletEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::tabletEvent(static_cast<QTabletEvent*>(event));
}

void Pilot42a7e9_UpdateDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::update();
}

void Pilot42a7e9_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::updateMicroFocus();
}

void Pilot42a7e9_WheelEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* Pilot42a7e9_PaintEngineDefault(void* ptr)
{
	return static_cast<Pilot42a7e9*>(ptr)->QMainWindow::paintEngine();
}

void* Pilot42a7e9_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<Pilot42a7e9*>(ptr)->QMainWindow::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* Pilot42a7e9_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<Pilot42a7e9*>(ptr)->QMainWindow::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* Pilot42a7e9_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<Pilot42a7e9*>(ptr)->QMainWindow::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char Pilot42a7e9_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<Pilot42a7e9*>(ptr)->QMainWindow::hasHeightForWidth();
}

int Pilot42a7e9_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<Pilot42a7e9*>(ptr)->QMainWindow::heightForWidth(w);
}

int Pilot42a7e9_MetricDefault(void* ptr, long long m)
{
	return static_cast<Pilot42a7e9*>(ptr)->QMainWindow::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char Pilot42a7e9_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Pilot42a7e9*>(ptr)->QMainWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Pilot42a7e9_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::childEvent(static_cast<QChildEvent*>(event));
}

void Pilot42a7e9_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Pilot42a7e9_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::customEvent(static_cast<QEvent*>(event));
}

void Pilot42a7e9_DeleteLaterDefault(void* ptr)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::deleteLater();
}

void Pilot42a7e9_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Pilot42a7e9_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Pilot42a7e9*>(ptr)->QMainWindow::timerEvent(static_cast<QTimerEvent*>(event));
}



int ProduktyOkno42a7e9_ProduktyOkno42a7e9_QRegisterMetaType()
{
	return qRegisterMetaType<ProduktyOkno42a7e9*>();
}

int ProduktyOkno42a7e9_ProduktyOkno42a7e9_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ProduktyOkno42a7e9*>(const_cast<const char*>(typeName));
}

int ProduktyOkno42a7e9_ProduktyOkno42a7e9_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ProduktyOkno42a7e9>();
#else
	return 0;
#endif
}

int ProduktyOkno42a7e9_ProduktyOkno42a7e9_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ProduktyOkno42a7e9>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* ProduktyOkno42a7e9___resizeDocks_docks_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyOkno42a7e9___resizeDocks_docks_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* ProduktyOkno42a7e9___resizeDocks_docks_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

int ProduktyOkno42a7e9___resizeDocks_sizes_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyOkno42a7e9___resizeDocks_sizes_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* ProduktyOkno42a7e9___resizeDocks_sizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* ProduktyOkno42a7e9___tabifiedDockWidgets_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyOkno42a7e9___tabifiedDockWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* ProduktyOkno42a7e9___tabifiedDockWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

void* ProduktyOkno42a7e9___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyOkno42a7e9___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* ProduktyOkno42a7e9___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* ProduktyOkno42a7e9___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyOkno42a7e9___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* ProduktyOkno42a7e9___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* ProduktyOkno42a7e9___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyOkno42a7e9___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* ProduktyOkno42a7e9___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* ProduktyOkno42a7e9___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ProduktyOkno42a7e9___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ProduktyOkno42a7e9___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ProduktyOkno42a7e9___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyOkno42a7e9___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ProduktyOkno42a7e9___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ProduktyOkno42a7e9___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyOkno42a7e9___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ProduktyOkno42a7e9___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ProduktyOkno42a7e9___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyOkno42a7e9___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ProduktyOkno42a7e9___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ProduktyOkno42a7e9___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyOkno42a7e9___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ProduktyOkno42a7e9___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ProduktyOkno42a7e9_NewProduktyOkno(void* parent, long long flags)
{
		return new ProduktyOkno42a7e9(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void ProduktyOkno42a7e9_DestroyProduktyOkno(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->~ProduktyOkno42a7e9();
}

void* ProduktyOkno42a7e9_CreatePopupMenuDefault(void* ptr)
{
	return static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::createPopupMenu();
}

char ProduktyOkno42a7e9_EventDefault(void* ptr, void* event)
{
	return static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::event(static_cast<QEvent*>(event));
}

void ProduktyOkno42a7e9_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void ProduktyOkno42a7e9_SetAnimatedDefault(void* ptr, char enabled)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::setAnimated(enabled != 0);
}

void ProduktyOkno42a7e9_SetDockNestingEnabledDefault(void* ptr, char enabled)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::setDockNestingEnabled(enabled != 0);
}

void ProduktyOkno42a7e9_SetUnifiedTitleAndToolBarOnMacDefault(void* ptr, char set)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::setUnifiedTitleAndToolBarOnMac(set != 0);
}

char ProduktyOkno42a7e9_CloseDefault(void* ptr)
{
	return static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::close();
}

char ProduktyOkno42a7e9_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::focusNextPrevChild(next != 0);
}

void ProduktyOkno42a7e9_ActionEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::actionEvent(static_cast<QActionEvent*>(event));
}

void ProduktyOkno42a7e9_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::changeEvent(static_cast<QEvent*>(event));
}

void ProduktyOkno42a7e9_CloseEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::closeEvent(static_cast<QCloseEvent*>(event));
}

void ProduktyOkno42a7e9_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void ProduktyOkno42a7e9_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void ProduktyOkno42a7e9_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void ProduktyOkno42a7e9_DropEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::dropEvent(static_cast<QDropEvent*>(event));
}

void ProduktyOkno42a7e9_EnterEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::enterEvent(static_cast<QEvent*>(event));
}

void ProduktyOkno42a7e9_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::focusInEvent(static_cast<QFocusEvent*>(event));
}

void ProduktyOkno42a7e9_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void ProduktyOkno42a7e9_HideDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::hide();
}

void ProduktyOkno42a7e9_HideEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::hideEvent(static_cast<QHideEvent*>(event));
}

void ProduktyOkno42a7e9_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void ProduktyOkno42a7e9_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void ProduktyOkno42a7e9_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void ProduktyOkno42a7e9_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::leaveEvent(static_cast<QEvent*>(event));
}

void ProduktyOkno42a7e9_LowerDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::lower();
}

void ProduktyOkno42a7e9_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void ProduktyOkno42a7e9_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void ProduktyOkno42a7e9_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void ProduktyOkno42a7e9_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void ProduktyOkno42a7e9_MoveEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::moveEvent(static_cast<QMoveEvent*>(event));
}

void ProduktyOkno42a7e9_PaintEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::paintEvent(static_cast<QPaintEvent*>(event));
}

void ProduktyOkno42a7e9_RaiseDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::raise();
}

void ProduktyOkno42a7e9_RepaintDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::repaint();
}

void ProduktyOkno42a7e9_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::resizeEvent(static_cast<QResizeEvent*>(event));
}

void ProduktyOkno42a7e9_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::setDisabled(disable != 0);
}

void ProduktyOkno42a7e9_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::setEnabled(vbo != 0);
}

void ProduktyOkno42a7e9_SetFocus2Default(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::setFocus();
}

void ProduktyOkno42a7e9_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::setHidden(hidden != 0);
}

void ProduktyOkno42a7e9_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void ProduktyOkno42a7e9_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::setVisible(visible != 0);
}

void ProduktyOkno42a7e9_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::setWindowModified(vbo != 0);
}

void ProduktyOkno42a7e9_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void ProduktyOkno42a7e9_ShowDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::show();
}

void ProduktyOkno42a7e9_ShowEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::showEvent(static_cast<QShowEvent*>(event));
}

void ProduktyOkno42a7e9_ShowFullScreenDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::showFullScreen();
}

void ProduktyOkno42a7e9_ShowMaximizedDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::showMaximized();
}

void ProduktyOkno42a7e9_ShowMinimizedDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::showMinimized();
}

void ProduktyOkno42a7e9_ShowNormalDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::showNormal();
}

void ProduktyOkno42a7e9_TabletEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::tabletEvent(static_cast<QTabletEvent*>(event));
}

void ProduktyOkno42a7e9_UpdateDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::update();
}

void ProduktyOkno42a7e9_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::updateMicroFocus();
}

void ProduktyOkno42a7e9_WheelEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* ProduktyOkno42a7e9_PaintEngineDefault(void* ptr)
{
	return static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::paintEngine();
}

void* ProduktyOkno42a7e9_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* ProduktyOkno42a7e9_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* ProduktyOkno42a7e9_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char ProduktyOkno42a7e9_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::hasHeightForWidth();
}

int ProduktyOkno42a7e9_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::heightForWidth(w);
}

int ProduktyOkno42a7e9_MetricDefault(void* ptr, long long m)
{
	return static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char ProduktyOkno42a7e9_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void ProduktyOkno42a7e9_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::childEvent(static_cast<QChildEvent*>(event));
}

void ProduktyOkno42a7e9_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ProduktyOkno42a7e9_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::customEvent(static_cast<QEvent*>(event));
}

void ProduktyOkno42a7e9_DeleteLaterDefault(void* ptr)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::deleteLater();
}

void ProduktyOkno42a7e9_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void ProduktyOkno42a7e9_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyOkno42a7e9*>(ptr)->QMainWindow::timerEvent(static_cast<QTimerEvent*>(event));
}



int WydzialyOkno42a7e9_WydzialyOkno42a7e9_QRegisterMetaType()
{
	return qRegisterMetaType<WydzialyOkno42a7e9*>();
}

int WydzialyOkno42a7e9_WydzialyOkno42a7e9_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WydzialyOkno42a7e9*>(const_cast<const char*>(typeName));
}

int WydzialyOkno42a7e9_WydzialyOkno42a7e9_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WydzialyOkno42a7e9>();
#else
	return 0;
#endif
}

int WydzialyOkno42a7e9_WydzialyOkno42a7e9_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WydzialyOkno42a7e9>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* WydzialyOkno42a7e9___resizeDocks_docks_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WydzialyOkno42a7e9___resizeDocks_docks_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* WydzialyOkno42a7e9___resizeDocks_docks_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

int WydzialyOkno42a7e9___resizeDocks_sizes_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WydzialyOkno42a7e9___resizeDocks_sizes_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WydzialyOkno42a7e9___resizeDocks_sizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WydzialyOkno42a7e9___tabifiedDockWidgets_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WydzialyOkno42a7e9___tabifiedDockWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* WydzialyOkno42a7e9___tabifiedDockWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

void* WydzialyOkno42a7e9___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WydzialyOkno42a7e9___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* WydzialyOkno42a7e9___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* WydzialyOkno42a7e9___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WydzialyOkno42a7e9___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* WydzialyOkno42a7e9___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* WydzialyOkno42a7e9___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WydzialyOkno42a7e9___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* WydzialyOkno42a7e9___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* WydzialyOkno42a7e9___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WydzialyOkno42a7e9___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WydzialyOkno42a7e9___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WydzialyOkno42a7e9___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WydzialyOkno42a7e9___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WydzialyOkno42a7e9___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WydzialyOkno42a7e9___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WydzialyOkno42a7e9___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WydzialyOkno42a7e9___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WydzialyOkno42a7e9___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WydzialyOkno42a7e9___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WydzialyOkno42a7e9___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WydzialyOkno42a7e9___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WydzialyOkno42a7e9___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WydzialyOkno42a7e9___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WydzialyOkno42a7e9_NewWydzialyOkno(void* parent, long long flags)
{
		return new WydzialyOkno42a7e9(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void WydzialyOkno42a7e9_DestroyWydzialyOkno(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->~WydzialyOkno42a7e9();
}

void* WydzialyOkno42a7e9_CreatePopupMenuDefault(void* ptr)
{
	return static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::createPopupMenu();
}

char WydzialyOkno42a7e9_EventDefault(void* ptr, void* event)
{
	return static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::event(static_cast<QEvent*>(event));
}

void WydzialyOkno42a7e9_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void WydzialyOkno42a7e9_SetAnimatedDefault(void* ptr, char enabled)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::setAnimated(enabled != 0);
}

void WydzialyOkno42a7e9_SetDockNestingEnabledDefault(void* ptr, char enabled)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::setDockNestingEnabled(enabled != 0);
}

void WydzialyOkno42a7e9_SetUnifiedTitleAndToolBarOnMacDefault(void* ptr, char set)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::setUnifiedTitleAndToolBarOnMac(set != 0);
}

char WydzialyOkno42a7e9_CloseDefault(void* ptr)
{
	return static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::close();
}

char WydzialyOkno42a7e9_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::focusNextPrevChild(next != 0);
}

void WydzialyOkno42a7e9_ActionEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::actionEvent(static_cast<QActionEvent*>(event));
}

void WydzialyOkno42a7e9_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::changeEvent(static_cast<QEvent*>(event));
}

void WydzialyOkno42a7e9_CloseEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::closeEvent(static_cast<QCloseEvent*>(event));
}

void WydzialyOkno42a7e9_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void WydzialyOkno42a7e9_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void WydzialyOkno42a7e9_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void WydzialyOkno42a7e9_DropEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::dropEvent(static_cast<QDropEvent*>(event));
}

void WydzialyOkno42a7e9_EnterEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::enterEvent(static_cast<QEvent*>(event));
}

void WydzialyOkno42a7e9_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::focusInEvent(static_cast<QFocusEvent*>(event));
}

void WydzialyOkno42a7e9_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void WydzialyOkno42a7e9_HideDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::hide();
}

void WydzialyOkno42a7e9_HideEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::hideEvent(static_cast<QHideEvent*>(event));
}

void WydzialyOkno42a7e9_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void WydzialyOkno42a7e9_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void WydzialyOkno42a7e9_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void WydzialyOkno42a7e9_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::leaveEvent(static_cast<QEvent*>(event));
}

void WydzialyOkno42a7e9_LowerDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::lower();
}

void WydzialyOkno42a7e9_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void WydzialyOkno42a7e9_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void WydzialyOkno42a7e9_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void WydzialyOkno42a7e9_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void WydzialyOkno42a7e9_MoveEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::moveEvent(static_cast<QMoveEvent*>(event));
}

void WydzialyOkno42a7e9_PaintEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::paintEvent(static_cast<QPaintEvent*>(event));
}

void WydzialyOkno42a7e9_RaiseDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::raise();
}

void WydzialyOkno42a7e9_RepaintDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::repaint();
}

void WydzialyOkno42a7e9_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::resizeEvent(static_cast<QResizeEvent*>(event));
}

void WydzialyOkno42a7e9_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::setDisabled(disable != 0);
}

void WydzialyOkno42a7e9_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::setEnabled(vbo != 0);
}

void WydzialyOkno42a7e9_SetFocus2Default(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::setFocus();
}

void WydzialyOkno42a7e9_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::setHidden(hidden != 0);
}

void WydzialyOkno42a7e9_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void WydzialyOkno42a7e9_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::setVisible(visible != 0);
}

void WydzialyOkno42a7e9_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::setWindowModified(vbo != 0);
}

void WydzialyOkno42a7e9_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void WydzialyOkno42a7e9_ShowDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::show();
}

void WydzialyOkno42a7e9_ShowEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::showEvent(static_cast<QShowEvent*>(event));
}

void WydzialyOkno42a7e9_ShowFullScreenDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::showFullScreen();
}

void WydzialyOkno42a7e9_ShowMaximizedDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::showMaximized();
}

void WydzialyOkno42a7e9_ShowMinimizedDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::showMinimized();
}

void WydzialyOkno42a7e9_ShowNormalDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::showNormal();
}

void WydzialyOkno42a7e9_TabletEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::tabletEvent(static_cast<QTabletEvent*>(event));
}

void WydzialyOkno42a7e9_UpdateDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::update();
}

void WydzialyOkno42a7e9_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::updateMicroFocus();
}

void WydzialyOkno42a7e9_WheelEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* WydzialyOkno42a7e9_PaintEngineDefault(void* ptr)
{
	return static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::paintEngine();
}

void* WydzialyOkno42a7e9_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* WydzialyOkno42a7e9_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* WydzialyOkno42a7e9_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char WydzialyOkno42a7e9_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::hasHeightForWidth();
}

int WydzialyOkno42a7e9_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::heightForWidth(w);
}

int WydzialyOkno42a7e9_MetricDefault(void* ptr, long long m)
{
	return static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char WydzialyOkno42a7e9_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WydzialyOkno42a7e9_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::childEvent(static_cast<QChildEvent*>(event));
}

void WydzialyOkno42a7e9_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WydzialyOkno42a7e9_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::customEvent(static_cast<QEvent*>(event));
}

void WydzialyOkno42a7e9_DeleteLaterDefault(void* ptr)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::deleteLater();
}

void WydzialyOkno42a7e9_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void WydzialyOkno42a7e9_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WydzialyOkno42a7e9*>(ptr)->QMainWindow::timerEvent(static_cast<QTimerEvent*>(event));
}



int NormyOkno42a7e9_NormyOkno42a7e9_QRegisterMetaType()
{
	return qRegisterMetaType<NormyOkno42a7e9*>();
}

int NormyOkno42a7e9_NormyOkno42a7e9_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<NormyOkno42a7e9*>(const_cast<const char*>(typeName));
}

int NormyOkno42a7e9_NormyOkno42a7e9_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<NormyOkno42a7e9>();
#else
	return 0;
#endif
}

int NormyOkno42a7e9_NormyOkno42a7e9_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<NormyOkno42a7e9>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* NormyOkno42a7e9___resizeDocks_docks_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void NormyOkno42a7e9___resizeDocks_docks_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* NormyOkno42a7e9___resizeDocks_docks_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

int NormyOkno42a7e9___resizeDocks_sizes_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void NormyOkno42a7e9___resizeDocks_sizes_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* NormyOkno42a7e9___resizeDocks_sizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* NormyOkno42a7e9___tabifiedDockWidgets_atList(void* ptr, int i)
{
	return ({QDockWidget * tmp = static_cast<QList<QDockWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QDockWidget *>*>(ptr)->size()-1) { static_cast<QList<QDockWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void NormyOkno42a7e9___tabifiedDockWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QDockWidget *>*>(ptr)->append(static_cast<QDockWidget*>(i));
}

void* NormyOkno42a7e9___tabifiedDockWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QDockWidget *>();
}

void* NormyOkno42a7e9___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void NormyOkno42a7e9___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* NormyOkno42a7e9___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* NormyOkno42a7e9___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void NormyOkno42a7e9___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* NormyOkno42a7e9___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* NormyOkno42a7e9___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void NormyOkno42a7e9___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* NormyOkno42a7e9___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* NormyOkno42a7e9___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void NormyOkno42a7e9___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* NormyOkno42a7e9___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* NormyOkno42a7e9___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void NormyOkno42a7e9___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* NormyOkno42a7e9___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* NormyOkno42a7e9___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void NormyOkno42a7e9___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* NormyOkno42a7e9___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* NormyOkno42a7e9___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void NormyOkno42a7e9___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* NormyOkno42a7e9___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* NormyOkno42a7e9___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void NormyOkno42a7e9___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* NormyOkno42a7e9___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* NormyOkno42a7e9_NewNormyOkno(void* parent, long long flags)
{
		return new NormyOkno42a7e9(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void NormyOkno42a7e9_DestroyNormyOkno(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->~NormyOkno42a7e9();
}

void* NormyOkno42a7e9_CreatePopupMenuDefault(void* ptr)
{
	return static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::createPopupMenu();
}

char NormyOkno42a7e9_EventDefault(void* ptr, void* event)
{
	return static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::event(static_cast<QEvent*>(event));
}

void NormyOkno42a7e9_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void NormyOkno42a7e9_SetAnimatedDefault(void* ptr, char enabled)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::setAnimated(enabled != 0);
}

void NormyOkno42a7e9_SetDockNestingEnabledDefault(void* ptr, char enabled)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::setDockNestingEnabled(enabled != 0);
}

void NormyOkno42a7e9_SetUnifiedTitleAndToolBarOnMacDefault(void* ptr, char set)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::setUnifiedTitleAndToolBarOnMac(set != 0);
}

char NormyOkno42a7e9_CloseDefault(void* ptr)
{
	return static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::close();
}

char NormyOkno42a7e9_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::focusNextPrevChild(next != 0);
}

void NormyOkno42a7e9_ActionEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::actionEvent(static_cast<QActionEvent*>(event));
}

void NormyOkno42a7e9_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::changeEvent(static_cast<QEvent*>(event));
}

void NormyOkno42a7e9_CloseEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::closeEvent(static_cast<QCloseEvent*>(event));
}

void NormyOkno42a7e9_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void NormyOkno42a7e9_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void NormyOkno42a7e9_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void NormyOkno42a7e9_DropEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::dropEvent(static_cast<QDropEvent*>(event));
}

void NormyOkno42a7e9_EnterEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::enterEvent(static_cast<QEvent*>(event));
}

void NormyOkno42a7e9_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::focusInEvent(static_cast<QFocusEvent*>(event));
}

void NormyOkno42a7e9_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void NormyOkno42a7e9_HideDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::hide();
}

void NormyOkno42a7e9_HideEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::hideEvent(static_cast<QHideEvent*>(event));
}

void NormyOkno42a7e9_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void NormyOkno42a7e9_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void NormyOkno42a7e9_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void NormyOkno42a7e9_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::leaveEvent(static_cast<QEvent*>(event));
}

void NormyOkno42a7e9_LowerDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::lower();
}

void NormyOkno42a7e9_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void NormyOkno42a7e9_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void NormyOkno42a7e9_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void NormyOkno42a7e9_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void NormyOkno42a7e9_MoveEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::moveEvent(static_cast<QMoveEvent*>(event));
}

void NormyOkno42a7e9_PaintEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::paintEvent(static_cast<QPaintEvent*>(event));
}

void NormyOkno42a7e9_RaiseDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::raise();
}

void NormyOkno42a7e9_RepaintDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::repaint();
}

void NormyOkno42a7e9_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::resizeEvent(static_cast<QResizeEvent*>(event));
}

void NormyOkno42a7e9_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::setDisabled(disable != 0);
}

void NormyOkno42a7e9_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::setEnabled(vbo != 0);
}

void NormyOkno42a7e9_SetFocus2Default(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::setFocus();
}

void NormyOkno42a7e9_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::setHidden(hidden != 0);
}

void NormyOkno42a7e9_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void NormyOkno42a7e9_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::setVisible(visible != 0);
}

void NormyOkno42a7e9_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::setWindowModified(vbo != 0);
}

void NormyOkno42a7e9_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void NormyOkno42a7e9_ShowDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::show();
}

void NormyOkno42a7e9_ShowEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::showEvent(static_cast<QShowEvent*>(event));
}

void NormyOkno42a7e9_ShowFullScreenDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::showFullScreen();
}

void NormyOkno42a7e9_ShowMaximizedDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::showMaximized();
}

void NormyOkno42a7e9_ShowMinimizedDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::showMinimized();
}

void NormyOkno42a7e9_ShowNormalDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::showNormal();
}

void NormyOkno42a7e9_TabletEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::tabletEvent(static_cast<QTabletEvent*>(event));
}

void NormyOkno42a7e9_UpdateDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::update();
}

void NormyOkno42a7e9_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::updateMicroFocus();
}

void NormyOkno42a7e9_WheelEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* NormyOkno42a7e9_PaintEngineDefault(void* ptr)
{
	return static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::paintEngine();
}

void* NormyOkno42a7e9_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* NormyOkno42a7e9_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* NormyOkno42a7e9_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char NormyOkno42a7e9_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::hasHeightForWidth();
}

int NormyOkno42a7e9_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::heightForWidth(w);
}

int NormyOkno42a7e9_MetricDefault(void* ptr, long long m)
{
	return static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char NormyOkno42a7e9_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void NormyOkno42a7e9_ChildEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::childEvent(static_cast<QChildEvent*>(event));
}

void NormyOkno42a7e9_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void NormyOkno42a7e9_CustomEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::customEvent(static_cast<QEvent*>(event));
}

void NormyOkno42a7e9_DeleteLaterDefault(void* ptr)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::deleteLater();
}

void NormyOkno42a7e9_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void NormyOkno42a7e9_TimerEventDefault(void* ptr, void* event)
{
	static_cast<NormyOkno42a7e9*>(ptr)->QMainWindow::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"

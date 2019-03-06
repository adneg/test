

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractListModel>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QQuickItem>
#include <QRadioData>
#include <QSize>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>


typedef QMap<qint32, QByteArray> type378cdd;
typedef QMap<qint32, QByteArray> type378cdd;
typedef QMap<qint32, QByteArray> type378cdd;
class Produktd7c2e5: public QObject
{
Q_OBJECT
Q_PROPERTY(QString produktName READ produktName WRITE setProduktName NOTIFY produktNameChanged)
Q_PROPERTY(float produktW READ produktW WRITE setProduktW NOTIFY produktWChanged)
Q_PROPERTY(qint32 produktId READ produktId WRITE setProduktId NOTIFY produktIdChanged)
public:
	Produktd7c2e5(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Produktd7c2e5_Produktd7c2e5_QRegisterMetaType();Produktd7c2e5_Produktd7c2e5_QRegisterMetaTypes();callbackProduktd7c2e5_Constructor(this);};
	QString produktName() { return ({ Moc_PackedString tempVal = callbackProduktd7c2e5_ProduktName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setProduktName(QString produktName) { QByteArray t463293 = produktName.toUtf8(); Moc_PackedString produktNamePacked = { const_cast<char*>(t463293.prepend("WHITESPACE").constData()+10), t463293.size()-10 };callbackProduktd7c2e5_SetProduktName(this, produktNamePacked); };
	void Signal_ProduktNameChanged(QString produktName) { QByteArray t463293 = produktName.toUtf8(); Moc_PackedString produktNamePacked = { const_cast<char*>(t463293.prepend("WHITESPACE").constData()+10), t463293.size()-10 };callbackProduktd7c2e5_ProduktNameChanged(this, produktNamePacked); };
	float produktW() { return callbackProduktd7c2e5_ProduktW(this); };
	void setProduktW(float produktW) { callbackProduktd7c2e5_SetProduktW(this, produktW); };
	void Signal_ProduktWChanged(float produktW) { callbackProduktd7c2e5_ProduktWChanged(this, produktW); };
	qint32 produktId() { return callbackProduktd7c2e5_ProduktId(this); };
	void setProduktId(qint32 produktId) { callbackProduktd7c2e5_SetProduktId(this, produktId); };
	void Signal_ProduktIdChanged(qint32 produktId) { callbackProduktd7c2e5_ProduktIdChanged(this, produktId); };
	 ~Produktd7c2e5() { callbackProduktd7c2e5_DestroyProdukt(this); };
	bool event(QEvent * e) { return callbackProduktd7c2e5_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackProduktd7c2e5_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackProduktd7c2e5_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackProduktd7c2e5_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackProduktd7c2e5_CustomEvent(this, event); };
	void deleteLater() { callbackProduktd7c2e5_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackProduktd7c2e5_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackProduktd7c2e5_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackProduktd7c2e5_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackProduktd7c2e5_TimerEvent(this, event); };
	QString produktNameDefault() { return _produktName; };
	void setProduktNameDefault(QString p) { if (p != _produktName) { _produktName = p; produktNameChanged(_produktName); } };
	float produktWDefault() { return _produktW; };
	void setProduktWDefault(float p) { if (p != _produktW) { _produktW = p; produktWChanged(_produktW); } };
	qint32 produktIdDefault() { return _produktId; };
	void setProduktIdDefault(qint32 p) { if (p != _produktId) { _produktId = p; produktIdChanged(_produktId); } };
signals:
	void produktNameChanged(QString produktName);
	void produktWChanged(float produktW);
	void produktIdChanged(qint32 produktId);
public slots:
private:
	QString _produktName;
	float _produktW;
	qint32 _produktId;
};

Q_DECLARE_METATYPE(Produktd7c2e5*)


void Produktd7c2e5_Produktd7c2e5_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class ProduktyModeld7c2e5: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<Produktd7c2e5*> produktitem READ produktitem WRITE setProduktitem NOTIFY produktitemChanged)
public:
	ProduktyModeld7c2e5(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");ProduktyModeld7c2e5_ProduktyModeld7c2e5_QRegisterMetaType();ProduktyModeld7c2e5_ProduktyModeld7c2e5_QRegisterMetaTypes();callbackProduktyModeld7c2e5_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackProduktyModeld7c2e5_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackProduktyModeld7c2e5_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackProduktyModeld7c2e5_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<Produktd7c2e5*> produktitem() { return ({ QList<Produktd7c2e5*>* tmpP = static_cast<QList<Produktd7c2e5*>*>(callbackProduktyModeld7c2e5_Produktitem(this)); QList<Produktd7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setProduktitem(QList<Produktd7c2e5*> produktitem) { callbackProduktyModeld7c2e5_SetProduktitem(this, ({ QList<Produktd7c2e5*>* tmpValue = new QList<Produktd7c2e5*>(produktitem); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_ProduktitemChanged(QList<Produktd7c2e5*> produktitem) { callbackProduktyModeld7c2e5_ProduktitemChanged(this, ({ QList<Produktd7c2e5*>* tmpValue = new QList<Produktd7c2e5*>(produktitem); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~ProduktyModeld7c2e5() { callbackProduktyModeld7c2e5_DestroyProduktyModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackProduktyModeld7c2e5_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackProduktyModeld7c2e5_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackProduktyModeld7c2e5_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackProduktyModeld7c2e5_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackProduktyModeld7c2e5_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackProduktyModeld7c2e5_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackProduktyModeld7c2e5_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackProduktyModeld7c2e5_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackProduktyModeld7c2e5_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackProduktyModeld7c2e5_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackProduktyModeld7c2e5_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackProduktyModeld7c2e5_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackProduktyModeld7c2e5_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackProduktyModeld7c2e5_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackProduktyModeld7c2e5_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackProduktyModeld7c2e5_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackProduktyModeld7c2e5_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackProduktyModeld7c2e5_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackProduktyModeld7c2e5_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackProduktyModeld7c2e5_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackProduktyModeld7c2e5_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackProduktyModeld7c2e5_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackProduktyModeld7c2e5_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackProduktyModeld7c2e5_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackProduktyModeld7c2e5_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackProduktyModeld7c2e5_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackProduktyModeld7c2e5_ModelReset(this); };
	void resetInternalData() { callbackProduktyModeld7c2e5_ResetInternalData(this); };
	void revert() { callbackProduktyModeld7c2e5_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackProduktyModeld7c2e5_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackProduktyModeld7c2e5_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackProduktyModeld7c2e5_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackProduktyModeld7c2e5_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackProduktyModeld7c2e5_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackProduktyModeld7c2e5_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackProduktyModeld7c2e5_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackProduktyModeld7c2e5_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackProduktyModeld7c2e5_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackProduktyModeld7c2e5_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackProduktyModeld7c2e5_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackProduktyModeld7c2e5_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackProduktyModeld7c2e5_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackProduktyModeld7c2e5_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackProduktyModeld7c2e5_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackProduktyModeld7c2e5_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackProduktyModeld7c2e5_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackProduktyModeld7c2e5_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackProduktyModeld7c2e5_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackProduktyModeld7c2e5_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackProduktyModeld7c2e5_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackProduktyModeld7c2e5_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackProduktyModeld7c2e5_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackProduktyModeld7c2e5_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackProduktyModeld7c2e5_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackProduktyModeld7c2e5_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackProduktyModeld7c2e5_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackProduktyModeld7c2e5_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackProduktyModeld7c2e5_CustomEvent(this, event); };
	void deleteLater() { callbackProduktyModeld7c2e5_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackProduktyModeld7c2e5_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackProduktyModeld7c2e5_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackProduktyModeld7c2e5_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackProduktyModeld7c2e5_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<Produktd7c2e5*> produktitemDefault() { return _produktitem; };
	void setProduktitemDefault(QList<Produktd7c2e5*> p) { if (p != _produktitem) { _produktitem = p; produktitemChanged(_produktitem); } };
signals:
	void rolesChanged(type378cdd roles);
	void produktitemChanged(QList<Produktd7c2e5*> produktitem);
public slots:
	void addProdukt(Produktd7c2e5* v0) { callbackProduktyModeld7c2e5_AddProdukt(this, v0); };
	void editProdukt(qint32 row, QString produktName, float produktW, qint32 produktId) { QByteArray t463293 = produktName.toUtf8(); Moc_PackedString produktNamePacked = { const_cast<char*>(t463293.prepend("WHITESPACE").constData()+10), t463293.size()-10 };callbackProduktyModeld7c2e5_EditProdukt(this, row, produktNamePacked, produktW, produktId); };
	void removeProdukt(qint32 row) { callbackProduktyModeld7c2e5_RemoveProdukt(this, row); };
private:
	type378cdd _roles;
	QList<Produktd7c2e5*> _produktitem;
};

Q_DECLARE_METATYPE(ProduktyModeld7c2e5*)


void ProduktyModeld7c2e5_ProduktyModeld7c2e5_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
}

class Elementd7c2e5: public QObject
{
Q_OBJECT
Q_PROPERTY(QString elementGName READ elementGName WRITE setElementGName NOTIFY elementGNameChanged)
Q_PROPERTY(float elementStan READ elementStan WRITE setElementStan NOTIFY elementStanChanged)
Q_PROPERTY(QString elementPName READ elementPName WRITE setElementPName NOTIFY elementPNameChanged)
Q_PROPERTY(qint32 elementIlosc READ elementIlosc WRITE setElementIlosc NOTIFY elementIloscChanged)
Q_PROPERTY(QString elementData READ elementData WRITE setElementData NOTIFY elementDataChanged)
Q_PROPERTY(qint32 elementId READ elementId WRITE setElementId NOTIFY elementIdChanged)
public:
	Elementd7c2e5(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Elementd7c2e5_Elementd7c2e5_QRegisterMetaType();Elementd7c2e5_Elementd7c2e5_QRegisterMetaTypes();callbackElementd7c2e5_Constructor(this);};
	QString elementGName() { return ({ Moc_PackedString tempVal = callbackElementd7c2e5_ElementGName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setElementGName(QString elementGName) { QByteArray t4038c7 = elementGName.toUtf8(); Moc_PackedString elementGNamePacked = { const_cast<char*>(t4038c7.prepend("WHITESPACE").constData()+10), t4038c7.size()-10 };callbackElementd7c2e5_SetElementGName(this, elementGNamePacked); };
	void Signal_ElementGNameChanged(QString elementGName) { QByteArray t4038c7 = elementGName.toUtf8(); Moc_PackedString elementGNamePacked = { const_cast<char*>(t4038c7.prepend("WHITESPACE").constData()+10), t4038c7.size()-10 };callbackElementd7c2e5_ElementGNameChanged(this, elementGNamePacked); };
	float elementStan() { return callbackElementd7c2e5_ElementStan(this); };
	void setElementStan(float elementStan) { callbackElementd7c2e5_SetElementStan(this, elementStan); };
	void Signal_ElementStanChanged(float elementStan) { callbackElementd7c2e5_ElementStanChanged(this, elementStan); };
	QString elementPName() { return ({ Moc_PackedString tempVal = callbackElementd7c2e5_ElementPName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setElementPName(QString elementPName) { QByteArray tc66e4a = elementPName.toUtf8(); Moc_PackedString elementPNamePacked = { const_cast<char*>(tc66e4a.prepend("WHITESPACE").constData()+10), tc66e4a.size()-10 };callbackElementd7c2e5_SetElementPName(this, elementPNamePacked); };
	void Signal_ElementPNameChanged(QString elementPName) { QByteArray tc66e4a = elementPName.toUtf8(); Moc_PackedString elementPNamePacked = { const_cast<char*>(tc66e4a.prepend("WHITESPACE").constData()+10), tc66e4a.size()-10 };callbackElementd7c2e5_ElementPNameChanged(this, elementPNamePacked); };
	qint32 elementIlosc() { return callbackElementd7c2e5_ElementIlosc(this); };
	void setElementIlosc(qint32 elementIlosc) { callbackElementd7c2e5_SetElementIlosc(this, elementIlosc); };
	void Signal_ElementIloscChanged(qint32 elementIlosc) { callbackElementd7c2e5_ElementIloscChanged(this, elementIlosc); };
	QString elementData() { return ({ Moc_PackedString tempVal = callbackElementd7c2e5_ElementData(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setElementData(QString elementData) { QByteArray te1326f = elementData.toUtf8(); Moc_PackedString elementDataPacked = { const_cast<char*>(te1326f.prepend("WHITESPACE").constData()+10), te1326f.size()-10 };callbackElementd7c2e5_SetElementData(this, elementDataPacked); };
	void Signal_ElementDataChanged(QString elementData) { QByteArray te1326f = elementData.toUtf8(); Moc_PackedString elementDataPacked = { const_cast<char*>(te1326f.prepend("WHITESPACE").constData()+10), te1326f.size()-10 };callbackElementd7c2e5_ElementDataChanged(this, elementDataPacked); };
	qint32 elementId() { return callbackElementd7c2e5_ElementId(this); };
	void setElementId(qint32 elementId) { callbackElementd7c2e5_SetElementId(this, elementId); };
	void Signal_ElementIdChanged(qint32 elementId) { callbackElementd7c2e5_ElementIdChanged(this, elementId); };
	 ~Elementd7c2e5() { callbackElementd7c2e5_DestroyElement(this); };
	bool event(QEvent * e) { return callbackElementd7c2e5_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackElementd7c2e5_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackElementd7c2e5_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackElementd7c2e5_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackElementd7c2e5_CustomEvent(this, event); };
	void deleteLater() { callbackElementd7c2e5_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackElementd7c2e5_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackElementd7c2e5_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackElementd7c2e5_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackElementd7c2e5_TimerEvent(this, event); };
	QString elementGNameDefault() { return _elementGName; };
	void setElementGNameDefault(QString p) { if (p != _elementGName) { _elementGName = p; elementGNameChanged(_elementGName); } };
	float elementStanDefault() { return _elementStan; };
	void setElementStanDefault(float p) { if (p != _elementStan) { _elementStan = p; elementStanChanged(_elementStan); } };
	QString elementPNameDefault() { return _elementPName; };
	void setElementPNameDefault(QString p) { if (p != _elementPName) { _elementPName = p; elementPNameChanged(_elementPName); } };
	qint32 elementIloscDefault() { return _elementIlosc; };
	void setElementIloscDefault(qint32 p) { if (p != _elementIlosc) { _elementIlosc = p; elementIloscChanged(_elementIlosc); } };
	QString elementDataDefault() { return _elementData; };
	void setElementDataDefault(QString p) { if (p != _elementData) { _elementData = p; elementDataChanged(_elementData); } };
	qint32 elementIdDefault() { return _elementId; };
	void setElementIdDefault(qint32 p) { if (p != _elementId) { _elementId = p; elementIdChanged(_elementId); } };
signals:
	void elementGNameChanged(QString elementGName);
	void elementStanChanged(float elementStan);
	void elementPNameChanged(QString elementPName);
	void elementIloscChanged(qint32 elementIlosc);
	void elementDataChanged(QString elementData);
	void elementIdChanged(qint32 elementId);
public slots:
private:
	QString _elementGName;
	float _elementStan;
	QString _elementPName;
	qint32 _elementIlosc;
	QString _elementData;
	qint32 _elementId;
};

Q_DECLARE_METATYPE(Elementd7c2e5*)


void Elementd7c2e5_Elementd7c2e5_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class ElementyModeld7c2e5: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<Elementd7c2e5*> elementitem READ elementitem WRITE setElementitem NOTIFY elementitemChanged)
public:
	ElementyModeld7c2e5(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");ElementyModeld7c2e5_ElementyModeld7c2e5_QRegisterMetaType();ElementyModeld7c2e5_ElementyModeld7c2e5_QRegisterMetaTypes();callbackElementyModeld7c2e5_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackElementyModeld7c2e5_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackElementyModeld7c2e5_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackElementyModeld7c2e5_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<Elementd7c2e5*> elementitem() { return ({ QList<Elementd7c2e5*>* tmpP = static_cast<QList<Elementd7c2e5*>*>(callbackElementyModeld7c2e5_Elementitem(this)); QList<Elementd7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setElementitem(QList<Elementd7c2e5*> elementitem) { callbackElementyModeld7c2e5_SetElementitem(this, ({ QList<Elementd7c2e5*>* tmpValue = new QList<Elementd7c2e5*>(elementitem); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_ElementitemChanged(QList<Elementd7c2e5*> elementitem) { callbackElementyModeld7c2e5_ElementitemChanged(this, ({ QList<Elementd7c2e5*>* tmpValue = new QList<Elementd7c2e5*>(elementitem); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~ElementyModeld7c2e5() { callbackElementyModeld7c2e5_DestroyElementyModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackElementyModeld7c2e5_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackElementyModeld7c2e5_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackElementyModeld7c2e5_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackElementyModeld7c2e5_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackElementyModeld7c2e5_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackElementyModeld7c2e5_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackElementyModeld7c2e5_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackElementyModeld7c2e5_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackElementyModeld7c2e5_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackElementyModeld7c2e5_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackElementyModeld7c2e5_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackElementyModeld7c2e5_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackElementyModeld7c2e5_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackElementyModeld7c2e5_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackElementyModeld7c2e5_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackElementyModeld7c2e5_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackElementyModeld7c2e5_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackElementyModeld7c2e5_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackElementyModeld7c2e5_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackElementyModeld7c2e5_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackElementyModeld7c2e5_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackElementyModeld7c2e5_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackElementyModeld7c2e5_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackElementyModeld7c2e5_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackElementyModeld7c2e5_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackElementyModeld7c2e5_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackElementyModeld7c2e5_ModelReset(this); };
	void resetInternalData() { callbackElementyModeld7c2e5_ResetInternalData(this); };
	void revert() { callbackElementyModeld7c2e5_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackElementyModeld7c2e5_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackElementyModeld7c2e5_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackElementyModeld7c2e5_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackElementyModeld7c2e5_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackElementyModeld7c2e5_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackElementyModeld7c2e5_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackElementyModeld7c2e5_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackElementyModeld7c2e5_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackElementyModeld7c2e5_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackElementyModeld7c2e5_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackElementyModeld7c2e5_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackElementyModeld7c2e5_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackElementyModeld7c2e5_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackElementyModeld7c2e5_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackElementyModeld7c2e5_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackElementyModeld7c2e5_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackElementyModeld7c2e5_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackElementyModeld7c2e5_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackElementyModeld7c2e5_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackElementyModeld7c2e5_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackElementyModeld7c2e5_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackElementyModeld7c2e5_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackElementyModeld7c2e5_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackElementyModeld7c2e5_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackElementyModeld7c2e5_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackElementyModeld7c2e5_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackElementyModeld7c2e5_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackElementyModeld7c2e5_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackElementyModeld7c2e5_CustomEvent(this, event); };
	void deleteLater() { callbackElementyModeld7c2e5_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackElementyModeld7c2e5_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackElementyModeld7c2e5_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackElementyModeld7c2e5_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackElementyModeld7c2e5_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<Elementd7c2e5*> elementitemDefault() { return _elementitem; };
	void setElementitemDefault(QList<Elementd7c2e5*> p) { if (p != _elementitem) { _elementitem = p; elementitemChanged(_elementitem); } };
signals:
	void rolesChanged(type378cdd roles);
	void elementitemChanged(QList<Elementd7c2e5*> elementitem);
public slots:
	void addElement(Elementd7c2e5* v0) { callbackElementyModeld7c2e5_AddElement(this, v0); };
	void editElement(qint32 row, QString elementGName, float elementStan, QString elementPName, qint32 elementIlosc, QString elementData, qint32 elementId) { QByteArray t4038c7 = elementGName.toUtf8(); Moc_PackedString elementGNamePacked = { const_cast<char*>(t4038c7.prepend("WHITESPACE").constData()+10), t4038c7.size()-10 };QByteArray tc66e4a = elementPName.toUtf8(); Moc_PackedString elementPNamePacked = { const_cast<char*>(tc66e4a.prepend("WHITESPACE").constData()+10), tc66e4a.size()-10 };QByteArray te1326f = elementData.toUtf8(); Moc_PackedString elementDataPacked = { const_cast<char*>(te1326f.prepend("WHITESPACE").constData()+10), te1326f.size()-10 };callbackElementyModeld7c2e5_EditElement(this, row, elementGNamePacked, elementStan, elementPNamePacked, elementIlosc, elementDataPacked, elementId); };
	void removeElement(qint32 row) { callbackElementyModeld7c2e5_RemoveElement(this, row); };
	void removeElementId(qint32 idin) { callbackElementyModeld7c2e5_RemoveElementId(this, idin); };
private:
	type378cdd _roles;
	QList<Elementd7c2e5*> _elementitem;
};

Q_DECLARE_METATYPE(ElementyModeld7c2e5*)


void ElementyModeld7c2e5_ElementyModeld7c2e5_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
}

class Grupad7c2e5: public QObject
{
Q_OBJECT
Q_PROPERTY(QString grupaName READ grupaName WRITE setGrupaName NOTIFY grupaNameChanged)
Q_PROPERTY(qint32 grupaId READ grupaId WRITE setGrupaId NOTIFY grupaIdChanged)
public:
	Grupad7c2e5(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Grupad7c2e5_Grupad7c2e5_QRegisterMetaType();Grupad7c2e5_Grupad7c2e5_QRegisterMetaTypes();callbackGrupad7c2e5_Constructor(this);};
	QString grupaName() { return ({ Moc_PackedString tempVal = callbackGrupad7c2e5_GrupaName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setGrupaName(QString grupaName) { QByteArray t335901 = grupaName.toUtf8(); Moc_PackedString grupaNamePacked = { const_cast<char*>(t335901.prepend("WHITESPACE").constData()+10), t335901.size()-10 };callbackGrupad7c2e5_SetGrupaName(this, grupaNamePacked); };
	void Signal_GrupaNameChanged(QString grupaName) { QByteArray t335901 = grupaName.toUtf8(); Moc_PackedString grupaNamePacked = { const_cast<char*>(t335901.prepend("WHITESPACE").constData()+10), t335901.size()-10 };callbackGrupad7c2e5_GrupaNameChanged(this, grupaNamePacked); };
	qint32 grupaId() { return callbackGrupad7c2e5_GrupaId(this); };
	void setGrupaId(qint32 grupaId) { callbackGrupad7c2e5_SetGrupaId(this, grupaId); };
	void Signal_GrupaIdChanged(qint32 grupaId) { callbackGrupad7c2e5_GrupaIdChanged(this, grupaId); };
	 ~Grupad7c2e5() { callbackGrupad7c2e5_DestroyGrupa(this); };
	bool event(QEvent * e) { return callbackGrupad7c2e5_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackGrupad7c2e5_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackGrupad7c2e5_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackGrupad7c2e5_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackGrupad7c2e5_CustomEvent(this, event); };
	void deleteLater() { callbackGrupad7c2e5_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackGrupad7c2e5_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackGrupad7c2e5_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackGrupad7c2e5_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackGrupad7c2e5_TimerEvent(this, event); };
	QString grupaNameDefault() { return _grupaName; };
	void setGrupaNameDefault(QString p) { if (p != _grupaName) { _grupaName = p; grupaNameChanged(_grupaName); } };
	qint32 grupaIdDefault() { return _grupaId; };
	void setGrupaIdDefault(qint32 p) { if (p != _grupaId) { _grupaId = p; grupaIdChanged(_grupaId); } };
signals:
	void grupaNameChanged(QString grupaName);
	void grupaIdChanged(qint32 grupaId);
public slots:
private:
	QString _grupaName;
	qint32 _grupaId;
};

Q_DECLARE_METATYPE(Grupad7c2e5*)


void Grupad7c2e5_Grupad7c2e5_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class GrupyModeld7c2e5: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<Grupad7c2e5*> grupaitem READ grupaitem WRITE setGrupaitem NOTIFY grupaitemChanged)
public:
	GrupyModeld7c2e5(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");GrupyModeld7c2e5_GrupyModeld7c2e5_QRegisterMetaType();GrupyModeld7c2e5_GrupyModeld7c2e5_QRegisterMetaTypes();callbackGrupyModeld7c2e5_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackGrupyModeld7c2e5_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackGrupyModeld7c2e5_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackGrupyModeld7c2e5_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<Grupad7c2e5*> grupaitem() { return ({ QList<Grupad7c2e5*>* tmpP = static_cast<QList<Grupad7c2e5*>*>(callbackGrupyModeld7c2e5_Grupaitem(this)); QList<Grupad7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setGrupaitem(QList<Grupad7c2e5*> grupaitem) { callbackGrupyModeld7c2e5_SetGrupaitem(this, ({ QList<Grupad7c2e5*>* tmpValue = new QList<Grupad7c2e5*>(grupaitem); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_GrupaitemChanged(QList<Grupad7c2e5*> grupaitem) { callbackGrupyModeld7c2e5_GrupaitemChanged(this, ({ QList<Grupad7c2e5*>* tmpValue = new QList<Grupad7c2e5*>(grupaitem); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~GrupyModeld7c2e5() { callbackGrupyModeld7c2e5_DestroyGrupyModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackGrupyModeld7c2e5_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackGrupyModeld7c2e5_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackGrupyModeld7c2e5_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackGrupyModeld7c2e5_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackGrupyModeld7c2e5_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackGrupyModeld7c2e5_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackGrupyModeld7c2e5_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackGrupyModeld7c2e5_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackGrupyModeld7c2e5_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackGrupyModeld7c2e5_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackGrupyModeld7c2e5_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackGrupyModeld7c2e5_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackGrupyModeld7c2e5_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackGrupyModeld7c2e5_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackGrupyModeld7c2e5_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackGrupyModeld7c2e5_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackGrupyModeld7c2e5_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackGrupyModeld7c2e5_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackGrupyModeld7c2e5_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackGrupyModeld7c2e5_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackGrupyModeld7c2e5_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackGrupyModeld7c2e5_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackGrupyModeld7c2e5_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackGrupyModeld7c2e5_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackGrupyModeld7c2e5_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackGrupyModeld7c2e5_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackGrupyModeld7c2e5_ModelReset(this); };
	void resetInternalData() { callbackGrupyModeld7c2e5_ResetInternalData(this); };
	void revert() { callbackGrupyModeld7c2e5_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackGrupyModeld7c2e5_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackGrupyModeld7c2e5_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackGrupyModeld7c2e5_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackGrupyModeld7c2e5_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackGrupyModeld7c2e5_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackGrupyModeld7c2e5_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackGrupyModeld7c2e5_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackGrupyModeld7c2e5_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackGrupyModeld7c2e5_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackGrupyModeld7c2e5_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackGrupyModeld7c2e5_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackGrupyModeld7c2e5_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackGrupyModeld7c2e5_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackGrupyModeld7c2e5_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackGrupyModeld7c2e5_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackGrupyModeld7c2e5_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackGrupyModeld7c2e5_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackGrupyModeld7c2e5_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackGrupyModeld7c2e5_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackGrupyModeld7c2e5_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackGrupyModeld7c2e5_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackGrupyModeld7c2e5_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackGrupyModeld7c2e5_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackGrupyModeld7c2e5_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackGrupyModeld7c2e5_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackGrupyModeld7c2e5_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackGrupyModeld7c2e5_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackGrupyModeld7c2e5_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackGrupyModeld7c2e5_CustomEvent(this, event); };
	void deleteLater() { callbackGrupyModeld7c2e5_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackGrupyModeld7c2e5_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackGrupyModeld7c2e5_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackGrupyModeld7c2e5_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackGrupyModeld7c2e5_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<Grupad7c2e5*> grupaitemDefault() { return _grupaitem; };
	void setGrupaitemDefault(QList<Grupad7c2e5*> p) { if (p != _grupaitem) { _grupaitem = p; grupaitemChanged(_grupaitem); } };
signals:
	void rolesChanged(type378cdd roles);
	void grupaitemChanged(QList<Grupad7c2e5*> grupaitem);
public slots:
	void addGrupa(Grupad7c2e5* v0) { callbackGrupyModeld7c2e5_AddGrupa(this, v0); };
	void editGrupa(qint32 row, QString grupaName, qint32 grupaId) { QByteArray t335901 = grupaName.toUtf8(); Moc_PackedString grupaNamePacked = { const_cast<char*>(t335901.prepend("WHITESPACE").constData()+10), t335901.size()-10 };callbackGrupyModeld7c2e5_EditGrupa(this, row, grupaNamePacked, grupaId); };
	void removeGrupa(qint32 row) { callbackGrupyModeld7c2e5_RemoveGrupa(this, row); };
private:
	type378cdd _roles;
	QList<Grupad7c2e5*> _grupaitem;
};

Q_DECLARE_METATYPE(GrupyModeld7c2e5*)


void GrupyModeld7c2e5_GrupyModeld7c2e5_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
}

void GrupyModeld7c2e5_AddGrupa(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<GrupyModeld7c2e5*>(ptr), "addGrupa", Q_ARG(Grupad7c2e5*, static_cast<Grupad7c2e5*>(v0)));
}

void GrupyModeld7c2e5_EditGrupa(void* ptr, int row, struct Moc_PackedString grupaName, int grupaId)
{
	QMetaObject::invokeMethod(static_cast<GrupyModeld7c2e5*>(ptr), "editGrupa", Q_ARG(qint32, row), Q_ARG(QString, QString::fromUtf8(grupaName.data, grupaName.len)), Q_ARG(qint32, grupaId));
}

void GrupyModeld7c2e5_RemoveGrupa(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<GrupyModeld7c2e5*>(ptr), "removeGrupa", Q_ARG(qint32, row));
}

struct Moc_PackedList GrupyModeld7c2e5_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<GrupyModeld7c2e5*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList GrupyModeld7c2e5_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<GrupyModeld7c2e5*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void GrupyModeld7c2e5_SetRoles(void* ptr, void* roles)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void GrupyModeld7c2e5_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void GrupyModeld7c2e5_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<GrupyModeld7c2e5*>(ptr), static_cast<void (GrupyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&GrupyModeld7c2e5::rolesChanged), static_cast<GrupyModeld7c2e5*>(ptr), static_cast<void (GrupyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&GrupyModeld7c2e5::Signal_RolesChanged));
}

void GrupyModeld7c2e5_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<GrupyModeld7c2e5*>(ptr), static_cast<void (GrupyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&GrupyModeld7c2e5::rolesChanged), static_cast<GrupyModeld7c2e5*>(ptr), static_cast<void (GrupyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&GrupyModeld7c2e5::Signal_RolesChanged));
}

void GrupyModeld7c2e5_RolesChanged(void* ptr, void* roles)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList GrupyModeld7c2e5_Grupaitem(void* ptr)
{
	return ({ QList<Grupad7c2e5*>* tmpValue = new QList<Grupad7c2e5*>(static_cast<GrupyModeld7c2e5*>(ptr)->grupaitem()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList GrupyModeld7c2e5_GrupaitemDefault(void* ptr)
{
	return ({ QList<Grupad7c2e5*>* tmpValue = new QList<Grupad7c2e5*>(static_cast<GrupyModeld7c2e5*>(ptr)->grupaitemDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void GrupyModeld7c2e5_SetGrupaitem(void* ptr, void* grupaitem)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->setGrupaitem(({ QList<Grupad7c2e5*>* tmpP = static_cast<QList<Grupad7c2e5*>*>(grupaitem); QList<Grupad7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void GrupyModeld7c2e5_SetGrupaitemDefault(void* ptr, void* grupaitem)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->setGrupaitemDefault(({ QList<Grupad7c2e5*>* tmpP = static_cast<QList<Grupad7c2e5*>*>(grupaitem); QList<Grupad7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void GrupyModeld7c2e5_ConnectGrupaitemChanged(void* ptr)
{
	QObject::connect(static_cast<GrupyModeld7c2e5*>(ptr), static_cast<void (GrupyModeld7c2e5::*)(QList<Grupad7c2e5*>)>(&GrupyModeld7c2e5::grupaitemChanged), static_cast<GrupyModeld7c2e5*>(ptr), static_cast<void (GrupyModeld7c2e5::*)(QList<Grupad7c2e5*>)>(&GrupyModeld7c2e5::Signal_GrupaitemChanged));
}

void GrupyModeld7c2e5_DisconnectGrupaitemChanged(void* ptr)
{
	QObject::disconnect(static_cast<GrupyModeld7c2e5*>(ptr), static_cast<void (GrupyModeld7c2e5::*)(QList<Grupad7c2e5*>)>(&GrupyModeld7c2e5::grupaitemChanged), static_cast<GrupyModeld7c2e5*>(ptr), static_cast<void (GrupyModeld7c2e5::*)(QList<Grupad7c2e5*>)>(&GrupyModeld7c2e5::Signal_GrupaitemChanged));
}

void GrupyModeld7c2e5_GrupaitemChanged(void* ptr, void* grupaitem)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->grupaitemChanged(({ QList<Grupad7c2e5*>* tmpP = static_cast<QList<Grupad7c2e5*>*>(grupaitem); QList<Grupad7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int GrupyModeld7c2e5_GrupyModeld7c2e5_QRegisterMetaType()
{
	return qRegisterMetaType<GrupyModeld7c2e5*>();
}

int GrupyModeld7c2e5_GrupyModeld7c2e5_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<GrupyModeld7c2e5*>(const_cast<const char*>(typeName));
}

int GrupyModeld7c2e5_GrupyModeld7c2e5_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<GrupyModeld7c2e5>();
#else
	return 0;
#endif
}

int GrupyModeld7c2e5_GrupyModeld7c2e5_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<GrupyModeld7c2e5>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int GrupyModeld7c2e5_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* GrupyModeld7c2e5_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int GrupyModeld7c2e5_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* GrupyModeld7c2e5_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int GrupyModeld7c2e5_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* GrupyModeld7c2e5_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* GrupyModeld7c2e5___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* GrupyModeld7c2e5___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList GrupyModeld7c2e5___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* GrupyModeld7c2e5___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* GrupyModeld7c2e5___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* GrupyModeld7c2e5___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* GrupyModeld7c2e5___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int GrupyModeld7c2e5___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* GrupyModeld7c2e5___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* GrupyModeld7c2e5___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* GrupyModeld7c2e5___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* GrupyModeld7c2e5___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* GrupyModeld7c2e5___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* GrupyModeld7c2e5___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* GrupyModeld7c2e5___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList GrupyModeld7c2e5___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* GrupyModeld7c2e5___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* GrupyModeld7c2e5___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList GrupyModeld7c2e5___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* GrupyModeld7c2e5___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* GrupyModeld7c2e5___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* GrupyModeld7c2e5___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* GrupyModeld7c2e5___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* GrupyModeld7c2e5___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* GrupyModeld7c2e5___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int GrupyModeld7c2e5_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* GrupyModeld7c2e5_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int GrupyModeld7c2e5_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* GrupyModeld7c2e5_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* GrupyModeld7c2e5___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* GrupyModeld7c2e5___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* GrupyModeld7c2e5___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* GrupyModeld7c2e5___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* GrupyModeld7c2e5___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* GrupyModeld7c2e5___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* GrupyModeld7c2e5___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* GrupyModeld7c2e5___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* GrupyModeld7c2e5___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* GrupyModeld7c2e5___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* GrupyModeld7c2e5___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* GrupyModeld7c2e5___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList GrupyModeld7c2e5___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* GrupyModeld7c2e5___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* GrupyModeld7c2e5___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList GrupyModeld7c2e5___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* GrupyModeld7c2e5___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void GrupyModeld7c2e5___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* GrupyModeld7c2e5___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList GrupyModeld7c2e5___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* GrupyModeld7c2e5___grupaitem_atList(void* ptr, int i)
{
	return ({Grupad7c2e5* tmp = static_cast<QList<Grupad7c2e5*>*>(ptr)->at(i); if (i == static_cast<QList<Grupad7c2e5*>*>(ptr)->size()-1) { static_cast<QList<Grupad7c2e5*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5___grupaitem_setList(void* ptr, void* i)
{
	static_cast<QList<Grupad7c2e5*>*>(ptr)->append(static_cast<Grupad7c2e5*>(i));
}

void* GrupyModeld7c2e5___grupaitem_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Grupad7c2e5*>();
}

void* GrupyModeld7c2e5___setGrupaitem_grupaitem_atList(void* ptr, int i)
{
	return ({Grupad7c2e5* tmp = static_cast<QList<Grupad7c2e5*>*>(ptr)->at(i); if (i == static_cast<QList<Grupad7c2e5*>*>(ptr)->size()-1) { static_cast<QList<Grupad7c2e5*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5___setGrupaitem_grupaitem_setList(void* ptr, void* i)
{
	static_cast<QList<Grupad7c2e5*>*>(ptr)->append(static_cast<Grupad7c2e5*>(i));
}

void* GrupyModeld7c2e5___setGrupaitem_grupaitem_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Grupad7c2e5*>();
}

void* GrupyModeld7c2e5___grupaitemChanged_grupaitem_atList(void* ptr, int i)
{
	return ({Grupad7c2e5* tmp = static_cast<QList<Grupad7c2e5*>*>(ptr)->at(i); if (i == static_cast<QList<Grupad7c2e5*>*>(ptr)->size()-1) { static_cast<QList<Grupad7c2e5*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5___grupaitemChanged_grupaitem_setList(void* ptr, void* i)
{
	static_cast<QList<Grupad7c2e5*>*>(ptr)->append(static_cast<Grupad7c2e5*>(i));
}

void* GrupyModeld7c2e5___grupaitemChanged_grupaitem_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Grupad7c2e5*>();
}

int GrupyModeld7c2e5_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* GrupyModeld7c2e5_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int GrupyModeld7c2e5_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* GrupyModeld7c2e5_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int GrupyModeld7c2e5_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void GrupyModeld7c2e5_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* GrupyModeld7c2e5_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* GrupyModeld7c2e5_NewGrupyModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new GrupyModeld7c2e5(static_cast<QWindow*>(parent));
	} else {
		return new GrupyModeld7c2e5(static_cast<QObject*>(parent));
	}
}

void GrupyModeld7c2e5_DestroyGrupyModel(void* ptr)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->~GrupyModeld7c2e5();
}

void GrupyModeld7c2e5_DestroyGrupyModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char GrupyModeld7c2e5_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* GrupyModeld7c2e5_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* GrupyModeld7c2e5_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long GrupyModeld7c2e5_FlagsDefault(void* ptr, void* index)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}



char GrupyModeld7c2e5_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char GrupyModeld7c2e5_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char GrupyModeld7c2e5_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char GrupyModeld7c2e5_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char GrupyModeld7c2e5_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char GrupyModeld7c2e5_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char GrupyModeld7c2e5_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char GrupyModeld7c2e5_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char GrupyModeld7c2e5_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char GrupyModeld7c2e5_SubmitDefault(void* ptr)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::submit();
}

void GrupyModeld7c2e5_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void GrupyModeld7c2e5_ResetInternalDataDefault(void* ptr)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::resetInternalData();
}

void GrupyModeld7c2e5_RevertDefault(void* ptr)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::revert();
}

void GrupyModeld7c2e5_SortDefault(void* ptr, int column, long long order)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList GrupyModeld7c2e5_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList GrupyModeld7c2e5_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* GrupyModeld7c2e5_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* GrupyModeld7c2e5_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* GrupyModeld7c2e5_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList GrupyModeld7c2e5_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* GrupyModeld7c2e5_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString GrupyModeld7c2e5_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray td21773 = static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(td21773.prepend("WHITESPACE").constData()+10), td21773.size()-10 }; });
}

void* GrupyModeld7c2e5_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* GrupyModeld7c2e5_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long GrupyModeld7c2e5_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long GrupyModeld7c2e5_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::supportedDropActions();
}

char GrupyModeld7c2e5_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char GrupyModeld7c2e5_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char GrupyModeld7c2e5_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int GrupyModeld7c2e5_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int GrupyModeld7c2e5_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char GrupyModeld7c2e5_EventDefault(void* ptr, void* e)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char GrupyModeld7c2e5_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void GrupyModeld7c2e5_ChildEventDefault(void* ptr, void* event)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void GrupyModeld7c2e5_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void GrupyModeld7c2e5_CustomEventDefault(void* ptr, void* event)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void GrupyModeld7c2e5_DeleteLaterDefault(void* ptr)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::deleteLater();
}

void GrupyModeld7c2e5_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void GrupyModeld7c2e5_TimerEventDefault(void* ptr, void* event)
{
	static_cast<GrupyModeld7c2e5*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString Produktd7c2e5_ProduktName(void* ptr)
{
	return ({ QByteArray t0dae3e = static_cast<Produktd7c2e5*>(ptr)->produktName().toUtf8(); Moc_PackedString { const_cast<char*>(t0dae3e.prepend("WHITESPACE").constData()+10), t0dae3e.size()-10 }; });
}

struct Moc_PackedString Produktd7c2e5_ProduktNameDefault(void* ptr)
{
	return ({ QByteArray t0ee7fb = static_cast<Produktd7c2e5*>(ptr)->produktNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t0ee7fb.prepend("WHITESPACE").constData()+10), t0ee7fb.size()-10 }; });
}

void Produktd7c2e5_SetProduktName(void* ptr, struct Moc_PackedString produktName)
{
	static_cast<Produktd7c2e5*>(ptr)->setProduktName(QString::fromUtf8(produktName.data, produktName.len));
}

void Produktd7c2e5_SetProduktNameDefault(void* ptr, struct Moc_PackedString produktName)
{
	static_cast<Produktd7c2e5*>(ptr)->setProduktNameDefault(QString::fromUtf8(produktName.data, produktName.len));
}

void Produktd7c2e5_ConnectProduktNameChanged(void* ptr)
{
	QObject::connect(static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(QString)>(&Produktd7c2e5::produktNameChanged), static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(QString)>(&Produktd7c2e5::Signal_ProduktNameChanged));
}

void Produktd7c2e5_DisconnectProduktNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(QString)>(&Produktd7c2e5::produktNameChanged), static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(QString)>(&Produktd7c2e5::Signal_ProduktNameChanged));
}

void Produktd7c2e5_ProduktNameChanged(void* ptr, struct Moc_PackedString produktName)
{
	static_cast<Produktd7c2e5*>(ptr)->produktNameChanged(QString::fromUtf8(produktName.data, produktName.len));
}

float Produktd7c2e5_ProduktW(void* ptr)
{
	return static_cast<Produktd7c2e5*>(ptr)->produktW();
}

float Produktd7c2e5_ProduktWDefault(void* ptr)
{
	return static_cast<Produktd7c2e5*>(ptr)->produktWDefault();
}

void Produktd7c2e5_SetProduktW(void* ptr, float produktW)
{
	static_cast<Produktd7c2e5*>(ptr)->setProduktW(produktW);
}

void Produktd7c2e5_SetProduktWDefault(void* ptr, float produktW)
{
	static_cast<Produktd7c2e5*>(ptr)->setProduktWDefault(produktW);
}

void Produktd7c2e5_ConnectProduktWChanged(void* ptr)
{
	QObject::connect(static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(float)>(&Produktd7c2e5::produktWChanged), static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(float)>(&Produktd7c2e5::Signal_ProduktWChanged));
}

void Produktd7c2e5_DisconnectProduktWChanged(void* ptr)
{
	QObject::disconnect(static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(float)>(&Produktd7c2e5::produktWChanged), static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(float)>(&Produktd7c2e5::Signal_ProduktWChanged));
}

void Produktd7c2e5_ProduktWChanged(void* ptr, float produktW)
{
	static_cast<Produktd7c2e5*>(ptr)->produktWChanged(produktW);
}

int Produktd7c2e5_ProduktId(void* ptr)
{
	return static_cast<Produktd7c2e5*>(ptr)->produktId();
}

int Produktd7c2e5_ProduktIdDefault(void* ptr)
{
	return static_cast<Produktd7c2e5*>(ptr)->produktIdDefault();
}

void Produktd7c2e5_SetProduktId(void* ptr, int produktId)
{
	static_cast<Produktd7c2e5*>(ptr)->setProduktId(produktId);
}

void Produktd7c2e5_SetProduktIdDefault(void* ptr, int produktId)
{
	static_cast<Produktd7c2e5*>(ptr)->setProduktIdDefault(produktId);
}

void Produktd7c2e5_ConnectProduktIdChanged(void* ptr)
{
	QObject::connect(static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(qint32)>(&Produktd7c2e5::produktIdChanged), static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(qint32)>(&Produktd7c2e5::Signal_ProduktIdChanged));
}

void Produktd7c2e5_DisconnectProduktIdChanged(void* ptr)
{
	QObject::disconnect(static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(qint32)>(&Produktd7c2e5::produktIdChanged), static_cast<Produktd7c2e5*>(ptr), static_cast<void (Produktd7c2e5::*)(qint32)>(&Produktd7c2e5::Signal_ProduktIdChanged));
}

void Produktd7c2e5_ProduktIdChanged(void* ptr, int produktId)
{
	static_cast<Produktd7c2e5*>(ptr)->produktIdChanged(produktId);
}

int Produktd7c2e5_Produktd7c2e5_QRegisterMetaType()
{
	return qRegisterMetaType<Produktd7c2e5*>();
}

int Produktd7c2e5_Produktd7c2e5_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Produktd7c2e5*>(const_cast<const char*>(typeName));
}

int Produktd7c2e5_Produktd7c2e5_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Produktd7c2e5>();
#else
	return 0;
#endif
}

int Produktd7c2e5_Produktd7c2e5_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Produktd7c2e5>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Produktd7c2e5___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Produktd7c2e5___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Produktd7c2e5___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Produktd7c2e5___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Produktd7c2e5___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Produktd7c2e5___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Produktd7c2e5___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Produktd7c2e5___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Produktd7c2e5___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Produktd7c2e5___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Produktd7c2e5___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Produktd7c2e5___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Produktd7c2e5___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Produktd7c2e5___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Produktd7c2e5___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Produktd7c2e5_NewProdukt(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Produktd7c2e5(static_cast<QWindow*>(parent));
	} else {
		return new Produktd7c2e5(static_cast<QObject*>(parent));
	}
}

void Produktd7c2e5_DestroyProdukt(void* ptr)
{
	static_cast<Produktd7c2e5*>(ptr)->~Produktd7c2e5();
}

void Produktd7c2e5_DestroyProduktDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Produktd7c2e5_EventDefault(void* ptr, void* e)
{
	return static_cast<Produktd7c2e5*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Produktd7c2e5_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Produktd7c2e5*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Produktd7c2e5_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Produktd7c2e5*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Produktd7c2e5_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Produktd7c2e5*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Produktd7c2e5_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Produktd7c2e5*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Produktd7c2e5_DeleteLaterDefault(void* ptr)
{
	static_cast<Produktd7c2e5*>(ptr)->QObject::deleteLater();
}

void Produktd7c2e5_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Produktd7c2e5*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Produktd7c2e5_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Produktd7c2e5*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void ProduktyModeld7c2e5_AddProdukt(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<ProduktyModeld7c2e5*>(ptr), "addProdukt", Q_ARG(Produktd7c2e5*, static_cast<Produktd7c2e5*>(v0)));
}

void ProduktyModeld7c2e5_EditProdukt(void* ptr, int row, struct Moc_PackedString produktName, float produktW, int produktId)
{
	QMetaObject::invokeMethod(static_cast<ProduktyModeld7c2e5*>(ptr), "editProdukt", Q_ARG(qint32, row), Q_ARG(QString, QString::fromUtf8(produktName.data, produktName.len)), Q_ARG(float, produktW), Q_ARG(qint32, produktId));
}

void ProduktyModeld7c2e5_RemoveProdukt(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<ProduktyModeld7c2e5*>(ptr), "removeProdukt", Q_ARG(qint32, row));
}

struct Moc_PackedList ProduktyModeld7c2e5_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<ProduktyModeld7c2e5*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList ProduktyModeld7c2e5_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<ProduktyModeld7c2e5*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void ProduktyModeld7c2e5_SetRoles(void* ptr, void* roles)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void ProduktyModeld7c2e5_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void ProduktyModeld7c2e5_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<ProduktyModeld7c2e5*>(ptr), static_cast<void (ProduktyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&ProduktyModeld7c2e5::rolesChanged), static_cast<ProduktyModeld7c2e5*>(ptr), static_cast<void (ProduktyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&ProduktyModeld7c2e5::Signal_RolesChanged));
}

void ProduktyModeld7c2e5_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<ProduktyModeld7c2e5*>(ptr), static_cast<void (ProduktyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&ProduktyModeld7c2e5::rolesChanged), static_cast<ProduktyModeld7c2e5*>(ptr), static_cast<void (ProduktyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&ProduktyModeld7c2e5::Signal_RolesChanged));
}

void ProduktyModeld7c2e5_RolesChanged(void* ptr, void* roles)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList ProduktyModeld7c2e5_Produktitem(void* ptr)
{
	return ({ QList<Produktd7c2e5*>* tmpValue = new QList<Produktd7c2e5*>(static_cast<ProduktyModeld7c2e5*>(ptr)->produktitem()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList ProduktyModeld7c2e5_ProduktitemDefault(void* ptr)
{
	return ({ QList<Produktd7c2e5*>* tmpValue = new QList<Produktd7c2e5*>(static_cast<ProduktyModeld7c2e5*>(ptr)->produktitemDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void ProduktyModeld7c2e5_SetProduktitem(void* ptr, void* produktitem)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->setProduktitem(({ QList<Produktd7c2e5*>* tmpP = static_cast<QList<Produktd7c2e5*>*>(produktitem); QList<Produktd7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void ProduktyModeld7c2e5_SetProduktitemDefault(void* ptr, void* produktitem)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->setProduktitemDefault(({ QList<Produktd7c2e5*>* tmpP = static_cast<QList<Produktd7c2e5*>*>(produktitem); QList<Produktd7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void ProduktyModeld7c2e5_ConnectProduktitemChanged(void* ptr)
{
	QObject::connect(static_cast<ProduktyModeld7c2e5*>(ptr), static_cast<void (ProduktyModeld7c2e5::*)(QList<Produktd7c2e5*>)>(&ProduktyModeld7c2e5::produktitemChanged), static_cast<ProduktyModeld7c2e5*>(ptr), static_cast<void (ProduktyModeld7c2e5::*)(QList<Produktd7c2e5*>)>(&ProduktyModeld7c2e5::Signal_ProduktitemChanged));
}

void ProduktyModeld7c2e5_DisconnectProduktitemChanged(void* ptr)
{
	QObject::disconnect(static_cast<ProduktyModeld7c2e5*>(ptr), static_cast<void (ProduktyModeld7c2e5::*)(QList<Produktd7c2e5*>)>(&ProduktyModeld7c2e5::produktitemChanged), static_cast<ProduktyModeld7c2e5*>(ptr), static_cast<void (ProduktyModeld7c2e5::*)(QList<Produktd7c2e5*>)>(&ProduktyModeld7c2e5::Signal_ProduktitemChanged));
}

void ProduktyModeld7c2e5_ProduktitemChanged(void* ptr, void* produktitem)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->produktitemChanged(({ QList<Produktd7c2e5*>* tmpP = static_cast<QList<Produktd7c2e5*>*>(produktitem); QList<Produktd7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int ProduktyModeld7c2e5_ProduktyModeld7c2e5_QRegisterMetaType()
{
	return qRegisterMetaType<ProduktyModeld7c2e5*>();
}

int ProduktyModeld7c2e5_ProduktyModeld7c2e5_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ProduktyModeld7c2e5*>(const_cast<const char*>(typeName));
}

int ProduktyModeld7c2e5_ProduktyModeld7c2e5_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ProduktyModeld7c2e5>();
#else
	return 0;
#endif
}

int ProduktyModeld7c2e5_ProduktyModeld7c2e5_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ProduktyModeld7c2e5>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int ProduktyModeld7c2e5_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* ProduktyModeld7c2e5_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int ProduktyModeld7c2e5_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* ProduktyModeld7c2e5_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int ProduktyModeld7c2e5_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* ProduktyModeld7c2e5_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* ProduktyModeld7c2e5___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* ProduktyModeld7c2e5___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList ProduktyModeld7c2e5___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ProduktyModeld7c2e5___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* ProduktyModeld7c2e5___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* ProduktyModeld7c2e5___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* ProduktyModeld7c2e5___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int ProduktyModeld7c2e5___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* ProduktyModeld7c2e5___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* ProduktyModeld7c2e5___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* ProduktyModeld7c2e5___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* ProduktyModeld7c2e5___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* ProduktyModeld7c2e5___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* ProduktyModeld7c2e5___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* ProduktyModeld7c2e5___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList ProduktyModeld7c2e5___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ProduktyModeld7c2e5___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* ProduktyModeld7c2e5___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList ProduktyModeld7c2e5___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ProduktyModeld7c2e5___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* ProduktyModeld7c2e5___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* ProduktyModeld7c2e5___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* ProduktyModeld7c2e5___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* ProduktyModeld7c2e5___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* ProduktyModeld7c2e5___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int ProduktyModeld7c2e5_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* ProduktyModeld7c2e5_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int ProduktyModeld7c2e5_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* ProduktyModeld7c2e5_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* ProduktyModeld7c2e5___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ProduktyModeld7c2e5___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ProduktyModeld7c2e5___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ProduktyModeld7c2e5___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ProduktyModeld7c2e5___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ProduktyModeld7c2e5___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ProduktyModeld7c2e5___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ProduktyModeld7c2e5___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ProduktyModeld7c2e5___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ProduktyModeld7c2e5___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ProduktyModeld7c2e5___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* ProduktyModeld7c2e5___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList ProduktyModeld7c2e5___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ProduktyModeld7c2e5___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* ProduktyModeld7c2e5___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList ProduktyModeld7c2e5___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ProduktyModeld7c2e5___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void ProduktyModeld7c2e5___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* ProduktyModeld7c2e5___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList ProduktyModeld7c2e5___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ProduktyModeld7c2e5___produktitem_atList(void* ptr, int i)
{
	return ({Produktd7c2e5* tmp = static_cast<QList<Produktd7c2e5*>*>(ptr)->at(i); if (i == static_cast<QList<Produktd7c2e5*>*>(ptr)->size()-1) { static_cast<QList<Produktd7c2e5*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5___produktitem_setList(void* ptr, void* i)
{
	static_cast<QList<Produktd7c2e5*>*>(ptr)->append(static_cast<Produktd7c2e5*>(i));
}

void* ProduktyModeld7c2e5___produktitem_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Produktd7c2e5*>();
}

void* ProduktyModeld7c2e5___setProduktitem_produktitem_atList(void* ptr, int i)
{
	return ({Produktd7c2e5* tmp = static_cast<QList<Produktd7c2e5*>*>(ptr)->at(i); if (i == static_cast<QList<Produktd7c2e5*>*>(ptr)->size()-1) { static_cast<QList<Produktd7c2e5*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5___setProduktitem_produktitem_setList(void* ptr, void* i)
{
	static_cast<QList<Produktd7c2e5*>*>(ptr)->append(static_cast<Produktd7c2e5*>(i));
}

void* ProduktyModeld7c2e5___setProduktitem_produktitem_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Produktd7c2e5*>();
}

void* ProduktyModeld7c2e5___produktitemChanged_produktitem_atList(void* ptr, int i)
{
	return ({Produktd7c2e5* tmp = static_cast<QList<Produktd7c2e5*>*>(ptr)->at(i); if (i == static_cast<QList<Produktd7c2e5*>*>(ptr)->size()-1) { static_cast<QList<Produktd7c2e5*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5___produktitemChanged_produktitem_setList(void* ptr, void* i)
{
	static_cast<QList<Produktd7c2e5*>*>(ptr)->append(static_cast<Produktd7c2e5*>(i));
}

void* ProduktyModeld7c2e5___produktitemChanged_produktitem_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Produktd7c2e5*>();
}

int ProduktyModeld7c2e5_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* ProduktyModeld7c2e5_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int ProduktyModeld7c2e5_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* ProduktyModeld7c2e5_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int ProduktyModeld7c2e5_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ProduktyModeld7c2e5_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* ProduktyModeld7c2e5_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* ProduktyModeld7c2e5_NewProduktyModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new ProduktyModeld7c2e5(static_cast<QWindow*>(parent));
	} else {
		return new ProduktyModeld7c2e5(static_cast<QObject*>(parent));
	}
}

void ProduktyModeld7c2e5_DestroyProduktyModel(void* ptr)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->~ProduktyModeld7c2e5();
}

void ProduktyModeld7c2e5_DestroyProduktyModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char ProduktyModeld7c2e5_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* ProduktyModeld7c2e5_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* ProduktyModeld7c2e5_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long ProduktyModeld7c2e5_FlagsDefault(void* ptr, void* index)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}



char ProduktyModeld7c2e5_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char ProduktyModeld7c2e5_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char ProduktyModeld7c2e5_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char ProduktyModeld7c2e5_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char ProduktyModeld7c2e5_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char ProduktyModeld7c2e5_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char ProduktyModeld7c2e5_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char ProduktyModeld7c2e5_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char ProduktyModeld7c2e5_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char ProduktyModeld7c2e5_SubmitDefault(void* ptr)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::submit();
}

void ProduktyModeld7c2e5_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void ProduktyModeld7c2e5_ResetInternalDataDefault(void* ptr)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::resetInternalData();
}

void ProduktyModeld7c2e5_RevertDefault(void* ptr)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::revert();
}

void ProduktyModeld7c2e5_SortDefault(void* ptr, int column, long long order)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList ProduktyModeld7c2e5_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList ProduktyModeld7c2e5_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ProduktyModeld7c2e5_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* ProduktyModeld7c2e5_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* ProduktyModeld7c2e5_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList ProduktyModeld7c2e5_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ProduktyModeld7c2e5_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString ProduktyModeld7c2e5_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray tfd3acd = static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(tfd3acd.prepend("WHITESPACE").constData()+10), tfd3acd.size()-10 }; });
}

void* ProduktyModeld7c2e5_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* ProduktyModeld7c2e5_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long ProduktyModeld7c2e5_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long ProduktyModeld7c2e5_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::supportedDropActions();
}

char ProduktyModeld7c2e5_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char ProduktyModeld7c2e5_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char ProduktyModeld7c2e5_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int ProduktyModeld7c2e5_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int ProduktyModeld7c2e5_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char ProduktyModeld7c2e5_EventDefault(void* ptr, void* e)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char ProduktyModeld7c2e5_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void ProduktyModeld7c2e5_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void ProduktyModeld7c2e5_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ProduktyModeld7c2e5_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void ProduktyModeld7c2e5_DeleteLaterDefault(void* ptr)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::deleteLater();
}

void ProduktyModeld7c2e5_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void ProduktyModeld7c2e5_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ProduktyModeld7c2e5*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString Elementd7c2e5_ElementGName(void* ptr)
{
	return ({ QByteArray t5f47b4 = static_cast<Elementd7c2e5*>(ptr)->elementGName().toUtf8(); Moc_PackedString { const_cast<char*>(t5f47b4.prepend("WHITESPACE").constData()+10), t5f47b4.size()-10 }; });
}

struct Moc_PackedString Elementd7c2e5_ElementGNameDefault(void* ptr)
{
	return ({ QByteArray t597a32 = static_cast<Elementd7c2e5*>(ptr)->elementGNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t597a32.prepend("WHITESPACE").constData()+10), t597a32.size()-10 }; });
}

void Elementd7c2e5_SetElementGName(void* ptr, struct Moc_PackedString elementGName)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementGName(QString::fromUtf8(elementGName.data, elementGName.len));
}

void Elementd7c2e5_SetElementGNameDefault(void* ptr, struct Moc_PackedString elementGName)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementGNameDefault(QString::fromUtf8(elementGName.data, elementGName.len));
}

void Elementd7c2e5_ConnectElementGNameChanged(void* ptr)
{
	QObject::connect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::elementGNameChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::Signal_ElementGNameChanged));
}

void Elementd7c2e5_DisconnectElementGNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::elementGNameChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::Signal_ElementGNameChanged));
}

void Elementd7c2e5_ElementGNameChanged(void* ptr, struct Moc_PackedString elementGName)
{
	static_cast<Elementd7c2e5*>(ptr)->elementGNameChanged(QString::fromUtf8(elementGName.data, elementGName.len));
}

float Elementd7c2e5_ElementStan(void* ptr)
{
	return static_cast<Elementd7c2e5*>(ptr)->elementStan();
}

float Elementd7c2e5_ElementStanDefault(void* ptr)
{
	return static_cast<Elementd7c2e5*>(ptr)->elementStanDefault();
}

void Elementd7c2e5_SetElementStan(void* ptr, float elementStan)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementStan(elementStan);
}

void Elementd7c2e5_SetElementStanDefault(void* ptr, float elementStan)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementStanDefault(elementStan);
}

void Elementd7c2e5_ConnectElementStanChanged(void* ptr)
{
	QObject::connect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(float)>(&Elementd7c2e5::elementStanChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(float)>(&Elementd7c2e5::Signal_ElementStanChanged));
}

void Elementd7c2e5_DisconnectElementStanChanged(void* ptr)
{
	QObject::disconnect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(float)>(&Elementd7c2e5::elementStanChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(float)>(&Elementd7c2e5::Signal_ElementStanChanged));
}

void Elementd7c2e5_ElementStanChanged(void* ptr, float elementStan)
{
	static_cast<Elementd7c2e5*>(ptr)->elementStanChanged(elementStan);
}

struct Moc_PackedString Elementd7c2e5_ElementPName(void* ptr)
{
	return ({ QByteArray t124c12 = static_cast<Elementd7c2e5*>(ptr)->elementPName().toUtf8(); Moc_PackedString { const_cast<char*>(t124c12.prepend("WHITESPACE").constData()+10), t124c12.size()-10 }; });
}

struct Moc_PackedString Elementd7c2e5_ElementPNameDefault(void* ptr)
{
	return ({ QByteArray tdbf014 = static_cast<Elementd7c2e5*>(ptr)->elementPNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tdbf014.prepend("WHITESPACE").constData()+10), tdbf014.size()-10 }; });
}

void Elementd7c2e5_SetElementPName(void* ptr, struct Moc_PackedString elementPName)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementPName(QString::fromUtf8(elementPName.data, elementPName.len));
}

void Elementd7c2e5_SetElementPNameDefault(void* ptr, struct Moc_PackedString elementPName)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementPNameDefault(QString::fromUtf8(elementPName.data, elementPName.len));
}

void Elementd7c2e5_ConnectElementPNameChanged(void* ptr)
{
	QObject::connect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::elementPNameChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::Signal_ElementPNameChanged));
}

void Elementd7c2e5_DisconnectElementPNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::elementPNameChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::Signal_ElementPNameChanged));
}

void Elementd7c2e5_ElementPNameChanged(void* ptr, struct Moc_PackedString elementPName)
{
	static_cast<Elementd7c2e5*>(ptr)->elementPNameChanged(QString::fromUtf8(elementPName.data, elementPName.len));
}

int Elementd7c2e5_ElementIlosc(void* ptr)
{
	return static_cast<Elementd7c2e5*>(ptr)->elementIlosc();
}

int Elementd7c2e5_ElementIloscDefault(void* ptr)
{
	return static_cast<Elementd7c2e5*>(ptr)->elementIloscDefault();
}

void Elementd7c2e5_SetElementIlosc(void* ptr, int elementIlosc)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementIlosc(elementIlosc);
}

void Elementd7c2e5_SetElementIloscDefault(void* ptr, int elementIlosc)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementIloscDefault(elementIlosc);
}

void Elementd7c2e5_ConnectElementIloscChanged(void* ptr)
{
	QObject::connect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(qint32)>(&Elementd7c2e5::elementIloscChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(qint32)>(&Elementd7c2e5::Signal_ElementIloscChanged));
}

void Elementd7c2e5_DisconnectElementIloscChanged(void* ptr)
{
	QObject::disconnect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(qint32)>(&Elementd7c2e5::elementIloscChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(qint32)>(&Elementd7c2e5::Signal_ElementIloscChanged));
}

void Elementd7c2e5_ElementIloscChanged(void* ptr, int elementIlosc)
{
	static_cast<Elementd7c2e5*>(ptr)->elementIloscChanged(elementIlosc);
}

struct Moc_PackedString Elementd7c2e5_ElementData(void* ptr)
{
	return ({ QByteArray t38559e = static_cast<Elementd7c2e5*>(ptr)->elementData().toUtf8(); Moc_PackedString { const_cast<char*>(t38559e.prepend("WHITESPACE").constData()+10), t38559e.size()-10 }; });
}

struct Moc_PackedString Elementd7c2e5_ElementDataDefault(void* ptr)
{
	return ({ QByteArray t614780 = static_cast<Elementd7c2e5*>(ptr)->elementDataDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t614780.prepend("WHITESPACE").constData()+10), t614780.size()-10 }; });
}

void Elementd7c2e5_SetElementData(void* ptr, struct Moc_PackedString elementData)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementData(QString::fromUtf8(elementData.data, elementData.len));
}

void Elementd7c2e5_SetElementDataDefault(void* ptr, struct Moc_PackedString elementData)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementDataDefault(QString::fromUtf8(elementData.data, elementData.len));
}

void Elementd7c2e5_ConnectElementDataChanged(void* ptr)
{
	QObject::connect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::elementDataChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::Signal_ElementDataChanged));
}

void Elementd7c2e5_DisconnectElementDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::elementDataChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(QString)>(&Elementd7c2e5::Signal_ElementDataChanged));
}

void Elementd7c2e5_ElementDataChanged(void* ptr, struct Moc_PackedString elementData)
{
	static_cast<Elementd7c2e5*>(ptr)->elementDataChanged(QString::fromUtf8(elementData.data, elementData.len));
}

int Elementd7c2e5_ElementId(void* ptr)
{
	return static_cast<Elementd7c2e5*>(ptr)->elementId();
}

int Elementd7c2e5_ElementIdDefault(void* ptr)
{
	return static_cast<Elementd7c2e5*>(ptr)->elementIdDefault();
}

void Elementd7c2e5_SetElementId(void* ptr, int elementId)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementId(elementId);
}

void Elementd7c2e5_SetElementIdDefault(void* ptr, int elementId)
{
	static_cast<Elementd7c2e5*>(ptr)->setElementIdDefault(elementId);
}

void Elementd7c2e5_ConnectElementIdChanged(void* ptr)
{
	QObject::connect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(qint32)>(&Elementd7c2e5::elementIdChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(qint32)>(&Elementd7c2e5::Signal_ElementIdChanged));
}

void Elementd7c2e5_DisconnectElementIdChanged(void* ptr)
{
	QObject::disconnect(static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(qint32)>(&Elementd7c2e5::elementIdChanged), static_cast<Elementd7c2e5*>(ptr), static_cast<void (Elementd7c2e5::*)(qint32)>(&Elementd7c2e5::Signal_ElementIdChanged));
}

void Elementd7c2e5_ElementIdChanged(void* ptr, int elementId)
{
	static_cast<Elementd7c2e5*>(ptr)->elementIdChanged(elementId);
}

int Elementd7c2e5_Elementd7c2e5_QRegisterMetaType()
{
	return qRegisterMetaType<Elementd7c2e5*>();
}

int Elementd7c2e5_Elementd7c2e5_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Elementd7c2e5*>(const_cast<const char*>(typeName));
}

int Elementd7c2e5_Elementd7c2e5_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Elementd7c2e5>();
#else
	return 0;
#endif
}

int Elementd7c2e5_Elementd7c2e5_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Elementd7c2e5>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Elementd7c2e5___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Elementd7c2e5___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Elementd7c2e5___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Elementd7c2e5___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Elementd7c2e5___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Elementd7c2e5___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Elementd7c2e5___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Elementd7c2e5___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Elementd7c2e5___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Elementd7c2e5___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Elementd7c2e5___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Elementd7c2e5___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Elementd7c2e5___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Elementd7c2e5___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Elementd7c2e5___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Elementd7c2e5_NewElement(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Elementd7c2e5(static_cast<QWindow*>(parent));
	} else {
		return new Elementd7c2e5(static_cast<QObject*>(parent));
	}
}

void Elementd7c2e5_DestroyElement(void* ptr)
{
	static_cast<Elementd7c2e5*>(ptr)->~Elementd7c2e5();
}

void Elementd7c2e5_DestroyElementDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Elementd7c2e5_EventDefault(void* ptr, void* e)
{
	return static_cast<Elementd7c2e5*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Elementd7c2e5_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Elementd7c2e5*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Elementd7c2e5_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Elementd7c2e5*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Elementd7c2e5_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Elementd7c2e5*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Elementd7c2e5_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Elementd7c2e5*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Elementd7c2e5_DeleteLaterDefault(void* ptr)
{
	static_cast<Elementd7c2e5*>(ptr)->QObject::deleteLater();
}

void Elementd7c2e5_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Elementd7c2e5*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Elementd7c2e5_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Elementd7c2e5*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void ElementyModeld7c2e5_AddElement(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<ElementyModeld7c2e5*>(ptr), "addElement", Q_ARG(Elementd7c2e5*, static_cast<Elementd7c2e5*>(v0)));
}

void ElementyModeld7c2e5_EditElement(void* ptr, int row, struct Moc_PackedString elementGName, float elementStan, struct Moc_PackedString elementPName, int elementIlosc, struct Moc_PackedString elementData, int elementId)
{
	QMetaObject::invokeMethod(static_cast<ElementyModeld7c2e5*>(ptr), "editElement", Q_ARG(qint32, row), Q_ARG(QString, QString::fromUtf8(elementGName.data, elementGName.len)), Q_ARG(float, elementStan), Q_ARG(QString, QString::fromUtf8(elementPName.data, elementPName.len)), Q_ARG(qint32, elementIlosc), Q_ARG(QString, QString::fromUtf8(elementData.data, elementData.len)), Q_ARG(qint32, elementId));
}

void ElementyModeld7c2e5_RemoveElement(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<ElementyModeld7c2e5*>(ptr), "removeElement", Q_ARG(qint32, row));
}

void ElementyModeld7c2e5_RemoveElementId(void* ptr, int idin)
{
	QMetaObject::invokeMethod(static_cast<ElementyModeld7c2e5*>(ptr), "removeElementId", Q_ARG(qint32, idin));
}

struct Moc_PackedList ElementyModeld7c2e5_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<ElementyModeld7c2e5*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList ElementyModeld7c2e5_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<ElementyModeld7c2e5*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void ElementyModeld7c2e5_SetRoles(void* ptr, void* roles)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void ElementyModeld7c2e5_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void ElementyModeld7c2e5_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<ElementyModeld7c2e5*>(ptr), static_cast<void (ElementyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&ElementyModeld7c2e5::rolesChanged), static_cast<ElementyModeld7c2e5*>(ptr), static_cast<void (ElementyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&ElementyModeld7c2e5::Signal_RolesChanged));
}

void ElementyModeld7c2e5_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<ElementyModeld7c2e5*>(ptr), static_cast<void (ElementyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&ElementyModeld7c2e5::rolesChanged), static_cast<ElementyModeld7c2e5*>(ptr), static_cast<void (ElementyModeld7c2e5::*)(QMap<qint32, QByteArray>)>(&ElementyModeld7c2e5::Signal_RolesChanged));
}

void ElementyModeld7c2e5_RolesChanged(void* ptr, void* roles)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList ElementyModeld7c2e5_Elementitem(void* ptr)
{
	return ({ QList<Elementd7c2e5*>* tmpValue = new QList<Elementd7c2e5*>(static_cast<ElementyModeld7c2e5*>(ptr)->elementitem()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList ElementyModeld7c2e5_ElementitemDefault(void* ptr)
{
	return ({ QList<Elementd7c2e5*>* tmpValue = new QList<Elementd7c2e5*>(static_cast<ElementyModeld7c2e5*>(ptr)->elementitemDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void ElementyModeld7c2e5_SetElementitem(void* ptr, void* elementitem)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->setElementitem(({ QList<Elementd7c2e5*>* tmpP = static_cast<QList<Elementd7c2e5*>*>(elementitem); QList<Elementd7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void ElementyModeld7c2e5_SetElementitemDefault(void* ptr, void* elementitem)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->setElementitemDefault(({ QList<Elementd7c2e5*>* tmpP = static_cast<QList<Elementd7c2e5*>*>(elementitem); QList<Elementd7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void ElementyModeld7c2e5_ConnectElementitemChanged(void* ptr)
{
	QObject::connect(static_cast<ElementyModeld7c2e5*>(ptr), static_cast<void (ElementyModeld7c2e5::*)(QList<Elementd7c2e5*>)>(&ElementyModeld7c2e5::elementitemChanged), static_cast<ElementyModeld7c2e5*>(ptr), static_cast<void (ElementyModeld7c2e5::*)(QList<Elementd7c2e5*>)>(&ElementyModeld7c2e5::Signal_ElementitemChanged));
}

void ElementyModeld7c2e5_DisconnectElementitemChanged(void* ptr)
{
	QObject::disconnect(static_cast<ElementyModeld7c2e5*>(ptr), static_cast<void (ElementyModeld7c2e5::*)(QList<Elementd7c2e5*>)>(&ElementyModeld7c2e5::elementitemChanged), static_cast<ElementyModeld7c2e5*>(ptr), static_cast<void (ElementyModeld7c2e5::*)(QList<Elementd7c2e5*>)>(&ElementyModeld7c2e5::Signal_ElementitemChanged));
}

void ElementyModeld7c2e5_ElementitemChanged(void* ptr, void* elementitem)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->elementitemChanged(({ QList<Elementd7c2e5*>* tmpP = static_cast<QList<Elementd7c2e5*>*>(elementitem); QList<Elementd7c2e5*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int ElementyModeld7c2e5_ElementyModeld7c2e5_QRegisterMetaType()
{
	return qRegisterMetaType<ElementyModeld7c2e5*>();
}

int ElementyModeld7c2e5_ElementyModeld7c2e5_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ElementyModeld7c2e5*>(const_cast<const char*>(typeName));
}

int ElementyModeld7c2e5_ElementyModeld7c2e5_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ElementyModeld7c2e5>();
#else
	return 0;
#endif
}

int ElementyModeld7c2e5_ElementyModeld7c2e5_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ElementyModeld7c2e5>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int ElementyModeld7c2e5_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* ElementyModeld7c2e5_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int ElementyModeld7c2e5_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* ElementyModeld7c2e5_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int ElementyModeld7c2e5_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* ElementyModeld7c2e5_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* ElementyModeld7c2e5___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* ElementyModeld7c2e5___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList ElementyModeld7c2e5___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ElementyModeld7c2e5___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* ElementyModeld7c2e5___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* ElementyModeld7c2e5___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* ElementyModeld7c2e5___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int ElementyModeld7c2e5___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* ElementyModeld7c2e5___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* ElementyModeld7c2e5___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* ElementyModeld7c2e5___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* ElementyModeld7c2e5___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* ElementyModeld7c2e5___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* ElementyModeld7c2e5___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* ElementyModeld7c2e5___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList ElementyModeld7c2e5___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ElementyModeld7c2e5___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* ElementyModeld7c2e5___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList ElementyModeld7c2e5___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ElementyModeld7c2e5___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* ElementyModeld7c2e5___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* ElementyModeld7c2e5___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* ElementyModeld7c2e5___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* ElementyModeld7c2e5___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* ElementyModeld7c2e5___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int ElementyModeld7c2e5_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* ElementyModeld7c2e5_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int ElementyModeld7c2e5_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* ElementyModeld7c2e5_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* ElementyModeld7c2e5___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ElementyModeld7c2e5___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ElementyModeld7c2e5___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ElementyModeld7c2e5___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ElementyModeld7c2e5___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ElementyModeld7c2e5___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ElementyModeld7c2e5___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ElementyModeld7c2e5___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ElementyModeld7c2e5___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ElementyModeld7c2e5___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ElementyModeld7c2e5___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* ElementyModeld7c2e5___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList ElementyModeld7c2e5___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ElementyModeld7c2e5___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* ElementyModeld7c2e5___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList ElementyModeld7c2e5___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ElementyModeld7c2e5___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void ElementyModeld7c2e5___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* ElementyModeld7c2e5___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList ElementyModeld7c2e5___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ElementyModeld7c2e5___elementitem_atList(void* ptr, int i)
{
	return ({Elementd7c2e5* tmp = static_cast<QList<Elementd7c2e5*>*>(ptr)->at(i); if (i == static_cast<QList<Elementd7c2e5*>*>(ptr)->size()-1) { static_cast<QList<Elementd7c2e5*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5___elementitem_setList(void* ptr, void* i)
{
	static_cast<QList<Elementd7c2e5*>*>(ptr)->append(static_cast<Elementd7c2e5*>(i));
}

void* ElementyModeld7c2e5___elementitem_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Elementd7c2e5*>();
}

void* ElementyModeld7c2e5___setElementitem_elementitem_atList(void* ptr, int i)
{
	return ({Elementd7c2e5* tmp = static_cast<QList<Elementd7c2e5*>*>(ptr)->at(i); if (i == static_cast<QList<Elementd7c2e5*>*>(ptr)->size()-1) { static_cast<QList<Elementd7c2e5*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5___setElementitem_elementitem_setList(void* ptr, void* i)
{
	static_cast<QList<Elementd7c2e5*>*>(ptr)->append(static_cast<Elementd7c2e5*>(i));
}

void* ElementyModeld7c2e5___setElementitem_elementitem_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Elementd7c2e5*>();
}

void* ElementyModeld7c2e5___elementitemChanged_elementitem_atList(void* ptr, int i)
{
	return ({Elementd7c2e5* tmp = static_cast<QList<Elementd7c2e5*>*>(ptr)->at(i); if (i == static_cast<QList<Elementd7c2e5*>*>(ptr)->size()-1) { static_cast<QList<Elementd7c2e5*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5___elementitemChanged_elementitem_setList(void* ptr, void* i)
{
	static_cast<QList<Elementd7c2e5*>*>(ptr)->append(static_cast<Elementd7c2e5*>(i));
}

void* ElementyModeld7c2e5___elementitemChanged_elementitem_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Elementd7c2e5*>();
}

int ElementyModeld7c2e5_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* ElementyModeld7c2e5_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int ElementyModeld7c2e5_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* ElementyModeld7c2e5_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int ElementyModeld7c2e5_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ElementyModeld7c2e5_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* ElementyModeld7c2e5_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* ElementyModeld7c2e5_NewElementyModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new ElementyModeld7c2e5(static_cast<QWindow*>(parent));
	} else {
		return new ElementyModeld7c2e5(static_cast<QObject*>(parent));
	}
}

void ElementyModeld7c2e5_DestroyElementyModel(void* ptr)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->~ElementyModeld7c2e5();
}

void ElementyModeld7c2e5_DestroyElementyModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char ElementyModeld7c2e5_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* ElementyModeld7c2e5_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* ElementyModeld7c2e5_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long ElementyModeld7c2e5_FlagsDefault(void* ptr, void* index)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}



char ElementyModeld7c2e5_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char ElementyModeld7c2e5_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char ElementyModeld7c2e5_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char ElementyModeld7c2e5_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char ElementyModeld7c2e5_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char ElementyModeld7c2e5_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char ElementyModeld7c2e5_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char ElementyModeld7c2e5_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char ElementyModeld7c2e5_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char ElementyModeld7c2e5_SubmitDefault(void* ptr)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::submit();
}

void ElementyModeld7c2e5_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void ElementyModeld7c2e5_ResetInternalDataDefault(void* ptr)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::resetInternalData();
}

void ElementyModeld7c2e5_RevertDefault(void* ptr)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::revert();
}

void ElementyModeld7c2e5_SortDefault(void* ptr, int column, long long order)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList ElementyModeld7c2e5_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList ElementyModeld7c2e5_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ElementyModeld7c2e5_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* ElementyModeld7c2e5_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* ElementyModeld7c2e5_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList ElementyModeld7c2e5_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ElementyModeld7c2e5_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString ElementyModeld7c2e5_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray ta7ef81 = static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(ta7ef81.prepend("WHITESPACE").constData()+10), ta7ef81.size()-10 }; });
}

void* ElementyModeld7c2e5_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* ElementyModeld7c2e5_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long ElementyModeld7c2e5_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long ElementyModeld7c2e5_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::supportedDropActions();
}

char ElementyModeld7c2e5_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char ElementyModeld7c2e5_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char ElementyModeld7c2e5_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int ElementyModeld7c2e5_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int ElementyModeld7c2e5_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char ElementyModeld7c2e5_EventDefault(void* ptr, void* e)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char ElementyModeld7c2e5_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void ElementyModeld7c2e5_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void ElementyModeld7c2e5_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ElementyModeld7c2e5_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void ElementyModeld7c2e5_DeleteLaterDefault(void* ptr)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::deleteLater();
}

void ElementyModeld7c2e5_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void ElementyModeld7c2e5_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ElementyModeld7c2e5*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString Grupad7c2e5_GrupaName(void* ptr)
{
	return ({ QByteArray t5f48e5 = static_cast<Grupad7c2e5*>(ptr)->grupaName().toUtf8(); Moc_PackedString { const_cast<char*>(t5f48e5.prepend("WHITESPACE").constData()+10), t5f48e5.size()-10 }; });
}

struct Moc_PackedString Grupad7c2e5_GrupaNameDefault(void* ptr)
{
	return ({ QByteArray td4f2be = static_cast<Grupad7c2e5*>(ptr)->grupaNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(td4f2be.prepend("WHITESPACE").constData()+10), td4f2be.size()-10 }; });
}

void Grupad7c2e5_SetGrupaName(void* ptr, struct Moc_PackedString grupaName)
{
	static_cast<Grupad7c2e5*>(ptr)->setGrupaName(QString::fromUtf8(grupaName.data, grupaName.len));
}

void Grupad7c2e5_SetGrupaNameDefault(void* ptr, struct Moc_PackedString grupaName)
{
	static_cast<Grupad7c2e5*>(ptr)->setGrupaNameDefault(QString::fromUtf8(grupaName.data, grupaName.len));
}

void Grupad7c2e5_ConnectGrupaNameChanged(void* ptr)
{
	QObject::connect(static_cast<Grupad7c2e5*>(ptr), static_cast<void (Grupad7c2e5::*)(QString)>(&Grupad7c2e5::grupaNameChanged), static_cast<Grupad7c2e5*>(ptr), static_cast<void (Grupad7c2e5::*)(QString)>(&Grupad7c2e5::Signal_GrupaNameChanged));
}

void Grupad7c2e5_DisconnectGrupaNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<Grupad7c2e5*>(ptr), static_cast<void (Grupad7c2e5::*)(QString)>(&Grupad7c2e5::grupaNameChanged), static_cast<Grupad7c2e5*>(ptr), static_cast<void (Grupad7c2e5::*)(QString)>(&Grupad7c2e5::Signal_GrupaNameChanged));
}

void Grupad7c2e5_GrupaNameChanged(void* ptr, struct Moc_PackedString grupaName)
{
	static_cast<Grupad7c2e5*>(ptr)->grupaNameChanged(QString::fromUtf8(grupaName.data, grupaName.len));
}

int Grupad7c2e5_GrupaId(void* ptr)
{
	return static_cast<Grupad7c2e5*>(ptr)->grupaId();
}

int Grupad7c2e5_GrupaIdDefault(void* ptr)
{
	return static_cast<Grupad7c2e5*>(ptr)->grupaIdDefault();
}

void Grupad7c2e5_SetGrupaId(void* ptr, int grupaId)
{
	static_cast<Grupad7c2e5*>(ptr)->setGrupaId(grupaId);
}

void Grupad7c2e5_SetGrupaIdDefault(void* ptr, int grupaId)
{
	static_cast<Grupad7c2e5*>(ptr)->setGrupaIdDefault(grupaId);
}

void Grupad7c2e5_ConnectGrupaIdChanged(void* ptr)
{
	QObject::connect(static_cast<Grupad7c2e5*>(ptr), static_cast<void (Grupad7c2e5::*)(qint32)>(&Grupad7c2e5::grupaIdChanged), static_cast<Grupad7c2e5*>(ptr), static_cast<void (Grupad7c2e5::*)(qint32)>(&Grupad7c2e5::Signal_GrupaIdChanged));
}

void Grupad7c2e5_DisconnectGrupaIdChanged(void* ptr)
{
	QObject::disconnect(static_cast<Grupad7c2e5*>(ptr), static_cast<void (Grupad7c2e5::*)(qint32)>(&Grupad7c2e5::grupaIdChanged), static_cast<Grupad7c2e5*>(ptr), static_cast<void (Grupad7c2e5::*)(qint32)>(&Grupad7c2e5::Signal_GrupaIdChanged));
}

void Grupad7c2e5_GrupaIdChanged(void* ptr, int grupaId)
{
	static_cast<Grupad7c2e5*>(ptr)->grupaIdChanged(grupaId);
}

int Grupad7c2e5_Grupad7c2e5_QRegisterMetaType()
{
	return qRegisterMetaType<Grupad7c2e5*>();
}

int Grupad7c2e5_Grupad7c2e5_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Grupad7c2e5*>(const_cast<const char*>(typeName));
}

int Grupad7c2e5_Grupad7c2e5_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Grupad7c2e5>();
#else
	return 0;
#endif
}

int Grupad7c2e5_Grupad7c2e5_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Grupad7c2e5>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Grupad7c2e5___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Grupad7c2e5___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Grupad7c2e5___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Grupad7c2e5___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Grupad7c2e5___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Grupad7c2e5___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Grupad7c2e5___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Grupad7c2e5___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Grupad7c2e5___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Grupad7c2e5___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Grupad7c2e5___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Grupad7c2e5___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Grupad7c2e5___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Grupad7c2e5___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Grupad7c2e5___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Grupad7c2e5_NewGrupa(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Grupad7c2e5(static_cast<QWindow*>(parent));
	} else {
		return new Grupad7c2e5(static_cast<QObject*>(parent));
	}
}

void Grupad7c2e5_DestroyGrupa(void* ptr)
{
	static_cast<Grupad7c2e5*>(ptr)->~Grupad7c2e5();
}

void Grupad7c2e5_DestroyGrupaDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Grupad7c2e5_EventDefault(void* ptr, void* e)
{
	return static_cast<Grupad7c2e5*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Grupad7c2e5_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Grupad7c2e5*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Grupad7c2e5_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Grupad7c2e5*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Grupad7c2e5_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Grupad7c2e5*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Grupad7c2e5_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Grupad7c2e5*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Grupad7c2e5_DeleteLaterDefault(void* ptr)
{
	static_cast<Grupad7c2e5*>(ptr)->QObject::deleteLater();
}

void Grupad7c2e5_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Grupad7c2e5*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Grupad7c2e5_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Grupad7c2e5*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"

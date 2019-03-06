/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.12.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#include <QtCore/QList>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.12.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_Produktd7c2e5_t {
    QByteArrayData data[8];
    char stringdata0[98];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_Produktd7c2e5_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_Produktd7c2e5_t qt_meta_stringdata_Produktd7c2e5 = {
    {
QT_MOC_LITERAL(0, 0, 13), // "Produktd7c2e5"
QT_MOC_LITERAL(1, 14, 18), // "produktNameChanged"
QT_MOC_LITERAL(2, 33, 0), // ""
QT_MOC_LITERAL(3, 34, 11), // "produktName"
QT_MOC_LITERAL(4, 46, 15), // "produktWChanged"
QT_MOC_LITERAL(5, 62, 8), // "produktW"
QT_MOC_LITERAL(6, 71, 16), // "produktIdChanged"
QT_MOC_LITERAL(7, 88, 9) // "produktId"

    },
    "Produktd7c2e5\0produktNameChanged\0\0"
    "produktName\0produktWChanged\0produktW\0"
    "produktIdChanged\0produktId"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_Produktd7c2e5[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       3,   14, // methods
       3,   38, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       3,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   29,    2, 0x06 /* Public */,
       4,    1,   32,    2, 0x06 /* Public */,
       6,    1,   35,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::Float,    5,
    QMetaType::Void, QMetaType::Int,    7,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::Float, 0x00495103,
       7, QMetaType::Int, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,

       0        // eod
};

void Produktd7c2e5::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        Produktd7c2e5 *_t = static_cast<Produktd7c2e5 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->produktNameChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->produktWChanged((*reinterpret_cast< float(*)>(_a[1]))); break;
        case 2: _t->produktIdChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (Produktd7c2e5::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Produktd7c2e5::produktNameChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (Produktd7c2e5::*)(float );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Produktd7c2e5::produktWChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (Produktd7c2e5::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Produktd7c2e5::produktIdChanged)) {
                *result = 2;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        Produktd7c2e5 *_t = static_cast<Produktd7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->produktName(); break;
        case 1: *reinterpret_cast< float*>(_v) = _t->produktW(); break;
        case 2: *reinterpret_cast< qint32*>(_v) = _t->produktId(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        Produktd7c2e5 *_t = static_cast<Produktd7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setProduktName(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setProduktW(*reinterpret_cast< float*>(_v)); break;
        case 2: _t->setProduktId(*reinterpret_cast< qint32*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject Produktd7c2e5::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_Produktd7c2e5.data,
    qt_meta_data_Produktd7c2e5,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *Produktd7c2e5::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *Produktd7c2e5::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_Produktd7c2e5.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int Produktd7c2e5::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 3)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 3)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 3;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 3;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void Produktd7c2e5::produktNameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void Produktd7c2e5::produktWChanged(float _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void Produktd7c2e5::produktIdChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_ProduktyModeld7c2e5_t {
    QByteArrayData data[17];
    char stringdata0[194];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ProduktyModeld7c2e5_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ProduktyModeld7c2e5_t qt_meta_stringdata_ProduktyModeld7c2e5 = {
    {
QT_MOC_LITERAL(0, 0, 19), // "ProduktyModeld7c2e5"
QT_MOC_LITERAL(1, 20, 12), // "rolesChanged"
QT_MOC_LITERAL(2, 33, 0), // ""
QT_MOC_LITERAL(3, 34, 10), // "type378cdd"
QT_MOC_LITERAL(4, 45, 5), // "roles"
QT_MOC_LITERAL(5, 51, 18), // "produktitemChanged"
QT_MOC_LITERAL(6, 70, 21), // "QList<Produktd7c2e5*>"
QT_MOC_LITERAL(7, 92, 11), // "produktitem"
QT_MOC_LITERAL(8, 104, 10), // "addProdukt"
QT_MOC_LITERAL(9, 115, 14), // "Produktd7c2e5*"
QT_MOC_LITERAL(10, 130, 2), // "v0"
QT_MOC_LITERAL(11, 133, 11), // "editProdukt"
QT_MOC_LITERAL(12, 145, 3), // "row"
QT_MOC_LITERAL(13, 149, 11), // "produktName"
QT_MOC_LITERAL(14, 161, 8), // "produktW"
QT_MOC_LITERAL(15, 170, 9), // "produktId"
QT_MOC_LITERAL(16, 180, 13) // "removeProdukt"

    },
    "ProduktyModeld7c2e5\0rolesChanged\0\0"
    "type378cdd\0roles\0produktitemChanged\0"
    "QList<Produktd7c2e5*>\0produktitem\0"
    "addProdukt\0Produktd7c2e5*\0v0\0editProdukt\0"
    "row\0produktName\0produktW\0produktId\0"
    "removeProdukt"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ProduktyModeld7c2e5[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       5,   14, // methods
       2,   60, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   39,    2, 0x06 /* Public */,
       5,    1,   42,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       8,    1,   45,    2, 0x0a /* Public */,
      11,    4,   48,    2, 0x0a /* Public */,
      16,    1,   57,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, 0x80000000 | 6,    7,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 9,   10,
    QMetaType::Void, QMetaType::Int, QMetaType::QString, QMetaType::Float, QMetaType::Int,   12,   13,   14,   15,
    QMetaType::Void, QMetaType::Int,   12,

 // properties: name, type, flags
       4, 0x80000000 | 3, 0x0049510b,
       7, 0x80000000 | 6, 0x0049510b,

 // properties: notify_signal_id
       0,
       1,

       0        // eod
};

void ProduktyModeld7c2e5::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ProduktyModeld7c2e5 *_t = static_cast<ProduktyModeld7c2e5 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 1: _t->produktitemChanged((*reinterpret_cast< QList<Produktd7c2e5*>(*)>(_a[1]))); break;
        case 2: _t->addProdukt((*reinterpret_cast< Produktd7c2e5*(*)>(_a[1]))); break;
        case 3: _t->editProdukt((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< float(*)>(_a[3])),(*reinterpret_cast< qint32(*)>(_a[4]))); break;
        case 4: _t->removeProdukt((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<Produktd7c2e5*> >(); break;
            }
            break;
        case 2:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< Produktd7c2e5* >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ProduktyModeld7c2e5::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ProduktyModeld7c2e5::rolesChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ProduktyModeld7c2e5::*)(QList<Produktd7c2e5*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ProduktyModeld7c2e5::produktitemChanged)) {
                *result = 1;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<Produktd7c2e5*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ProduktyModeld7c2e5 *_t = static_cast<ProduktyModeld7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<Produktd7c2e5*>*>(_v) = _t->produktitem(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        ProduktyModeld7c2e5 *_t = static_cast<ProduktyModeld7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setProduktitem(*reinterpret_cast< QList<Produktd7c2e5*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject ProduktyModeld7c2e5::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_ProduktyModeld7c2e5.data,
    qt_meta_data_ProduktyModeld7c2e5,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *ProduktyModeld7c2e5::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ProduktyModeld7c2e5::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ProduktyModeld7c2e5.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int ProduktyModeld7c2e5::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 5)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 5;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 5)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 5;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 2;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void ProduktyModeld7c2e5::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ProduktyModeld7c2e5::produktitemChanged(QList<Produktd7c2e5*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
struct qt_meta_stringdata_Elementd7c2e5_t {
    QByteArrayData data[14];
    char stringdata0[203];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_Elementd7c2e5_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_Elementd7c2e5_t qt_meta_stringdata_Elementd7c2e5 = {
    {
QT_MOC_LITERAL(0, 0, 13), // "Elementd7c2e5"
QT_MOC_LITERAL(1, 14, 19), // "elementGNameChanged"
QT_MOC_LITERAL(2, 34, 0), // ""
QT_MOC_LITERAL(3, 35, 12), // "elementGName"
QT_MOC_LITERAL(4, 48, 18), // "elementStanChanged"
QT_MOC_LITERAL(5, 67, 11), // "elementStan"
QT_MOC_LITERAL(6, 79, 19), // "elementPNameChanged"
QT_MOC_LITERAL(7, 99, 12), // "elementPName"
QT_MOC_LITERAL(8, 112, 19), // "elementIloscChanged"
QT_MOC_LITERAL(9, 132, 12), // "elementIlosc"
QT_MOC_LITERAL(10, 145, 18), // "elementDataChanged"
QT_MOC_LITERAL(11, 164, 11), // "elementData"
QT_MOC_LITERAL(12, 176, 16), // "elementIdChanged"
QT_MOC_LITERAL(13, 193, 9) // "elementId"

    },
    "Elementd7c2e5\0elementGNameChanged\0\0"
    "elementGName\0elementStanChanged\0"
    "elementStan\0elementPNameChanged\0"
    "elementPName\0elementIloscChanged\0"
    "elementIlosc\0elementDataChanged\0"
    "elementData\0elementIdChanged\0elementId"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_Elementd7c2e5[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       6,   14, // methods
       6,   62, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       6,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   44,    2, 0x06 /* Public */,
       4,    1,   47,    2, 0x06 /* Public */,
       6,    1,   50,    2, 0x06 /* Public */,
       8,    1,   53,    2, 0x06 /* Public */,
      10,    1,   56,    2, 0x06 /* Public */,
      12,    1,   59,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::Float,    5,
    QMetaType::Void, QMetaType::QString,    7,
    QMetaType::Void, QMetaType::Int,    9,
    QMetaType::Void, QMetaType::QString,   11,
    QMetaType::Void, QMetaType::Int,   13,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::Float, 0x00495103,
       7, QMetaType::QString, 0x00495103,
       9, QMetaType::Int, 0x00495103,
      11, QMetaType::QString, 0x00495103,
      13, QMetaType::Int, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,
       3,
       4,
       5,

       0        // eod
};

void Elementd7c2e5::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        Elementd7c2e5 *_t = static_cast<Elementd7c2e5 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->elementGNameChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->elementStanChanged((*reinterpret_cast< float(*)>(_a[1]))); break;
        case 2: _t->elementPNameChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->elementIloscChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 4: _t->elementDataChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->elementIdChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (Elementd7c2e5::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Elementd7c2e5::elementGNameChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (Elementd7c2e5::*)(float );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Elementd7c2e5::elementStanChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (Elementd7c2e5::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Elementd7c2e5::elementPNameChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (Elementd7c2e5::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Elementd7c2e5::elementIloscChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (Elementd7c2e5::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Elementd7c2e5::elementDataChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (Elementd7c2e5::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Elementd7c2e5::elementIdChanged)) {
                *result = 5;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        Elementd7c2e5 *_t = static_cast<Elementd7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->elementGName(); break;
        case 1: *reinterpret_cast< float*>(_v) = _t->elementStan(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->elementPName(); break;
        case 3: *reinterpret_cast< qint32*>(_v) = _t->elementIlosc(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->elementData(); break;
        case 5: *reinterpret_cast< qint32*>(_v) = _t->elementId(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        Elementd7c2e5 *_t = static_cast<Elementd7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setElementGName(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setElementStan(*reinterpret_cast< float*>(_v)); break;
        case 2: _t->setElementPName(*reinterpret_cast< QString*>(_v)); break;
        case 3: _t->setElementIlosc(*reinterpret_cast< qint32*>(_v)); break;
        case 4: _t->setElementData(*reinterpret_cast< QString*>(_v)); break;
        case 5: _t->setElementId(*reinterpret_cast< qint32*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject Elementd7c2e5::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_Elementd7c2e5.data,
    qt_meta_data_Elementd7c2e5,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *Elementd7c2e5::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *Elementd7c2e5::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_Elementd7c2e5.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int Elementd7c2e5::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 6)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 6)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 6;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 6;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void Elementd7c2e5::elementGNameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void Elementd7c2e5::elementStanChanged(float _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void Elementd7c2e5::elementPNameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void Elementd7c2e5::elementIloscChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void Elementd7c2e5::elementDataChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void Elementd7c2e5::elementIdChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}
struct qt_meta_stringdata_ElementyModeld7c2e5_t {
    QByteArrayData data[22];
    char stringdata0[257];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ElementyModeld7c2e5_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ElementyModeld7c2e5_t qt_meta_stringdata_ElementyModeld7c2e5 = {
    {
QT_MOC_LITERAL(0, 0, 19), // "ElementyModeld7c2e5"
QT_MOC_LITERAL(1, 20, 12), // "rolesChanged"
QT_MOC_LITERAL(2, 33, 0), // ""
QT_MOC_LITERAL(3, 34, 10), // "type378cdd"
QT_MOC_LITERAL(4, 45, 5), // "roles"
QT_MOC_LITERAL(5, 51, 18), // "elementitemChanged"
QT_MOC_LITERAL(6, 70, 21), // "QList<Elementd7c2e5*>"
QT_MOC_LITERAL(7, 92, 11), // "elementitem"
QT_MOC_LITERAL(8, 104, 10), // "addElement"
QT_MOC_LITERAL(9, 115, 14), // "Elementd7c2e5*"
QT_MOC_LITERAL(10, 130, 2), // "v0"
QT_MOC_LITERAL(11, 133, 11), // "editElement"
QT_MOC_LITERAL(12, 145, 3), // "row"
QT_MOC_LITERAL(13, 149, 12), // "elementGName"
QT_MOC_LITERAL(14, 162, 11), // "elementStan"
QT_MOC_LITERAL(15, 174, 12), // "elementPName"
QT_MOC_LITERAL(16, 187, 12), // "elementIlosc"
QT_MOC_LITERAL(17, 200, 11), // "elementData"
QT_MOC_LITERAL(18, 212, 9), // "elementId"
QT_MOC_LITERAL(19, 222, 13), // "removeElement"
QT_MOC_LITERAL(20, 236, 15), // "removeElementId"
QT_MOC_LITERAL(21, 252, 4) // "idin"

    },
    "ElementyModeld7c2e5\0rolesChanged\0\0"
    "type378cdd\0roles\0elementitemChanged\0"
    "QList<Elementd7c2e5*>\0elementitem\0"
    "addElement\0Elementd7c2e5*\0v0\0editElement\0"
    "row\0elementGName\0elementStan\0elementPName\0"
    "elementIlosc\0elementData\0elementId\0"
    "removeElement\0removeElementId\0idin"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ElementyModeld7c2e5[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       6,   14, // methods
       2,   74, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   44,    2, 0x06 /* Public */,
       5,    1,   47,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       8,    1,   50,    2, 0x0a /* Public */,
      11,    7,   53,    2, 0x0a /* Public */,
      19,    1,   68,    2, 0x0a /* Public */,
      20,    1,   71,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, 0x80000000 | 6,    7,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 9,   10,
    QMetaType::Void, QMetaType::Int, QMetaType::QString, QMetaType::Float, QMetaType::QString, QMetaType::Int, QMetaType::QString, QMetaType::Int,   12,   13,   14,   15,   16,   17,   18,
    QMetaType::Void, QMetaType::Int,   12,
    QMetaType::Void, QMetaType::Int,   21,

 // properties: name, type, flags
       4, 0x80000000 | 3, 0x0049510b,
       7, 0x80000000 | 6, 0x0049510b,

 // properties: notify_signal_id
       0,
       1,

       0        // eod
};

void ElementyModeld7c2e5::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ElementyModeld7c2e5 *_t = static_cast<ElementyModeld7c2e5 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 1: _t->elementitemChanged((*reinterpret_cast< QList<Elementd7c2e5*>(*)>(_a[1]))); break;
        case 2: _t->addElement((*reinterpret_cast< Elementd7c2e5*(*)>(_a[1]))); break;
        case 3: _t->editElement((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< float(*)>(_a[3])),(*reinterpret_cast< QString(*)>(_a[4])),(*reinterpret_cast< qint32(*)>(_a[5])),(*reinterpret_cast< QString(*)>(_a[6])),(*reinterpret_cast< qint32(*)>(_a[7]))); break;
        case 4: _t->removeElement((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 5: _t->removeElementId((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<Elementd7c2e5*> >(); break;
            }
            break;
        case 2:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< Elementd7c2e5* >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ElementyModeld7c2e5::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ElementyModeld7c2e5::rolesChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ElementyModeld7c2e5::*)(QList<Elementd7c2e5*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ElementyModeld7c2e5::elementitemChanged)) {
                *result = 1;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<Elementd7c2e5*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ElementyModeld7c2e5 *_t = static_cast<ElementyModeld7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<Elementd7c2e5*>*>(_v) = _t->elementitem(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        ElementyModeld7c2e5 *_t = static_cast<ElementyModeld7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setElementitem(*reinterpret_cast< QList<Elementd7c2e5*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject ElementyModeld7c2e5::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_ElementyModeld7c2e5.data,
    qt_meta_data_ElementyModeld7c2e5,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *ElementyModeld7c2e5::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ElementyModeld7c2e5::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ElementyModeld7c2e5.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int ElementyModeld7c2e5::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 6)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 6)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 2;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void ElementyModeld7c2e5::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ElementyModeld7c2e5::elementitemChanged(QList<Elementd7c2e5*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
struct qt_meta_stringdata_Grupad7c2e5_t {
    QByteArrayData data[6];
    char stringdata0[63];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_Grupad7c2e5_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_Grupad7c2e5_t qt_meta_stringdata_Grupad7c2e5 = {
    {
QT_MOC_LITERAL(0, 0, 11), // "Grupad7c2e5"
QT_MOC_LITERAL(1, 12, 16), // "grupaNameChanged"
QT_MOC_LITERAL(2, 29, 0), // ""
QT_MOC_LITERAL(3, 30, 9), // "grupaName"
QT_MOC_LITERAL(4, 40, 14), // "grupaIdChanged"
QT_MOC_LITERAL(5, 55, 7) // "grupaId"

    },
    "Grupad7c2e5\0grupaNameChanged\0\0grupaName\0"
    "grupaIdChanged\0grupaId"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_Grupad7c2e5[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       2,   14, // methods
       2,   30, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   24,    2, 0x06 /* Public */,
       4,    1,   27,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::Int,    5,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::Int, 0x00495103,

 // properties: notify_signal_id
       0,
       1,

       0        // eod
};

void Grupad7c2e5::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        Grupad7c2e5 *_t = static_cast<Grupad7c2e5 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->grupaNameChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->grupaIdChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (Grupad7c2e5::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Grupad7c2e5::grupaNameChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (Grupad7c2e5::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&Grupad7c2e5::grupaIdChanged)) {
                *result = 1;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        Grupad7c2e5 *_t = static_cast<Grupad7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->grupaName(); break;
        case 1: *reinterpret_cast< qint32*>(_v) = _t->grupaId(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        Grupad7c2e5 *_t = static_cast<Grupad7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setGrupaName(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setGrupaId(*reinterpret_cast< qint32*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject Grupad7c2e5::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_Grupad7c2e5.data,
    qt_meta_data_Grupad7c2e5,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *Grupad7c2e5::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *Grupad7c2e5::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_Grupad7c2e5.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int Grupad7c2e5::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 2)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 2)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 2;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 2;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void Grupad7c2e5::grupaNameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void Grupad7c2e5::grupaIdChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
struct qt_meta_stringdata_GrupyModeld7c2e5_t {
    QByteArrayData data[16];
    char stringdata0[164];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_GrupyModeld7c2e5_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_GrupyModeld7c2e5_t qt_meta_stringdata_GrupyModeld7c2e5 = {
    {
QT_MOC_LITERAL(0, 0, 16), // "GrupyModeld7c2e5"
QT_MOC_LITERAL(1, 17, 12), // "rolesChanged"
QT_MOC_LITERAL(2, 30, 0), // ""
QT_MOC_LITERAL(3, 31, 10), // "type378cdd"
QT_MOC_LITERAL(4, 42, 5), // "roles"
QT_MOC_LITERAL(5, 48, 16), // "grupaitemChanged"
QT_MOC_LITERAL(6, 65, 19), // "QList<Grupad7c2e5*>"
QT_MOC_LITERAL(7, 85, 9), // "grupaitem"
QT_MOC_LITERAL(8, 95, 8), // "addGrupa"
QT_MOC_LITERAL(9, 104, 12), // "Grupad7c2e5*"
QT_MOC_LITERAL(10, 117, 2), // "v0"
QT_MOC_LITERAL(11, 120, 9), // "editGrupa"
QT_MOC_LITERAL(12, 130, 3), // "row"
QT_MOC_LITERAL(13, 134, 9), // "grupaName"
QT_MOC_LITERAL(14, 144, 7), // "grupaId"
QT_MOC_LITERAL(15, 152, 11) // "removeGrupa"

    },
    "GrupyModeld7c2e5\0rolesChanged\0\0"
    "type378cdd\0roles\0grupaitemChanged\0"
    "QList<Grupad7c2e5*>\0grupaitem\0addGrupa\0"
    "Grupad7c2e5*\0v0\0editGrupa\0row\0grupaName\0"
    "grupaId\0removeGrupa"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_GrupyModeld7c2e5[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       5,   14, // methods
       2,   58, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   39,    2, 0x06 /* Public */,
       5,    1,   42,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       8,    1,   45,    2, 0x0a /* Public */,
      11,    3,   48,    2, 0x0a /* Public */,
      15,    1,   55,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, 0x80000000 | 6,    7,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 9,   10,
    QMetaType::Void, QMetaType::Int, QMetaType::QString, QMetaType::Int,   12,   13,   14,
    QMetaType::Void, QMetaType::Int,   12,

 // properties: name, type, flags
       4, 0x80000000 | 3, 0x0049510b,
       7, 0x80000000 | 6, 0x0049510b,

 // properties: notify_signal_id
       0,
       1,

       0        // eod
};

void GrupyModeld7c2e5::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        GrupyModeld7c2e5 *_t = static_cast<GrupyModeld7c2e5 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 1: _t->grupaitemChanged((*reinterpret_cast< QList<Grupad7c2e5*>(*)>(_a[1]))); break;
        case 2: _t->addGrupa((*reinterpret_cast< Grupad7c2e5*(*)>(_a[1]))); break;
        case 3: _t->editGrupa((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< qint32(*)>(_a[3]))); break;
        case 4: _t->removeGrupa((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<Grupad7c2e5*> >(); break;
            }
            break;
        case 2:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< Grupad7c2e5* >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (GrupyModeld7c2e5::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&GrupyModeld7c2e5::rolesChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (GrupyModeld7c2e5::*)(QList<Grupad7c2e5*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&GrupyModeld7c2e5::grupaitemChanged)) {
                *result = 1;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<Grupad7c2e5*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        GrupyModeld7c2e5 *_t = static_cast<GrupyModeld7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<Grupad7c2e5*>*>(_v) = _t->grupaitem(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        GrupyModeld7c2e5 *_t = static_cast<GrupyModeld7c2e5 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setGrupaitem(*reinterpret_cast< QList<Grupad7c2e5*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject GrupyModeld7c2e5::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_GrupyModeld7c2e5.data,
    qt_meta_data_GrupyModeld7c2e5,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *GrupyModeld7c2e5::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *GrupyModeld7c2e5::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_GrupyModeld7c2e5.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int GrupyModeld7c2e5::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 5)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 5;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 5)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 5;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 2;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void GrupyModeld7c2e5::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void GrupyModeld7c2e5::grupaitemChanged(QList<Grupad7c2e5*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE

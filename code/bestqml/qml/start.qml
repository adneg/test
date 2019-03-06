import QtQuick 2.6
import QtQuick.Controls 2.2

import QtQuick.Layouts 1.3
import QtQuick.Controls.Styles 1.4
import CustomQmlTypes 1.0

//import QtGraphicalEffects 1.0
//Default, Fusion, Imagine, Universal, Material
import QtQuick.Controls.Material 2.2
//import Qt.labs.settings 1.0

ApplicationWindow {
    property BridgeQML template: BridgeQML {
    }
    id: window
    //    width: 1280
    //    height: 1024
    //    color: "transparent"
    visibility: "FullScreen"
    visible: true
//    Settings {
//        id: settings
//        property string style: "Default"
//    }
    Menu {
        id: optionsMenu
        transformOrigin: Menu.TopLeft

//        MenuItem {
//            font.pixelSize: template.bSize
//            text: "Ustawienia"

//            onTriggered: {
//                tlo.visible = true

//                settingsPopup.open()
//            }
//        }
        MenuItem {
            text: "Wyłacz"
            font.pixelSize: template.bSize
            onTriggered: {
                tlo.visible = true
                window.close()
            }
        }
    }

    Rectangle {
        id: header

        anchors.top: parent.top

        //            anchors.bottom: test.top
        //anchors.verticalCenter:  parent.verticalCenter
        //anchors.fill: parent
        width: parent.width
        height: (parent.height / 10 * 1.5)
        color: "white"
        border.color: "black"
        border.width: 1
        z: 150
        // radius: 10
        Rectangle {
            id: logo
            anchors.top: parent.top
            anchors.left: parent.left

            //            anchors.bottom: test.top
            //anchors.verticalCenter:  parent.verticalCenter
            //anchors.fill: parent
            width: parent.width / 2
            height: parent.height
            color: "white"
            //            border.color: "black"
            //            border.width: 1
            ToolButton {
                id: tb
                anchors.verticalCenter: parent.verticalCenter
                anchors.left: parent.left
                font.pixelSize: (logo.width / 8)
                text: qsTr("⋮")
                onClicked: optionsMenu.open()
            }

            Text {
//                anchors.left: tb.right
                                anchors.horizontalCenter: parent.horizontalCenter
                anchors.verticalCenter: parent.verticalCenter
                text: "Akwizycja Danych"
                font.family: "Helvetica"
                font.italic: true
                font.pixelSize: (logo.width / 16)
                color: "black"
            }
            implicitWidth: 100
            implicitHeight: 40
            border.color: "#999"
            gradient: Gradient {
                GradientStop { position: 0 ; color: "#fff" }
                GradientStop { position: 1 ; color: "#eee" }
            }
            // radius: 10
        }
        Rectangle {
            id: stan
            anchors.top: parent.top
            anchors.right: parent.right

            //            anchors.bottom: test.top
            //anchors.verticalCenter:  parent.verticalCenter
            //anchors.fill: parent
            width: parent.width / 2
            height: parent.height
            color: "lightskyblue"
           // color: "lightgreen"
            // border.color: "red"
            // border.width: 1
            // radius: 10
            Text {
                id: percent
                anchors.horizontalCenter: parent.horizontalCenter
                anchors.verticalCenter: parent.verticalCenter
                text: "__%"
                font.family: "Helvetica"
                font.pixelSize: (logo.width / 6)
                //color: "black"
                color: "black"
            }
            //            MouseArea {
            //                anchors.fill: parent
            //                onClicked: {
            //                    tlo.visible = true
            //                    settingsPopup.open()
            //                }
            //            }
        }
    }

    Rectangle {
        id: labels
        anchors.top: header.bottom
        height: (parent.height / 10)
        z: 100
        width: parent.width
        radius: 10
        Rectangle {
            id: labelW
            anchors.margins: 2
            anchors.top: labels.top
            anchors.bottom: labels.bottom
            anchors.left: labels.left
            height: (parent.height / 10)
            width: (parent.width / 5) - 4
            Text {

                anchors.horizontalCenter: parent.horizontalCenter
                anchors.verticalCenter: parent.verticalCenter
                text: "WYDZIAŁ"
                font.bold: true
                font.pixelSize: template.lSize
            }
        }

        Rectangle {
            id: labelSO
            anchors.margins: 2
            anchors.top: labels.top
            anchors.bottom: labels.bottom
            anchors.left: labelW.right
            height: (parent.height / 10)
            width: (parent.width / 5) - 4
            Text {

                anchors.horizontalCenter: parent.horizontalCenter
                anchors.verticalCenter: parent.verticalCenter
                text: "STAN OSOBOWY"
                font.bold: true
                font.pixelSize: template.lSize
            }
        }
        Rectangle {
            id: labelP
            anchors.margins: 2
            anchors.top: labels.top
            anchors.bottom: labels.bottom
            anchors.left: labelSO.right
            height: (parent.height / 10)
            width: (parent.width / 5) - 4
            Text {

                anchors.horizontalCenter: parent.horizontalCenter
                anchors.verticalCenter: parent.verticalCenter
                text: "PRODUKT"
                font.bold: true
                font.pixelSize: template.lSize
            }
        }

        Rectangle {
            id: labelI
            anchors.margins: 2
            anchors.top: labels.top
            anchors.bottom: labels.bottom
            anchors.left: labelP.right
            height: (parent.height / 10)
            width: (parent.width / 5) - 4
            Text {

                anchors.horizontalCenter: parent.horizontalCenter
                anchors.verticalCenter: parent.verticalCenter
                text: "ILOŚĆ"
                font.bold: true
                font.pixelSize: template.lSize
            }
        }
    }

    Rectangle {
        id: body
        anchors.top: labels.bottom
        height: (parent.height / 10)
        z: 100
        //anchors.: food.top

        //            anchors.horizontalCenter: parent.horizontalCenter
        //            anchors.verticalCenter:  parent.verticalCenter
        //anchors.fill: parent
        width: parent.width

        //height: 100
        //        color: "white"
        //        border.color: "black"
        //border.width: 5
        radius: 10

        ComboBox {
            anchors.margins: 2
            id: grupa
            width: (parent.width / 5) - 4
            height: parent.height
            anchors.top: parent.top
            anchors.left: parent.left
            model: GrupyModel
            textRole: "grupaName"
            currentIndex: -1
            displayText: {
                if (grupa.currentIndex > -1) {
                    currentText
                } else {
                    "Wydział:"
                }
            }
            font.bold: true
            font.pixelSize: template.cSize
            contentItem: Text {
                leftPadding: 5
                rightPadding: grupa.indicator.width + grupa.spacing

                text: grupa.displayText
                font: grupa.font

                color: grupa.pressed ? "red" : "blue"
                verticalAlignment: Text.AlignVCenter
                elide: Text.ElideRight
            }

            background: Rectangle {
                implicitWidth: 120
                implicitHeight: 40
                border.color: grupa.pressed ? "red" : "black"
                border.width: grupa.visualFocus ? 2 : 1
                radius: 2
            }
            popup: Popup {
                y: grupa.height - 1
                width: grupa.width
                implicitHeight: contentItem.implicitHeight
                padding: 1
                font.pixelSize: template.cSize - 2
                contentItem: ListView {
                    clip: true
                    implicitHeight: contentHeight
                    model: grupa.popup.visible ? grupa.delegateModel : null
                    currentIndex: grupa.highlightedIndex

                    ScrollIndicator.vertical: ScrollIndicator {
                    }
                }

                background: Rectangle {
                    border.color: "black"
                    radius: 2
                }
            }

            onCurrentIndexChanged: {

                //              console.log("Grupy current GrupaId Name: ", grupa.model.data(grupa.model.index(grupa.currentIndex, 0), Qt.UserRole + 1)) //P name
                //              console.log("Grupy current GrupaId Id: ", grupa.model.data(grupa.model.index(grupa.currentIndex, 0), Qt.UserRole + 2)) //P W
            }
        }
        ComboBox {
            anchors.margins: 2

            id: stany
            currentIndex: 0
            width: (parent.width / 5) - 4
            height: parent.height
            anchors.top: parent.top
            anchors.left: grupa.right
            font.bold: true

            model: ListModel {
                id: listModel
                Component.onCompleted: {
                    for (var i = 0; i < 800; i++)
                        listModel.append({
                                             "Name": "" + i / 2
                                         })
                    //listModel.append({"Name": "Item " + i})
                }
            }
            displayText: {

                if (stany.currentIndex > 0) {
                    currentText
                } else {
                    "Stan:"
                }
            }

            font.pixelSize: template.cSize
            contentItem: Text {
                leftPadding: 5
                rightPadding: stany.indicator.width + stany.spacing

                text: stany.displayText
                font: stany.font
                color: stany.pressed ? "red" : "blue"
                verticalAlignment: Text.AlignVCenter
                elide: Text.ElideRight
            }

            background: Rectangle {
                implicitWidth: 120
                implicitHeight: 40
                border.color: stany.pressed ? "red" : "black"
                border.width: stany.visualFocus ? 2 : 1
                radius: 2
            }
            popup: Popup {
                y: stany.height - 1
                width: stany.width
                implicitHeight: contentItem.implicitHeight
                padding: 1
                font.pixelSize: template.cSize - 2
                contentItem: ListView {
                    clip: true
                    implicitHeight: contentHeight
                    model: stany.popup.visible ? stany.delegateModel : null
                    currentIndex: stany.highlightedIndex

                    ScrollIndicator.vertical: ScrollIndicator {
                    }
                }

                background: Rectangle {
                    border.color: "black"
                    radius: 2
                }
            }
        }
        ComboBox {
            anchors.margins: 2
            id: produkt
            currentIndex: -1
            width: (parent.width / 5) - 4
            height: parent.height
            anchors.top: parent.top
            anchors.left: stany.right
            font.bold: true
            model: ProduktyModel
            textRole: "produktName"
            displayText: {

                if (produkt.currentIndex > -1) {
                    currentText
                } else {
                    "Produkt:"
                }
            }
            font.pixelSize: template.cSize
            contentItem: Text {
                leftPadding: 5
                rightPadding: produkt.indicator.width + produkt.spacing

                text: produkt.displayText
                font: produkt.font

                color: produkt.pressed ? "red" : "blue"
                verticalAlignment: Text.AlignVCenter
                elide: Text.ElideRight
            }

            background: Rectangle {
                implicitWidth: 120
                implicitHeight: 40
                border.color: produkt.pressed ? "red" : "black"
                border.width: produkt.visualFocus ? 2 : 1
                radius: 2
            }
            popup: Popup {
                y: produkt.height - 1
                width: produkt.width
                implicitHeight: contentItem.implicitHeight
                padding: 1
                font.pixelSize: template.cSize - 2
                contentItem: ListView {
                    clip: true
                    implicitHeight: contentHeight
                    model: produkt.popup.visible ? produkt.delegateModel : null
                    currentIndex: produkt.highlightedIndex

                    ScrollIndicator.vertical: ScrollIndicator {
                    }
                }

                background: Rectangle {
                    border.color: "black"
                    radius: 2
                }
            }
            onCurrentIndexChanged: {

                //              console.log("Prodkty current Produkt Name: ", produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 1)) //P name
                //              console.log("Prodkty current Produkt W: ", produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 2)) //P W
                //              console.log("Prodkty current Produkt Id: ", produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 3)) //P W
            }
        }

        ComboBox {
            anchors.margins: 2
            id: ilosc
            currentIndex: 0
            width: (parent.width / 5) - 4
            height: parent.height
            anchors.top: parent.top
            anchors.left: produkt.right
            model: 10000
            font.bold: true

            displayText: {

                if (ilosc.currentIndex > 0) {
                    currentText
                } else {
                    "Ilość:"
                }
            }
            font.pixelSize: template.cSize
            contentItem: Text {
                leftPadding: 5
                rightPadding: ilosc.indicator.width + ilosc.spacing

                text: ilosc.displayText
                font: ilosc.font

                color: ilosc.pressed ? "red" : "blue"
                verticalAlignment: Text.AlignVCenter
                elide: Text.ElideRight
            }

            background: Rectangle {
                implicitWidth: 120
                implicitHeight: 40
                border.color: ilosc.pressed ? "red" : "black"
                border.width: ilosc.visualFocus ? 2 : 1
                radius: 2
            }
            popup: Popup {
                y: ilosc.height - 1
                width: ilosc.width

                //                implicitHeight: contentItem.implicitHeight
                //                padding: 1

                //                contentItem: ListView {
                //                    clip: true
                //                    implicitHeight: contentHeight
                //                    model: ilosc.popup.visible ? ilosc.delegateModel : null
                //                    currentIndex: ilosc.highlightedIndex

                //                    ScrollIndicator.vertical: ScrollIndicator {
                //                    }
                //                }
                background: Rectangle {
                    border.color: "black"
                    radius: 2
                }
            }
            onPressedChanged: if (pressed) {
                                  tlo.visible = true
                                  ddialog_calc.open()
                              }
        }

        Button {
            anchors.margins: 2

            id: zatwierdz
            width: (parent.width / 5) - 4
            height: parent.height
            text: qsTr("Zatwierdz")
            anchors.top: parent.top
            anchors.left: ilosc.right
            font.bold: true
            font.pixelSize: template.bSize
            highlighted: true

            //            contentItem: Text {
            //                text: zatwierdz.text
            //                font: zatwierdz.font
            //                opacity: enabled ? 1.0 : 0.3
            //                color: zatwierdz.down ? "blue" : "red"
            //                horizontalAlignment: Text.AlignHCenter
            //                verticalAlignment: Text.AlignVCenter
            //                elide: Text.ElideRight
            //            }

            //            background: Rectangle {
            //                implicitWidth: 100
            //                implicitHeight: 40
            //                opacity: enabled ? 1 : 0.3
            //                border.color: zatwierdz.down ? "#17a81a" : "#21be2b"
            //                border.width: 1
            //                radius: 2
            //            }
            //            background: Rectangle {
            //                implicitWidth: 100
            //                implicitHeight: 25
            //                border.width: zatwierdz.activeFocus ? 2 : 1
            //                border.color: "#888"
            //                radius: 4
            //                gradient: Gradient {
            //                    GradientStop { position: 0 ; color: zatwierdz.pressed ? "#ccc" : "#eee" }
            //                    GradientStop { position: 1 ; color: zatwierdz.pressed ? "#aaa" : "#ccc" }
            //                }
            //            }
            onClicked: {
                //                console.log("Grupy current GrupaId Name: ", grupa.model.data(grupa.model.index(grupa.currentIndex, 0), Qt.UserRole + 1)) //P name
                //                console.log("Grupy current GrupaId Id: ", grupa.model.data(grupa.model.index(grupa.currentIndex, 0), Qt.UserRole + 2)) //P W

                //                console.log("Stan osobowy:",stany.currentIndex/2)

                //                console.log("Prodkty current Produkt Name: ", produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 1)) //P name
                //                console.log("Prodkty current Produkt W: ", produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 2)) //P W
                //                console.log("Prodkty current Produkt Id: ", produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 3)) //P W
                //                console.log("Ilosc:",ilosc.currentIndex)
                if (grupa.currentIndex < 0 || stany.currentIndex < 1
                        || ilosc.currentIndex < 1 || produkt.currentIndex < 0) {
                    console.log("stany.currentIndex:", stany.currentIndex,
                                "ilosc.currentIndex:", ilosc.currentIndex,
                                "produkt.currentIndex", produkt.currentIndex)
                    tlo.visible = true
                    ddialog_e.open()
                    return
                }
                if (grupa.currentIndex > -1 && ilosc.currentIndex > 0
                        && produkt.currentIndex > -1
                        && stany.currentIndex > 0) {
                    window.template.sendZatwierdz(
                                grupa.model.data(grupa.model.index(grupa.currentIndex, 0),
                                                 Qt.UserRole + 1), grupa.model.data(grupa.model.index(grupa.currentIndex, 0), Qt.UserRole + 2),
                                (stany.currentIndex / 2), produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 1),
                                produkt.model.data(produkt.model.index(produkt.currentIndex, 0),
                                                   Qt.UserRole + 2), produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 3),
                                ilosc.currentIndex)
                    ilosc.currentIndex = 0
                    produkt.currentIndex = -1
                }
            }

            //width: 3* zatwierdz.height
        }
    }

    Rectangle {
        id: food
        anchors.margins: 5
        //height: 100
        anchors.top: body.bottom
        anchors.bottom: parent.bottom

        //            anchors.horizontalCenter: parent.horizontalCenter
        //            anchors.verticalCenter:  parent.verticalCenter
        //anchors.fill: parent
        width: parent.width

        //        height: (parent.height / 10 * 2)
        //                                height: 100
        //        color: "green"
        //        border.color: "black"
        //border.width: 1
        //        radius: 10
        ListView {
            id: next
            property int indexToRemove: -10
            property int idToRemove: -10

            anchors.top: parent.top
            anchors.bottom: parent.bottom
            anchors.left: parent.left
            anchors.right: parent.right
            currentIndex: -1
            anchors.margins: 2
            width: parent.width / 5
            rotation: 180

            model: ElementyModel

            delegate: ItemDelegate {
                Rectangle {
                    id: rmarg
                    border.color: next.indexToRemove == index ? "red" : "white"
                    border.width: 1
                    anchors.fill: parent
                }

                rotation: 180
                width: parent.width
                height: window.height / 10

                //text: model.title
                Text {
                    padding: height / 8
                    anchors.margins: 2
                    width: (parent.width / 5) - 4
                    height: parent.height
                    anchors.top: parent.top
                    anchors.left: parent.left

                    font.bold: true

                    id: tx1
                    text: model.elementGName
                    font.family: "Helvetica"
                    //font.pointSize: 24
                    //color: "green"
                    font.pixelSize: template.eSize
                }
                Text {
                    padding: height / 8
                    anchors.margins: 2
                    width: (parent.width / 5) - 4
                    height: parent.height
                    anchors.top: parent.top
                    anchors.left: tx1.right

                    font.bold: true

                    id: tx2
                    text: model.elementStan
                    font.family: "Helvetica"
                    font.pixelSize: template.eSize
                    //font.pointSize: 24
                    //color: "green"
                }

                Text {
                    padding: height / 8
                    anchors.margins: 2
                    width: (parent.width / 5) - 4
                    height: parent.height
                    anchors.top: parent.top
                    anchors.left: tx2.right

                    font.bold: true

                    id: tx3
                    text: model.elementPName
                    font.family: "Helvetica"
                    font.pixelSize: template.eSize
                    //font.pointSize: 24
                    //color: "green"
                }
                Text {
                    padding: height / 8
                    anchors.margins: 2
                    width: (parent.width / 5) - 4
                    height: parent.height
                    anchors.top: parent.top
                    anchors.left: tx3.right

                    font.bold: true

                    id: tx4
                    text: model.elementIlosc
                    font.family: "Helvetica"
                    font.pixelSize: template.eSize
                    //font.pointSize: 24
                    //color: "green"
                }
                Button {
                    anchors.margins: 2

                    id: usun
                    width: (parent.width / 5) - 4

                    height: parent.height - (parent.height / 12)
                    text: "Usuń\n" + model.elementData
                    anchors.top: parent.top
                    anchors.left: tx4.right
                    font.bold: true
                    font.pixelSize: template.bSize
                    onClicked: {
                        next.idToRemove = model.elementId
                        selectremove()
                    }
                }

                //        Button {
                //            z: 50
                //            anchors.margins: 2
                //            id: usun
                //            width: (parent.width/5)-4

                //            height:parent.height-(parent.height/12)
                //            anchors.top: parent.top
                //            anchors.left: tx4.right
                //            font.bold: true
                //            font.pixelSize: template.bSize
                //            //text: qsTr("Usuń\n2019-01-01 10:01")
                //            Text {
                //                            anchors.horizontalCenter: parent.horizontalCenter
                //                            anchors.verticalCenter:  parent.verticalCenter
                //                //padding: height/8
                //                anchors.margins: 2

                //                font.bold: true
                //                font.pixelSize: template.bSize
                //                id: txusun
                //                text: "Usuń\n"+model.elementData
                //                font.family: "Helvetica"
                //                //font.pointSize: 24
                //               //color: "green"
                //            }
                //            onClicked: {
                //                next.idToRemove = model.elementId
                //                selectremove()

                //                //ddialog.open()
                ////                window.template.sendRemoveE(model.currentIndex)

                //                    //next.currentIndex = index
                //            }

                //        }
                function selectremove() {
                    model.currentIndex = index
                    next.indexToRemove = index
                    tlo.visible = true
                    ddialog.open()
                }
            }
            ScrollIndicator.vertical: ScrollIndicator {
            }
            onContentHeightChanged: {

                //                currentIndex = count - 1
                positionViewAtEnd()
            }
        }
    }
    Popup {

        id: ddialog_e
        z: 200

        closePolicy: Popup.NoAutoClose
        width: window.width / 2
        height: window.height / 2
        x: (window.width - width) / 2
        y: (window.height - height) / 2

        //          background: Rectangle {
        //                border.color: "red"
        //                border.width: 3
        //              anchors.fill: parent

        //          }
        Label {
            anchors.horizontalCenter: parent.horizontalCenter
            //anchors. verticalCenter:  parent.verticalCenter
            anchors.top: parent.top
            anchors.margins: 5

            id: label_e
            text: "Za niskie wartośći, sprawdź pola:\n "
            color: "blue"
            font.bold: true
            font.pixelSize: template.iSize
        }

        Label {
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.top: label_e.bottom
            //              anchors. verticalCenter:  parent.verticalCenter
            //             anchors.top:label_e.bottom
            anchors.margins: 5

            id: label2_e
            text: "WYDZIAŁ\nSTAN\nPRODUKT\nILOŚĆ"
            color: "blue"
            font.bold: true
            font.pixelSize: template.iSize
        }

        Button {
            id: anuluj
            anchors.horizontalCenter: parent.horizontalCenter
            highlighted: true

            anchors.bottom: parent.bottom

            width: (parent.width / 5) - 4
            height: (parent.height / 5) - 4

            text: "Anuluj"
            anchors.leftMargin: 7
            anchors.rightMargin: 7
            anchors.bottomMargin: 15
            font.pixelSize: template.bSize
            onClicked: {
                tlo.visible = false
                ddialog_e.close()
            }
        }
    }

    Popup {

        id: ddialog_calc
        z: 200

        closePolicy: Popup.NoAutoClose
        width: window.width / 1.5
        height: window.height / 1.5
        x: (window.width - width) / 2
        y: (window.height - height) / 2

        //          background: Rectangle {
        //                border.color: "blue"
        //                border.width: 3
        //              anchors.fill: parent
        //          }
        Column {
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.verticalCenter: parent.verticalCenter
            spacing: 5
            TextField {
                id: nilo
                placeholderText: qsTr("ILOŚĆ")
                width: grid.width
                height: ddialog_calc.height / 6
                font.bold: true

                readOnly: true
                font.pixelSize: template.bSize
                verticalAlignment: Qt.AlignVCenter
                horizontalAlignment: Qt.AlignHCenter
                background: Rectangle {
                    implicitWidth: 200
                    implicitHeight: 40
                    color: "white"
                    border.color: "black"
                }
            }

            Grid {
                id: grid
                rows: 4
                columns: 3
                spacing: 6

                //               anchors.fill: ddialog_calc
                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "7"
                    font.pixelSize: template.bSize
                    font.bold: true
                    onClicked: {
                        if (nilo.text.length < 4) {
                            nilo.text = nilo.text + text
                        }
                    }
                }
                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "8"
                    font.pixelSize: template.bSize
                    font.bold: true
                    onClicked: {
                        if (nilo.text.length < 4) {
                            nilo.text = nilo.text + text
                        }
                    }
                }
                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "9"
                    font.pixelSize: template.bSize
                    font.bold: true
                    onClicked: {
                        if (nilo.text.length < 4) {
                            nilo.text = nilo.text + text
                        }
                    }
                }

                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "4"
                    font.pixelSize: template.bSize
                    font.bold: true
                    onClicked: {
                        if (nilo.text.length < 4) {
                            nilo.text = nilo.text + text
                        }
                    }
                }
                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "5"
                    font.pixelSize: template.bSize
                    font.bold: true
                    onClicked: {
                        if (nilo.text.length < 4) {
                            nilo.text = nilo.text + text
                        }
                    }
                }
                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "6"
                    font.pixelSize: template.bSize
                    font.bold: true
                    onClicked: {
                        if (nilo.text.length < 4) {
                            nilo.text = nilo.text + text
                        }
                    }
                }
                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "1"
                    font.pixelSize: template.bSize
                    font.bold: true
                    onClicked: {
                        if (nilo.text.length < 4) {
                            nilo.text = nilo.text + text
                        }
                    }
                }
                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "2"
                    font.pixelSize: template.bSize
                    font.bold: true
                    onClicked: {
                        if (nilo.text.length < 4) {
                            nilo.text = nilo.text + text
                        }
                    }
                }
                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "3"
                    font.pixelSize: template.bSize
                    font.bold: true
                    onClicked: {
                        if (nilo.text.length < 4) {
                            nilo.text = nilo.text + text
                        }
                    }
                }
                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "OK"
                    font.pixelSize: template.bSize
                    font.bold: true
                    onClicked: {

                        tlo.visible = false
                        ddialog_calc.close()
                        ilosc.currentIndex = parseInt(nilo.text)
                        nilo.text = ""
                    }
                    highlighted: true
                }

                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "0"
                    font.pixelSize: template.bSize
                    font.bold: true
                    onClicked: {
                        if (text == "0" && nilo.text.length < 1) {
                            return
                        }
                        if (nilo.text.length < 4) {
                            nilo.text = nilo.text + text
                        }
                    }
                }

                Button {
                    width: ddialog_calc.width / 6
                    height: ddialog_calc.height / 6
                    text: "<-"
                    font.pixelSize: template.bSize
                    onClicked: {
                        nilo.text = nilo.text.substring(0, nilo.text.length - 1)
                    }
                }
            }
        }
    }

    Popup {

        id: ddialog_i
        z: 200

        closePolicy: Popup.NoAutoClose
        width: window.width / 2.5
        height: window.height / 2.5
        x: (window.width - width) / 2
        y: (window.height - height) / 2

        //          background: Rectangle {
        //                border.color: "green"
        //                border.width: 3
        //              anchors.fill: parent

        //          }
        Label {
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.verticalCenter: parent.verticalCenter
            anchors.top: parent.top
            anchors.margins: 5

            id: label_i
            text: "Potwierdzam dodanie REKORDU"
            color: "blue"
            font.bold: true
            font.pixelSize: template.iSize
        }

        Button {
            id: ok
            anchors.horizontalCenter: parent.horizontalCenter
            highlighted: true
            anchors.leftMargin: 7
            anchors.rightMargin: 7
            anchors.bottomMargin: 15

            anchors.bottom: parent.bottom

            width: (parent.width / 5) - 4
            height: (parent.height / 5) - 4

            text: "OK"

            font.pixelSize: template.bSize

            onClicked: {
                tlo.visible = false
                ddialog_i.close()
            }
        }
    }




    Popup {

        id: ddialog_u
        z: 200

        closePolicy: Popup.NoAutoClose
        width: window.width / 2.5
        height: window.height / 2.5
        x: (window.width - width) / 2
        y: (window.height - height) / 2

        //          background: Rectangle {
        //                border.color: "green"
        //                border.width: 3
        //              anchors.fill: parent

        //          }
        Label {
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.verticalCenter: parent.verticalCenter
            anchors.top: parent.top
            anchors.margins: 5

            id: label_u
            text: "Potwierdzam usunięcie REKORDU"
            color: "blue"
            font.bold: true
            font.pixelSize: template.iSize
        }

        Button {
            id: ok_u
            anchors.horizontalCenter: parent.horizontalCenter
            highlighted: true
            anchors.leftMargin: 7
            anchors.rightMargin: 7
            anchors.bottomMargin: 15

            anchors.bottom: parent.bottom

            width: (parent.width / 5) - 4
            height: (parent.height / 5) - 4

            text: "OK"

            font.pixelSize: template.bSize

            onClicked: {
                tlo.visible = false
                ddialog_u.close()
            }
        }
    }


    Popup {

        id: ddialog
        z: 200

        closePolicy: Popup.NoAutoClose
        width: window.width / 2.5
        height: window.height / 2.5
        x: (window.width - width) / 2
        y: (window.height - height) / 2

        //          background: Rectangle {
        //                border.color: "blue"
        //                border.width: 3
        //              anchors.fill: parent

        //          }
        Label {
            anchors.horizontalCenter: parent.horizontalCenter
            //anchors. verticalCenter:  parent.verticalCenter
            anchors.top: parent.top
            anchors.margins: 5

            id: label
            text: "Czy napewno usunąć REKORD?"

            color: "blue"
            font.bold: true
            font.pixelSize: template.iSize
        }

        Button {
            id: yes
            highlighted: true

            anchors.left: parent.left
            //               anchors.right:no.left
            anchors.bottom: parent.bottom

            anchors.leftMargin: 7
            anchors.rightMargin: 7
            anchors.bottomMargin: 15
            width: (parent.width / 5) - 4
            height: (parent.height / 5) - 4
            text: "TAK"

            font.pixelSize: template.bSize
            onClicked: {

                window.template.sendRemoveE(next.indexToRemove, next.idToRemove)
                next.indexToRemove = -100
                next.idToRemove = -100
                tlo.visible = false
                ddialog.close()
            }
        }
        Button {
            id: no
            anchors.right: parent.right
            //            anchors.left:yes.right
            anchors.bottom: parent.bottom
            anchors.leftMargin: 7
            anchors.rightMargin: 7
            anchors.bottomMargin: 15
            width: (parent.width / 5) - 4
            height: (parent.height / 5) - 4
            text: "Nie"

            font.pixelSize: template.bSize
            onClicked: {
                next.indexToRemove = -100
                next.idToRemove = -100
                tlo.visible = false
                ddialog.close()
            }
        }
    }
    Rectangle {
        visible: false
        id: tlo
        z: 150
        width: window.width
        height: window.height

        color: "#90a9a9a9"
        MouseArea {
            anchors.fill: parent
            onClicked: {
                console.log("nic nie zrobi")
            }
        }
        //         FastBlur {
        //             anchors.fill: parent
        //             source: window
        //             radius: 32
        //         }
    }




    Popup {

        id: ddialog_ie
        z: 200

        closePolicy: Popup.NoAutoClose
        width: window.width / 2.5
        height: window.height / 2.5
        x: (window.width - width) / 2
        y: (window.height - height) / 2

        //          background: Rectangle {
        //                border.color: "green"
        //                border.width: 3
        //              anchors.fill: parent

        //          }
        Label {
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.verticalCenter: parent.verticalCenter
            anchors.top: parent.top
            anchors.margins: 5

            id: label_i_ie
            text: "REKORD NIE ZOSTAŁ DODANY\nBŁĄD POŁĄCZENIA Z SERWEREM\nKONIECZNY RESTART APLIKACJI!"
            color: "red"
            font.bold: true
            font.pixelSize: template.iSize
        }

        Button {
            id: ok_ie
            anchors.horizontalCenter: parent.horizontalCenter
            highlighted: true
            anchors.leftMargin: 7
            anchors.rightMargin: 7
            anchors.bottomMargin: 15

            anchors.bottom: parent.bottom

            width: (parent.width / 5) - 4
            height: (parent.height / 5) - 4

            text: "OK"

            font.pixelSize: template.bSize

            onClicked: {
//                tlo.visible = false
//                ddialog_ie.close()
                window.close()
            }
        }
    }
    Popup {

        id: ddialog_ied
        z: 200

        closePolicy: Popup.NoAutoClose
        width: window.width / 2.5
        height: window.height / 2.5
        x: (window.width - width) / 2
        y: (window.height - height) / 2

        //          background: Rectangle {
        //                border.color: "green"
        //                border.width: 3
        //              anchors.fill: parent

        //          }
        Label {
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.verticalCenter: parent.verticalCenter
            anchors.top: parent.top
            anchors.margins: 5

            id: label_i_ied
            text: "REKORD NIE ZOSTAŁ USUNIĘTY\nBŁĄD POŁĄCZENIA Z SERWEREM\nKONIECZNY RESTART APLIKACJI!"
            color: "red"
            font.bold: true
            font.pixelSize: template.iSize
        }

        Button {
            id: ok_ied
            anchors.horizontalCenter: parent.horizontalCenter
            highlighted: true
            anchors.leftMargin: 7
            anchors.rightMargin: 7
            anchors.bottomMargin: 15

            anchors.bottom: parent.bottom

            width: (parent.width / 5) - 4
            height: (parent.height / 5) - 4

            text: "OK"

            font.pixelSize: template.bSize

            onClicked: {
//                tlo.visible = false
//                ddialog_ied.close()
                window.close()
            }
        }
    }

    Popup {

        id: ddialog_critical
        z: 200

        closePolicy: Popup.NoAutoClose
        width: window.width
        height: window.height
        x: window.x
        y: window.y

        //          background: Rectangle {
        //                border.color: "green"
        //                border.width: 3
        //              anchors.fill: parent

        //          }
        Label {
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.verticalCenter: parent.verticalCenter
            anchors.top: parent.top
            anchors.margins: 5

            id: label_critical
            text: "!!! NIE MA POŁĄCZENIA Z SERWEREM\nKONIECZNY RESTART APLIKACJI !!!"
            color: "red"
            font.bold: true
            font.pixelSize: template.iSize
        }

        Button {
            id: ok_critical
            anchors.horizontalCenter: parent.horizontalCenter
            highlighted: true
            anchors.leftMargin: 7
            anchors.rightMargin: 7
            anchors.bottomMargin: 15

            anchors.bottom: parent.bottom

            width: (parent.width / 5) - 4
            height: (parent.height / 5) - 4

            text: "OK"

            font.pixelSize: template.bSize

            onClicked: {
                window.close()
            }
        }
    }
//    Popup {

//        id: settingsPopup
//        z: 200

//        closePolicy: Popup.NoAutoClose
//        width: window.width / 2.5
//        height: window.height / 2.5
//        x: (window.width - width) / 2
//        y: (window.height - height) / 2

//        Label {
//            text: "Ustawienia Stylów: "
//            font.bold: true
//            anchors.horizontalCenter: parent.horizontalCenter
//            //anchors. verticalCenter:  parent.verticalCenter
//            anchors.top: parent.top
//            anchors.margins: 5
//            font.pixelSize: template.iSize
//        }

//        Label {
//            anchors.verticalCenter: parent.verticalCenter
//            anchors.right: styleBox.left
//            anchors.margins: 5
//            font.pixelSize: template.iSize

//            text: "Style:"
//        }
//        ComboBox {
//            anchors.margins: 2
//            id: styleBox
//            //                     width: (parent.width/5)-4
//            property int styleIndex: -1
//            font.pixelSize: template.cSize
//            model: ["Default", "Material", "Universal", "Fusion", "Imagine"]
//            Component.onCompleted: {
//                styleIndex = find(settings.style, Qt.MatchFixedString)
//                if (styleIndex !== -1)
//                    currentIndex = styleIndex
//            }

//            popup: Popup {
//                z: 210
//                y: styleBox.height - 1
//                width: styleBox.width
//                implicitHeight: contentItem.implicitHeight
//                padding: 1
//                font.pixelSize: template.cSize - 2
//                contentItem: ListView {
//                    clip: true
//                    implicitHeight: contentHeight
//                    model: styleBox.popup.visible ? styleBox.delegateModel : null

//                    ScrollIndicator.vertical: ScrollIndicator {
//                    }
//                }

//                background: Rectangle {
//                    border.color: "black"
//                    radius: 2
//                }
//            }
//            //                     Layout.fillWidth: true
//            anchors.horizontalCenter: parent.horizontalCenter
//            anchors.verticalCenter: parent.verticalCenter
//        }

//        Label {
//            anchors.horizontalCenter: parent.horizontalCenter
//            anchors.top: styleBox.bottom

//            text: "Konieczne ponowne uruchmienie aplikacji"
//            color: "#e41e25"
//            opacity: styleBox.currentIndex !== styleBox.styleIndex ? 1.0 : 0.0
//            font.pixelSize: template.iSize
//        }

//        Button {
//            id: okButton
//            font.pixelSize: template.bSize
//            highlighted: true
//            anchors.left: parent.left
//            anchors.bottom: parent.bottom
//            anchors.leftMargin: 7
//            anchors.rightMargin: 7
//            anchors.bottomMargin: 15
//            width: (parent.width / 5) - 4
//            height: (parent.height / 5) - 4
//            text: "Ustaw"
//            onClicked: {
//                settings.style = styleBox.displayText
//                tlo.visible = false
//                settingsPopup.close()
//            }
//        }

//        Button {
//            id: cancelButton

//            font.pixelSize: template.bSize
//            anchors.right: parent.right
//            //            anchors.left:yes.right
//            anchors.bottom: parent.bottom
//            anchors.leftMargin: 7
//            anchors.rightMargin: 7
//            anchors.bottomMargin: 15

//            width: (parent.width / 5) - 4
//            height: (parent.height / 5) - 4
//            text: "Anuluj"
//            onClicked: {
//                styleBox.currentIndex = styleBox.styleIndex
//                tlo.visible = false
//                settingsPopup.close()
//            }
//        }
//    }

    Connections {
        target: window.template
        onReturnStatusOn: {
            if (stat == 1) {
                                            next.indexToRemove = -100
                                            next.idToRemove = -100

                ddialog.close()

                ddialog_ie.close()
                ddialog_ied.close()
                ddialog_u.close()
                ddialog_calc.close()
                                            tlo.visible = true
                ddialog_i.open()


            } else if (stat == 0) {
                                            next.indexToRemove = -100
                                            next.idToRemove = -100


                ddialog.close()
                ddialog_i.close()
                ddialog_ied.close()
                ddialog_u.close()
                ddialog_calc.close()
                                            tlo.visible = true
                   ddialog_ie.open()

            } else if (stat == -1) {
                                            next.indexToRemove = -100
                                            next.idToRemove = -100

//                ddialog_ied.open()
                ddialog.close()
                ddialog_i.close()
                ddialog_ie.close()
                ddialog_u.close()
                ddialog_calc.close()
                                             tlo.visible = true
                ddialog_ied.open()

            } else if (stat == -2) {
                next.indexToRemove = -100
                next.idToRemove = -100

                ddialog.close()
                ddialog_i.close()
                ddialog_ie.close()
                ddialog_ied.close()
                 ddialog_calc.close()
                tlo.visible = true
                ddialog_u.open()

           } else if (stat > 1000) {


                if ((stat -1000) >= 100) {
                    stan.color = "lightgreen"
                } else if ((stat -1000) < 100) {
                    stan.color = "indianred"
                }
                percent.text = (stat-1000).toString() +"%"

            } else if (stat == -100) {
                percent.text="__%"; stan.color = "lightskyblue"

            } else if (stat == -1000) {
                ddialog.close()
                ddialog_i.close()
                ddialog_ie.close()
                ddialog_ied.close()
                 ddialog_calc.close()
                ddialog_u.close()
                ddialog_critical.open()


            }


        }

        }
} //}

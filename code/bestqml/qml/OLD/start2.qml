import QtQuick 2.6
import QtQuick.Controls 2.2
import QtQuick.Controls.Material 2.2
import QtQuick.Layouts 1.3
import QtQuick.Controls.Styles 1.4
import CustomQmlTypes 1.0
import QtQml.Models 2.1

ApplicationWindow {
    property BridgeQML template: BridgeQML {
    }
    id: window
//    width: 1280
//    height: 720
//    color: "transparent"
  visibility: "FullScreen"
    visible: true
    //    Rectangle  {
    //        id: test
    //flags: Qt.WindowStaysOnTopHint


    //        anchors.fill: parent
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
        z: 100
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
            Text {

                anchors.horizontalCenter: parent.horizontalCenter
                anchors.verticalCenter: parent.verticalCenter
                text: "LOGO"
                font.family: "Helvetica"
                font.pixelSize: (logo.width / 5)
                color: "black"
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
            color: "green"
            // border.color: "red"
            // border.width: 1
            // radius: 10
            Text {

                anchors.horizontalCenter: parent.horizontalCenter
                anchors.verticalCenter: parent.verticalCenter
                text: "100%"
                font.family: "Helvetica"
                font.pixelSize: (logo.width / 5)
                //color: "black"
                color: "red"
            }
        }
    }

    Rectangle {
        id: body
        anchors.top: header.bottom
        height: (parent.height / 10 * 1)
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
            width: (parent.width/5)-4
           height:parent.height
            anchors.top: parent.top
            anchors.left: parent.left
            model: GrupyModel
            textRole: "grupaName"
            displayText: "Grupa: " + currentText
            font.bold: true
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
            currentIndex: 0
            id: stany
            width: (parent.width/5)-4
            height:parent.height
            anchors.top: parent.top
            anchors.left: grupa.right
            font.bold: true
            model: ListModel {
                id: listModel
                Component.onCompleted: {
                    for (var i = 0; i < 800; i++)
                        listModel.append({
                                             Name: "" + i / 2
                                         })
                    //listModel.append({"Name": "Item " + i})
                }
            }
            displayText: "Stan: " + currentIndex / 2
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
            width: (parent.width/5)-4
            height:parent.height
            anchors.top: parent.top
            anchors.left: stany.right
            font.bold: true
            model: ProduktyModel
            textRole: "produktName"
            displayText: "Produkt: " + currentText
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
            width: (parent.width/5)-4
            height:parent.height
            anchors.top: parent.top
            anchors.left: produkt.right
            model: 8000
            font.bold: true
            displayText: "Ilość: " + currentText
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
                implicitHeight: contentItem.implicitHeight
                padding: 1

                contentItem: ListView {
                    clip: true
                    implicitHeight: contentHeight
                    model: ilosc.popup.visible ? ilosc.delegateModel : null
                    currentIndex: ilosc.highlightedIndex

                    ScrollIndicator.vertical: ScrollIndicator {
                    }
                }

                background: Rectangle {
                    border.color: "black"
                    radius: 2
                }
            }
        }
        Button {
            anchors.margins: 2

            id: zatwierdz
            width: (parent.width/5)-4
            height:parent.height
            text: qsTr("Zatwierdz")
            anchors.top: parent.top
            anchors.left: ilosc.right
            font.bold: true
            contentItem: Text {
                text: zatwierdz.text
                font: zatwierdz.font
                opacity: enabled ? 1.0 : 0.3
                color: zatwierdz.down ? "blue" : "red"
                horizontalAlignment: Text.AlignHCenter
                verticalAlignment: Text.AlignVCenter
                elide: Text.ElideRight
            }

//            background: Rectangle {
//                implicitWidth: 100
//                implicitHeight: 40
//                opacity: enabled ? 1 : 0.3
//                border.color: zatwierdz.down ? "#17a81a" : "#21be2b"
//                border.width: 1
//                radius: 2
//            }
            background: Rectangle {
                implicitWidth: 100
                implicitHeight: 25
                border.width: zatwierdz.activeFocus ? 2 : 1
                border.color: "#888"
                radius: 4
                gradient: Gradient {
                    GradientStop { position: 0 ; color: zatwierdz.pressed ? "#ccc" : "#eee" }
                    GradientStop { position: 1 ; color: zatwierdz.pressed ? "#aaa" : "#ccc" }
                }
            }
            onClicked: {
//                console.log("Grupy current GrupaId Name: ", grupa.model.data(grupa.model.index(grupa.currentIndex, 0), Qt.UserRole + 1)) //P name
//                console.log("Grupy current GrupaId Id: ", grupa.model.data(grupa.model.index(grupa.currentIndex, 0), Qt.UserRole + 2)) //P W

//                console.log("Stan osobowy:",stany.currentIndex/2)

//                console.log("Prodkty current Produkt Name: ", produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 1)) //P name
//                console.log("Prodkty current Produkt W: ", produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 2)) //P W
//                console.log("Prodkty current Produkt Id: ", produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 3)) //P W
//                console.log("Ilosc:",ilosc.currentIndex)

                window.template.sendZatwierdz(
                            grupa.model.data(grupa.model.index(grupa.currentIndex, 0), Qt.UserRole + 1),
                                              grupa.model.data(grupa.model.index(grupa.currentIndex, 0), Qt.UserRole + 2),
                                              (stany.currentIndex/2),
                                              produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 1),
                                              produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 2),
                                              produkt.model.data(produkt.model.index(produkt.currentIndex, 0), Qt.UserRole + 3),
                                              ilosc.currentIndex
                                              )


            }

            //width: 3* zatwierdz.height
        }
        //        ComboBox {
        //            id:sztuki
        //            anchors.top: parent.top
        //            anchors.left: parent.left
        //            model: 1000
        //            //model: ["First", "Second", "Third","First", "Second", "Third","First", "Second", "Third","First", "Second", "Third","First", "Second", "Third","First", "Second", "Third","First", "Second", "Third","First", "Second", "Third"]
        ////            width: Math.max(implicitWidth, Math.min(implicitWidth * 2, page.availableWidth / 3))
        //            anchors.horizontalCenter: parent.horizontalCenter
        //        }
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
    property int indexToRemove: -1

    anchors.top: parent.top
      anchors.bottom: parent.bottom
    anchors.left: parent.left
    anchors.right:  parent.right
    currentIndex :-1
    anchors.margins: 2
    width: parent.width/5
    rotation: 180



    model: DelegateModel {
        id: myDelegateModel
        model: ElementyModel
        delegate:  ItemDelegate {

        Rectangle {

            id: rmarg
            border.color :"white"
            border.width:  1
            anchors.fill: parent
        }

        rotation: 180
        width: parent.width
        height: window.height/10
        //text: model.title

        Text {
            padding: height/8
            anchors.margins: 2
            width: (parent.width/5)-4
            height:parent.height
            anchors.top: parent.top
            anchors.left: parent.left

            font.bold: true

            id: tx1
            text: model.elementGName
            font.family: "Helvetica"
            //font.pointSize: 24
            //color: "green"
        }
        Text {
            padding: height/8
            anchors.margins: 2
            width: (parent.width/5)-4
            height:parent.height
            anchors.top: parent.top
            anchors.left: tx1.right

            font.bold: true

            id: tx2
            text: model.elementStan
            font.family: "Helvetica"
            //font.pointSize: 24
            //color: "green"
        }

        Text {
            padding: height/8
            anchors.margins: 2
            width: (parent.width/5)-4
            height:parent.height
            anchors.top: parent.top
            anchors.left: tx2.right

            font.bold: true

            id: tx3
            text: model.elementPName
            font.family: "Helvetica"
            //font.pointSize: 24
            //color: "green"
        }
        Text {
            padding: height/8
            anchors.margins: 2
            width: (parent.width/5)-4
            height:parent.height
            anchors.top: parent.top
            anchors.left: tx3.right

            font.bold: true

            id: tx4
            text: model.elementIlosc
            font.family: "Helvetica"
            //font.pointSize: 24
           //color: "green"
        }
        Button {
            z: 50
            anchors.margins: 2
            id: usun
            width: (parent.width/5)-4
            height:parent.height-(parent.height/10)
            anchors.top: parent.top
            anchors.left: tx4.right
            font.bold: true
            //text: qsTr("Usuń\n2019-01-01 10:01")
            Text {
                            anchors.horizontalCenter: parent.horizontalCenter
                            anchors.verticalCenter:  parent.verticalCenter
                //padding: height/8
                anchors.margins: 2


                font.bold: true

                id: txusun
                text: "Usuń\n"+model.elementData
                font.family: "Helvetica"
                //font.pointSize: 24
               //color: "green"
            }
            onClicked: {
                selectremove()

                //ddialog.open()
//                window.template.sendRemoveE(model.currentIndex)

                    //next.currentIndex = index
            }


        }

        function selectremove(){

            model.currentIndex = index
            next.indexToRemove = index
            rmarg.border.color ="red"
            rmarg.border.width= 1
            ddialog.open()






        }
//        function selectremoveYes(){

//window.template.sendRemoveE(model.currentIndex)

//        }




    }
   }
    ScrollIndicator.vertical: ScrollIndicator {
    }
    onContentHeightChanged: {

//                currentIndex = count - 1

        positionViewAtEnd()
    }
   function selectremovNo() {
            // Update the currently-selected item
//            currentSelectedItem = myDelegateModel.items.get(next.indexToRemove).model.display

            // Log the Display Role
            console.log(myDelegateModel.get(next.indexToRemove).attributes.get(1))
    }


}


    }
    Dialog {
          id: ddialog
          //closePolicy: Popup.NoAutoClose
          width: window.width/4
          title: "Usuwanie Elemetu"
          //closePolicy: Popup.CloseOnEscape
          height:window.height/4
          x: (window.width - width) / 2
          y: (window.height - height) / 2



          Rectangle {

              anchors.fill: parent
              id: col1
              Label {
                    anchors.horizontalCenter: parent.horizontalCenter
                  //anchors. verticalCenter:  parent.verticalCenter
                  anchors.margins: 5
                   id:label
                  text: "Czy napewno usunąć element?"
                  color: "blue"
                  font.bold: true
              }
              Button {

                  highlighted: true
                   anchors.top: label.bottom
                   anchors.left:parent.left
                   anchors.right:no.left
                   anchors.margins: 5
                  text: "TAK"
                  id: yes
                  onClicked: {
                      window.template.sendRemoveE(next.indexToRemove)
                      //window.next.delegate.selectremoveYes()
//                      window.template.afterWake_afterHalt(1)

                      ddialog.close()
//                      listView.visible = true
                  }
              }
              Button {
                  anchors.top: label.bottom
                anchors.right: parent.right
                 anchors.left:yes.right
                 anchors.margins: 5

                  text: "Nie"
                  id: no
                  onClicked: {
                      next.selectremovNo()
//                     next.delegate.selectremoveNo()
//                      next.model.children[next.indexToRemove].selectremoveNo()
//                      window.next.model.selectremoveNo()
//                      window.template.afterWake_afterHalt(1)
                      ddialog.close()
//                      listView.visible = true
                  }
              }
          }
      }


} //}

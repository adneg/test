import QtQuick 2.6
import QtQuick.Controls 2.2

import QtQuick.Layouts 1.3
import QtQuick.Controls.Styles 1.4
//import CustomQmlTypes 1.0

//import QtGraphicalEffects 1.0
//Default, Fusion, Imagine, Universal, Material
import QtQuick.Controls.Material 2.2
//import Qt.labs.settings 1.0

ApplicationWindow {

    id: window
    visibility: "FullScreen"
    visible: true
	
	        Label {
            id:lberror1
           text: "BRAK PLIKU KONFIGURACYJNEGO: \"settings.ini\""
           font.bold: true
           anchors.horizontalCenter: parent.horizontalCenter
//           anchors.verticalCenter:  parent.verticalCenter
           anchors.bottom: berror.top
           anchors.margins: 10
           font.pixelSize: (berror.width / 16)
       }
	
	        Button {
            id: berror
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.verticalCenter:  parent.verticalCenter
            highlighted: true

            //anchors.top: lberror1.bottom

            width: (parent.width / 5) - 4
            height: (parent.height / 5) - 4

            text: "Anuluj"
            anchors.leftMargin: 7
            anchors.rightMargin: 7
            anchors.bottomMargin: 15
            font.pixelSize: (berror.width / 16)
            onClicked: {
				window.close()
            }
        }

} 

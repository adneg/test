package main

import (

	//"fmt"
	"os"
	//"os/exec"

	//"time"
	"ui"

	//	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	//	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	widgets.NewQApplication(len(os.Args), os.Args)

	pilot := ui.IntPilot()

	pilot.Start()
	widgets.QApplication_Exec()

}

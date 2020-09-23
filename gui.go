package main

// import (
// 	"fmt"
// 	"image/color"

// 	"fyne.io/fyne"
// 	"fyne.io/fyne/app"
// 	"fyne.io/fyne/canvas"
// 	"fyne.io/fyne/dialog"
// 	"fyne.io/fyne/layout"
// 	"fyne.io/fyne/theme"
// 	"fyne.io/fyne/widget"
// )

// type GUI struct {
// 	daemon     *daemon
// 	app        fyne.App
// 	mainWindow *fyne.Window
// }

// func NewGUI(daemon *daemon) (*GUI, error) {
// 	a := app.New()
// 	return &GUI{
// 		daemon: daemon,
// 		app:    a,
// 	}, nil
// }

// func (g *GUI) ShowAndRun() {
// 	g.showPrimaryScene()
// 	g.app.Run()
// }

// func (g *GUI) showPrimaryScene() {
// 	logln("showing primary scene")
// 	menuItem := fyne.NewMenuItem("Item1", func() { fmt.Println("menu item 1") })

// 	preferencesItem := fyne.NewMenuItem("Preferences", func() { fmt.Println("settings") })

// 	mainWindow := g.app.NewWindow("golinks daemon")
// 	mainMenu := fyne.NewMainMenu(
// 		fyne.NewMenu("File", menuItem),
// 		fyne.NewMenu("Edit", preferencesItem),
// 	)

// 	mainWindow.SetMainMenu(mainMenu)

// 	tabs := widget.NewTabContainer(
// 		widget.NewTabItemWithIcon("Home", theme.HomeIcon(), g.homeScene()),
// 		widget.NewTabItemWithIcon("Workers", theme.ComputerIcon(), g.workersScene()),
// 	)

// 	tabs.SetTabLocation(widget.TabLocationLeading)
// 	tabs.SelectTabIndex(0)
// 	mainWindow.SetContent(tabs)
// 	mainWindow.Resize(fyne.NewSize(800, 600))

// 	g.mainWindow = &mainWindow
// 	mainWindow.Show()
// }

// func (g *GUI) homeScene() fyne.CanvasObject {
// 	vbox := widget.NewVBox(
// 		widget.NewLabelWithStyle("Home screen", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
// 		widget.NewHBox(layout.NewSpacer()),
// 	)
// 	return vbox
// }

// func (g *GUI) workersScene() fyne.CanvasObject {
// 	var workerItems []fyne.CanvasObject
// 	for _, worker := range g.daemon.workerManager.WorkerConfig.Workers {
// 		workerItems = append(workerItems, widget.NewButton(worker.RootPath, nil))
// 	}

// 	circleSpacer := canvas.NewCircle(color.White)
// 	circleSpacer.StrokeWidth = 3
// 	return widget.NewVBox(
// 		widget.NewLabelWithStyle("Workers", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
// 		widget.NewScrollContainer(widget.NewVBox(workerItems...)),
// 		circleSpacer,
// 		widget.NewButtonWithIcon("Add Worker", theme.ContentAddIcon(), g.addWorkerAction),
// 	)
// }

// func (g *GUI) makeWorkerListEntry(worker *Worker) fyne.CanvasObject {
// 	browserButton := widget.NewButton("Open", func() {
// 		dialog.ShowFileOpen(func(closer fyne.URIReadCloser, err error) {
// 			if err != nil {
// 				errln("worker root file opener error", err)
// 				return
// 			}

// 			if closer != nil {
// 				logln("closer Name", closer.Name())
// 			}
// 		}, *g.mainWindow)
// 	})

// 	form := &widget.Form{
// 		Items: []*widget.FormItem{
// 			{Text: "Root", Widget: browserButton},
// 		},
// 	}

// 	return form
// }

// func (g *GUI) addWorkerAction() {
// 	//TODO directory selection not yet supported https://github.com/fyne-io/fyne/issues/941
// 	//TODO hidden directories not visible https://github.com/fyne-io/fyne/issues/1278
// 	fi := dialog.NewFileOpen(func(uri fyne.URIReadCloser, err error) {
// 		if err != nil {
// 			errln("failed to get worker diretory", err)
// 			return
// 		}
// 		if uri == nil {
// 			logln("add worker canceled")
// 		}
// 		logln("uri", uri.URI().String())
// 		logln("name", uri.Name())
// 		worker := &Worker{
// 			daemon:           g.daemon,
// 			RootPath:         uri.URI().String(),
// 			GenerationPeriod: 30000,
// 		}
// 		g.daemon.workerManager.WorkerConfig.Workers = append(g.daemon.workerManager.WorkerConfig.Workers, worker)
// 		g.daemon.workerManager.startNewWorkers()
// 	}, *g.mainWindow)
// 	fi.Show()
// }

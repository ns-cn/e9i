package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/faiface/beep/speaker"
	"github.com/ns-cn/e9i/command"
	"github.com/ns-cn/e9i/loop"
	"github.com/ns-cn/e9i/source"
	"log"
	"math/rand"

	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var application fyne.Window

func createApp() {
	myApp := app.New()
	myApp.Settings().SetTheme(&MyTheme{})
	application = myApp.NewWindow("E9I, like 163")
}

var tabs *container.AppTabs
var input *widget.Entry

func Ui() fyne.CanvasObject {

	//toolbar := widget.NewToolbar(
	//	widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {}),
	//	widget.NewToolbarSeparator(),
	//	widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
	//	widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
	//	widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
	//	widget.NewToolbarSpacer(),
	//	widget.NewToolbarAction(theme.HelpIcon(), func() {
	//		tabs.SelectIndex(1)
	//	}),
	//)

	tabs = container.NewAppTabs(
		UiSearchTab(),
		UiStarTab(),
		UiPlayListTab(),
		UiInfoTab(),
	)

	loopTitle := widget.NewButton(loop.Title[loopType], func() {
		loopType = (loopType + 1) % 4
		RefreshSearchList()
	})
	playPre := widget.NewButton("上一首", func() {
		refreshNextIndex(-2)
		actionChan <- command.PLAYNEXT
	})
	pauseTitle := ""
	if !paused && OnProcessing {
		pauseTitle = "暂停"
	} else {
		pauseTitle = "播放"
	}
	pause := widget.NewButton(pauseTitle, func() {
		if !OnProcessing {
			size := len(playlist)
			current = (current + size - 1) % size
			actionChan <- command.PLAYNEXT
		} else {
			if paused {
				paused = false
				speaker.Unlock()
			} else {
				paused = true
				speaker.Lock()
			}
		}
		RefreshSearchList()
	})
	playNext := widget.NewButton("下一首", func() {
		actionChan <- command.PLAYNEXT
	})
	controlBar := container.NewGridWithColumns(4, loopTitle, playPre, pause, playNext)

	tabs.SetTabLocation(container.TabLocationTop)
	content := container.NewBorder(nil, controlBar, nil, nil, tabs)
	return content
}

func UiSearchTab() *container.TabItem {
	if input == nil {
		input = widget.NewEntry()
		input.SetPlaceHolder("Enter text...")
		input.Resize(fyne.NewSize(300, 300))
		input.OnSubmitted = func(text string) {
			songs := Search(text)
			playlist = playlist[:0]
			playlist = append(playlist, songs...)
			RefreshSearchList()
			log.Println("Search: ", input.Text)
		}
	}
	if input.Text == "" {
		input.SetText(source.Recommand[rand.Intn(len(source.Recommand))])
	}
	btnRandom := widget.NewButton("", func() {
		input.SetText(source.Recommand[rand.Intn(len(source.Recommand))])
	})
	btnRandom.SetIcon(theme.ViewRefreshIcon())
	btnSearch := widget.NewButton("搜索", func() {
		songs := Search(input.Text)
		playlist = playlist[:0]
		playlist = append(playlist, songs...)
		RefreshSearchList()
		log.Println("Search: ", input.Text)
	})
	btnSearch.SetIcon(theme.SearchIcon())
	toolBar := container.NewBorder(nil, nil, btnRandom, btnSearch, input)
	searchResult := widget.NewList(
		func() int {
			return len(playlist)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("template", func() {
			})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			button := o.(*widget.Button)
			if i == current && len(playlist) > current && playlist[current].ID == currentSong.ID {
				button.SetIcon(theme.ViewRefreshIcon())
			} else {
				button.SetIcon(nil)
			}
			song := playlist[i]
			button.SetText(fmt.Sprintf("%s(%s)", song.Name, song.GetArtistDisplay()))
			button.OnTapped = func() {
				current = i
				go Play(playlist[i])
				RefreshSearchList()
			}
		})
	itemContent := container.NewBorder(toolBar, nil, nil, nil, searchResult)
	searchItem := container.NewTabItemWithIcon("搜索", theme.SearchIcon(), itemContent)
	return searchItem
}

func UiStarTab() *container.TabItem {
	content := container.NewCenter(container.NewVBox(
		container.NewCenter(widget.NewLabel("E9I, faker of 163")),
		container.NewCenter(widget.NewLabel(fmt.Sprintf("Version: %s", source.Version))),
	))
	return container.NewTabItemWithIcon("收藏", theme.MediaMusicIcon(), content)
}

func UiPlayListTab() *container.TabItem {
	content := container.NewCenter(container.NewVBox(
		container.NewCenter(widget.NewLabel("E9I, faker of 163")),
		container.NewCenter(widget.NewLabel(fmt.Sprintf("Version: %s", source.Version))),
	))
	return container.NewTabItemWithIcon("当前", theme.MediaPlayIcon(), content)
}

func UiInfoTab() *container.TabItem {
	content := container.NewCenter(container.NewVBox(
		container.NewCenter(widget.NewLabel("E9I, faker of 163")),
		container.NewCenter(widget.NewLabel(fmt.Sprintf("Version: %s", source.Version))),
	))
	return container.NewTabItemWithIcon("关于", theme.InfoIcon(), content)
}

func RefreshSearchList() {
	application.SetContent(Ui())
	tabs.SelectIndex(0)
}

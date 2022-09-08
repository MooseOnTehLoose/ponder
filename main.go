package main

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/pkg/browser"
)

//generate a manifest and compile via rsrc -manifest <name>.manifest -o rsrc.syso
//then compile with go build -ldflags="-H windowsgui"

type MainWin struct {
	*walk.MainWindow
	url             *walk.TextEdit
	pasteBox        *walk.TextEdit
	formattedBox    *walk.TextEdit
	copyToClipboard *walk.PushButton
	openURL         *walk.PushButton
}

type Deck struct {
	Name              string          `json:"name"`
	DateUpdated       int             `json:"date_updated"`
	Slug              string          `json:"slug"`
	ResourceURI       string          `json:"resource_uri"`
	URL               string          `json:"url"`
	Inventory         [][]interface{} `json:"inventory"`
	User              string          `json:"user"`
	UserDisplay       string          `json:"user_display"`
	FeaturedCard      string          `json:"featured_card"`
	Score             int             `json:"score"`
	ThumbnailURL      string          `json:"thumbnail_url"`
	SmallThumbnailURL string          `json:"small_thumbnail_url"`
}

func main() {

	mainWin, _ := NewMainWin()

	mainWin.Run()

}

func NewMainWin() (*MainWin, error) {

	//Create a new struct instance
	mainWin := new(MainWin)

	//Call the create function on it
	err := MainWindow{
		AssignTo: &mainWin.MainWindow,
		Title:    "Ponder",
		MaxSize:  Size{Width: 900, Height: 600},
		MinSize:  Size{Width: 900, Height: 300},
		Layout:   VBox{},

		//Main Window
		Children: []Widget{

			//URL Paste Box and Open Browser Button
			HSplitter{
				MaxSize: Size{Height: 30},
				Children: []Widget{TextEdit{
					Background:         nil,
					Font:               Font{},
					MaxSize:            Size{Width: 300},
					MinSize:            Size{Width: 300},
					AlwaysConsumeSpace: true,
					Name:               "tappedout.net deck URL",
					OnBoundsChanged: func() {
					},
					OnKeyDown: func(key walk.Key) {
					},
					OnKeyPress: func(key walk.Key) {
					},
					OnKeyUp: func(key walk.Key) {
					},
					OnMouseDown: func(x int, y int, button walk.MouseButton) {
					},
					OnMouseMove: func(x int, y int, button walk.MouseButton) {
					},
					OnMouseUp: func(x int, y int, button walk.MouseButton) {
					},
					OnSizeChanged: func() {
					},
					Persistent:      false,
					ToolTipText:     "Paste your tappedout.net Deck URL Here!",
					Alignment:       0,
					Column:          0,
					ColumnSpan:      0,
					GraphicsEffects: []walk.WidgetGraphicsEffect{},
					Row:             0,
					RowSpan:         0,
					StretchFactor:   0,
					AssignTo:        &mainWin.url,
					CompactHeight:   false,
					HScroll:         false,
					MaxLength:       0,
					OnTextChanged: func() {
					},
					ReadOnly:      nil,
					Text:          nil,
					TextAlignment: 0,
					TextColor:     0,
					VScroll:       false,
				}, PushButton{
					Accessibility:      Accessibility{},
					Background:         nil,
					ContextMenuItems:   []MenuItem{},
					DoubleBuffering:    false,
					Enabled:            nil,
					Font:               Font{},
					MaxSize:            Size{Width: 300},
					MinSize:            Size{Width: 300},
					AlwaysConsumeSpace: true,
					Name:               "",
					OnBoundsChanged: func() {
					},
					OnKeyDown: func(key walk.Key) {
					},
					OnKeyPress: func(key walk.Key) {
					},
					OnKeyUp: func(key walk.Key) {
					},
					OnMouseDown: func(x int, y int, button walk.MouseButton) {
					},
					OnMouseMove: func(x int, y int, button walk.MouseButton) {
					},
					OnMouseUp: func(x int, y int, button walk.MouseButton) {
					},
					OnSizeChanged: func() {
					},
					Persistent:         false,
					RightToLeftReading: false,
					ToolTipText:        "Open the JSON Data for the given URL",
					Visible:            nil,
					Alignment:          0,
					Column:             0,
					ColumnSpan:         0,
					GraphicsEffects:    []walk.WidgetGraphicsEffect{},
					Row:                0,
					RowSpan:            0,
					StretchFactor:      0,
					Image:              nil,
					OnClicked: func() {
						browser.OpenURL(urlSwitch(mainWin.url.Text()))
					},
					Text:           "Fetch Deck",
					AssignTo:       &mainWin.openURL,
					ImageAboveText: false,
				}, PushButton{
					Accessibility:      Accessibility{},
					Background:         nil,
					ContextMenuItems:   []MenuItem{},
					DoubleBuffering:    false,
					Enabled:            nil,
					Font:               Font{},
					MaxSize:            Size{Width: 300},
					MinSize:            Size{Width: 300},
					AlwaysConsumeSpace: true,
					Name:               "",
					OnBoundsChanged: func() {
					},
					OnKeyDown: func(key walk.Key) {
					},
					OnKeyPress: func(key walk.Key) {
					},
					OnKeyUp: func(key walk.Key) {
					},
					OnMouseDown: func(x int, y int, button walk.MouseButton) {
					},
					OnMouseMove: func(x int, y int, button walk.MouseButton) {
					},
					OnMouseUp: func(x int, y int, button walk.MouseButton) {
					},
					OnSizeChanged: func() {
					},
					Persistent:         false,
					RightToLeftReading: false,
					ToolTipText:        "Grab the Formatted Decklist for TableTop",
					Visible:            nil,
					Alignment:          0,
					Column:             0,
					ColumnSpan:         0,
					GraphicsEffects:    []walk.WidgetGraphicsEffect{},
					Row:                0,
					RowSpan:            0,
					StretchFactor:      0,
					Image:              nil,
					OnClicked: func() {
						walk.Clipboard().SetText(mainWin.formattedBox.Text())
					},
					Text:           "Copy To ClipBoard",
					AssignTo:       &mainWin.copyToClipboard,
					ImageAboveText: false,
				}},
				HandleWidth: 30,
			},
			//Paste Deck JSON and Formatted Deck List
			HSplitter{
				Children: []Widget{
					TextEdit{
						AssignTo: &mainWin.pasteBox,
						VScroll:  true,
						OnTextChanged: func() {
							mainWin.formattedBox.SetText(convert(mainWin.pasteBox.Text()))
						},
					},
					TextEdit{
						AssignTo:    &mainWin.formattedBox,
						ReadOnly:    true,
						ToolTipText: "Formatted for ScryFall",
						VScroll:     true,
					},
				},
			},
		},
	}.Create()

	return mainWin, err

}

func convert(deck string) string {
	incoming := Deck{}
	err := json.Unmarshal([]byte(deck), &incoming)
	if err != nil {
		return err.Error()
	}

	cards := ""
	//1x Ancient Copper Dragon (CLB:396)
	//1x Ancient Tomb (TMP)

	for A := range incoming.Inventory {

		quantity := "1x "
		name := incoming.Inventory[A][0].(string)
		set := ""

		cardMap := incoming.Inventory[A][1].(map[string]interface{})

		if numVal, exists := cardMap["qty"]; exists {
			quantity = strconv.Itoa(int(numVal.(float64))) + "x "
		}
		if setVal, exists := cardMap["tla"]; exists {
			set = " (" + setVal.(string) + ")"
		}
		if varVal, exists := cardMap["variation"]; exists {
			set = " (" + cardMap["tla"].(string) + ") " + varVal.(string)
		}
		cards = cards + quantity + name + set + "\r\n"
	}

	return cards
}

func urlSwitch(url string) string {
	collectionURL := strings.ReplaceAll(url, "https://tappedout.net/mtg-decks/", "http://tappedout.net/api/collection/collection:deck/")
	return collectionURL
}

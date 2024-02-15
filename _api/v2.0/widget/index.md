---
layout: page
tags: [api]
title: Fyne API "widget"
---

# widget
---
```go
import "fyne.io/fyne/v2/widget"
```

Package widget defines the UI widgets within the Fyne toolkit

## Usage

#### func  ShowModalPopUp

```go
func ShowModalPopUp(content fyne.CanvasObject, canvas fyne.Canvas)
```
ShowModalPopUp creates a new popUp for the specified content and displays it on the passed canvas. A modal PopUp blocks interactions with underlying elements, covered with a semi-transparent overlay.

#### func  ShowPopUp

```go
func ShowPopUp(content fyne.CanvasObject, canvas fyne.Canvas)
```
ShowPopUp creates a new popUp for the specified content and displays it on the passed canvas.

#### func  ShowPopUpAtPosition

```go
func ShowPopUpAtPosition(content fyne.CanvasObject, canvas fyne.Canvas, pos fyne.Position)
```
ShowPopUpAtPosition creates a new popUp for the specified content at the specified absolute position. It will then display the popup on the passed canvas.

#### func  ShowPopUpMenuAtPosition

```go
func ShowPopUpMenuAtPosition(menu *fyne.Menu, c fyne.Canvas, pos fyne.Position)
```
ShowPopUpMenuAtPosition creates a PopUp menu populated with items from the passed menu structure. It will automatically be positioned at the provided location and shown as an overlay on the specified canvas.

#### types

 * [Accordion](accordion.html)
 * [AccordionItem](accordionitem.html)
 * [BaseWidget](basewidget.html)
 * [Button](button.html)
 * [ButtonAlign](buttonalign.html)
 * [ButtonIconPlacement](buttoniconplacement.html)
 * [ButtonImportance](buttonimportance.html)
 * [ButtonStyle](buttonstyle.html)
 * [Card](card.html)
 * [Check](check.html)
 * [CustomTextGridStyle](customtextgridstyle.html)
 * [DisableableWidget](disableablewidget.html)
 * [Entry](entry.html)
 * [FileIcon](fileicon.html)
 * [Form](form.html)
 * [FormItem](formitem.html)
 * [Hyperlink](hyperlink.html)
 * [Icon](icon.html)
 * [Label](label.html)
 * [List](list.html)
 * [ListItemID](listitemid.html)
 * [Menu](menu.html)
 * [Orientation](orientation.html)
 * [PopUp](popup.html)
 * [PopUpMenu](popupmenu.html)
 * [ProgressBar](progressbar.html)
 * [ProgressBarInfinite](progressbarinfinite.html)
 * [RadioGroup](radiogroup.html)
 * [Select](select.html)
 * [SelectEntry](selectentry.html)
 * [Separator](separator.html)
 * [Slider](slider.html)
 * [Table](table.html)
 * [TableCellID](tablecellid.html)
 * [TextGrid](textgrid.html)
 * [TextGridCell](textgridcell.html)
 * [TextGridRow](textgridrow.html)
 * [TextGridStyle](textgridstyle.html)
 * [Toolbar](toolbar.html)
 * [ToolbarAction](toolbaraction.html)
 * [ToolbarItem](toolbaritem.html)
 * [ToolbarSeparator](toolbarseparator.html)
 * [ToolbarSpacer](toolbarspacer.html)
 * [Tree](tree.html)
 * [TreeNodeID](treenodeid.html)

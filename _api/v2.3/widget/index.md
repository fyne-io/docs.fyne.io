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

Package widget defines the UI widgets within the Fyne toolkit.

## Usage

```go
var (
	// RichTextStyleBlockquote represents a quote presented in an indented block.
	//
	// Since: 2.1
	RichTextStyleBlockquote = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    false,
		SizeName:  theme.SizeNameText,
		TextStyle: fyne.TextStyle{Italic: true},
	}
	// RichTextStyleCodeBlock represents a code blog segment.
	//
	// Since: 2.1
	RichTextStyleCodeBlock = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    false,
		SizeName:  theme.SizeNameText,
		TextStyle: fyne.TextStyle{Monospace: true},
	}
	// RichTextStyleCodeInline represents an inline code segment.
	//
	// Since: 2.1
	RichTextStyleCodeInline = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameText,
		TextStyle: fyne.TextStyle{Monospace: true},
	}
	// RichTextStyleEmphasis represents regular text with emphasis.
	//
	// Since: 2.1
	RichTextStyleEmphasis = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameText,
		TextStyle: fyne.TextStyle{Italic: true},
	}
	// RichTextStyleHeading represents a heading text that stands on its own line.
	//
	// Since: 2.1
	RichTextStyleHeading = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    false,
		SizeName:  theme.SizeNameHeadingText,
		TextStyle: fyne.TextStyle{Bold: true},
	}
	// RichTextStyleInline represents standard text that can be surrounded by other elements.
	//
	// Since: 2.1
	RichTextStyleInline = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameText,
	}
	// RichTextStyleParagraph represents standard text that should appear separate from other text.
	//
	// Since: 2.1
	RichTextStyleParagraph = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    false,
		SizeName:  theme.SizeNameText,
	}
	// RichTextStylePassword represents standard sized text where the characters are obscured.
	//
	// Since: 2.1
	RichTextStylePassword = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameText,
	}
	// RichTextStyleStrong represents regular text with a strong emphasis.
	//
	// Since: 2.1
	RichTextStyleStrong = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    true,
		SizeName:  theme.SizeNameText,
		TextStyle: fyne.TextStyle{Bold: true},
	}
	// RichTextStyleSubHeading represents a sub-heading text that stands on its own line.
	//
	// Since: 2.1
	RichTextStyleSubHeading = RichTextStyle{
		ColorName: theme.ColorNameForeground,
		Inline:    false,
		SizeName:  theme.SizeNameSubHeadingText,
		TextStyle: fyne.TextStyle{Bold: true},
	}
)
```

#### func  NewSimpleRenderer

```go
func NewSimpleRenderer(object fyne.CanvasObject) fyne.WidgetRenderer
```
NewSimpleRenderer creates a new SimpleRenderer to render a widget using a single fyne.CanvasObject.


<div class="since">Since: <code>
2.1</code></div>

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
 * [CheckGroup](checkgroup.html)
 * [CustomTextGridStyle](customtextgridstyle.html)
 * [DisableableWidget](disableablewidget.html)
 * [Entry](entry.html)
 * [FileIcon](fileicon.html)
 * [Form](form.html)
 * [FormItem](formitem.html)
 * [Hyperlink](hyperlink.html)
 * [HyperlinkSegment](hyperlinksegment.html)
 * [Icon](icon.html)
 * [ImageSegment](imagesegment.html)
 * [Label](label.html)
 * [List](list.html)
 * [ListItemID](listitemid.html)
 * [ListSegment](listsegment.html)
 * [Menu](menu.html)
 * [Orientation](orientation.html)
 * [ParagraphSegment](paragraphsegment.html)
 * [PopUp](popup.html)
 * [PopUpMenu](popupmenu.html)
 * [ProgressBar](progressbar.html)
 * [ProgressBarInfinite](progressbarinfinite.html)
 * [RadioGroup](radiogroup.html)
 * [RichText](richtext.html)
 * [RichTextBlock](richtextblock.html)
 * [RichTextSegment](richtextsegment.html)
 * [RichTextStyle](richtextstyle.html)
 * [Select](select.html)
 * [SelectEntry](selectentry.html)
 * [Separator](separator.html)
 * [SeparatorSegment](separatorsegment.html)
 * [Slider](slider.html)
 * [Table](table.html)
 * [TableCellID](tablecellid.html)
 * [TextGrid](textgrid.html)
 * [TextGridCell](textgridcell.html)
 * [TextGridRow](textgridrow.html)
 * [TextGridStyle](textgridstyle.html)
 * [TextSegment](textsegment.html)
 * [Toolbar](toolbar.html)
 * [ToolbarAction](toolbaraction.html)
 * [ToolbarItem](toolbaritem.html)
 * [ToolbarSeparator](toolbarseparator.html)
 * [ToolbarSpacer](toolbarspacer.html)
 * [Tree](tree.html)
 * [TreeNodeID](treenodeid.html)

---
layout: page
tags: [api]
title: Fyne API theme
---

# theme
---
```go
import "fyne.io/fyne/theme"
```

Package theme defines how a Fyne app should look when rendered

## Usage

#### func  BackgroundColor

```go
func BackgroundColor() color.Color
```
BackgroundColor returns the theme's background color

#### func  ButtonColor

```go
func ButtonColor() color.Color
```
ButtonColor returns the theme's standard button color

#### func  CancelIcon

```go
func CancelIcon() fyne.Resource
```
CancelIcon returns a resource containing the standard cancel icon for the current theme

#### func  CheckButtonCheckedIcon

```go
func CheckButtonCheckedIcon() fyne.Resource
```
CheckButtonCheckedIcon returns a resource containing the standard checkbox checked icon for the current theme

#### func  CheckButtonIcon

```go
func CheckButtonIcon() fyne.Resource
```
CheckButtonIcon returns a resource containing the standard checkbox icon for the current theme

#### func  ConfirmIcon

```go
func ConfirmIcon() fyne.Resource
```
ConfirmIcon returns a resource containing the standard confirm icon for the current theme

#### func  ContentAddIcon

```go
func ContentAddIcon() fyne.Resource
```
ContentAddIcon returns a resource containing the standard content add icon for the current theme

#### func  ContentClearIcon

```go
func ContentClearIcon() fyne.Resource
```
ContentClearIcon returns a resource containing the standard content clear icon for the current theme

#### func  ContentCopyIcon

```go
func ContentCopyIcon() fyne.Resource
```
ContentCopyIcon returns a resource containing the standard content copy icon for the current theme

#### func  ContentCutIcon

```go
func ContentCutIcon() fyne.Resource
```
ContentCutIcon returns a resource containing the standard content cut icon for the current theme

#### func  ContentPasteIcon

```go
func ContentPasteIcon() fyne.Resource
```
ContentPasteIcon returns a resource containing the standard content paste icon for the current theme

#### func  ContentRedoIcon

```go
func ContentRedoIcon() fyne.Resource
```
ContentRedoIcon returns a resource containing the standard content redo icon for the current theme

#### func  ContentRemoveIcon

```go
func ContentRemoveIcon() fyne.Resource
```
ContentRemoveIcon returns a resource containing the standard content remove icon for the current theme

#### func  ContentUndoIcon

```go
func ContentUndoIcon() fyne.Resource
```
ContentUndoIcon returns a resource containing the standard content undo icon for the current theme

#### func  DarkTheme

```go
func DarkTheme() fyne.Theme
```
DarkTheme defines the built in dark theme colors and sizes

#### func  DefaultTextBoldFont

```go
func DefaultTextBoldFont() fyne.Resource
```
DefaultTextBoldFont retutns the font resource for the built-in bold font style

#### func  DefaultTextBoldItalicFont

```go
func DefaultTextBoldItalicFont() fyne.Resource
```
DefaultTextBoldItalicFont returns the font resource for the built-in bold and italic font style

#### func  DefaultTextFont

```go
func DefaultTextFont() fyne.Resource
```
DefaultTextFont returns the font resource for the built-in regular font style

#### func  DefaultTextItalicFont

```go
func DefaultTextItalicFont() fyne.Resource
```
DefaultTextItalicFont returns the font resource for the built-in italic font style

#### func  DefaultTextMonospaceFont

```go
func DefaultTextMonospaceFont() fyne.Resource
```
DefaultTextMonospaceFont retutns the font resource for the built-in monospace font face

#### func  DeleteIcon

```go
func DeleteIcon() fyne.Resource
```
DeleteIcon returns a resource containing the standard delete icon for the current theme

#### func  DisabledButtonColor

```go
func DisabledButtonColor() color.Color
```
DisabledButtonColor returns the theme's disabled button color

#### func  DisabledIconColor

```go
func DisabledIconColor() color.Color
```
DisabledIconColor returns the color for a disabledIcon UI element

#### func  DisabledTextColor

```go
func DisabledTextColor() color.Color
```
DisabledTextColor returns the color for a disabledIcon UI element

#### func  DocumentCreateIcon

```go
func DocumentCreateIcon() fyne.Resource
```
DocumentCreateIcon returns a resource containing the standard document create icon for the current theme

#### func  DocumentPrintIcon

```go
func DocumentPrintIcon() fyne.Resource
```
DocumentPrintIcon returns a resource containing the standard document print icon for the current theme

#### func  DocumentSaveIcon

```go
func DocumentSaveIcon() fyne.Resource
```
DocumentSaveIcon returns a resource containing the standard document save icon for the current theme

#### func  FileApplicationIcon

```go
func FileApplicationIcon() fyne.Resource
```
FileApplicationIcon returns a resource containing the file icon representing application files for the current theme

#### func  FileAudioIcon

```go
func FileAudioIcon() fyne.Resource
```
FileAudioIcon returns a resource containing the file icon representing audio files for the current theme

#### func  FileIcon

```go
func FileIcon() fyne.Resource
```
FileIcon returns a resource containing the appropriate file icon for the current theme

#### func  FileImageIcon

```go
func FileImageIcon() fyne.Resource
```
FileImageIcon returns a resource containing the file icon representing image files for the current theme

#### func  FileTextIcon

```go
func FileTextIcon() fyne.Resource
```
FileTextIcon returns a resource containing the file icon representing text files for the current theme

#### func  FileVideoIcon

```go
func FileVideoIcon() fyne.Resource
```
FileVideoIcon returns a resource containing the file icon representing video files for the current theme

#### func  FocusColor

```go
func FocusColor() color.Color
```
FocusColor returns the color used to highlight a focussed widget

#### func  FolderIcon

```go
func FolderIcon() fyne.Resource
```
FolderIcon returns a resource containing the standard folder icon for the current theme

#### func  FolderNewIcon

```go
func FolderNewIcon() fyne.Resource
```
FolderNewIcon returns a resource containing the standard folder creation icon for the current theme

#### func  FolderOpenIcon

```go
func FolderOpenIcon() fyne.Resource
```
FolderOpenIcon returns a resource containing the standard folder open icon for the current theme

#### func  FyneLogo

```go
func FyneLogo() fyne.Resource
```
FyneLogo returns a resource containing the Fyne logo

#### func  HelpIcon

```go
func HelpIcon() fyne.Resource
```
HelpIcon returns a resource containing the standard help icon for the current theme

#### func  HomeIcon

```go
func HomeIcon() fyne.Resource
```
HomeIcon returns a resource containing the standard home folder icon for the current theme

#### func  HoverColor

```go
func HoverColor() color.Color
```
HoverColor returns the color used to highlight interactive elements currently under a cursor

#### func  HyperlinkColor

```go
func HyperlinkColor() color.Color
```
HyperlinkColor returns the theme's standard hyperlink color

#### func  IconColor

```go
func IconColor() color.Color
```
IconColor returns the theme's standard text color

#### func  IconInlineSize

```go
func IconInlineSize() int
```
IconInlineSize is the standard size of icons which appear within buttons, labels etc.

#### func  InfoIcon

```go
func InfoIcon() fyne.Resource
```
InfoIcon returns a resource containing the standard dialog info icon for the current theme

#### func  LightTheme

```go
func LightTheme() fyne.Theme
```
LightTheme defines the built in light theme colors and sizes

#### func  MailAttachmentIcon

```go
func MailAttachmentIcon() fyne.Resource
```
MailAttachmentIcon returns a resource containing the standard mail attachment icon for the current theme

#### func  MailComposeIcon

```go
func MailComposeIcon() fyne.Resource
```
MailComposeIcon returns a resource containing the standard mail compose icon for the current theme

#### func  MailForwardIcon

```go
func MailForwardIcon() fyne.Resource
```
MailForwardIcon returns a resource containing the standard mail forward icon for the current theme

#### func  MailReplyAllIcon

```go
func MailReplyAllIcon() fyne.Resource
```
MailReplyAllIcon returns a resource containing the standard mail reply all icon for the current theme

#### func  MailReplyIcon

```go
func MailReplyIcon() fyne.Resource
```
MailReplyIcon returns a resource containing the standard mail reply icon for the current theme

#### func  MailSendIcon

```go
func MailSendIcon() fyne.Resource
```
MailSendIcon returns a resource containing the standard mail send icon for the current theme

#### func  MediaFastForwardIcon

```go
func MediaFastForwardIcon() fyne.Resource
```
MediaFastForwardIcon returns a resource containing the standard media fast-forward icon for the current theme

#### func  MediaFastRewindIcon

```go
func MediaFastRewindIcon() fyne.Resource
```
MediaFastRewindIcon returns a resource containing the standard media fast-rewind icon for the current theme

#### func  MediaPauseIcon

```go
func MediaPauseIcon() fyne.Resource
```
MediaPauseIcon returns a resource containing the standard media pause icon for the current theme

#### func  MediaPlayIcon

```go
func MediaPlayIcon() fyne.Resource
```
MediaPlayIcon returns a resource containing the standard media play icon for the current theme

#### func  MediaRecordIcon

```go
func MediaRecordIcon() fyne.Resource
```
MediaRecordIcon returns a resource containing the standard media record icon for the current theme

#### func  MediaReplayIcon

```go
func MediaReplayIcon() fyne.Resource
```
MediaReplayIcon returns a resource containing the standard media replay icon for the current theme

#### func  MediaSkipNextIcon

```go
func MediaSkipNextIcon() fyne.Resource
```
MediaSkipNextIcon returns a resource containing the standard media skip next icon for the current theme

#### func  MediaSkipPreviousIcon

```go
func MediaSkipPreviousIcon() fyne.Resource
```
MediaSkipPreviousIcon returns a resource containing the standard media skip previous icon for the current theme

#### func  MenuDropDownIcon

```go
func MenuDropDownIcon() fyne.Resource
```
MenuDropDownIcon returns a resource containing the standard menu drop down icon for the current theme

#### func  MenuDropUpIcon

```go
func MenuDropUpIcon() fyne.Resource
```
MenuDropUpIcon returns a resource containing the standard menu drop up icon for the current theme

#### func  MenuExpandIcon

```go
func MenuExpandIcon() fyne.Resource
```
MenuExpandIcon returns a resource containing the standard (mobile) expand "submenu icon for the current theme

#### func  MenuIcon

```go
func MenuIcon() fyne.Resource
```
MenuIcon returns a resource containing the standard (mobile) menu icon for the current theme

#### func  MoveDownIcon

```go
func MoveDownIcon() fyne.Resource
```
MoveDownIcon returns a resource containing the standard down arrow icon for the current theme

#### func  MoveUpIcon

```go
func MoveUpIcon() fyne.Resource
```
MoveUpIcon returns a resource containing the standard up arrow icon for the current theme

#### func  NavigateBackIcon

```go
func NavigateBackIcon() fyne.Resource
```
NavigateBackIcon returns a resource containing the standard backward navigation icon for the current theme

#### func  NavigateNextIcon

```go
func NavigateNextIcon() fyne.Resource
```
NavigateNextIcon returns a resource containing the standard forward navigation icon for the current theme

#### func  Padding

```go
func Padding() int
```
Padding is the standard gap between elements and the border around interface elements

#### func  PlaceHolderColor

```go
func PlaceHolderColor() color.Color
```
PlaceHolderColor returns the theme's standard text color

#### func  PrimaryColor

```go
func PrimaryColor() color.Color
```
PrimaryColor returns the color used to highlight primary features

#### func  QuestionIcon

```go
func QuestionIcon() fyne.Resource
```
QuestionIcon returns a resource containing the standard dialog question icon for the current theme

#### func  RadioButtonCheckedIcon

```go
func RadioButtonCheckedIcon() fyne.Resource
```
RadioButtonCheckedIcon returns a resource containing the standard radio button checked icon for the current theme

#### func  RadioButtonIcon

```go
func RadioButtonIcon() fyne.Resource
```
RadioButtonIcon returns a resource containing the standard radio button icon for the current theme

#### func  ScrollBarColor

```go
func ScrollBarColor() color.Color
```
ScrollBarColor returns the color (and translucency) for a scrollBar

#### func  ScrollBarSize

```go
func ScrollBarSize() int
```
ScrollBarSize is the width (or height) of the bars on a ScrollContainer

#### func  ScrollBarSmallSize

```go
func ScrollBarSmallSize() int
```
ScrollBarSmallSize is the width (or height) of the minimized bars on a ScrollContainer

#### func  SearchIcon

```go
func SearchIcon() fyne.Resource
```
SearchIcon returns a resource containing the standard search icon for the current theme

#### func  SearchReplaceIcon

```go
func SearchReplaceIcon() fyne.Resource
```
SearchReplaceIcon returns a resource containing the standard search and replace icon for the current theme

#### func  SettingsIcon

```go
func SettingsIcon() fyne.Resource
```
SettingsIcon returns a resource containing the standard settings icon for the current theme

#### func  ShadowColor

```go
func ShadowColor() color.Color
```
ShadowColor returns the color (and translucency) for shadows used for indicating elevation

#### func  TextBoldFont

```go
func TextBoldFont() fyne.Resource
```
TextBoldFont retutns the font resource for the bold font style

#### func  TextBoldItalicFont

```go
func TextBoldItalicFont() fyne.Resource
```
TextBoldItalicFont returns the font resource for the bold and italic font style

#### func  TextColor

```go
func TextColor() color.Color
```
TextColor returns the theme's standard text color

#### func  TextFont

```go
func TextFont() fyne.Resource
```
TextFont returns the font resource for the regular font style

#### func  TextItalicFont

```go
func TextItalicFont() fyne.Resource
```
TextItalicFont returns the font resource for the italic font style

#### func  TextMonospaceFont

```go
func TextMonospaceFont() fyne.Resource
```
TextMonospaceFont retutns the font resource for the monospace font face

#### func  TextSize

```go
func TextSize() int
```
TextSize returns the standard text size

#### func  ViewFullScreenIcon

```go
func ViewFullScreenIcon() fyne.Resource
```
ViewFullScreenIcon returns a resource containing the standard fullscreen icon for the current theme

#### func  ViewRefreshIcon

```go
func ViewRefreshIcon() fyne.Resource
```
ViewRefreshIcon returns a resource containing the standard refresh icon for the current theme

#### func  ViewRestoreIcon

```go
func ViewRestoreIcon() fyne.Resource
```
ViewRestoreIcon returns a resource containing the standard exit fullscreen icon for the current theme

#### func  VisibilityIcon

```go
func VisibilityIcon() fyne.Resource
```
VisibilityIcon returns a resource containing the standard visibity icon for the current theme

#### func  VisibilityOffIcon

```go
func VisibilityOffIcon() fyne.Resource
```
VisibilityOffIcon returns a resource containing the standard visibity off icon for the current theme

#### func  VolumeDownIcon

```go
func VolumeDownIcon() fyne.Resource
```
VolumeDownIcon returns a resource containing the standard volume down icon for the current theme

#### func  VolumeMuteIcon

```go
func VolumeMuteIcon() fyne.Resource
```
VolumeMuteIcon returns a resource containing the standard volume mute icon for the current theme

#### func  VolumeUpIcon

```go
func VolumeUpIcon() fyne.Resource
```
VolumeUpIcon returns a resource containing the standard volume up icon for the current theme

#### func  WarningIcon

```go
func WarningIcon() fyne.Resource
```
WarningIcon returns a resource containing the standard dialog warning icon for the current theme

#### func  ZoomFitIcon

```go
func ZoomFitIcon() fyne.Resource
```
ZoomFitIcon returns a resource containing the standard zoom fit icon for the current theme

#### func  ZoomInIcon

```go
func ZoomInIcon() fyne.Resource
```
ZoomInIcon returns a resource containing the standard zoom in icon for the current theme

#### func  ZoomOutIcon

```go
func ZoomOutIcon() fyne.Resource
```
ZoomOutIcon returns a resource containing the standard zoom out icon for the current theme

#### types

 * [DisabledResource](disabledresource.html)
 * [ThemedResource](themedresource.html)

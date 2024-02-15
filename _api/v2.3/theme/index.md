---
layout: page
tags: [api]
title: Fyne API "theme"
---

# theme
---
```go
import "fyne.io/fyne/v2/theme"
```

Package theme defines how a Fyne app should look when rendered.

## Usage

```go
const (
	// IconNameCancel is the name of theme lookup for cancel icon.
	//
	// Since: 2.0
	IconNameCancel fyne.ThemeIconName = "cancel"

	// IconNameConfirm is the name of theme lookup for confirm icon.
	//
	// Since: 2.0
	IconNameConfirm fyne.ThemeIconName = "confirm"

	// IconNameDelete is the name of theme lookup for delete icon.
	//
	// Since: 2.0
	IconNameDelete fyne.ThemeIconName = "delete"

	// IconNameSearch is the name of theme lookup for search icon.
	//
	// Since: 2.0
	IconNameSearch fyne.ThemeIconName = "search"

	// IconNameSearchReplace is the name of theme lookup for search and replace icon.
	//
	// Since: 2.0
	IconNameSearchReplace fyne.ThemeIconName = "searchReplace"

	// IconNameMenu is the name of theme lookup for menu icon.
	//
	// Since: 2.0
	IconNameMenu fyne.ThemeIconName = "menu"

	// IconNameMenuExpand is the name of theme lookup for menu expansion icon.
	//
	// Since: 2.0
	IconNameMenuExpand fyne.ThemeIconName = "menuExpand"

	// IconNameCheckButtonChecked is the name of theme lookup for checked check button icon.
	//
	// Since: 2.0
	IconNameCheckButtonChecked fyne.ThemeIconName = "checked"

	// IconNameCheckButton is the name of theme lookup for  unchecked check button icon.
	//
	// Since: 2.0
	IconNameCheckButton fyne.ThemeIconName = "unchecked"

	// IconNameRadioButton is the name of theme lookup for radio button unchecked icon.
	//
	// Since: 2.0
	IconNameRadioButton fyne.ThemeIconName = "radioButton"

	// IconNameRadioButtonChecked is the name of theme lookup for radio button checked icon.
	//
	// Since: 2.0
	IconNameRadioButtonChecked fyne.ThemeIconName = "radioButtonChecked"

	// IconNameColorAchromatic is the name of theme lookup for greyscale color icon.
	//
	// Since: 2.0
	IconNameColorAchromatic fyne.ThemeIconName = "colorAchromatic"

	// IconNameColorChromatic is the name of theme lookup for full color icon.
	//
	// Since: 2.0
	IconNameColorChromatic fyne.ThemeIconName = "colorChromatic"

	// IconNameColorPalette is the name of theme lookup for color palette icon.
	//
	// Since: 2.0
	IconNameColorPalette fyne.ThemeIconName = "colorPalette"

	// IconNameContentAdd is the name of theme lookup for content add icon.
	//
	// Since: 2.0
	IconNameContentAdd fyne.ThemeIconName = "contentAdd"

	// IconNameContentRemove is the name of theme lookup for content remove icon.
	//
	// Since: 2.0
	IconNameContentRemove fyne.ThemeIconName = "contentRemove"

	// IconNameContentCut is the name of theme lookup for content cut icon.
	//
	// Since: 2.0
	IconNameContentCut fyne.ThemeIconName = "contentCut"

	// IconNameContentCopy is the name of theme lookup for content copy icon.
	//
	// Since: 2.0
	IconNameContentCopy fyne.ThemeIconName = "contentCopy"

	// IconNameContentPaste is the name of theme lookup for content paste icon.
	//
	// Since: 2.0
	IconNameContentPaste fyne.ThemeIconName = "contentPaste"

	// IconNameContentClear is the name of theme lookup for content clear icon.
	//
	// Since: 2.0
	IconNameContentClear fyne.ThemeIconName = "contentClear"

	// IconNameContentRedo is the name of theme lookup for content redo icon.
	//
	// Since: 2.0
	IconNameContentRedo fyne.ThemeIconName = "contentRedo"

	// IconNameContentUndo is the name of theme lookup for content undo icon.
	//
	// Since: 2.0
	IconNameContentUndo fyne.ThemeIconName = "contentUndo"

	// IconNameInfo is the name of theme lookup for info icon.
	//
	// Since: 2.0
	IconNameInfo fyne.ThemeIconName = "info"

	// IconNameQuestion is the name of theme lookup for question icon.
	//
	// Since: 2.0
	IconNameQuestion fyne.ThemeIconName = "question"

	// IconNameWarning is the name of theme lookup for warning icon.
	//
	// Since: 2.0
	IconNameWarning fyne.ThemeIconName = "warning"

	// IconNameError is the name of theme lookup for error icon.
	//
	// Since: 2.0
	IconNameError fyne.ThemeIconName = "error"

	// IconNameDocument is the name of theme lookup for document icon.
	//
	// Since: 2.0
	IconNameDocument fyne.ThemeIconName = "document"

	// IconNameDocumentCreate is the name of theme lookup for document create icon.
	//
	// Since: 2.0
	IconNameDocumentCreate fyne.ThemeIconName = "documentCreate"

	// IconNameDocumentPrint is the name of theme lookup for document print icon.
	//
	// Since: 2.0
	IconNameDocumentPrint fyne.ThemeIconName = "documentPrint"

	// IconNameDocumentSave is the name of theme lookup for document save icon.
	//
	// Since: 2.0
	IconNameDocumentSave fyne.ThemeIconName = "documentSave"

	// IconNameMoreHorizontal is the name of theme lookup for horizontal more.
	//
	// Since 2.0
	IconNameMoreHorizontal fyne.ThemeIconName = "moreHorizontal"

	// IconNameMoreVertical is the name of theme lookup for vertical more.
	//
	// Since 2.0
	IconNameMoreVertical fyne.ThemeIconName = "moreVertical"

	// IconNameMailAttachment is the name of theme lookup for mail attachment icon.
	//
	// Since: 2.0
	IconNameMailAttachment fyne.ThemeIconName = "mailAttachment"

	// IconNameMailCompose is the name of theme lookup for mail compose icon.
	//
	// Since: 2.0
	IconNameMailCompose fyne.ThemeIconName = "mailCompose"

	// IconNameMailForward is the name of theme lookup for mail forward icon.
	//
	// Since: 2.0
	IconNameMailForward fyne.ThemeIconName = "mailForward"

	// IconNameMailReply is the name of theme lookup for mail reply icon.
	//
	// Since: 2.0
	IconNameMailReply fyne.ThemeIconName = "mailReply"

	// IconNameMailReplyAll is the name of theme lookup for mail reply-all icon.
	//
	// Since: 2.0
	IconNameMailReplyAll fyne.ThemeIconName = "mailReplyAll"

	// IconNameMailSend is the name of theme lookup for mail send icon.
	//
	// Since: 2.0
	IconNameMailSend fyne.ThemeIconName = "mailSend"

	// IconNameMediaMusic is the name of theme lookup for media music icon.
	//
	// Since: 2.1
	IconNameMediaMusic fyne.ThemeIconName = "mediaMusic"

	// IconNameMediaPhoto is the name of theme lookup for media photo icon.
	//
	// Since: 2.1
	IconNameMediaPhoto fyne.ThemeIconName = "mediaPhoto"

	// IconNameMediaVideo is the name of theme lookup for media video icon.
	//
	// Since: 2.1
	IconNameMediaVideo fyne.ThemeIconName = "mediaVideo"

	// IconNameMediaFastForward is the name of theme lookup for media fast-forward icon.
	//
	// Since: 2.0
	IconNameMediaFastForward fyne.ThemeIconName = "mediaFastForward"

	// IconNameMediaFastRewind is the name of theme lookup for media fast-rewind icon.
	//
	// Since: 2.0
	IconNameMediaFastRewind fyne.ThemeIconName = "mediaFastRewind"

	// IconNameMediaPause is the name of theme lookup for media pause icon.
	//
	// Since: 2.0
	IconNameMediaPause fyne.ThemeIconName = "mediaPause"

	// IconNameMediaPlay is the name of theme lookup for media play icon.
	//
	// Since: 2.0
	IconNameMediaPlay fyne.ThemeIconName = "mediaPlay"

	// IconNameMediaRecord is the name of theme lookup for media record icon.
	//
	// Since: 2.0
	IconNameMediaRecord fyne.ThemeIconName = "mediaRecord"

	// IconNameMediaReplay is the name of theme lookup for media replay icon.
	//
	// Since: 2.0
	IconNameMediaReplay fyne.ThemeIconName = "mediaReplay"

	// IconNameMediaSkipNext is the name of theme lookup for media skip next icon.
	//
	// Since: 2.0
	IconNameMediaSkipNext fyne.ThemeIconName = "mediaSkipNext"

	// IconNameMediaSkipPrevious is the name of theme lookup for media skip previous icon.
	//
	// Since: 2.0
	IconNameMediaSkipPrevious fyne.ThemeIconName = "mediaSkipPrevious"

	// IconNameMediaStop is the name of theme lookup for media stop icon.
	//
	// Since: 2.0
	IconNameMediaStop fyne.ThemeIconName = "mediaStop"

	// IconNameMoveDown is the name of theme lookup for move down icon.
	//
	// Since: 2.0
	IconNameMoveDown fyne.ThemeIconName = "arrowDown"

	// IconNameMoveUp is the name of theme lookup for move up icon.
	//
	// Since: 2.0
	IconNameMoveUp fyne.ThemeIconName = "arrowUp"

	// IconNameNavigateBack is the name of theme lookup for navigate back icon.
	//
	// Since: 2.0
	IconNameNavigateBack fyne.ThemeIconName = "arrowBack"

	// IconNameNavigateNext is the name of theme lookup for navigate next icon.
	//
	// Since: 2.0
	IconNameNavigateNext fyne.ThemeIconName = "arrowForward"

	// IconNameArrowDropDown is the name of theme lookup for drop-down arrow icon.
	//
	// Since: 2.0
	IconNameArrowDropDown fyne.ThemeIconName = "arrowDropDown"

	// IconNameArrowDropUp is the name of theme lookup for drop-up arrow icon.
	//
	// Since: 2.0
	IconNameArrowDropUp fyne.ThemeIconName = "arrowDropUp"

	// IconNameFile is the name of theme lookup for file icon.
	//
	// Since: 2.0
	IconNameFile fyne.ThemeIconName = "file"

	// IconNameFileApplication is the name of theme lookup for file application icon.
	//
	// Since: 2.0
	IconNameFileApplication fyne.ThemeIconName = "fileApplication"

	// IconNameFileAudio is the name of theme lookup for file audio icon.
	//
	// Since: 2.0
	IconNameFileAudio fyne.ThemeIconName = "fileAudio"

	// IconNameFileImage is the name of theme lookup for file image icon.
	//
	// Since: 2.0
	IconNameFileImage fyne.ThemeIconName = "fileImage"

	// IconNameFileText is the name of theme lookup for file text icon.
	//
	// Since: 2.0
	IconNameFileText fyne.ThemeIconName = "fileText"

	// IconNameFileVideo is the name of theme lookup for file video icon.
	//
	// Since: 2.0
	IconNameFileVideo fyne.ThemeIconName = "fileVideo"

	// IconNameFolder is the name of theme lookup for folder icon.
	//
	// Since: 2.0
	IconNameFolder fyne.ThemeIconName = "folder"

	// IconNameFolderNew is the name of theme lookup for folder new icon.
	//
	// Since: 2.0
	IconNameFolderNew fyne.ThemeIconName = "folderNew"

	// IconNameFolderOpen is the name of theme lookup for folder open icon.
	//
	// Since: 2.0
	IconNameFolderOpen fyne.ThemeIconName = "folderOpen"

	// IconNameHelp is the name of theme lookup for help icon.
	//
	// Since: 2.0
	IconNameHelp fyne.ThemeIconName = "help"

	// IconNameHistory is the name of theme lookup for history icon.
	//
	// Since: 2.0
	IconNameHistory fyne.ThemeIconName = "history"

	// IconNameHome is the name of theme lookup for home icon.
	//
	// Since: 2.0
	IconNameHome fyne.ThemeIconName = "home"

	// IconNameSettings is the name of theme lookup for settings icon.
	//
	// Since: 2.0
	IconNameSettings fyne.ThemeIconName = "settings"

	// IconNameStorage is the name of theme lookup for storage icon.
	//
	// Since: 2.0
	IconNameStorage fyne.ThemeIconName = "storage"

	// IconNameUpload is the name of theme lookup for upload icon.
	//
	// Since: 2.0
	IconNameUpload fyne.ThemeIconName = "upload"

	// IconNameViewFullScreen is the name of theme lookup for view fullscreen icon.
	//
	// Since: 2.0
	IconNameViewFullScreen fyne.ThemeIconName = "viewFullScreen"

	// IconNameViewRefresh is the name of theme lookup for view refresh icon.
	//
	// Since: 2.0
	IconNameViewRefresh fyne.ThemeIconName = "viewRefresh"

	// IconNameViewZoomFit is the name of theme lookup for view zoom fit icon.
	//
	// Since: 2.0
	IconNameViewZoomFit fyne.ThemeIconName = "viewZoomFit"

	// IconNameViewZoomIn is the name of theme lookup for view zoom in icon.
	//
	// Since: 2.0
	IconNameViewZoomIn fyne.ThemeIconName = "viewZoomIn"

	// IconNameViewZoomOut is the name of theme lookup for view zoom out icon.
	//
	// Since: 2.0
	IconNameViewZoomOut fyne.ThemeIconName = "viewZoomOut"

	// IconNameViewRestore is the name of theme lookup for view restore icon.
	//
	// Since: 2.0
	IconNameViewRestore fyne.ThemeIconName = "viewRestore"

	// IconNameVisibility is the name of theme lookup for visibility icon.
	//
	// Since: 2.0
	IconNameVisibility fyne.ThemeIconName = "visibility"

	// IconNameVisibilityOff is the name of theme lookup for invisibility icon.
	//
	// Since: 2.0
	IconNameVisibilityOff fyne.ThemeIconName = "visibilityOff"

	// IconNameVolumeDown is the name of theme lookup for volume down icon.
	//
	// Since: 2.0
	IconNameVolumeDown fyne.ThemeIconName = "volumeDown"

	// IconNameVolumeMute is the name of theme lookup for volume mute icon.
	//
	// Since: 2.0
	IconNameVolumeMute fyne.ThemeIconName = "volumeMute"

	// IconNameVolumeUp is the name of theme lookup for volume up icon.
	//
	// Since: 2.0
	IconNameVolumeUp fyne.ThemeIconName = "volumeUp"

	// IconNameDownload is the name of theme lookup for download icon.
	//
	// Since: 2.0
	IconNameDownload fyne.ThemeIconName = "download"

	// IconNameComputer is the name of theme lookup for computer icon.
	//
	// Since: 2.0
	IconNameComputer fyne.ThemeIconName = "computer"

	// IconNameAccount is the name of theme lookup for account icon.
	//
	// Since: 2.1
	IconNameAccount fyne.ThemeIconName = "account"

	// IconNameLogin is the name of theme lookup for login icon.
	//
	// Since: 2.1
	IconNameLogin fyne.ThemeIconName = "login"

	// IconNameLogout is the name of theme lookup for logout icon.
	//
	// Since: 2.1
	IconNameLogout fyne.ThemeIconName = "logout"

	// IconNameList is the name of theme lookup for list icon.
	//
	// Since: 2.1
	IconNameList fyne.ThemeIconName = "list"

	// IconNameGrid is the name of theme lookup for grid icon.
	//
	// Since: 2.1
	IconNameGrid fyne.ThemeIconName = "grid"
)
```

```go
const (
	// ColorRed is the red primary color name.
	//
	// Since: 1.4
	ColorRed = "red"
	// ColorOrange is the orange primary color name.
	//
	// Since: 1.4
	ColorOrange = "orange"
	// ColorYellow is the yellow primary color name.
	//
	// Since: 1.4
	ColorYellow = "yellow"
	// ColorGreen is the green primary color name.
	//
	// Since: 1.4
	ColorGreen = "green"
	// ColorBlue is the blue primary color name.
	//
	// Since: 1.4
	ColorBlue = "blue"
	// ColorPurple is the purple primary color name.
	//
	// Since: 1.4
	ColorPurple = "purple"
	// ColorBrown is the brown primary color name.
	//
	// Since: 1.4
	ColorBrown = "brown"
	// ColorGray is the gray primary color name.
	//
	// Since: 1.4
	ColorGray = "gray"

	// ColorNameBackground is the name of theme lookup for background color.
	//
	// Since: 2.0
	ColorNameBackground fyne.ThemeColorName = "background"

	// ColorNameButton is the name of theme lookup for button color.
	//
	// Since: 2.0
	ColorNameButton fyne.ThemeColorName = "button"

	// ColorNameDisabledButton is the name of theme lookup for disabled button color.
	//
	// Since: 2.0
	ColorNameDisabledButton fyne.ThemeColorName = "disabledButton"

	// ColorNameDisabled is the name of theme lookup for disabled foreground color.
	//
	// Since: 2.0
	ColorNameDisabled fyne.ThemeColorName = "disabled"

	// ColorNameError is the name of theme lookup for foreground error color.
	//
	// Since: 2.0
	ColorNameError fyne.ThemeColorName = "error"

	// ColorNameFocus is the name of theme lookup for focus color.
	//
	// Since: 2.0
	ColorNameFocus fyne.ThemeColorName = "focus"

	// ColorNameForeground is the name of theme lookup for foreground color.
	//
	// Since: 2.0
	ColorNameForeground fyne.ThemeColorName = "foreground"

	// ColorNameHover is the name of theme lookup for hover color.
	//
	// Since: 2.0
	ColorNameHover fyne.ThemeColorName = "hover"

	// ColorNameInputBackground is the name of theme lookup for background color of an input field.
	//
	// Since: 2.0
	ColorNameInputBackground fyne.ThemeColorName = "inputBackground"

	// ColorNameInputBorder is the name of theme lookup for border color of an input field.
	//
	// Since: 2.3
	ColorNameInputBorder fyne.ThemeColorName = "inputBorder"

	// ColorNameMenuBackground is the name of theme lookup for background color of menus.
	//
	// Since: 2.3
	ColorNameMenuBackground fyne.ThemeColorName = "menuBackground"

	// ColorNameOverlayBackground is the name of theme lookup for background color of overlays like dialogs.
	//
	// Since: 2.3
	ColorNameOverlayBackground fyne.ThemeColorName = "overlayBackground"

	// ColorNamePlaceHolder is the name of theme lookup for placeholder text color.
	//
	// Since: 2.0
	ColorNamePlaceHolder fyne.ThemeColorName = "placeholder"

	// ColorNamePressed is the name of theme lookup for the tap overlay color.
	//
	// Since: 2.0
	ColorNamePressed fyne.ThemeColorName = "pressed"

	// ColorNamePrimary is the name of theme lookup for primary color.
	//
	// Since: 2.0
	ColorNamePrimary fyne.ThemeColorName = "primary"

	// ColorNameScrollBar is the name of theme lookup for scrollbar color.
	//
	// Since: 2.0
	ColorNameScrollBar fyne.ThemeColorName = "scrollBar"

	// ColorNameSelection is the name of theme lookup for selection color.
	//
	// Since: 2.1
	ColorNameSelection fyne.ThemeColorName = "selection"

	// ColorNameSeparator is the name of theme lookup for separator bars.
	//
	// Since: 2.3
	ColorNameSeparator fyne.ThemeColorName = "separator"

	// ColorNameShadow is the name of theme lookup for shadow color.
	//
	// Since: 2.0
	ColorNameShadow fyne.ThemeColorName = "shadow"

	// ColorNameSuccess is the name of theme lookup for foreground success color.
	//
	// Since: 2.3
	ColorNameSuccess fyne.ThemeColorName = "success"

	// ColorNameWarning is the name of theme lookup for foreground warning color.
	//
	// Since: 2.3
	ColorNameWarning fyne.ThemeColorName = "warning"

	// SizeNameCaptionText is the name of theme lookup for helper text size, normally smaller than regular text size.
	//
	// Since: 2.0
	SizeNameCaptionText fyne.ThemeSizeName = "helperText"

	// SizeNameInlineIcon is the name of theme lookup for inline icons size.
	//
	// Since: 2.0
	SizeNameInlineIcon fyne.ThemeSizeName = "iconInline"

	// SizeNameInnerPadding is the name of theme lookup for internal widget padding size.
	//
	// Since: 2.3
	SizeNameInnerPadding fyne.ThemeSizeName = "innerPadding"

	// SizeNameLineSpacing is the name of theme lookup for between text line spacing.
	//
	// Since: 2.3
	SizeNameLineSpacing fyne.ThemeSizeName = "lineSpacing"

	// SizeNamePadding is the name of theme lookup for padding size.
	//
	// Since: 2.0
	SizeNamePadding fyne.ThemeSizeName = "padding"

	// SizeNameScrollBar is the name of theme lookup for the scrollbar size.
	//
	// Since: 2.0
	SizeNameScrollBar fyne.ThemeSizeName = "scrollBar"

	// SizeNameScrollBarSmall is the name of theme lookup for the shrunk scrollbar size.
	//
	// Since: 2.0
	SizeNameScrollBarSmall fyne.ThemeSizeName = "scrollBarSmall"

	// SizeNameSeparatorThickness is the name of theme lookup for the thickness of a separator.
	//
	// Since: 2.0
	SizeNameSeparatorThickness fyne.ThemeSizeName = "separator"

	// SizeNameText is the name of theme lookup for text size.
	//
	// Since: 2.0
	SizeNameText fyne.ThemeSizeName = "text"

	// SizeNameHeadingText is the name of theme lookup for text size of a heading.
	//
	// Since: 2.1
	SizeNameHeadingText fyne.ThemeSizeName = "headingText"

	// SizeNameSubHeadingText is the name of theme lookup for text size of a sub-heading.
	//
	// Since: 2.1
	SizeNameSubHeadingText fyne.ThemeSizeName = "subHeadingText"

	// SizeNameInputBorder is the name of theme lookup for input border size.
	//
	// Since: 2.0
	SizeNameInputBorder fyne.ThemeSizeName = "inputBorder"

	// VariantDark is the version of a theme that satisfies a user preference for a light look.
	//
	// Since: 2.0
	VariantDark fyne.ThemeVariant = 0

	// VariantLight is the version of a theme that satisfies a user preference for a dark look.
	//
	// Since: 2.0
	VariantLight fyne.ThemeVariant = 1
)
```

#### func  AccountIcon

```go
func AccountIcon() fyne.Resource
```
AccountIcon returns a resource containing the standard account icon for the current theme

#### func  BackgroundColor

```go
func BackgroundColor() color.Color
```
BackgroundColor returns the theme's background color.

#### func  ButtonColor

```go
func ButtonColor() color.Color
```
ButtonColor returns the theme's standard button color.

#### func  CancelIcon

```go
func CancelIcon() fyne.Resource
```
CancelIcon returns a resource containing the standard cancel icon for the current theme

#### func  CaptionTextSize

```go
func CaptionTextSize() float32
```
CaptionTextSize returns the size for caption text.

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

#### func  ColorAchromaticIcon

```go
func ColorAchromaticIcon() fyne.Resource
```
ColorAchromaticIcon returns a resource containing the standard achromatic color icon for the current theme

#### func  ColorChromaticIcon

```go
func ColorChromaticIcon() fyne.Resource
```
ColorChromaticIcon returns a resource containing the standard chromatic color icon for the current theme

#### func  ColorPaletteIcon

```go
func ColorPaletteIcon() fyne.Resource
```
ColorPaletteIcon returns a resource containing the standard color palette icon for the current theme

#### func  ComputerIcon

```go
func ComputerIcon() fyne.Resource
```
ComputerIcon returns a resource containing the standard computer icon for the current theme

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
DarkTheme defines the built-in dark theme colors and sizes.


<div class="deprecated">
Deprecated: This method ignores user preference and should not be used, it will be removed in v3.0.</div>

#### func  DefaultSymbolFont

```go
func DefaultSymbolFont() fyne.Resource
```
DefaultSymbolFont returns the font resource for the built-in symbol font.


<div class="since">Since: <code>
2.2</code></div>

#### func  DefaultTextBoldFont

```go
func DefaultTextBoldFont() fyne.Resource
```
DefaultTextBoldFont returns the font resource for the built-in bold font style.

#### func  DefaultTextBoldItalicFont

```go
func DefaultTextBoldItalicFont() fyne.Resource
```
DefaultTextBoldItalicFont returns the font resource for the built-in bold and italic font style.

#### func  DefaultTextFont

```go
func DefaultTextFont() fyne.Resource
```
DefaultTextFont returns the font resource for the built-in regular font style.

#### func  DefaultTextItalicFont

```go
func DefaultTextItalicFont() fyne.Resource
```
DefaultTextItalicFont returns the font resource for the built-in italic font style.

#### func  DefaultTextMonospaceFont

```go
func DefaultTextMonospaceFont() fyne.Resource
```
DefaultTextMonospaceFont returns the font resource for the built-in monospace font face.

#### func  DefaultTheme

```go
func DefaultTheme() fyne.Theme
```
DefaultTheme returns a built-in theme that can adapt to the user preference of light or dark colors.


<div class="since">Since: <code>
2.0</code></div>

#### func  DeleteIcon

```go
func DeleteIcon() fyne.Resource
```
DeleteIcon returns a resource containing the standard delete icon for the current theme

#### func  DisabledButtonColor

```go
func DisabledButtonColor() color.Color
```
DisabledButtonColor returns the theme's disabled button color.

#### func  DisabledColor

```go
func DisabledColor() color.Color
```
DisabledColor returns the foreground color for a disabled UI element.


<div class="since">Since: <code>
2.0</code></div>

#### func  DisabledTextColor

```go
func DisabledTextColor() color.Color
```
DisabledTextColor returns the theme's disabled text color - this is actually the disabled color since 1.4.


<div class="deprecated">
Deprecated: Use theme.DisabledColor() colour instead.</div>

#### func  DocumentCreateIcon

```go
func DocumentCreateIcon() fyne.Resource
```
DocumentCreateIcon returns a resource containing the standard document create icon for the current theme

#### func  DocumentIcon

```go
func DocumentIcon() fyne.Resource
```
DocumentIcon returns a resource containing the standard document icon for the current theme

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

#### func  DownloadIcon

```go
func DownloadIcon() fyne.Resource
```
DownloadIcon returns a resource containing the standard download icon for the current theme

#### func  ErrorColor

```go
func ErrorColor() color.Color
```
ErrorColor returns the theme's error foreground color.


<div class="since">Since: <code>
2.0</code></div>

#### func  ErrorIcon

```go
func ErrorIcon() fyne.Resource
```
ErrorIcon returns a resource containing the standard dialog error icon for the current theme

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
FocusColor returns the color used to highlight a focused widget.

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

#### func  ForegroundColor

```go
func ForegroundColor() color.Color
```
ForegroundColor returns the theme's standard foreground color for text and icons.


<div class="since">Since: <code>
2.0</code></div>

#### func  FromJSON

```go
func FromJSON(data string) (fyne.Theme, error)
```
FromJSON returns a Theme created from the given JSON metadata. Any values not present in the data will fall back to the default theme. If a parse error occurs it will be returned along with a default theme.


<div class="since">Since: <code>
2.2</code></div>

#### func  FromJSONReader

```go
func FromJSONReader(r io.Reader) (fyne.Theme, error)
```
FromJSONReader returns a Theme created from the given JSON metadata through the reader. Any values not present in the data will fall back to the default theme. If a parse error occurs it will be returned along with a default theme.


<div class="since">Since: <code>
2.2</code></div>

#### func  FromLegacy

```go
func FromLegacy(t fyne.LegacyTheme) fyne.Theme
```
FromLegacy returns a 2.0 Theme created from the given LegacyTheme data. This is a transition path and will be removed in the future (probably version 3.0).


<div class="since">Since: <code>
2.0</code></div>

#### func  FyneLogo

```go
func FyneLogo() fyne.Resource
```
FyneLogo returns a resource containing the Fyne logo

#### func  GridIcon

```go
func GridIcon() fyne.Resource
```
GridIcon returns a resource containing the standard grid icon for the current theme

#### func  HelpIcon

```go
func HelpIcon() fyne.Resource
```
HelpIcon returns a resource containing the standard help icon for the current theme

#### func  HistoryIcon

```go
func HistoryIcon() fyne.Resource
```
HistoryIcon returns a resource containing the standard history icon for the current theme

#### func  HomeIcon

```go
func HomeIcon() fyne.Resource
```
HomeIcon returns a resource containing the standard home folder icon for the current theme

#### func  HoverColor

```go
func HoverColor() color.Color
```
HoverColor returns the color used to highlight interactive elements currently under a cursor.

#### func  IconInlineSize

```go
func IconInlineSize() float32
```
IconInlineSize is the standard size of icons which appear within buttons, labels etc.

#### func  InfoIcon

```go
func InfoIcon() fyne.Resource
```
InfoIcon returns a resource containing the standard dialog info icon for the current theme

#### func  InnerPadding

```go
func InnerPadding() float32
```
InnerPadding is the standard gap between element content and the outside edge of a widget.


<div class="since">Since: <code>
2.3</code></div>

#### func  InputBackgroundColor

```go
func InputBackgroundColor() color.Color
```
InputBackgroundColor returns the color used to draw underneath input elements.

#### func  InputBorderColor

```go
func InputBorderColor() color.Color
```
InputBorderColor returns the color used to draw underneath input elements.


<div class="since">Since: <code>
2.3</code></div>

#### func  InputBorderSize

```go
func InputBorderSize() float32
```
InputBorderSize returns the input border size (or underline size for an entry).


<div class="since">Since: <code>
2.0</code></div>

#### func  LightTheme

```go
func LightTheme() fyne.Theme
```
LightTheme defines the built-in light theme colors and sizes.


<div class="deprecated">
Deprecated: This method ignores user preference and should not be used, it will be removed in v3.0.</div>

#### func  LineSpacing

```go
func LineSpacing() float32
```
LineSpacing is the default gap between multiple lines of text.


<div class="since">Since: <code>
2.3</code></div>

#### func  ListIcon

```go
func ListIcon() fyne.Resource
```
ListIcon returns a resource containing the standard list icon for the current theme

#### func  LoginIcon

```go
func LoginIcon() fyne.Resource
```
LoginIcon returns a resource containing the standard login icon for the current theme

#### func  LogoutIcon

```go
func LogoutIcon() fyne.Resource
```
LogoutIcon returns a resource containing the standard logout icon for the current theme

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

#### func  MediaMusicIcon

```go
func MediaMusicIcon() fyne.Resource
```
MediaMusicIcon returns a resource containing the standard media music icon for the current theme


<div class="since">Since: <code>
2.1</code></div>

#### func  MediaPauseIcon

```go
func MediaPauseIcon() fyne.Resource
```
MediaPauseIcon returns a resource containing the standard media pause icon for the current theme

#### func  MediaPhotoIcon

```go
func MediaPhotoIcon() fyne.Resource
```
MediaPhotoIcon returns a resource containing the standard media photo icon for the current theme


<div class="since">Since: <code>
2.1</code></div>

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

#### func  MediaStopIcon

```go
func MediaStopIcon() fyne.Resource
```
MediaStopIcon returns a resource containing the standard media stop icon for the current theme

#### func  MediaVideoIcon

```go
func MediaVideoIcon() fyne.Resource
```
MediaVideoIcon returns a resource containing the standard media video icon for the current theme


<div class="since">Since: <code>
2.1</code></div>

#### func  MenuBackgroundColor

```go
func MenuBackgroundColor() color.Color
```
MenuBackgroundColor returns the theme's background color for menus.


<div class="since">Since: <code>
2.3</code></div>

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

#### func  MoreHorizontalIcon

```go
func MoreHorizontalIcon() fyne.Resource
```
MoreHorizontalIcon returns a resource containing the standard horizontal more icon for the current theme

#### func  MoreVerticalIcon

```go
func MoreVerticalIcon() fyne.Resource
```
MoreVerticalIcon returns a resource containing the standard vertical more icon for the current theme

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

#### func  OverlayBackgroundColor

```go
func OverlayBackgroundColor() color.Color
```
OverlayBackgroundColor returns the theme's background color for overlays like dialogs.


<div class="since">Since: <code>
2.3</code></div>

#### func  Padding

```go
func Padding() float32
```
Padding is the standard gap between elements and the border around interface elements.

#### func  PlaceHolderColor

```go
func PlaceHolderColor() color.Color
```
PlaceHolderColor returns the theme's standard text color.

#### func  PressedColor

```go
func PressedColor() color.Color
```
PressedColor returns the color used to overlap tapped features.


<div class="since">Since: <code>
2.0</code></div>

#### func  PrimaryColor

```go
func PrimaryColor() color.Color
```
PrimaryColor returns the color used to highlight primary features.

#### func  PrimaryColorNamed

```go
func PrimaryColorNamed(name string) color.Color
```
PrimaryColorNamed returns a theme specific color value for a named primary color.


<div class="since">Since: <code>
1.4</code></div>

#### func  PrimaryColorNames

```go
func PrimaryColorNames() []string
```
PrimaryColorNames returns a list of the standard primary color options.


<div class="since">Since: <code>
1.4</code></div>

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
ScrollBarColor returns the color (and translucency) for a scrollBar.

#### func  ScrollBarSize

```go
func ScrollBarSize() float32
```
ScrollBarSize is the width (or height) of the bars on a ScrollContainer.

#### func  ScrollBarSmallSize

```go
func ScrollBarSmallSize() float32
```
ScrollBarSmallSize is the width (or height) of the minimized bars on a ScrollContainer.

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

#### func  SelectionColor

```go
func SelectionColor() color.Color
```
SelectionColor returns the color for a selected element.


<div class="since">Since: <code>
2.1</code></div>

#### func  SeparatorColor

```go
func SeparatorColor() color.Color
```
SeparatorColor returns the color for the separator element.


<div class="since">Since: <code>
2.3</code></div>

#### func  SeparatorThicknessSize

```go
func SeparatorThicknessSize() float32
```
SeparatorThicknessSize is the standard thickness of the separator widget.


<div class="since">Since: <code>
2.0</code></div>

#### func  SettingsIcon

```go
func SettingsIcon() fyne.Resource
```
SettingsIcon returns a resource containing the standard settings icon for the current theme

#### func  ShadowColor

```go
func ShadowColor() color.Color
```
ShadowColor returns the color (and translucency) for shadows used for indicating elevation.

#### func  StorageIcon

```go
func StorageIcon() fyne.Resource
```
StorageIcon returns a resource containing the standard storage icon for the current theme

#### func  SuccessColor

```go
func SuccessColor() color.Color
```
SuccessColor returns the theme's success foreground color.


<div class="since">Since: <code>
2.3</code></div>

#### func  TextBoldFont

```go
func TextBoldFont() fyne.Resource
```
TextBoldFont returns the font resource for the bold font style.

#### func  TextBoldItalicFont

```go
func TextBoldItalicFont() fyne.Resource
```
TextBoldItalicFont returns the font resource for the bold and italic font style.

#### func  TextColor

```go
func TextColor() color.Color
```
TextColor returns the theme's standard text color - this is actually the foreground color since 1.4.


<div class="deprecated">
Deprecated: Use theme.ForegroundColor() colour instead.</div>

#### func  TextFont

```go
func TextFont() fyne.Resource
```
TextFont returns the font resource for the regular font style.

#### func  TextHeadingSize

```go
func TextHeadingSize() float32
```
TextHeadingSize returns the text size for header text.


<div class="since">Since: <code>
2.1</code></div>

#### func  TextItalicFont

```go
func TextItalicFont() fyne.Resource
```
TextItalicFont returns the font resource for the italic font style.

#### func  TextMonospaceFont

```go
func TextMonospaceFont() fyne.Resource
```
TextMonospaceFont returns the font resource for the monospace font face.

#### func  TextSize

```go
func TextSize() float32
```
TextSize returns the standard text size.

#### func  TextSubHeadingSize

```go
func TextSubHeadingSize() float32
```
TextSubHeadingSize returns the text size for sub-header text.


<div class="since">Since: <code>
2.1</code></div>

#### func  UploadIcon

```go
func UploadIcon() fyne.Resource
```
UploadIcon returns a resource containing the standard upload icon for the current theme

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
VisibilityIcon returns a resource containing the standard visibility icon for the current theme

#### func  VisibilityOffIcon

```go
func VisibilityOffIcon() fyne.Resource
```
VisibilityOffIcon returns a resource containing the standard visibility off icon for the current theme

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

#### func  WarningColor

```go
func WarningColor() color.Color
```
WarningColor returns the theme's warning foreground color.


<div class="since">Since: <code>
2.3</code></div>

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
 * [ErrorThemedResource](errorthemedresource.html)
 * [InvertedThemedResource](invertedthemedresource.html)
 * [PrimaryThemedResource](primarythemedresource.html)
 * [ThemedResource](themedresource.html)

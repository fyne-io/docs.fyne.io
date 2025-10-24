---
tags: [api]
title: theme (package)
slug: (package)

aliases:
- /api/v2.7/theme

package: fyne.io/fyne/v2/theme
---


---
```go
import "fyne.io/fyne/v2/theme"
```

Package theme defines how a Fyne app should look when rendered.

#

```go
const (
	// ColorRed is the red primary color name.
	//
	// Since: 1.4
	ColorRed = internaltheme.ColorRed
	// ColorOrange is the orange primary color name.
	//
	// Since: 1.4
	ColorOrange = internaltheme.ColorOrange
	// ColorYellow is the yellow primary color name.
	//
	// Since: 1.4
	ColorYellow = internaltheme.ColorYellow
	// ColorGreen is the green primary color name.
	//
	// Since: 1.4
	ColorGreen = internaltheme.ColorGreen
	// ColorBlue is the blue primary color name.
	//
	// Since: 1.4
	ColorBlue = internaltheme.ColorBlue
	// ColorPurple is the purple primary color name.
	//
	// Since: 1.4
	ColorPurple = internaltheme.ColorPurple
	// ColorBrown is the brown primary color name.
	//
	// Since: 1.4
	ColorBrown = internaltheme.ColorBrown
	// ColorGray is the gray primary color name.
	//
	// Since: 1.4
	ColorGray = internaltheme.ColorGray

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

	// ColorNameError is the name of theme lookup for error color.
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

	// ColorNameForegroundOnError is the name of theme lookup for a contrast color to the error color.
	//
	// Since: 2.5
	ColorNameForegroundOnError fyne.ThemeColorName = "foregroundOnError"

	// ColorNameForegroundOnPrimary is the name of theme lookup for a contrast color to the primary color.
	//
	// Since: 2.5
	ColorNameForegroundOnPrimary fyne.ThemeColorName = "foregroundOnPrimary"

	// ColorNameForegroundOnSuccess is the name of theme lookup for a contrast color to the success color.
	//
	// Since: 2.5
	ColorNameForegroundOnSuccess fyne.ThemeColorName = "foregroundOnSuccess"

	// ColorNameForegroundOnWarning is the name of theme lookup for a contrast color to the warning color.
	//
	// Since: 2.5
	ColorNameForegroundOnWarning fyne.ThemeColorName = "foregroundOnWarning"

	// ColorNameHeaderBackground is the name of theme lookup for background color of a collection header.
	//
	// Since: 2.4
	ColorNameHeaderBackground fyne.ThemeColorName = "headerBackground"

	// ColorNameHover is the name of theme lookup for hover color.
	//
	// Since: 2.0
	ColorNameHover fyne.ThemeColorName = "hover"

	// ColorNameHyperlink is the name of theme lookup for hyperlink color.
	//
	// Since: 2.4
	ColorNameHyperlink fyne.ThemeColorName = "hyperlink"

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

	// ColorNameScrollBarBackground is the name of theme lookup for scrollbar background color.
	//
	// Since: 2.6
	ColorNameScrollBarBackground fyne.ThemeColorName = "scrollBarBackground"

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

	// ColorNameSuccess is the name of theme lookup for success color.
	//
	// Since: 2.3
	ColorNameSuccess fyne.ThemeColorName = "success"

	// ColorNameWarning is the name of theme lookup for warning color.
	//
	// Since: 2.3
	ColorNameWarning fyne.ThemeColorName = "warning"
)
```
Keep in mind to add new constants to the tests at test/theme.go.

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

	// IconNameCheckButton is the name of theme lookup for unchecked check button icon.
	//
	// Since: 2.0
	IconNameCheckButton fyne.ThemeIconName = "unchecked"

	// IconNameCheckButtonChecked is the name of theme lookup for checked check button icon.
	//
	// Since: 2.0
	IconNameCheckButtonChecked fyne.ThemeIconName = "checked"

	// IconNameCheckButtonFill is the name of theme lookup for filled check button icon.
	//
	// Since: 2.5
	IconNameCheckButtonFill fyne.ThemeIconName = "iconNameCheckButtonFill"

	// IconNameCheckButtonPartial is the name of theme lookup for "partially" checked check button icon.
	//
	// Since: 2.6
	IconNameCheckButtonPartial fyne.ThemeIconName = "partial"

	// IconNameRadioButton is the name of theme lookup for radio button unchecked icon.
	//
	// Since: 2.0
	IconNameRadioButton fyne.ThemeIconName = "radioButton"

	// IconNameRadioButtonChecked is the name of theme lookup for radio button checked icon.
	//
	// Since: 2.0
	IconNameRadioButtonChecked fyne.ThemeIconName = "radioButtonChecked"

	// IconNameRadioButtonFill is the name of theme lookup for filled radio button icon.
	//
	// Since: 2.5
	IconNameRadioButtonFill fyne.ThemeIconName = "iconNameRadioButtonFill"

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

	// IconNameBrokenImage is the name of the theme lookup for broken-image icon.
	//
	// Since: 2.4
	IconNameBrokenImage fyne.ThemeIconName = "broken-image"

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

	// IconNameDragCornerIndicator is the name of the icon used in inner windows to indicate a draggable corner.
	//
	// Since: 2.5
	IconNameDragCornerIndicator fyne.ThemeIconName = "dragCornerIndicator"

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

	// IconNameDesktop is the name of theme lookup for desktop icon.
	//
	// Since: 2.5
	IconNameDesktop fyne.ThemeIconName = "desktop"

	// IconNameAccount is the name of theme lookup for account icon.
	//
	// Since: 2.1
	IconNameAccount fyne.ThemeIconName = "account"

	// IconNameCalendar is the name of theme lookup for calendar icon.
	//
	// Since: 2.6
	IconNameCalendar fyne.ThemeIconName = "calendar"

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

	// IconNameWindowClose is the name of theme lookup for window close icon.
	//
	// Since: 2.5
	IconNameWindowClose fyne.ThemeIconName = "windowClose"

	// IconNameWindowMaximize is the name of theme lookup for window maximize icon.
	//
	// Since: 2.5
	IconNameWindowMaximize fyne.ThemeIconName = "windowMaximize"

	// IconNameWindowMinimize is the name of theme lookup for window minimize icon.
	//
	// Since: 2.5
	IconNameWindowMinimize fyne.ThemeIconName = "windowMinimize"
)
```

```go
const (
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

	// SizeNameInputRadius is the name of theme lookup for input corner radius.
	//
	// Since: 2.4
	SizeNameInputRadius fyne.ThemeSizeName = "inputRadius"

	// SizeNameSelectionRadius is the name of theme lookup for selection corner radius.
	//
	// Since: 2.4
	SizeNameSelectionRadius fyne.ThemeSizeName = "selectionRadius"

	// SizeNameScrollBarRadius is the name of theme lookup for the scroll bar corner radius.
	//
	// Since: 2.5
	SizeNameScrollBarRadius fyne.ThemeSizeName = "scrollBarRadius"

	// SizeNameWindowButtonHeight is the name of the height for an inner window titleBar button.
	//
	// Since: 2.6
	SizeNameWindowButtonHeight fyne.ThemeSizeName = "windowButtonHeight"

	// SizeNameWindowButtonRadius is the name of the radius for an inner window titleBar button.
	//
	// Since: 2.6
	SizeNameWindowButtonRadius fyne.ThemeSizeName = "windowButtonRadius"

	// SizeNameWindowButtonIcon is the name of the width of an inner window titleBar button.
	//
	// Since: 2.6
	SizeNameWindowButtonIcon fyne.ThemeSizeName = "windowButtonIcon"

	// SizeNameWindowTitleBarHeight is the height for inner window titleBars.
	//
	// Since: 2.6
	SizeNameWindowTitleBarHeight fyne.ThemeSizeName = "windowTitleBarHeight"
)
```

```go
const (
	// VariantDark is the version of a theme that satisfies a user preference for a dark look.
	//
	// Since: 2.0
	VariantDark = internaltheme.VariantDark

	// VariantLight is the version of a theme that satisfies a user preference for a light look.
	//
	// Since: 2.0
	VariantLight = internaltheme.VariantLight
)
```
Keep in mind to add new constants to the tests at test/theme.go.

###

```go
func AccountIcon() fyne.Resource
```
AccountIcon returns a resource containing the standard account icon for the current theme

###

```go
func BackgroundColor() color.Color
```
BackgroundColor returns the theme's background color.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameBackground) instead.</div>

###

```go
func BrokenImageIcon() fyne.Resource
```
BrokenImageIcon returns a resource containing an icon to specify a broken or missing image


<div class="since">Since: <code>
2.4</code></div>

###

```go
func ButtonColor() color.Color
```
ButtonColor returns the theme's standard button color.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameButton) instead.</div>

###

```go
func CalendarIcon() fyne.Resource
```
CalendarIcon returns a resource containing the standard account icon for the current theme


<div class="since">Since: <code>
2.6</code></div>

###

```go
func CancelIcon() fyne.Resource
```
CancelIcon returns a resource containing the standard cancel icon for the current theme

###

```go
func CaptionTextSize() float32
```
CaptionTextSize returns the size for caption text.

###

```go
func CheckButtonCheckedIcon() fyne.Resource
```
CheckButtonCheckedIcon returns a resource containing the standard checkbox checked icon for the current theme

###

```go
func CheckButtonFillIcon() fyne.Resource
```
CheckButtonFillIcon returns a resource containing the filled checkbox icon for the current theme.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func CheckButtonIcon() fyne.Resource
```
CheckButtonIcon returns a resource containing the standard checkbox icon for the current theme

###

```go
func Color(name fyne.ThemeColorName) color.Color
```
Color looks up the named colour for current theme and variant.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func ColorAchromaticIcon() fyne.Resource
```
ColorAchromaticIcon returns a resource containing the standard achromatic color icon for the current theme

###

```go
func ColorChromaticIcon() fyne.Resource
```
ColorChromaticIcon returns a resource containing the standard chromatic color icon for the current theme

###

```go
func ColorForWidget(name fyne.ThemeColorName, w fyne.Widget) color.Color
```
ColorForWidget looks up the named colour for the requested widget using the current theme and variant. If the widget theme has been overridden that theme will be used.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func ColorPaletteIcon() fyne.Resource
```
ColorPaletteIcon returns a resource containing the standard color palette icon for the current theme

###

```go
func ComputerIcon() fyne.Resource
```
ComputerIcon returns a resource containing the standard computer icon for the current theme

###

```go
func ConfirmIcon() fyne.Resource
```
ConfirmIcon returns a resource containing the standard confirm icon for the current theme

###

```go
func ContentAddIcon() fyne.Resource
```
ContentAddIcon returns a resource containing the standard content add icon for the current theme

###

```go
func ContentClearIcon() fyne.Resource
```
ContentClearIcon returns a resource containing the standard content clear icon for the current theme

###

```go
func ContentCopyIcon() fyne.Resource
```
ContentCopyIcon returns a resource containing the standard content copy icon for the current theme

###

```go
func ContentCutIcon() fyne.Resource
```
ContentCutIcon returns a resource containing the standard content cut icon for the current theme

###

```go
func ContentPasteIcon() fyne.Resource
```
ContentPasteIcon returns a resource containing the standard content paste icon for the current theme

###

```go
func ContentRedoIcon() fyne.Resource
```
ContentRedoIcon returns a resource containing the standard content redo icon for the current theme

###

```go
func ContentRemoveIcon() fyne.Resource
```
ContentRemoveIcon returns a resource containing the standard content remove icon for the current theme

###

```go
func ContentUndoIcon() fyne.Resource
```
ContentUndoIcon returns a resource containing the standard content undo icon for the current theme

###

```go
func Current() fyne.Theme
```
Current returns the theme that is currently used for the running application. It looks up based on user preferences and application configuration.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func CurrentForWidget(w fyne.CanvasObject) fyne.Theme
```
CurrentForWidget returns the theme that is currently used for the specified widget. It looks for widget overrides and falls back to the application's current theme.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func DarkTheme() fyne.Theme
```
DarkTheme defines the built-in dark theme colors and sizes.


<div class="deprecated">
Deprecated: This method ignores user preference and should not be used, it will be removed in v3.0.</div> If developers want to ignore user preference for theme variant they can set a custom theme.

###

```go
func DefaultEmojiFont() fyne.Resource
```
DefaultEmojiFont returns the font resource for the built-in emoji font. This may return nil if the application was packaged without an emoji font.


<div class="since">Since: <code>
2.4</code></div>

###

```go
func DefaultSymbolFont() fyne.Resource
```
DefaultSymbolFont returns the font resource for the built-in symbol font.


<div class="since">Since: <code>
2.2</code></div>

###

```go
func DefaultTextBoldFont() fyne.Resource
```
DefaultTextBoldFont returns the font resource for the built-in bold font style.

###

```go
func DefaultTextBoldItalicFont() fyne.Resource
```
DefaultTextBoldItalicFont returns the font resource for the built-in bold and italic font style.

###

```go
func DefaultTextFont() fyne.Resource
```
DefaultTextFont returns the font resource for the built-in regular font style.

###

```go
func DefaultTextItalicFont() fyne.Resource
```
DefaultTextItalicFont returns the font resource for the built-in italic font style.

###

```go
func DefaultTextMonospaceFont() fyne.Resource
```
DefaultTextMonospaceFont returns the font resource for the built-in monospace font face.

###

```go
func DefaultTheme() fyne.Theme
```
DefaultTheme returns a built-in theme that can adapt to the user preference of light or dark colors.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func DeleteIcon() fyne.Resource
```
DeleteIcon returns a resource containing the standard delete icon for the current theme

###

```go
func DesktopIcon() fyne.Resource
```
DesktopIcon returns a resource containing the standard desktop icon for the current theme

###

```go
func DisabledButtonColor() color.Color
```
DisabledButtonColor returns the theme's disabled button color.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameDisabledButton) instead.</div>

###

```go
func DisabledColor() color.Color
```
DisabledColor returns the foreground color for a disabled UI element.


<div class="since">Since: <code>
2.0</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameDisabled) instead.</div>

###

```go
func DisabledTextColor() color.Color
```
DisabledTextColor returns the theme's disabled text color - this is actually the disabled color since 1.4.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameDisabled) instead.</div>

###

```go
func DocumentCreateIcon() fyne.Resource
```
DocumentCreateIcon returns a resource containing the standard document create icon for the current theme

###

```go
func DocumentIcon() fyne.Resource
```
DocumentIcon returns a resource containing the standard document icon for the current theme

###

```go
func DocumentPrintIcon() fyne.Resource
```
DocumentPrintIcon returns a resource containing the standard document print icon for the current theme

###

```go
func DocumentSaveIcon() fyne.Resource
```
DocumentSaveIcon returns a resource containing the standard document save icon for the current theme

###

```go
func DownloadIcon() fyne.Resource
```
DownloadIcon returns a resource containing the standard download icon for the current theme

###

```go
func ErrorColor() color.Color
```
ErrorColor returns the theme's error foreground color.


<div class="since">Since: <code>
2.0</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameError) instead.</div>

###

```go
func ErrorIcon() fyne.Resource
```
ErrorIcon returns a resource containing the standard dialog error icon for the current theme

###

```go
func FileApplicationIcon() fyne.Resource
```
FileApplicationIcon returns a resource containing the file icon representing application files for the current theme

###

```go
func FileAudioIcon() fyne.Resource
```
FileAudioIcon returns a resource containing the file icon representing audio files for the current theme

###

```go
func FileIcon() fyne.Resource
```
FileIcon returns a resource containing the appropriate file icon for the current theme

###

```go
func FileImageIcon() fyne.Resource
```
FileImageIcon returns a resource containing the file icon representing image files for the current theme

###

```go
func FileTextIcon() fyne.Resource
```
FileTextIcon returns a resource containing the file icon representing text files for the current theme

###

```go
func FileVideoIcon() fyne.Resource
```
FileVideoIcon returns a resource containing the file icon representing video files for the current theme

###

```go
func FocusColor() color.Color
```
FocusColor returns the color used to highlight a focused widget.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameFocus) instead.</div>

###

```go
func FolderIcon() fyne.Resource
```
FolderIcon returns a resource containing the standard folder icon for the current theme

###

```go
func FolderNewIcon() fyne.Resource
```
FolderNewIcon returns a resource containing the standard folder creation icon for the current theme

###

```go
func FolderOpenIcon() fyne.Resource
```
FolderOpenIcon returns a resource containing the standard folder open icon for the current theme

###

```go
func Font(style fyne.TextStyle) fyne.Resource
```
Font looks up the font for current theme and text style.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func ForegroundColor() color.Color
```
ForegroundColor returns the theme's standard foreground color for text and icons.


<div class="since">Since: <code>
2.0</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameForeground) instead.</div>

###

```go
func FromJSON(data string) (fyne.Theme, error)
```
FromJSON returns a Theme created from the given JSON metadata. Any values not present in the data will fall back to the default theme. If a parse error occurs it will be returned along with a default theme.


<div class="since">Since: <code>
2.2</code></div>

###

```go
func FromJSONReader(r io.Reader) (fyne.Theme, error)
```
FromJSONReader returns a Theme created from the given JSON metadata through the reader. Any values not present in the data will fall back to the default theme. If a parse error occurs it will be returned along with a default theme.


<div class="since">Since: <code>
2.2</code></div>

###

```go
func FromJSONReaderWithFallback(r io.Reader, fallback fyne.Theme) (fyne.Theme, error)
```
FromJSONReaderWithFallback returns a Theme created from the given JSON metadata through the reader. Any values not present in the data will fall back to the specified theme. If a parse error occurs it will be returned along with a specified fallback theme.


<div class="since">Since: <code>
2.7</code></div>

###

```go
func FromJSONWithFallback(data string, fallback fyne.Theme) (fyne.Theme, error)
```
FromJSONWithFallback returns a Theme created from the given JSON metadata. Any values not present in the data will fall back to the specified theme. If a parse error occurs it will be returned along with a specified fallback theme.


<div class="since">Since: <code>
2.7</code></div>

###

```go
func FromLegacy(t fyne.LegacyTheme) fyne.Theme
```
FromLegacy returns a 2.0 Theme created from the given LegacyTheme data. This is a transition path and will be removed in the future (probably version 3.0).


<div class="since">Since: <code>
2.0</code></div>

###

```go
func FyneLogo() fyne.Resource
```
FyneLogo returns a resource containing the Fyne logo.


<div class="deprecated">
Deprecated: Applications should use their own icon in most cases.</div>

###

```go
func GridIcon() fyne.Resource
```
GridIcon returns a resource containing the standard grid icon for the current theme

###

```go
func HeaderBackgroundColor() color.Color
```
HeaderBackgroundColor returns the color used to draw underneath collection headers.


<div class="since">Since: <code>
2.4</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameHeaderBackground) instead.</div>

###

```go
func HelpIcon() fyne.Resource
```
HelpIcon returns a resource containing the standard help icon for the current theme

###

```go
func HistoryIcon() fyne.Resource
```
HistoryIcon returns a resource containing the standard history icon for the current theme

###

```go
func HomeIcon() fyne.Resource
```
HomeIcon returns a resource containing the standard home folder icon for the current theme

###

```go
func HoverColor() color.Color
```
HoverColor returns the color used to highlight interactive elements currently under a cursor.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameHover) instead.</div>

###

```go
func HyperlinkColor() color.Color
```
HyperlinkColor returns the color used for the Hyperlink widget and hyperlink text elements.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameHyperlink) instead.</div>

###

```go
func Icon(name fyne.ThemeIconName) fyne.Resource
```
Icon looks up the specified icon for current theme.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func IconForWidget(name fyne.ThemeIconName, w fyne.Widget) fyne.Resource
```
IconForWidget looks up the specified icon for requested widget using the current theme. If the widget theme has been overridden that theme will be used.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func IconInlineSize() float32
```
IconInlineSize is the standard size of icons which appear within buttons, labels etc.

###

```go
func InfoIcon() fyne.Resource
```
InfoIcon returns a resource containing the standard dialog info icon for the current theme

###

```go
func InnerPadding() float32
```
InnerPadding is the standard gap between element content and the outside edge of a widget.


<div class="since">Since: <code>
2.3</code></div>

###

```go
func InputBackgroundColor() color.Color
```
InputBackgroundColor returns the color used to draw underneath input elements.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameInputBackground) instead.</div>

###

```go
func InputBorderColor() color.Color
```
InputBorderColor returns the color used to draw underneath input elements.


<div class="since">Since: <code>
2.3</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameInputBorder) instead.</div>

###

```go
func InputBorderSize() float32
```
InputBorderSize returns the input border size (or underline size for an entry).


<div class="since">Since: <code>
2.0</code></div>

###

```go
func InputRadiusSize() float32
```
InputRadiusSize returns the input radius size.


<div class="since">Since: <code>
2.4</code></div>

###

```go
func LightTheme() fyne.Theme
```
LightTheme defines the built-in light theme colors and sizes.


<div class="deprecated">
Deprecated: This method ignores user preference and should not be used, it will be removed in v3.0.</div> If developers want to ignore user preference for theme variant they can set a custom theme.

###

```go
func LineSpacing() float32
```
LineSpacing is the default gap between multiple lines of text.


<div class="since">Since: <code>
2.3</code></div>

###

```go
func ListIcon() fyne.Resource
```
ListIcon returns a resource containing the standard list icon for the current theme

###

```go
func LoginIcon() fyne.Resource
```
LoginIcon returns a resource containing the standard login icon for the current theme

###

```go
func LogoutIcon() fyne.Resource
```
LogoutIcon returns a resource containing the standard logout icon for the current theme

###

```go
func MailAttachmentIcon() fyne.Resource
```
MailAttachmentIcon returns a resource containing the standard mail attachment icon for the current theme

###

```go
func MailComposeIcon() fyne.Resource
```
MailComposeIcon returns a resource containing the standard mail compose icon for the current theme

###

```go
func MailForwardIcon() fyne.Resource
```
MailForwardIcon returns a resource containing the standard mail forward icon for the current theme

###

```go
func MailReplyAllIcon() fyne.Resource
```
MailReplyAllIcon returns a resource containing the standard mail reply all icon for the current theme

###

```go
func MailReplyIcon() fyne.Resource
```
MailReplyIcon returns a resource containing the standard mail reply icon for the current theme

###

```go
func MailSendIcon() fyne.Resource
```
MailSendIcon returns a resource containing the standard mail send icon for the current theme

###

```go
func MediaFastForwardIcon() fyne.Resource
```
MediaFastForwardIcon returns a resource containing the standard media fast-forward icon for the current theme

###

```go
func MediaFastRewindIcon() fyne.Resource
```
MediaFastRewindIcon returns a resource containing the standard media fast-rewind icon for the current theme

###

```go
func MediaMusicIcon() fyne.Resource
```
MediaMusicIcon returns a resource containing the standard media music icon for the current theme


<div class="since">Since: <code>
2.1</code></div>

###

```go
func MediaPauseIcon() fyne.Resource
```
MediaPauseIcon returns a resource containing the standard media pause icon for the current theme

###

```go
func MediaPhotoIcon() fyne.Resource
```
MediaPhotoIcon returns a resource containing the standard media photo icon for the current theme


<div class="since">Since: <code>
2.1</code></div>

###

```go
func MediaPlayIcon() fyne.Resource
```
MediaPlayIcon returns a resource containing the standard media play icon for the current theme

###

```go
func MediaRecordIcon() fyne.Resource
```
MediaRecordIcon returns a resource containing the standard media record icon for the current theme

###

```go
func MediaReplayIcon() fyne.Resource
```
MediaReplayIcon returns a resource containing the standard media replay icon for the current theme

###

```go
func MediaSkipNextIcon() fyne.Resource
```
MediaSkipNextIcon returns a resource containing the standard media skip next icon for the current theme

###

```go
func MediaSkipPreviousIcon() fyne.Resource
```
MediaSkipPreviousIcon returns a resource containing the standard media skip previous icon for the current theme

###

```go
func MediaStopIcon() fyne.Resource
```
MediaStopIcon returns a resource containing the standard media stop icon for the current theme

###

```go
func MediaVideoIcon() fyne.Resource
```
MediaVideoIcon returns a resource containing the standard media video icon for the current theme


<div class="since">Since: <code>
2.1</code></div>

###

```go
func MenuBackgroundColor() color.Color
```
MenuBackgroundColor returns the theme's background color for menus.


<div class="since">Since: <code>
2.3</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameMenuBackground) instead.</div>

###

```go
func MenuDropDownIcon() fyne.Resource
```
MenuDropDownIcon returns a resource containing the standard menu drop down icon for the current theme

###

```go
func MenuDropUpIcon() fyne.Resource
```
MenuDropUpIcon returns a resource containing the standard menu drop up icon for the current theme

###

```go
func MenuExpandIcon() fyne.Resource
```
MenuExpandIcon returns a resource containing the standard (mobile) expand "submenu icon for the current theme

###

```go
func MenuIcon() fyne.Resource
```
MenuIcon returns a resource containing the standard (mobile) menu icon for the current theme

###

```go
func MoreHorizontalIcon() fyne.Resource
```
MoreHorizontalIcon returns a resource containing the standard horizontal more icon for the current theme

###

```go
func MoreVerticalIcon() fyne.Resource
```
MoreVerticalIcon returns a resource containing the standard vertical more icon for the current theme

###

```go
func MoveDownIcon() fyne.Resource
```
MoveDownIcon returns a resource containing the standard down arrow icon for the current theme

###

```go
func MoveUpIcon() fyne.Resource
```
MoveUpIcon returns a resource containing the standard up arrow icon for the current theme

###

```go
func NavigateBackIcon() fyne.Resource
```
NavigateBackIcon returns a resource containing the standard backward navigation icon for the current theme

###

```go
func NavigateNextIcon() fyne.Resource
```
NavigateNextIcon returns a resource containing the standard forward navigation icon for the current theme

###

```go
func OverlayBackgroundColor() color.Color
```
OverlayBackgroundColor returns the theme's background color for overlays like dialogs.


<div class="since">Since: <code>
2.3</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameOverlayBackground) instead.</div>

###

```go
func Padding() float32
```
Padding is the standard gap between elements and the border around interface elements.

###

```go
func PlaceHolderColor() color.Color
```
PlaceHolderColor returns the theme's standard text color.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNamePlaceHolder) instead.</div>

###

```go
func PressedColor() color.Color
```
PressedColor returns the color used to overlap tapped features.


<div class="since">Since: <code>
2.0</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNamePressed) instead.</div>

###

```go
func PrimaryColor() color.Color
```
PrimaryColor returns the color used to highlight primary features.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNamePrimary) instead.</div>

###

```go
func PrimaryColorNamed(name string) color.Color
```
PrimaryColorNamed returns a theme specific color value for a named primary color.


<div class="since">Since: <code>
1.4</code></div>


<div class="deprecated">
Deprecated: You should not access named primary colors but access the primary color using Color(theme.ColorNamePrimary) instead.</div>

###

```go
func PrimaryColorNames() []string
```
PrimaryColorNames returns a list of the standard primary color options.


<div class="since">Since: <code>
1.4</code></div>

###

```go
func QuestionIcon() fyne.Resource
```
QuestionIcon returns a resource containing the standard dialog question icon for the current theme

###

```go
func RadioButtonCheckedIcon() fyne.Resource
```
RadioButtonCheckedIcon returns a resource containing the standard radio button checked icon for the current theme

###

```go
func RadioButtonFillIcon() fyne.Resource
```
RadioButtonFillIcon returns a resource containing the filled checkbox icon for the current theme.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func RadioButtonIcon() fyne.Resource
```
RadioButtonIcon returns a resource containing the standard radio button icon for the current theme

###

```go
func ScrollBarColor() color.Color
```
ScrollBarColor returns the color (and translucency) for a scrollBar.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameScrollBar) instead.</div>

###

```go
func ScrollBarSize() float32
```
ScrollBarSize is the width (or height) of the bars on a ScrollContainer.

###

```go
func ScrollBarSmallSize() float32
```
ScrollBarSmallSize is the width (or height) of the minimized bars on a ScrollContainer.

###

```go
func SearchIcon() fyne.Resource
```
SearchIcon returns a resource containing the standard search icon for the current theme

###

```go
func SearchReplaceIcon() fyne.Resource
```
SearchReplaceIcon returns a resource containing the standard search and replace icon for the current theme

###

```go
func SelectionColor() color.Color
```
SelectionColor returns the color for a selected element.


<div class="since">Since: <code>
2.1</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameSelection) instead.</div>

###

```go
func SelectionRadiusSize() float32
```
SelectionRadiusSize returns the selection highlight radius size.


<div class="since">Since: <code>
2.4</code></div>

###

```go
func SeparatorColor() color.Color
```
SeparatorColor returns the color for the separator element.


<div class="since">Since: <code>
2.3</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameSeparator) instead.</div>

###

```go
func SeparatorThicknessSize() float32
```
SeparatorThicknessSize is the standard thickness of the separator widget.


<div class="since">Since: <code>
2.0</code></div>

###

```go
func SettingsIcon() fyne.Resource
```
SettingsIcon returns a resource containing the standard settings icon for the current theme

###

```go
func ShadowColor() color.Color
```
ShadowColor returns the color (and translucency) for shadows used for indicating elevation.


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameShadow) instead.</div>

###

```go
func Size(name fyne.ThemeSizeName) float32
```
Size looks up the specified size for current theme.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func SizeForWidget(name fyne.ThemeSizeName, w fyne.Widget) float32
```
SizeForWidget looks up the specified size for the requested widget using the current theme. If the widget theme has been overridden that theme will be used.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func StorageIcon() fyne.Resource
```
StorageIcon returns a resource containing the standard storage icon for the current theme

###

```go
func SuccessColor() color.Color
```
SuccessColor returns the theme's success foreground color.


<div class="since">Since: <code>
2.3</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameSuccess) instead.</div>

###

```go
func SymbolFont() fyne.Resource
```
SymbolFont returns the font resource for the symbol font style.


<div class="since">Since: <code>
2.4</code></div>

###

```go
func TextBoldFont() fyne.Resource
```
TextBoldFont returns the font resource for the bold font style.

###

```go
func TextBoldItalicFont() fyne.Resource
```
TextBoldItalicFont returns the font resource for the bold and italic font style.

###

```go
func TextColor() color.Color
```
TextColor returns the theme's standard text color - this is actually the foreground color since 1.4.


<div class="deprecated">
Deprecated: Use theme.ForegroundColor() colour instead.</div>

###

```go
func TextFont() fyne.Resource
```
TextFont returns the font resource for the regular font style.

###

```go
func TextHeadingSize() float32
```
TextHeadingSize returns the text size for header text.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func TextItalicFont() fyne.Resource
```
TextItalicFont returns the font resource for the italic font style.

###

```go
func TextMonospaceFont() fyne.Resource
```
TextMonospaceFont returns the font resource for the monospace font face.

###

```go
func TextSize() float32
```
TextSize returns the standard text size.

###

```go
func TextSubHeadingSize() float32
```
TextSubHeadingSize returns the text size for sub-header text.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func UploadIcon() fyne.Resource
```
UploadIcon returns a resource containing the standard upload icon for the current theme

###

```go
func ViewFullScreenIcon() fyne.Resource
```
ViewFullScreenIcon returns a resource containing the standard fullscreen icon for the current theme

###

```go
func ViewRefreshIcon() fyne.Resource
```
ViewRefreshIcon returns a resource containing the standard refresh icon for the current theme

###

```go
func ViewRestoreIcon() fyne.Resource
```
ViewRestoreIcon returns a resource containing the standard exit fullscreen icon for the current theme

###

```go
func VisibilityIcon() fyne.Resource
```
VisibilityIcon returns a resource containing the standard visibility icon for the current theme

###

```go
func VisibilityOffIcon() fyne.Resource
```
VisibilityOffIcon returns a resource containing the standard visibility off icon for the current theme

###

```go
func VolumeDownIcon() fyne.Resource
```
VolumeDownIcon returns a resource containing the standard volume down icon for the current theme

###

```go
func VolumeMuteIcon() fyne.Resource
```
VolumeMuteIcon returns a resource containing the standard volume mute icon for the current theme

###

```go
func VolumeUpIcon() fyne.Resource
```
VolumeUpIcon returns a resource containing the standard volume up icon for the current theme

###

```go
func WarningColor() color.Color
```
WarningColor returns the theme's warning foreground color.


<div class="since">Since: <code>
2.3</code></div>


<div class="deprecated">
Deprecated: Use Color(theme.ColorNameWarning) instead.</div>

###

```go
func WarningIcon() fyne.Resource
```
WarningIcon returns a resource containing the standard dialog warning icon for the current theme

###

```go
func WindowCloseIcon() fyne.Resource
```
WindowCloseIcon returns a resource containing the window close icon for the current theme


<div class="since">Since: <code>
2.5</code></div>

###

```go
func WindowMaximizeIcon() fyne.Resource
```
WindowMaximizeIcon returns a resource containing the window maximize icon for the current theme


<div class="since">Since: <code>
2.5</code></div>

###

```go
func WindowMinimizeIcon() fyne.Resource
```
WindowMinimizeIcon returns a resource containing the window minimize icon for the current theme


<div class="since">Since: <code>
2.5</code></div>

###

```go
func ZoomFitIcon() fyne.Resource
```
ZoomFitIcon returns a resource containing the standard zoom fit icon for the current theme

###

```go
func ZoomInIcon() fyne.Resource
```
ZoomInIcon returns a resource containing the standard zoom in icon for the current theme

###

```go
func ZoomOutIcon() fyne.Resource
```
ZoomOutIcon returns a resource containing the standard zoom out icon for the current theme

###

 * [DisabledResource](disabledresource.html)
 * [ErrorThemedResource](errorthemedresource.html)
 * [InvertedThemedResource](invertedthemedresource.html)
 * [PrimaryThemedResource](primarythemedresource.html)
 * [ThemedResource](themedresource.html)

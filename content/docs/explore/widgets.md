---
layout: page
title: Widget List
readonly: true
slug: widgets

weight: 2030

aliases:
- /explore/widgets.html
- /started/widgets
---

<style>
  html:not([data-dark-mode]) img.widget.light {
    display: visible;
  }

  html:not([data-dark-mode]) img.widget.dark {
    display: none;
  }

  html[data-dark-mode] img.widget.light {
    display: none;
  }

  html[data-dark-mode] img.widget.dark {
    display: visible;
  }
</style>

## Standard Widgets (in `widget` package)

---

### Accordion

Accordion displays a list of AccordionItems. Each item is represented by a button that reveals a detailed view when tapped.

{{% widget name="accordion" %}}

### Activity

Display an animated activity indicator.

{{% widget name="activity" %}}

### Button

[Button](/widget/button) widget has a text label and icon, both are optional.

{{% widget name="button" %}}

### Card

Card widget groups elements with a header and subheader, all are optional.

{{% widget name="card" %}}

### Check

Check widget has a text label and a checked (or unchecked) icon.

{{% widget name="check" %}}

### Entry

[Entry](/widget/entry) widget allows simple text to be input when focused.

{{% widget name="entry" %}}
{{% widget name="entry-valid" %}}
{{% widget name="entry-invalid" %}}

PasswordEntry widget hides text input and adds a button to display the text.

{{% widget name="password" %}}

### FileIcon

FileIcon provides helpful standard icons for various types of file.
It displays the type of file as an indicator icon and shows the extension of the file type.

{{% widget name="fileicon" %}}

### Form

[Form](/widget/form) widget is two column grid where each row has a label and a widget (usually an input). The last row of the grid will contain the appropriate form control buttons if any should be shown.

{{% widget name="form" %}}

### Hyperlink

Hyperlink widget is a text component with appropriate padding and layout. When clicked, the URL opens in your default web browser.

{{% widget name="hyperlink" %}}

### Icon

Icon widget is a basic image component that load's its resource to match the theme.

{{% widget name="icon" %}}

### Label

[Label](/widget/label) widget is a label component with appropriate padding and layout.

{{% widget name="label" %}}

### Progress bar

[ProgressBar](/widget/progressbar) widget creates a horizontal panel that indicates progress.

{{% widget name="progress" %}}

ProgressBarInfinite widget creates a horizontal panel that indicates waiting indefinitely An infinite progress bar loops 0% -> 100% repeatedly until Stop() is called.

{{% widget name="progressinf" %}}

### RadioGroup

RadioGroup widget has a list of text labels and radio check icons next to each.

{{% widget name="radiogroup" %}}

### RichText

RichText widget is a text component that shows various styles and embedded objects.
It supports markdown parsing to construct the widget with ease.

{{% widget name="richtext" %}}

### Select

Select widget has a list of options, with the current one shown, and triggers an event function when clicked.

{{% widget name="select" %}}

### SelectEntry

Select entry widget adds an editable component to the select widget.
Users can select an option or enter their own value.

{{% widget name="selectentry" %}}

### Separator

Separator widget shows a dividing line between other elements.

{{% widget name="separator" %}}

### Slider

Slider if a widget that can slide between two fixed values.

{{% widget name="slider" %}}

### TextGrid

TextGrid is a monospaced grid of characters. This is designed to be used by a text editor, code preview or terminal emulator.

{{% widget name="textgrid" %}}

### Toolbar

[Toolbar](/widget/toolbar) widget creates a horizontal list of tool buttons.

{{% widget name="toolbar" %}}


## Collection Widgets (in `widget` package)

Collection widgets provide advanced caching functionality to provide high performance rendering of massive data. This does lead to a more complex constructor,
but is a good balance for the outcome it enables.
Each of these widgets uses a series of callbacks, the minimum set is defined by their constructor function, which includes the data size, the creation of template items that can be re-used and finally the function that applies data to a widget as it is about to be added to the display.

### List

[List](/collection/list) provides a high performance vertical scroll of many sub-items.

{{% widget name="list" %}}

### Table

[Table](/collection/table) provides a high performance scrolled two dimensional display of many sub-items.

{{% widget name="table" %}}

### Tree

[Tree](/collection/tree) provides a high performance vertical scroll of items that can be expanded to reveal child elements..

{{% widget name="tree" %}}

### GridWrap

GridWrap is a collection widget that display each child item at the same size and wraps them to new lines as required.

{{% widget name="gridwrap" %}}


## Container Widgets (in `container` package)

Container widgets are like regular containers but they provide some additional functionality.

### AppTabs

[AppTabs](/container/apptabs) widget allows switching visible content from a list of TabItems. Each item is represented by a button at the top of the widget.

{{% widget name="apptabs" %}}

### Clip

Clip is a container that shows only a portion of the Content, as defined by the size of the clip.

{{% widget name="clip" %}}

### Navigation

The Navigation container provides a stack based navigation where items can be `Push`ed onto and the user can nagivate back and forwards.

{{% widget name="navigation" %}}

### Scroll

A Scroll container defines an area that is smaller than the Content, providing scrollbars as required.

{{% widget name="scroll" %}}

### Split

The Split container defines a container whose size is split between two children.

{{% widget name="split" %}}

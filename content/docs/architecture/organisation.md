---
title: Organisation and Packages
slug: organisation

weight: 9040

aliases:
- /architecture/organisation.html
- /tour/introduction/organisation
---

The Fyne project is split into many packages, each providing different
types of functionality. They are as follows:

`fyne.io/fyne/v2`
: This import provides the basic definitions common to all Fyne code 
: including data types and interfaces.

`fyne.io/fyne/v2/app`
: The app package provides the APIs that start a new application.
: Normally you only require [app.New()](/api/v2/app/pkg/#func--new) or [app.NewWithID()](/api/v2/app/pkg/#func--newwithid).

`fyne.io/fyne/v2/canvas`
: The canvas package provides all of the drawing APIs within Fyne.
: The complete Fyne toolkit is made up of these primitive graphical types.

`fyne.io/fyne/v2/container`
: The container package provides containers that are used to lay out and organise applications.

`fyne.io/fyne/v2/data/binding`
: The binding package contains ways of binding data sources to widgets.

`fyne.io/fyne/v2/data/validation`
: The validation package provides tooling for validating data inside widgets.

`fyne.io/fyne/v2/dialog`
: The dialog package contains dialogs such as confirm, error and file save/open.

`fyne.io/fyne/v2/layout`
: The layout package provides various layout implementations for use
: with containers (discussed in a later tutorial).

`fyne.io/fyne/v2/storage`
: The storage package provides storage access and management functionality. 

`fyne.io/fyne/v2/test`
: Applications can be tested more easily using the tools within the test
: package.

`fyne.io/fyne/v2/widget`
: Most graphical applications are created using a collection of widgets.
: All the widgets and interactive elements within Fyne are in this package.

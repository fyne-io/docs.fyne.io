---
layout: page
tags: [api]
title: Fyne API app
---

# app
---
```go
import "fyne.io/fyne/app"
```

Package app provides app implementations for working with Fyne graphical interfaces. The fastest way to get started is to call app.New() which will normally load a new desktop application. If the "ci" tag is passed to go (go run -tags ci myapp.go) it will run an in-memory application.

## Usage

#### func  New

```go
func New() fyne.App
```
New returns a new application instance with the default driver and no unique ID

#### func  NewAppWithDriver

```go
func NewAppWithDriver(d fyne.Driver, id string) fyne.App
```
NewAppWithDriver initialises a new Fyne application using the specified driver and returns a handle to that App. The id should be globally unique to this app Built in drivers are provided in the "driver" package.

<div class="deprecated"> Deprecated: Developers should not specify a driver manually but use NewAppWithID()</div>

#### func  NewWithID

```go
func NewWithID(id string) fyne.App
```
NewWithID returns a new app instance using the appropriate runtime driver. The ID string should be globally unique to this app.

#### types

 * [SettingsSchema](settingsschema.html)

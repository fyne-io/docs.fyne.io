---
layout: page
tags: [api]
title: Fyne API "app"
---

# app
---
```go
import "fyne.io/fyne/v2/app"
```

Package app provides app implementations for working with Fyne graphical interfaces. The fastest way to get started is to call app.New() which will normally load a new desktop application. If the "ci" tag is passed to go (go run -tags ci myapp.go) it will run an in-memory application.

## Usage

#### func  New

```go
func New() fyne.App
```
New returns a new application instance with the default driver and no unique ID

#### func  NewWithID

```go
func NewWithID(id string) fyne.App
```
NewWithID returns a new app instance using the test (headless) driver. The ID string should be globally unique to this app.

#### types

 * [SettingsSchema](settingsschema.html)

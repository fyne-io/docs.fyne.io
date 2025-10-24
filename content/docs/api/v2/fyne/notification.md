---
tags: [api]
title: fyne.Notification
slug: notification

aliases:
- /api/v2.0//notification
- /api/v2.1//notification
- /api/v2.2//notification
- /api/v2.3//notification
- /api/v2.4//notification
- /api/v2.5//notification
- /api/v2.6//notification
- /api/v2.7//notification

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Notification

```go
type Notification struct {
	Title, Content string
}
```

Notification represents a user notification that can be sent to the operating system.

#### func  NewNotification

```go
func NewNotification(title, content string) *Notification
```
NewNotification creates a notification that can be passed to [App.SendNotification].

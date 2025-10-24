---
tags: [api]
title: fyne.Notification
slug: notification

aliases:
- /api/v2.7//notification

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Notification struct {
	Title, Content string
}
```

Notification represents a user notification that can be sent to the operating system.

###

```go
func NewNotification(title, content string) *Notification
```
NewNotification creates a notification that can be passed to [App.SendNotification].

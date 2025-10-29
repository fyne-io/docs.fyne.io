---
tags: [api]
title: mobile.Driver
slug: driver

aliases:
- /api/driver/mobile/driver
- /api/driver/mobile/driver.html
- /api/v2.0/driver/mobile/driver
- /api/v2.0/driver/mobile/driver.html
- /api/v2.1/driver/mobile/driver
- /api/v2.1/driver/mobile/driver.html
- /api/v2.2/driver/mobile/driver
- /api/v2.2/driver/mobile/driver.html
- /api/v2.3/driver/mobile/driver
- /api/v2.3/driver/mobile/driver.html
- /api/v2.4/driver/mobile/driver
- /api/v2.4/driver/mobile/driver.html
- /api/v2.5/driver/mobile/driver
- /api/v2.5/driver/mobile/driver.html
- /api/v2.6/driver/mobile/driver
- /api/v2.6/driver/mobile/driver.html
- /api/v2.7/driver/mobile/driver
- /api/v2.7/driver/mobile/driver.html

package: fyne.io/fyne/v2/driver/mobile
---


---
```go
import "fyne.io/fyne/v2/driver/mobile"
```

## Usage

#### type Driver

```go
type Driver interface {
	// GoBack asks the OS to go to the previous app / activity, where supported
	GoBack()
}
```

Driver represents the extended capabilities of a mobile driver


<div class="since">Since: <code>
2.4</code></div>

---
tags: [api]
title: app.SettingsSchema
slug: settingsschema

aliases:
- /api/v2.0/app/settingsschema
- /api/v2.1/app/settingsschema
- /api/v2.2/app/settingsschema
- /api/v2.3/app/settingsschema
- /api/v2.4/app/settingsschema
- /api/v2.5/app/settingsschema
- /api/v2.6/app/settingsschema
- /api/v2.7/app/settingsschema

package: fyne.io/fyne/v2/app
---


---
```go
import "fyne.io/fyne/v2/app"
```

## Usage

#### type SettingsSchema

```go
type SettingsSchema struct {
	// these items are used for global settings load
	ThemeName         string  `json:"theme"`
	Scale             float32 `json:"scale"`
	PrimaryColor      string  `json:"primary_color"`
	CloudName         string  `json:"cloud_name"`
	CloudConfig       string  `json:"cloud_config"`
	DisableAnimations bool    `json:"no_animations"`
}
```

SettingsSchema is used for loading and storing global settings

#### func (*SettingsSchema) StoragePath

```go
func (sc *SettingsSchema) StoragePath() string
```
StoragePath returns the location of the settings storage

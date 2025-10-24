---
title: Geometry

weight: 9010
---

Fyne apps are based on 1 canvas per window.
Each canvas has a root CanvasObject which can be a single widget or a Container for many sub-objects whose size and position are controlled by a Layout.

## Position
---

Each canvas has its origin at the top left (0, 0) every element of the UI may be scaled depending on the output device and so the API does not describe pixels or exact measurements.
The position (10, 10) may be 10 pixels right and down from the origin on, for example, a 120DPI monitor but on a HiDPI (or "Retina") display this will probably be closer to 20 pixels.

Every position referenced by a CanvasObject is relative to it's parent.
This is important for layout algorithms but also for developers in situations such as the `Tappable.Tapped(PointEvent)` handlers.
Here the X/Y coordinates will be calculated from the top left of the button not the overall canvas.
This is designed to allow code to be as self-contained as possible.

## Pixel size
---

Like other vector based GUI libraries the Fyne coordinates need to be based
on some baseline monitor resolution. All [scaling](/architecture/scaling) is
relative to this value. For fyne that resolution is 120DPI.
This means that the sizes referred to in `fyne.Size` will be 1=1px when your monitor is 120DPI and scale values are all set to 1.
For a HiDPI screen, as mentioned above, the actual DPI may be closer to 240
and on mobile devices it could even be 360 or higher.
To manage handle this complexity the toolkit manages scaling internally so
your apps will always look the right size. 
If a user sets the scale to be smaller then their apps will always have
smaller than normal fonts, buttons etc, and when they specify larger then
your app will scale up to suit.

When compared to [Material Design](https://material.io) we can see that
their baseline DPI is [160](https://material.io/design/layout/pixel-density.html#pixel-density-on-android), although the maths is similar the 
actual numbers will be different. This means that device-independent 
sizes in Fyne use a smaller number to represent the same physical size.
For example an icon that is `18` tall in Fyne would be sized at `24` in a
standard material design (for example Android) app.
This does not matter when building your application, but may be important
when working with designers or experts with Material Design.

One time that pixel sizes will matter is if you start loading bitmaps images. Normally these scale appropriately, but if you specify
`FillMode=fyne.FillOriginal` then the actual image size will be different
on different devices, due to the pixel density. Normally this feature
would be used inside a `Scroll` container.
Fyne also defines a `canvas.Raster` primitive which will draw pixels exactly at the pixel density of the output device. This enables your code
to draw at the highest possible output resolution without knowing 
details of the device you are running on.
If for some reason you need "pixel perfect" positioning you need to multiply `CanvasObject.Size()` by `Canvas.Scale()`.


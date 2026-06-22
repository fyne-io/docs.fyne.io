---
title: Geometry

weight: 9010
---

Fyne apps are based on 1 canvas per window.
Each canvas has a root CanvasObject which can be a single widget or a Container for many sub-objects whose size and position are controlled by a Layout.

## Position

Each canvas has its origin at the top left (0, 0) every element of the UI may be scaled depending on the output device and so the API does not describe pixels or exact measurements.
The position (10, 10) may be 10 pixels right and down from the origin on, for example, a 120DPI monitor but on a HiDPI (or "Retina") display this will probably be closer to 20 pixels.

Every position referenced by a CanvasObject is relative to it's parent.
This is important for layout algorithms but also for developers in situations such as the `Tappable.Tapped(PointEvent)` handlers.
Here the X/Y coordinates will be calculated from the top left of the button not the overall canvas.
This is designed to allow code to be as self-contained as possible.

## Pixel size

Like other vector based GUI libraries the Fyne coordinates need to be based
on some baseline monitor resolution. All [scaling](/architecture/scaling) is
relative to this value. For fyne that value is system-dependent (matching the other apps
running on the computer) and falls back to 120DPI.
Another way to consider it is that a [fyne.Size](/api/v2/fyne/size/#type--size) value of 1
will be thje same as 1px when your monitor is a low pixel density and the system scale values are all set to 100%.
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

## Pixel Perfect Content

When you want to render content that is pixel precise instead of matching other content sizes
you can make use of the [canvas.Raster](/api/v2/canvas/raster/#type--raster) primitive.
This object draws pixels exactly at the pixel density of the output device.
This enables your code to draw at the highest possible output resolution without knowing 
details of the device you are running on. A raster object fills the available space like any other widget - the callback functions (the "generator func") is told how many pixels to draw.

If for some reason you need "pixel perfect" positioning or sizing of other objects you need to know how many pixels are used in a certain distance.
To do this you can make use of the `PixelCoordinateForPosition()` function on your `Canvas`.

One time that pixel sizes will matter is if you start loading bitmaps images. Normally these scale to fill the space given (like any other widget). However, if you specify
`FillMode=fyne.FillOriginal` then the image will be output so that 1px in the bitmap is 1px on your screen.
This means that the actual output size will vary on different devices, due to the pixel density. Due to that it is normally recommended that "original scale images" be used inside a `Scroll` container.

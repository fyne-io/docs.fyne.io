---
title: Scaling

weight: 9020
---

Fyne is built entirely using vector graphics, which means applications written with Fyne will scale to any size beautifully
(not just whole number increments).
This is a great benefit to the rising popularity of high density displays on mobile devices and high-end computers.
The default scale value is calculated to match your operating system - on some systems this is user configuration 
and on others from your screen's pixel density (DPI - dots per inch).
If a Fyne window is moved to another screen it will re-scale and adjust the window size accordingly!
We call this "auto scaling", and it is designed to keep an app user interface the same size as you change monitor.

You can adjust the size of applications using the `fyne_settings` app or by setting a specific scale using the `FYNE_SCALE` environment variable.
These values can make content larger or smaller than the system settings, using "1.5" will make things 50% larger
or setting 0.8 will make it 20% smaller.

<table style="text-align: center; margin: auto;"><tr>
<td><img src="/images/architecture/hello-normal.png" style="width: 207px;"  alt="Hello normal size" />
  <br />Standard size</td>
<td><img src="/images/architecture/hello-small.png" style="width: 160px;" alt="Hello small size" />
  <br />FYNE_SCALE=0.5</td>
<td><img src="/images/architecture/hello-large.png" style="width: 350px;"  alt="Hello large size" />
  <br />FYNE_SCALE=2.5</td>
</tr></table>

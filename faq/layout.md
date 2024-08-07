---
layout: page
title: Layout and Widget Size

order: 10
---

Intro

## Move and Resize

**Q: How can I move my widget to a different position or resize it?**

**A:** The position and size of elements in a Fyne app are controlled by the layout of the container that they are within. If the elements of your UI are too small consider using a different [layout](/started/layouts) or container.

A new `Window` will expand whatever element is passed to `SetContent()` to fill its size. Each time you add a container to this it will divide up the available space according to the layout. Layouts like `HBox` and `VBox` will shrink content to its `MinSize()` in one dimension or another to pack contents in. Layouts like `Max` or `Border` will expand content to fill the space. By writing a [custom layout](/extend/custom-layout) you could fully control the items within a container.

**Q: Why is my image so small?**

**A:** One of the difficulties in using a fully scalable user interface toolkit such as Fyne is that the coordinate system is device independent. This allows apps to draw at the right resolution or pixel density to get the best results based on the hardware connected. What this means for pixel based images is that their size could vary based on details not known at compile time.

Due to this complication an image loaded using `canvas.NewImageFromFile()` or similar calls will not have a size set, leading to it being very small or appearing to be hidden by default. When placed in an appropriate layout the image will stretch according to its `FillMode` property. If you desire the image to always be set to a certain size (or larger) you can call `Image.SetMinSize()` and specify a device independent size for the image.

## Containers and Layout 

**Q: How can I manually control the position of elements**

**A:** In some situations you may want to have complete control over the position and size of elements in a container. To do this you create a container without a layout.
The `container.NewWithoutLayout()` function will create a container for manual positioning - you should pass to that constructor a list of the graphical elements that you want to manage in this container.

Once set up then you can use `Move()` and `Resize()` on each element to position it as desired. When doing this be careful to note that it will not adjust as the available space changes - and it does not have an explicit minimum size either. To add either of those features you should replace your manual positioning with a [custom layout](/tutorial/custom-layout).

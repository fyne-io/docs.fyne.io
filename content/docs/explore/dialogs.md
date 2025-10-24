---
layout: page
title: Dialog List
readonly: true
slug: dialogs

weight: 2050

aliases:
- /started/dialogs
---

<style>
  html:not([data-dark-mode]) img.dialog.light {
    display: visible;
  }

  html:not([data-dark-mode]) img.dialog.dark {
    display: none;
  }

  html[data-dark-mode] img.dialog.light {
    display: none;
  }

  html[data-dark-mode] img.dialog.dark {
    display: visible;
  }
</style>

## Standard Dialogs

---

### Color

Allow users to pick a colour from a standard set (or any color in advanced mode).

{{% dialog name="color" %}}

### Confirm

Ask for conformation of an action.

{{% dialog name="confirm" %}}

### FileOpen

Present this to ask user to choose a file to use inside the app.
The actual dialog displayed will depend on the current operating system.

{{% dialog name="fileopen" %}}

### Form

Get various input elements in a dialog, with validation.

{{% dialog name="form" %}}

### Information

A simple way to present some information to the app user.

{{% dialog name="information" %}}

### Custom

Present any content inside a dialog container.

{{% dialog name="custom" %}}


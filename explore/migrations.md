---
title: Migrations
---

In some upgrades the Fyne toolkit may deliver a major feature that is a non-trivial upgrade
for larger applications. To allow this to happen without major API changes we have introduced
a migrations system. This allows apps to opt in to a new feature or way of working over time,
to test the changes carefully and then mark the project as migrated.

It is recommended to stay up to date with migrations within a year of them being introduced, 
as the toolkit will switch all these features on over time (following deprecation lifecycle).

In v2.6.0 Fyne introduced a new threading model which is the first time we used a migration.
If you have not already you should see the details of this feature at [Using Goroutines](/started/goroutines).
Once you have adopted this feature follow on to see how a migration is marked as complete.

## Opting in to a migration

Migrations are controlled by the new "[Migrations]" section inside `FyneApp.toml`.
Each line in this section is a name for the migration followed by a boolean value to switch it on or off.

You can see below an example metadata file where we have turned on the "fyneDo" migration. 
This should be done once you have updated your code so that any goroutines in your code use
the [fyne.Do] method to ensure the code runs on the toolkit's main goroutine.

```toml
[Details]
  Icon = "Icon.png"
  Name = "Example"
  ID = "com.example.app"

[Migrations]
  fyneDo = true
```

## Testing a migration

It is also possible to test a migration before you mark it as complete - there is a compiler
flag that will use this feature for the current build command.
For example, to test the fyne.Do migration you could build as follows:

```
go build -tags migrated_fynedo .
```

{% include youtube.html id="3uLhNMr0VaA" %}


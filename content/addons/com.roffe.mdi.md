---
addonid: com.roffe.mdi

title: Material Design Icons
home: https://roffe.nu
icon: https://github.com/roffe/mdi/raw/master/mdi.jpg
screenshots:
- url: https://github.com/roffe/mdi/raw/master/mdi.jpg

date:      2026-09-14 14:40:09
excerpt:   Material Design Icons (7447 SVGs, v7.4.47) as Fyne resources.
category:  icons
developer: Joakim, Karlsson

git: https://github.com/roffe/mdi
package: github.com/roffe/mdi
version: v1.0.0
layout: "addon"
---

Every icon is its own function backed by its own //go:embed variable, so the linker dead-code-eliminates the icons you don&#39;t call — your binary only contains the SVGs you actually use. Icons are themed: they follow the Fyne theme&#39;s foreground color.
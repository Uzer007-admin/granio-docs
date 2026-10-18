# Документация и каталоги Granio

Общий репозиторий документации и исходников каталогов экосистемы Granio.

**Основной исходник:** https://github.com/fyne-io/docs.fyne.io.
**Цель форка:** сопровождать документацию, каталог приложений и каталог расширений
в одном репозитории с отдельными конфигурациями сайтов.
Истории и сведения об исходных авторах сохранены; полный реестр —
[SOURCES.json](SOURCES.json).

| Сайт | Исходники | Исходник |
| --- | --- | --- |
| Документация | корень репозитория | [docs.fyne.io](https://github.com/fyne-io/docs.fyne.io) |
| Приложения | `sites/catalog` | [apps.fyne.io](https://github.com/fyne-io/apps.fyne.io) |
| Расширения | `sites/addons` | [addons.fyne.io](https://github.com/fyne-io/addons.fyne.io) |

Начало работы:

```sh
git clone --recurse-submodules --branch dev/granio https://github.com/Uzer007-admin/granio-docs.git
cd granio-docs
```

Сайты используют Hugo; каталоги сохраняют собственные темы и ресурсы.
Сборка из корня репозитория:

```sh
hugo --destination public/docs
hugo --source sites/catalog --destination ../../public/catalog
hugo --source sites/addons --destination ../../public/addons
```

Go-генераторы требуют Go 1.27.1+. Домен исходного проекта не подключается:
первоначальные значения сохранены в `CNAME.upstream`.
Развёртывание сайтов выполняется отдельно от проверки исходников.
[Материалы исходного проекта](UPSTREAM.md).

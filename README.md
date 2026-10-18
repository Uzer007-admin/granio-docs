# granio-docs

Компонент независимого проекта **Granio — Гранио**.

**Исходный репозиторий:** [docs.fyne.io](https://github.com/fyne-io/docs.fyne.io).
**Цель форка:** Документация независимого проекта Granio.

История изменений, лицензии и сведения об исходных авторах сохранены.
Go-модули переведены на собственные пути `github.com/Uzer007-admin/…`;
минимальная версия Go — **1.27.1**. Ядро и общие API подключаются из
[`granio/v2`](https://github.com/Uzer007-admin/granio).

Состояние переноса и результаты проверок приведены в
[реестре экосистемы](https://github.com/Uzer007-admin/granio/blob/dev/granio/docs/ECOSYSTEM.md).
Наличие форка само по себе не означает проверку всех платформ и сценариев.

---

## Документация исходного проекта

Команды установки и ссылки в сохранённом тексте ниже относятся к исходнику.
Для Granio используются новые пути модулей, перечисленные выше.

# docs.fyne.io Website

This repository contains the source of the https://docs.fyne.io website.

## Running the website locally

The website uses [hugo](https://gohugo.io/) to generate the website from markdown files.

### Install hugo
Follow the steps at [https://gohugo.io/installation/]

### Clone the Fyne docs repository
```bash
git clone https://github.com/fyne-io/docs.fyne.io.git
```
```bash
cd docs.fyne.io
```

### Start the local webserver
```bash
hugo server -D
```

## Run generator scripts

To generate new images etc. you can use the following commands:

```bash
cd _gen
go run genwidgets.go # generate widget images
go run genlayouts.go # generate layout images
go run gendialogs.go # generate dialog images
./genicons.sh # generate icon list html
./genapi.sh # generate api documentation structure
```


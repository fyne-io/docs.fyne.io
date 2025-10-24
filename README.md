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


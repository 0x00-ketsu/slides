# Slides

A terminal based preview tool for `markdown`.

![screenshot](./docs/assets/screenshot.png)

## Features

- Navigation: move to the next or previous slide
- Text Movement: move text up or down within a slide
- Scrolling: scroll up or down within a slide
- Search: search for a specific word
- Tagbar: preview of slides

## Install

### Go

```bash
go install github.com/0x00-ketsu/slides@latest
```

From source

```bash
git clone https://github.com/0x00-ketsu/slides.git
cd slides
make build
```

## Usage

### Quickly

Create(or copy an exist) a markdown file contains your slides, here's a simple example:

```markdown
# Welcome to Slides
A Terminal based preview tool for markdown

---

# h1
## h2
### h3
#### h4
##### h5
###### h6

---

# Markdown components
You can use everything in markdown!
* Like bulleted list
* You know the deal

1. Numbered lists too

---

# Tables

| Tables | Too    |
| ------ | ------ |
| Even   | Tables |
```

Then, run:

```bash
slides example.md 
```

`slides` is accepts input from `stdin`:

```bash
curl http://example.com/slides.md | slides
```

### Configuration

Copy the default configuration file to the your local home config directory.

```bash
slides config -c
```

### Keymaps

See [keymaps](./config/default.yaml) in default confguiration file.

## Inspiration by

- [lookatme](https://github.com/d0c-s4vage/lookatme) by James Johnson
- [slides](https://github.com/maaslalani/slides) by Maas Lalani

## License

MIT

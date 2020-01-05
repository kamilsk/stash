> # 🧺 stash
>
> Stash your bash history.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

## 💡 Idea

Port of [legacy](https://github.com/kamilsk/dotfiles/blob/master/bin/stash).

...

Full description of the idea is available [here][design.page].

## 🏆 Motivation

...

## 🤼‍♂️ How to

...

## 🧩 Installation

### Homebrew

```bash
$ brew install kamilsk/tap/stash
```

### Binary

```bash
$ curl -sSL https://bit.ly/install-stash | sh
# or
$ wget -qO- https://bit.ly/install-stash | sh
```

### Source

```bash
# use standard go tools
$ go get -u github.com/kamilsk/stash
# or use egg tool
$ egg tools add github.com/kamilsk/stash
```

> [egg][egg.page]<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

### Bash and Zsh completions

```bash
$ stash completion bash > /path/to/bash_completion.d/stash.sh
$ stash completion zsh  > /path/to/zsh-completions/_stash.zsh
```

## 🤲 Outcomes

...

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[build.icon]:       https://travis-ci.org/kamilsk/stash.svg?branch=master
[build.page]:       https://travis-ci.org/kamilsk/stash

[design.page]:      https://www.notion.so/33715348cc114ea79dd350a25d16e0b0?r=0b753cbf767346f5a6fd51194829a2f3

[promo.page]:       https://github.com/kamilsk/stash

[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[egg.page]:         https://github.com/kamilsk/egg

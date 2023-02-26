# git-saver

![img|200x200](hero-gopher.svg)

Command line tool to ease up process of saving files and folders to git repositories.

## Config

To create an initial config with example values in `$HOME/.git-saver.config.yml`.

```shell
git-saver init-config
```

Edit this file to your needs.

## Usage

A simple call to git-saver

```shell
git-saver
```

will copy all configured files and folders to their repositories and commit and push them.

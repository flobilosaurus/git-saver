# git-saver

<img src="https://github.com/flobilosaurus/git-saver/blob/d86adf0751b3a19702fe46d49fc1077bd084d2e9/hero-gopher.svg" width="200" height="200">

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

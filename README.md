# podgen

[Warning] don't use the project - it hasn't been properly deveopled yet. I shall publish it only after it works but I don't have any private repo quota...Thus everything listed here is not usable unless this notice is modified or removed!!!

Podgen is a tool for statically generate a podcast site with itunes enabled rss. After putting your generated online, you could register your rss against itunes podcast so that your awesome voice could be reached everywhere.

This project is highly inspired by go projects from hashicorp. The ``./scripts`` is from **consul** and the CLI is modified based on **vault** (this is not true anymore, now the CLI is using github.com/spf13/cobra).

## Installation

For Mac OSX, the easiest way to install it is through homebrew:

```
$ brew install podgen
```

You can also download the compiled version from github.

## Usage

To use podgen, first of all, you need to init a site:

```
$ mkdir programmer_life
$ cd programmer_life
$ podgen init
```

It will create configuration files for you to use:

```
$ ls
build     CNAME     channel.yml     items.yml   music
```

And it will create a ``gh-pages`` branch to store the build target. Just same as any other github powered static site.

Then edit ``channel.yml`` to put your channel information and add your podcast items into ``items.yml``. Finally copy your mp3 into ``music`` folder. You're almost set! Ah, if you want to use it against your customer domain, set the ``CNAME`` file!

The last step is to build the project:

```
$ podgen build [--template github.com/tyrchen/podgen-basic]
```

If you don't pass ``--template`` to the ``build`` command, it will use the ``podgen-basic`` template from github. The downloaded template is cached locally in ``~/.podgen/github.com/tyrchen/podgen-basic`` in case you need to use it next time.

Build should be done very quickly. It will generate the podcast site and rss, put them into ``build``, then push all the changes under ``build`` to ``gh-pages``. You shall then see the generated site in less than a minute. Meanwhile, your itunes podcast app shall get the latest rss. Try it and have fun!

![Place hold for my awesome programmer_life podcast](http://placehold.it/600x400)

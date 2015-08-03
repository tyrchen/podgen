# podgen

Podgen is a tool for statically generate a podcast site with itunes enabled rss. After putting your generated online, you could register your rss against itunes podcast so that your awesome voice could be reached everywhere.

This project is highly inspired by go projects from hashicorp. The ``./scripts`` is from **consul** and the CLI is modified based on **vault** (this is not true anymore, now the CLI is using github.com/spf13/cobra).

## Installation

For Mac OSX, the easiest way to download the current [release](https://github.com/tyrchen/podgen/releases/download/v0.2.0/podgen) from [tyrchen/podgen release](https://github.com/tyrchen/podgen/releases).

Later on we will support homebrew so that you could install directly with:

```
$ brew install podgen
```

For linux user, please download source code and compile it yourself.

podgen doesn't support windows at this stage. I will finish all the functionalities then consider windows support. Sorry.

## Usage

To use podgen, first of all, you need to init a site:

```
$ mkdir programmer_life
$ cd programmer_life
$ podgen init [--template github.com/tyrchen/podgen-basic]
```

[note] If you don't pass ``--template`` to the ``init`` command, it will use the ``tyrchen/podgen-basic`` template from github.

It will create configuration files for you to use:

```
$ ls
build     CNAME     channel.yml     items.yml   assets template
```

And it will create a ``gh-pages`` branch to store the build target. Just same as any other github powered static site.

Then edit ``channel.yml`` to put your channel information and add your podcast items into ``items.yml``. Finally copy your mp3 into ``assets`` folder. You're almost set! Ah, if you want to use it against your customer domain, set the ``CNAME`` file!

The last step is to build the project:

```
$ podgen build
```

Build should be done very quickly. It will generate the podcast site and rss, put them into ``build``, then push all the changes under ``build`` to ``gh-pages``. You shall then see the generated site in less than a minute.

If everything looks fine. You just need to issue "podgen push" to push the source and generated html to ``master`` and ``gh-pages``.

```
$ podgen push -m "add a new episode"
```

Meanwhile, your itunes podcast app shall get the latest rss. Try it and have fun!

You can always welcome to view a live demo of my podcast site: [programmer life](http://podcast.tchen.me)

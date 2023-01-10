# nust (not unique setup tool)

## TL;DR

**🚩 Problem**: you installed fresh linux on new PC and now you need to configure vim, fetch all your personal git repos, install bunch of software, enable couple of systemds. Man that sucks. 

**🎯 Solution**: No worries,`nust` is a really easy to understand setup tool exacly for these kinds of problems.

**🧐 How**: You can create your **nust tasks** for your purposes. Nust is just hub for all of your scripts you want to run on your new fresh installed linux. `nust` makes it really easy to fetch these scripts from remote location for e.g. GitHub. 

**🤔 How it works**: **nust task** is just a **makefile** with **nust_do**(required), **nust_undo**(optional) targets. nust tasks can be bundled into **nust package**. You can create nust task locally or you can give to `nust` a link to your GitHub repository that has **nust package**. 

**✨ Best practice**: make nust tasks as independent as possible. 

Oh also, have I mentioned that `nust` program should work on **any** linux distribution? 

## Ok you hooked me up, but actually where to start with `nust`? 

## Requirements

You have to have `git` and `make` commands on your system

## Installation

pls write here when your program is ready 

### basics

Executing one `nust` task:
```
nust do -t my_task.nust
```

Executing one `nust` task with arguments for `makefile`:
```
nust do -t my_task.nust -m "VERY_SECRET_KEY_FILE=secret_key_file_path EXAMPLE_VAR=example"
```

Executing `nust` package (you can see examples folder to get more details):
```
nust do -p my_desktop_setup.nustpack
```

Fetching and executing `nust` package from repo (again more examples in folder examples):
```
nust do -r https://github.com/example/example-nustpackage
```

### Naming your `nust` package repositories 


```
nustpack-<distro>-<name of nust package>
```

if your `nust` package is distro dependent then put name of a distro, otherwise you can omit it. 

For e.g
```
https://github.com/exampleuser/nustpack-manjaro-my-desktop-setup
https://github.com/exampleuser/nustpack-arch-my-desktop-setup
https://github.com/exampleuser/nustpack-setup-neovim-config
https://github.com/exampleuser/nustpack-ubuntu-virtualization-setup
https://github.com/exampleuser/nustpack-my-firewall-config
```

This is merely my suggestion you can completely ignore it if you want. But if you wrote nust package why wouldn't you make it useful for somebody else too ;)?
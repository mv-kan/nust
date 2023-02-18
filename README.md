# nust (not unique setup tool)

## TL;DR

**🚩 Problem**: you installed fresh linux on new PC and now you need to configure vim, fetch all your personal git repos, install bunch of software, enable couple of systemds. Man that sucks. 

**🎯 Solution**: No worries,`nust` is a really easy to understand setup tool exacly for these kinds of problems.

**🧐 How**: You can create your **nust tasks** for your purposes. Nust is just hub for all of your scripts you want to run on your new fresh installed linux. `nust` makes it really easy to fetch these scripts from remote location for e.g. GitHub. 

**🤔 How it works**: **nust task** is just a **makefile** with **nust_do**(required), **nust_undo**(optional) targets. nust tasks can be bundled into **nust package**. You can create nust task locally or you can give to `nust` a link to your GitHub repository that has **nust package**. 

**✨ Best practice**: make nust tasks as independent as possible. 

Oh also, have I mentioned that `nust` program should work on **any** linux distribution? 

## Ok you hooked me up, but actually where to start with `nust`? 

### Requirements

You have to have `make` command on your system

### Installation`

Download .tar.gz file unpack it and you will find `nust` binary. To install `nust` copy it to `/usr/bin/` directory like this:
```
sudo cp ./nust /usr/bin/
```

To uninstall use:
```
sudo rm /usr/bin/nust
```

## Using the damn thing

Paste here your ubuntu desktop setup

This is merely my suggestion you can completely ignore it if you want. But if you wrote nust package why wouldn't you make it useful for somebody else too ;)?
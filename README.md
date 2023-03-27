<img src="./static/NUST.svg"  width="25%"/>

# nust (not unique setup tool)

## TL;DR

**🚩 Problem**: you installed fresh linux on new PC and now you need to configure vim, fetch all your personal git repos, install bunch of software, enable couple of systemds. Man that sucks. 

**🎯 Solution**: Bruh just write `shell` scripts to do all of that 

**🤔 But if I have scripts why `nust` exists?**: `nust` is basically middle man between your `sh` or `bash` script and the `/bin/bash` program. `nust` can perform retries and save the outcome of runned script in `nust_exec_info.json` (true is not fail, false is fail). This is nothing out of ordinary, just quality of life program. Maybe in future I will add more of that qol things to make `nust` a viable program for someone to use. Maybe...

## Ok you hooked me up, but actually where to start with `nust`? 

### Installation

Download .tar.gz file unpack it and you will find `nust` binary. To install `nust` copy it to `/usr/bin/` directory like this:
```
sudo cp ./nust /usr/bin/
```

To uninstall use:
```
sudo rm /usr/bin/nust
```

#### Personally

Personally I would just bake in the nust program into my repo with setup like in https://github.com/mv-kan/nustpack-ubuntu2204-mysetup

## Using the damn thing

this is example how can you use `nust` to set up your workspace in ubuntu 22.04

```
https://github.com/mv-kan/nustpack-ubuntu2204-mysetup
```
# nust (not unique setup tool)

## TL;DR

**🚩 Problem**: you installed fresh linux on new PC and now you need to configure vim, fetch all your personal git repos, install bunch of software, enable couple of systemds. Man that sucks. 

**🎯 Solution**: No worries,`nust` is a really easy to understand setup tool exacly for these kinds of problems.

**🧐 How**: You can create your **nust homeworks** for your purposes. Nust is just hub for all of your scripts you want to run on your new fresh installed linux. `nust` makes it really easy to fetch these scripts from remote location for e.g. GitHub. 

**🤔 How it works**: **nust homework** is just a **makefile** with **nust_do**(required), **nust_undo**(optional) commands. You can create nust homework locally or you can give to `nust` a link to your GitHub repository that is **nust homework**.

**✨ Best practice**: make nust homeworks as independent as possible. 

Oh also, have I mentioned that `nust` program should work on **any** linux distribution? 

## Ok you hooked me up, but actually where to start with `nust`? 

### basics

nust from folder (scans for files with .nust extension for e.g. install_software.nust, configure_my_systemd.nust or "makefile" filename)
```
nust do -d ./folder
# undo (requires nust_undo command in makefile)
nust undo -d ./folder
```

nust with one file 
```
nust do -f ./makefile
# undo (requires nust_undo command in makefile)
nust undo -f ./makefile 
```

nust with url to repo (in repo there have to be file with "makefile" name or files with .nust extension in root dir of repo) 
```
nust do -r https://github.com/your-name/your-nust-package
```

### List of repos

Yes you can make multi do from one file.

```
# in my_linux_setup file

# install all software I need for work like vscode, calibre, steam, snap, flathub 
https://github.com/your-name/your-nust-package1

# enable couple of services in systemd
https://github.com/your-name/your-nust-package2

# copy to home directory all github repos I am working on
https://github.com/your-name/your-nust-package3
```

give nust list of repos
```
nust do -l my_linux_setup
```

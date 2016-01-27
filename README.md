# logfs - Simple Fuse FS

This is just an experiment to play with fuse and log the syscall when you surf the filesystem 

## Built

You will need install GO > 1.4

[golang](https://golang.org/doc/install) 

## Install  

go get github.com/ligadous/logfs

## Mount

You will need a empty folder in any place (example: /mnt/lfs)
and a second folder with some files (~/test)

`Usage of logfs:
  -folder string
    	origin folder with content
  -mount string
    	mount folder

logfs -folder=~/test -mount=/mnt/lfs
`

Open a second terminal and test some trival commands:

cd /mnt/lfs
ls
ls -la
echo "test1" > arq1.txt
rm arq1.txt

Look the messages in first terminal!

## Unmount filesystem

In the second terminal, leaves the mount folder and run 

umount /mnt/lfs

The program must finish!

*Thanks for http://bazil.org and Russ Cox (http://research.swtch.com/)*


package main

import (
	"./fs"
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"flag"
	"fmt"
	"log"
	"os"
)

var folder = flag.String("folder", "", "origin folder with content")
var mount = flag.String("mount", "", "mount folder")

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *mount == "" || *folder == "" {
		usage()
		os.Exit(2)
	}

	c, err := fuse.Mount(
		*mount,
		fuse.FSName("logfs"),
		fuse.Subtype("logfs"),
		fuse.VolumeName("Log FS"),
		//fuse.AllowOther(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	cfg := &fs.Config{}
	srv := fs.New(c, cfg)
	filesys := logfs.NewLogFS(*folder)

	if err := srv.Serve(filesys); err != nil {
		log.Fatal(err)
	}

	// Verify mount error
	<-c.Ready
	if err := c.MountError; err != nil {
		log.Fatal(err)
	}

	log.Println("Log FS unmounted")
}

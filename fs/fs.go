package logfs

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"golang.org/x/net/context"
	"log"
	"os"
	"sync/atomic"
	"syscall"
	"time"
)

type LogFS struct {
	root   *Dir
	nodeId uint64
	path   string
	size   int64
}

// Compile-time interface checks.
var _ fs.FS = (*LogFS)(nil)

//var _ fs.FSStatfser = (*LogFS)(nil)

const DEF_MODE = os.FileMode(int(0777))

func NewLogFS(path string) *LogFS {
	log.Printf("** Start mount Log FS (%s)", path)

	fs := &LogFS{
		path: path,
	}
	fs.root = fs.newDir(path, os.ModeDir|DEF_MODE)
	if fs.root.attr.Inode != 1 {
		panic("O Root must receive id 1")
	}
	return fs
}

func (m *LogFS) nextId() uint64 {
	nId := atomic.AddUint64(&m.nodeId, 1)
	log.Printf("** nextId (%d)", nId)
	return nId
}

func (m *LogFS) newDir(path string, mode os.FileMode) *Dir {
	log.Printf("** newDir (%s)", path)

	n := time.Now()
	return &Dir{
		attr: fuse.Attr{
			Inode:  m.nextId(),
			Atime:  n,
			Mtime:  n,
			Ctime:  n,
			Crtime: n,
			Mode:   os.ModeDir | mode,
		},
		path: path,
		fs:   m,
	}
}

func (m *LogFS) newFile(path string, mode os.FileMode) *File {
	log.Printf("** newFile (%s)", path)

	n := time.Now()
	return &File{
		attr: fuse.Attr{
			Inode:  m.nextId(),
			Atime:  n,
			Mtime:  n,
			Ctime:  n,
			Crtime: n,
			Mode:   mode,
		},
		path: path,
	}
}

func (f *LogFS) Root() (fs.Node, error) {
	log.Println("** Root")

	return f.root, nil
}

func (f *LogFS) Statfs(ctx context.Context, req *fuse.StatfsRequest, res *fuse.StatfsResponse) error {
	log.Println("** Statfs")

	s := syscall.Statfs_t{}
	err := syscall.Statfs(f.path, &s)
	if err != nil {
		log.Println("DRIVE) Statfs falha no syscall; ", err)
		return err
	}

	res.Blocks = s.Blocks
	res.Bfree = s.Bfree
	res.Bavail = s.Bavail
	res.Ffree = s.Ffree
	res.Bsize = s.Bsize
	return nil
}

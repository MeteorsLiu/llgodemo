package main

import (
	"unsafe"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/libuv"
)

// see https://github.com/phracker/MacOSX-SDKs/blob/master/MacOSX10.8.sdk/System/Library/Frameworks/Kernel.framework/Versions/A/Headers/sys/fcntl.h
const (
	O_RDONLY  = 0x0000 /* open for reading only */
	O_WRONLY  = 0x0001 /* open for writing only */
	O_RDWR    = 0x0002 /* open for reading and writing */
	O_ACCMODE = 0x0003
)

var (
	readReq  = &libuv.Fs{}
	closeReq = &libuv.Fs{}
	buffer   *libuv.Buf
)

type WriteReq struct {
	Req libuv.Write
	Buf libuv.Buf
}

func onRead(req *libuv.Fs) {
	switch req.GetResult() {
	case -1:
		panic("file doesn't exist")
	case 0:
		libuv.FsClose(libuv.DefaultLoop(), closeReq, libuv.File(req.GetResult()), nil)
		return
	}

	// read successfully
	c.Printf(c.Str("read: %s\n"), buffer.Base)
}

func onOpen(req *libuv.Fs) {
	if req.GetResult() == -1 {
		panic("file doesn't exist")
	}

	// async open successfully
	// start to read.
	libuv.FsRead(
		libuv.DefaultLoop(),
		readReq,
		libuv.File(req.GetResult()),
		buffer, 1,
		-1, onRead,
	)
}

func main() {
	// Initialize the default event loop

	req := &libuv.Fs{}
	loop := libuv.DefaultLoop()

	// allocate a memory
	base := c.Malloc(1024)

	// initalize the libuv buffer
	buf := libuv.InitBuf((*c.Char)(base), 1024)
	buffer = &buf

	// Zero the memory
	c.Memset(unsafe.Pointer(buffer.Base), 0, buffer.Len)

	// open a file
	libuv.FsOpen(loop, req, c.Str("test.txt"), O_RDONLY, 0, onOpen)

	loop.Run(libuv.RUN_DEFAULT)
}

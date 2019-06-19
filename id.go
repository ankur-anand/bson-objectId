package id

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

// objectIdCounter is atomically incremented when generating
// a new Object Id using NewObjectId() function.
// It's used as a counter part of an id.
var objectIDCounter uint32

// machineId  is used to Strore machine id generated once will be used
// in sebsequent calls.
var machineID = generateMachineID()

// readMachineID generates machine id and puts it into
// the machineId Global varibale. If this function fails
// to get the hostname, it will cause a runtime error.
func generateMachineID() []byte {
	var sum [3]byte // 3 byte Machine ID
	id := sum[:]
	hostname, err1 := os.Hostname()
	if err1 != nil {
		// if getting hostname failed
		// get a crypto random id and return
		_, err2 := io.ReadFull(rand.Reader, id)
		if err2 != nil {
			panic(fmt.Errorf("Cannot get hostname: %v, %v", err1, err2))
		}
		return id
	}
	hw := md5.New()
	// append hostname to the running hash
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}

// New returns a new unique ObjectId.
// 4 byte time,
// 3 byte Machine ID
// 2 byte pid
// 3 byte self increased id.
func New() string {
	var b [12]byte
	// TimeStamp, 4 bytes, big endian.
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	b[4] = machineID[0]
	b[5] = machineID[1]
	b[6] = machineID[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian
	pid := os.Getegid()
	b[7] = byte(pid >> 8)
	b[8] = byte(pid)
	// increment 3 bytes, big Endian
	i := atomic.AddUint32(&objectIDCounter, 1)
	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)
	return hex.EncodeToString(b[:])
}
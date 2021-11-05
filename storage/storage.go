package storage

import (
    "fmt"
    "io"
)

type Storage interface {
    Put(key string, r io.Reader, dataLength int64) error
    PutFile(key string, localFile string) error
    Get(key string) (io.ReadCloser, error)
    Rename(srcKey string, destKey string) error
    Copy(srcKey string, destKey string) error
    Exists(key string) (bool, error)
    Size(key string) (int64, error)
    Delete(key string) error
    Url(key string) string
}

var disks = make(map[DiskName]Storage)

func Register(name DiskName, disk Storage) {
    if disk == nil {
        panic("storage: Register disk is nil")
    }
    disks[name] = disk
}

func Disk(name DiskName) (Storage, error) {
    disk, exist := disks[name]
    if !exist {
        return nil, fmt.Errorf("storage: Unknown disk %q", name)
    }
    return disk, nil
}

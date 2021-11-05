package storage

import (
    "os"
    "strings"
)

func NormalizeKey(key string) string {
    key = strings.Replace(key, "\\", "/", -1)
    key = strings.Replace(key, " ", "", -1)
    key = filterNewLines(key)

    return key
}

func filterNewLines(s string) string {
    return strings.Map(func(r rune) rune {
        switch r {
        case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
            return -1
        default:
            return r
        }
    }, s)
}

func OpenAsReadOnly(key string) (*os.File, os.FileInfo, error) {
    fd, err := os.Open(key)
    if err != nil {
        if os.IsNotExist(err) {
            return nil, nil, FileNotFoundErr
        }
        if os.IsPermission(err) {
            return nil, nil, FileNoPermissionErr
        }
        return nil, nil, err
    }

    stat, err := fd.Stat()
    if err != nil {
        return nil, nil, err
    }

    return fd, stat, nil
}

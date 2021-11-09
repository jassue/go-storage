package kodo

import (
    "github.com/jassue/go-storage/storage"
    "github.com/qiniu/go-sdk/v7/sms/bytes"
    "io"
    "io/ioutil"
    "testing"
)

var disk storage.Storage

func TestMain(m *testing.M) {
    disk, _ = Init(Config{
        AccessKey: "",
        Bucket:    "",
        Domain:    "www.example.com",
        SecretKey: "",
        IsSsl:     true,
        IsPrivate: false,
    })

    m.Run()
}

func TestKodo_Put(t *testing.T) {
    data, _ := ioutil.ReadFile("../tests/accounts.txt")
    err := disk.Put("test_data/accounts.txt", bytes.NewReader(data), int64(len(data)))
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestKodo_PutFile(t *testing.T) {
    err := disk.PutFile("test_data/accounts2.txt", "../tests/accounts.txt")
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestKodo_Get(t *testing.T) {
    body, err := disk.Get("test_data/accounts.txt")
    if _, ok := body.(io.Closer); ok {
        defer body.Close()
    }
    if err != nil {
        t.Error(err.Error())
        return
    }

    data, err := ioutil.ReadAll(body)
    err = disk.Put("test_data/get_put_accounts.txt", bytes.NewReader(data), int64(len(data)))
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestKodo_Rename(t *testing.T) {
    err := disk.Rename("test_data/accounts2.txt", "test_data/rename_accounts.txt")
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestKodo_Copy(t *testing.T) {
    err := disk.Copy("test_data/accounts.txt", "test_data/copy_accounts.txt")
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestKodo_Exists(t *testing.T) {
    exists, err := disk.Exists("test_data/accounts.txt")
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Logf("isExisted : %v", exists)
    t.Log("success")
}

func TestKodo_Size(t *testing.T) {
    size, err := disk.Size("test_data/accounts.txt")
    if err != nil {
        t.Log(err.Error())
        return
    }
    t.Logf("size : %d", size)
    t.Log("success")
}

func TestKodo_Delete(t *testing.T) {
    err := disk.Delete("test_data/accounts.txt")
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestKodo_Url(t *testing.T) {
    url := disk.Url("test_data/accounts.txt")
    t.Log("url : " + url)
    t.Log("success")
}

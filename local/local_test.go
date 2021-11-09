package local

import (
    "github.com/jassue/go-storage/storage"
    "github.com/qiniu/go-sdk/v7/sms/bytes"
    "io"
    "io/ioutil"
    "testing"
)

var disk storage.Storage

func TestMain(m *testing.M)  {
    disk, _ = Init(Config{
        RootDir: "../tests",
        AppUrl: "http://localhost:8888/tests",
    })

    m.Run()
}

func TestLocal_Put(t *testing.T) {
    data, _ := ioutil.ReadFile("../tests/accounts.txt")
    err := disk.Put("local/accounts.txt", bytes.NewReader(data), int64(len(data)))
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestLocal_PutFile(t *testing.T) {
    err := disk.PutFile("local/put_file_accounts.txt", "../tests/accounts.txt")
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestLocal_Get(t *testing.T) {
    body, err := disk.Get("local/accounts.txt")
    if _, ok := body.(io.Closer); ok {
        defer body.Close()
    }
    if err != nil {
        t.Error(err.Error())
        return
    }

    data, err := ioutil.ReadAll(body)
    err = disk.Put("local/get_put_accounts.txt", bytes.NewReader(data), int64(len(data)))
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestLocal_Rename(t *testing.T) {
    err := disk.Rename("local/put_file_accounts.txt", "local2/rename_accounts.txt")
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestLocal_Copy(t *testing.T) {
    err := disk.Copy("local/accounts.txt", "local3/copy_accounts.txt")
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestLocal_Exists(t *testing.T) {
    exists, err := disk.Exists("local/accounts.txt")
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Logf("isExisted : %v", exists)
    t.Log("success")
}

func TestLocal_Size(t *testing.T) {
    size, err := disk.Size("local/accounts.txt")
    if err != nil {
        t.Log(err.Error())
        return
    }
    t.Logf("size : %d", size)
    t.Log("success")
}

func TestLocal_Delete(t *testing.T) {
    err := disk.Delete("local/accounts.txt")
    if err != nil {
        t.Error(err.Error())
        return
    }
    t.Log("success")
}

func TestLocal_Url(t *testing.T) {
    url := disk.Url("local/accounts.txt")
    t.Log("url : " + url)
    t.Log("success")
}

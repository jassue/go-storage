package oss

import (
    alioss "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "github.com/jassue/go-storage/storage"
    "io"
    "strconv"
    "sync"
)

type Config struct {
    AccessKeyId string `mapstructure:"access_key_id" json:"access_key_id" yaml:"access_key_id"`
    AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret" yaml:"access_key_secret"`
    Bucket string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
    Endpoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
    IsSsl bool `mapstructure:"is_ssl" json:"is_ssl" yaml:"is_ssl"`
    IsPrivate bool `mapstructure:"is_private" json:"is_private" yaml:"is_private"`
}

type oss struct {
    config *Config
    client *alioss.Client
    bucket *alioss.Bucket
}

var (
    o *oss
    once *sync.Once
    initErr error
)

func Init(config Config) (storage.Storage, error) {
    once = &sync.Once{}
    once.Do(func() {
        o = &oss{}
        o.config = &config

        o.client, initErr = alioss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
        if initErr != nil {
            return
        }

        o.bucket, initErr = o.client.Bucket(config.Bucket)
        if initErr != nil {
            return
        }

        storage.Register(storage.Oss, o)
    })
    if initErr != nil {
        return nil, initErr
    }
    return o, nil
}

func (o *oss) Put(key string, r io.Reader, dataLength int64) error {
    key = storage.NormalizeKey(key)

    err := o.bucket.PutObject(key, r)
    if err != nil {
        return err
    }

    return nil
}

func (o *oss) PutFile(key string, localFile string) error {
    key = storage.NormalizeKey(key)

    err := o.bucket.PutObjectFromFile(key, localFile)
    if err != nil {
        return err
    }

    return nil
}

func (o *oss) Get(key string) (io.ReadCloser, error) {
    key = storage.NormalizeKey(key)

    body, err := o.bucket.GetObject(key)
    if err != nil {
        return nil, err
    }

    return body, nil
}

func (o *oss) Rename(srcKey string, destKey string) error {
    srcKey = storage.NormalizeKey(srcKey)
    destKey = storage.NormalizeKey(destKey)

    _, err := o.bucket.CopyObject(srcKey, destKey)
    if err != nil {
        return err
    }

    err = o.Delete(srcKey)
    if err != nil {
        return err
    }

    return nil
}

func (o *oss) Copy(srcKey string, destKey string) error {
    srcKey = storage.NormalizeKey(srcKey)
    destKey = storage.NormalizeKey(destKey)

    _, err := o.bucket.CopyObject(srcKey, destKey)
    if err != nil {
        return err
    }

    return nil
}

func (o *oss) Exists(key string) (bool, error) {
    key = storage.NormalizeKey(key)

    return o.bucket.IsObjectExist(key)
}

func (o *oss) Size(key string) (int64, error) {
    key = storage.NormalizeKey(key)

    props, err := o.bucket.GetObjectDetailedMeta(key)
    if err != nil {
        return 0, err
    }

    size, err := strconv.ParseInt(props.Get("Content-Length"), 10, 64)
    if err != nil {
        return 0, err
    }

    return size, nil
}

func (o *oss) Delete(key string) error {
    key = storage.NormalizeKey(key)

    err := o.bucket.DeleteObject(key)
    if err != nil {
        return err
    }

    return nil
}

func (o *oss) Url(key string) string {
    var prefix string
    key = storage.NormalizeKey(key)

    if o.config.IsSsl {
        prefix = "https://"
    } else {
        prefix = "http://"
    }

    if o.config.IsPrivate {
        url, err := o.bucket.SignURL(key, alioss.HTTPGet, 3600)
        if err == nil {
            return url
        }
    }

    return prefix + o.config.Bucket + "." + o.config.Endpoint + "/" + key
}

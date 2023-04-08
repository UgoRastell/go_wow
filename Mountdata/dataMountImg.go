package data

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type MountKey struct {
    ID  int    `json:"id"`
    Key string `json:"key"`
}

type MountMedia struct {
    Assets []struct {
        Key   string `json:"key"`
        Value string `json:"value"`
    } `json:"assets"`
    ID int `json:"id"`
}

type MountMediaResponse struct {
    ID     int    `json:"id"`
    Assets []struct {
        Key   string `json:"key"`
        Value string `json:"value"`
    } `json:"assets"`
}

func DataMountMedia() ([]MountMedia, error) {
    mounts, err := MountIndex()
    if err != nil {
        fmt.Println("Erreur :", err)
        return nil, err
    }

    var mountMediaList []MountMedia
    for _, mount := range mounts {
        url := mount.URL.Href

        resp, err := http.Get(url)
        if err != nil {
            fmt.Println("Erreur :", err)
            return nil, err
        }
        defer resp.Body.Close()

        var mediaResponse MountMediaResponse
        err = json.NewDecoder(resp.Body).Decode(&mediaResponse)
        if err != nil {
            fmt.Println("Erreur :", err)
            return nil, err
        }

        mediaList := make([]struct {
            Key   string `json:"key"`
            Value string `json:"value"`
        }, 0)

        for _, asset := range mediaResponse.Assets {
            mediaList = append(mediaList, struct {
                Key   string `json:"key"`
                Value string `json:"value"`
            }{
                Key:   asset.Key,
                Value: asset.Value,
            })
        }

        mountMediaList = append(mountMediaList, MountMedia{
            ID:     mediaResponse.ID,
            Assets: mediaList,
        })
    }
    fmt.Println(mountMediaList)
    return mountMediaList, nil
}
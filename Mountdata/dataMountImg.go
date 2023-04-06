package data

import (
    "encoding/json"
    "fmt"
    "net/http"

    "wow/tokens"
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

func DataMountMedia() ([]MountMedia, error) {
    mounts, err := MountIndex()
    if err != nil {
        fmt.Println("Erreur :", err)
        return nil, err
    }

    var mountMediaList []MountMedia

    for _, mount := range mounts {
        url := mount.URL.Href

        resp, err := http.Get(url + "&access_token=" + token.Access())
        if err != nil {
            fmt.Println("Erreur :", err)
            return nil, err
        }
        defer resp.Body.Close()

        var mountMedia MountMedia
        err = json.NewDecoder(resp.Body).Decode(&mountMedia)
        if err != nil {
            fmt.Println("Erreur :", err)
            return nil, err
        }

        mediaList := make([]struct {
            ID    int
            Key   string
            Value string
        }, 0)

        for _, asset := range mountMedia.Assets {
            mediaList = append(mediaList, struct {
                ID    int
                Key   string
                Value string
            }{
                ID:    mountMedia.ID,
                Key:   asset.Key,
                Value: asset.Value,
            })
        }

        for _, media := range mediaList {
            fmt.Printf("%+v\n", media)
        }

        mountMediaList = append(mountMediaList, mountMedia)
    }

    return mountMediaList, nil
}

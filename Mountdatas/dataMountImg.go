package data

import (
	"fmt"
	"net/http"
	"encoding/json"

	"wow/tokens"
)

type MountImg struct {
	Assets         []urlImg        `json:"assets"`
	Id			   int             `json:"id"`
}

type urlImg struct {
    Value string `json:"value"`
}

func DataImg() ([]string, error){
    mounts, err := DataMount()
    if err != nil {
        fmt.Println("Erreur :", err)
        return nil, err
    }

    var imgURLs []string

    for _, mount := range mounts {
        if len(mount.CreatureDisplays) > 0 {
            url := mount.CreatureDisplays[0].Key.Href + "&access_token=" + token.Access()
            resp, err := http.Get(url)
            if err != nil {
                fmt.Println("Erreur :", err)
                return nil, err
            }
            defer resp.Body.Close()

            var mountImg MountImg

            err = json.NewDecoder(resp.Body).Decode(&mountImg)
            if err != nil {
                fmt.Println("Erreur :", err)
                return nil, err
            }

			newMountUrls := MountImg{
				Id: mountImg.Id,
				URL: url,
			}

			imgURLs = append(imgURLs, newMountUrls)
			fmt.Println(imgURLs)

            fmt.Println(mountImg.Id)
        }
    }
    return imgURLs, nil
}
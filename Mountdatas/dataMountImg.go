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
    Key   string `json:"key"`
    Value string `json:"value"`
}

func DataImg() ([]MountImg, error){
    mounts, err := DataMount()
    if err != nil {
        fmt.Println("Erreur :", err)
        return nil, err
    }

    var imgURLs []MountImg

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

            
            newMountURL := MountImg{
                Id: mountImg.Id,
                Assets: mountImg.Assets,
            }
            imgURLs = append(imgURLs, newMountURL)
            
        }
        
    }

    return imgURLs, nil
}


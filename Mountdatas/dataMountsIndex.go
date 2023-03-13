package data

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wow/tokens"
)

// Les données des montures sont stockées dans la propriété "mounts" de la réponse de l'API.

type Mount struct {
    URL  MountUrl `json:"key"`
}

type MountUrl struct {
    Href string `json:"href"`
}

type MountsResponse struct {
    Mounts []Mount `json:"mounts"`
}


func MountIndex() ([]Mount, error) {
    
    // Définir l'URL de l'API Blizzard
    url := "https://eu.api.blizzard.com/data/wow/mount/index?namespace=static-eu&locale=fr_FR&access_token=" + token.Access()

    // Envoyer une requête GET à l'API Blizzard
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Erreur :", err)
        return nil, err
    }
    defer resp.Body.Close()
    
    var mountsResponse MountsResponse

    
    err = json.NewDecoder(resp.Body).Decode(&mountsResponse)
    if err != nil {
        fmt.Println("Erreur :", err)
        return nil, err
    }

    for i := range mountsResponse.Mounts {
        mountsResponse.Mounts[i].URL.Href += "&access_token=" + token.Access()
    }

    mounts := make([]Mount, len(mountsResponse.Mounts))
    copy(mounts, mountsResponse.Mounts)

    return mounts, nil
}

package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wow/tokens"
)

type Creature struct {
	CreatureDisplays []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"creature_displays"`
	// Autres champs dans votre structure Creature
}

type MountMedia struct {
	Assets []struct {
		Key      string `json:"key"`
		Value    string `json:"value"`
		KeyFrom  string `json:"key_from,omitempty"`
	} `json:"assets"`
	ID int `json:"id"`
}

type MountMediaResponse struct {
	Assets []struct {
		Key      string `json:"key"`
		Value    string `json:"value"`
		Assets   []struct {
			Value string `json:"value"`
		} `json:"assets"`
		KeyFrom  string `json:"key_from,omitempty"`
	} `json:"assets"`
	ID int `json:"id"`
}



func DataMountMedia() ([]MountMedia, error) {
    creatures, err := DataMount()
    if err != nil {
        fmt.Println("Erreur :", err)
        return nil, err
    }

    var mountMediaList []MountMedia
    for _, creature := range creatures {
        url := creature.CreatureDisplays[0].Key.Href + "&access_token=" + token.Access()
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

        // Mise à jour des données dans la structure MountMedia
        var updatedAssets []struct {
            Key     string `json:"key"`
            Value   string `json:"value"`
            KeyFrom string `json:"key_from,omitempty"`
        }
        for _, asset := range mediaResponse.Assets {
            updatedAsset := struct {
                Key     string `json:"key"`
                Value   string `json:"value"`
                KeyFrom string `json:"key_from,omitempty"`
            }{
                Key:     asset.Key,
                Value:   asset.Value,
                KeyFrom: asset.KeyFrom,
            }
            updatedAssets = append(updatedAssets, updatedAsset)
        }

        // Ajout de la structure mise à jour à la liste
        mountMediaList = append(mountMediaList, MountMedia{
            Assets: updatedAssets,
            ID:     mediaResponse.ID,
        })
    }
    return mountMediaList, nil
}


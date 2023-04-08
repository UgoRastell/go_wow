package data

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type MountInfos struct {
    ID                 int                        `json:"id"`
    Name               map[string]string         `json:"name"`
    Description        map[string]string         `json:"description"`
    Faction            MountFactionRequirements  `json:"requirements"`
    Source             MountSource               `json:"source"`
    CreatureDisplays   []CreatureDisplay         `json:"creature_displays"`
}

type CreatureDisplay struct {
    Key struct {
        Href string `json:"href"`
    } `json:"key"`

    ID int `json:"id"`
}

type MountSource struct {
    Type string              `json:"type"`
    Name map[string]string  `json:"name"`
}

type MountFactionRequirements struct {
    Faction MountFaction `json:"faction"`
}

type MountFaction struct {
    Type string             `json:"type"`
    Name map[string]string  `json:"name"`
}

func DataMount() ([]MountInfos, error) {
    mounts, err := MountIndex()
    if err != nil {
        fmt.Println("Erreur :", err)
        return nil, err
    }

    var mountInfosList []MountInfos

    for _, mount := range mounts {
        url := mount.URL.Href

        resp, err := http.Get(url)
        if err != nil {
            fmt.Println("Erreur :", err)
            return nil, err
        }
        defer resp.Body.Close()

        var mountInfos MountInfos
        err = json.NewDecoder(resp.Body).Decode(&mountInfos)
        if err != nil {
            fmt.Println("Erreur :", err)
            return nil, err
        }

        creatureDisplays := mountInfos.CreatureDisplays
        if len(creatureDisplays) > 0 {
            newSource := MountSource{
                Name: map[string]string{
                    "source": mountInfos.Source.Name["fr_FR"],
                },
            }

            newMountInfos := MountInfos{
                ID: mountInfos.CreatureDisplays[0].ID,
                Name: map[string]string{
                    "name": mountInfos.Name["fr_FR"],
                },
                Description: map[string]string{
                    "description": mountInfos.Description["fr_FR"],
                },
                Faction: mountInfos.Faction,
                Source: newSource,
                CreatureDisplays: []CreatureDisplay{
                    {
                        Key: struct {
                            Href string `json:"href"`
                        }{
                            Href: mountInfos.CreatureDisplays[0].Key.Href,
                        },
                        ID: mountInfos.CreatureDisplays[0].ID,
                    },
                },
            }            

            mountInfosList = append(mountInfosList, newMountInfos)
        }
    }
    return mountInfosList, nil
}
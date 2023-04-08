package data

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strconv"
)

type MountData struct {
    ID               int                 `json:"id"`
    Name             map[string]string   `json:"name"`
    Description      map[string]string   `json:"description"`
    Faction          MountFaction        `json:"faction"`
    Source           MountSource         `json:"source"`
    CreatureDisplays []CreatureDisplay   `json:"creature_displays"`
    Assets           []MountMediaAsset   `json:"assets"`
}

type MountMediaAsset struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

func Data(mountInfosList []MountInfos, mountMediaList []MountMedia) []MountData {
    mountDataMap := make(map[int]MountData)

    for _, mountInfo := range mountInfosList {
        mountDataMap[mountInfo.ID] = MountData{
            ID:               mountInfo.ID,
            Name:             mountInfo.Name,
            Description:      mountInfo.Description,
            Faction:          mountInfo.Faction.Faction,
            Source:           mountInfo.Source,
            CreatureDisplays: mountInfo.CreatureDisplays,
        }
    }

    for _, mountMedia := range mountMediaList {
        mountData, ok := mountDataMap[mountMedia.ID]
        if ok {
            assets := make([]MountMediaAsset, len(mountMedia.Assets))
            for i, asset := range mountMedia.Assets {
                assets[i] = MountMediaAsset{
                    Key:   asset.Key,
                    Value: asset.Value,
                }
            }
            mountData.Assets = assets
            mountDataMap[mountMedia.ID] = mountData
        }
    }

    mountDataList := make([]MountData, len(mountDataMap))
    i := 0
    for _, mountData := range mountDataMap {
        mountDataList[i] = mountData
        i++
    }

    return mountDataList
}

func WriteMountDataToFile(mountDataList []MountData, filePath string) error {
    jsonData, err := json.Marshal(mountDataList)
    if err != nil {
        fmt.Println("Erreur :", err)
        return err
    }

    err = ioutil.WriteFile(filePath, jsonData, 0644)
    if err != nil {
        fmt.Println("Erreur :", err)
        return err
    }

    return nil
}

func Run() {
    mountInfosList, err := DataMount()
    if err != nil {
        fmt.Println("Erreur :", err)
        return
    }

    mountMediaList, err := DataMountMedia()
    if err != nil {
        fmt.Println("Erreur :", err)
        return
    }

    mountDataList := Data(mountInfosList, mountMediaList)

    err = WriteMountDataToFile(mountDataList, "mountData.json")
    if err != nil {
        fmt.Println("Erreur :", err)
        return
    }

    fmt.Println(strconv.Itoa(len(mountDataList)) + " montures fusionnées et écrites dans mountData.json")
}
package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	db "wow/Database"
)
type MountData struct {
    ID               int               `json:"id"`
    Name             map[string]string `json:"name"`
    Description      map[string]string `json:"description"`
    Faction          MountFaction      `json:"faction"`
    Source           MountSource       `json:"source"`
    Assets           []MountMediaAsset `json:"assets"`
}
type MountMediaAsset struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func data(mountInfosList []MountInfos, mountMediaList []MountMedia) []MountData {
    mountDataList := make([]MountData, 0, len(mountInfosList))
    for _, mountInfo := range mountInfosList {
        mountData := MountData{
            ID:               mountInfo.ID,
            Name:             mountInfo.Name,
            Description:      mountInfo.Description,
            Faction:          mountInfo.Faction.Faction,
            Source:           mountInfo.Source,
        }

        for _, mountMedia := range mountMediaList {
            if mountInfo.ID == mountMedia.ID {
                assets := make([]MountMediaAsset, len(mountMedia.Assets))
                for i, asset := range mountMedia.Assets {
                    assets[i] = MountMediaAsset{
                        Key:   asset.Key,
                        Value: asset.Value,
                    }
                }
                mountData.Assets = assets
            }
        }

        mountDataList = append(mountDataList, mountData)
    }

    client, err := db.ConnexionDatabase()
    if err != nil {
        fmt.Printf("Erreur lors de la connexion à la base de données : %v\n", err)
        return nil
    }
	for _, mountData := range mountDataList {
		err = db.InsertDocument(client, "gowow", "mounts", mountData)
		if err != nil {
			fmt.Printf("Erreur lors de l'insertion des données dans la collection : %v\n", err)
			return nil
		}
	}
	
    return mountDataList
}


func writeMountDataToFile(mountDataList []MountData, filePath string) error {
	jsonData, err := json.Marshal(mountDataList)
	if err != nil {
		return fmt.Errorf("Erreur lors de la conversion des données en JSON : %v", err)
	}

	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'écriture des données dans le fichier : %v", err)
	}

	return nil
}

func Run() {
	mountInfosList, err := DataMount()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des données MountInfos : %v\n", err)
		return
	}

	mountMediaList, err := DataMountMedia()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des données MountMedia : %v\n", err)
		return
	}

	mountDataList := data(mountInfosList, mountMediaList)


	err = writeMountDataToFile(mountDataList, "mountData.json")
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture des données dans le fichier : %v\n", err)
		return
	}

}

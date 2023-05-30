package data

import (
	"fmt"
	db "wow/Database"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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

func Data(mountInfosList []MountInfos, mountMediaList []MountMedia) []MountData {
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

	filter := bson.M{}

	err = db.DeleteDocuments(client, "gowow", "mounts", filter)
	if err != nil {
		log.Printf("Erreur lors de la suppression des documents : %v\n", err)
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


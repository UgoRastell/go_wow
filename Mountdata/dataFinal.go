package data

import (
	"fmt"
	"log"

	db "wow/Database"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

type MountData struct {
	ID          int               `json:"id"`
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description"`
	Faction     MountFaction      `json:"faction"`
	Source      MountSource       `json:"source"`
	Assets      []MountMediaAsset `json:"assets"`
}

type MountMediaAsset struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Data(mountInfosList []MountInfos, mountMediaList []MountMedia) []MountData {
	mountDataList := make([]MountData, 0, len(mountInfosList))
	for _, mountInfo := range mountInfosList {
		mountData := MountData{
			ID:          mountInfo.ID,
			Name:        mountInfo.Name,
			Description: mountInfo.Description,
			Faction:     mountInfo.Faction.Faction,
			Source:      mountInfo.Source,
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

func Run(s *discordgo.Session, i *discordgo.InteractionCreate) []MountData {
	mountInfosList, err := DataMount()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des données MountInfos : %v\n", err)
		return nil
	}

	mountMediaList, err := DataMountMedia()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des données MountMedia : %v\n", err)
		return nil
	}

	// Envoi du message de chargement
	loadingMessage, err := s.ChannelMessageSend(i.ChannelID, "Chargement des montures en cours...")
	if err != nil {
		log.Printf("Erreur lors de l'envoi du message de chargement : %v\n", err)
		return nil
	}

	// Exécution du chargement des données
	go func() {
		mountDataList := Data(mountInfosList, mountMediaList)

		if len(mountDataList) > 0 {
			// Mise à jour du message de chargement avec le nombre de montures chargées
			updateMessage := fmt.Sprintf("Chargement des montures terminé. Nombre de montures chargées : %d", len(mountDataList))
			_, err := s.ChannelMessageEdit(i.ChannelID, loadingMessage.ID, updateMessage)
			if err != nil {
				log.Printf("Erreur lors de la mise à jour du message de chargement : %v\n", err)
			}
		} else {
			// Suppression du message de chargement s'il n'y a pas de montures chargées
			err = s.ChannelMessageDelete(i.ChannelID, loadingMessage.ID)
			if err != nil {
				log.Printf("Erreur lors de la suppression du message de chargement : %v\n", err)
			}
		}
	}()

	return nil
}

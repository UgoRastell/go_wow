package commande

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"wow/Mountdata"
)

func Run(s *discordgo.Session, i *discordgo.InteractionCreate) {
	mountInfosList, err := data.DataMount()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des données MountInfos : %v\n", err)
		return 
	}

	mountMediaList, err := data.DataMountMedia()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des données MountMedia : %v\n", err)
		return 
	}

	mountDataList := data.Data(mountInfosList, mountMediaList)

	fmt.Println(mountDataList)

	return
}

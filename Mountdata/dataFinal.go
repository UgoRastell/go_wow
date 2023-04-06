package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Data récupère les informations des montures et des médias associés, et crée un fichier JSON par monture contenant toutes les informations.
func Data() {
	// Récupération des informations des montures
	mounts, err := DataMount()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	// Récupération des médias des montures
	medias, err := DataMountMedia()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	// Création d'une map pour associer les médias aux montures par id
	mediasMap := make(map[int][]MountMedia)
	for _, media := range medias {
		mediasMap[media.ID] = append(mediasMap[media.ID], media)
	}

	// Parcours des montures pour créer un fichier par id de monture
	for _, mount := range mounts {
		// Récupération des médias associés à cette monture
		mountMedias, ok := mediasMap[mount.ID]
		if !ok {
			fmt.Printf("Aucun média trouvé pour la monture avec l'id %d\n", mount.ID)
			continue
		}

		// Récupération de la bonne image pour chaque créature_display
		for i, creatureDisplay := range mount.CreatureDisplays {
			for _, media := range mountMedias {
				if media.Key.Href == creatureDisplay.Key.Href {
					mount.CreatureDisplays[i].ID = media.ID
					break
				}
			}
		}

		// Conversion de la monture et de ses médias en JSON
		mountWithMedia := struct {
			ID             int                    `json:"id"`
			Name           map[string]string     `json:"name"`
			Description    map[string]string     `json:"description"`
			Requirements   MountRequirements     `json:"requirements"`
			Source         MountSource           `json:"source"`
			CreatureDisplays []MountCreatureDisplay `json:"creature_displays"`
			Medias         []MountMedia           `json:"medias"`
		}{
			ID:             mount.ID,
			Name:           mount.Name,
			Description:    mount.Description,
			Requirements:   mount.Requirements,
			Source:         mount.Source,
			CreatureDisplays: mount.CreatureDisplays,
			Medias:         mountMedias,
		}

		// Écriture du fichier JSON
		mountWithMediaJSON, err := json.MarshalIndent(mountWithMedia, "", "    ")
		if err != nil {
			fmt.Printf("Erreur lors de la conversion de la monture avec l'id %d en JSON : %v\n", mount.ID, err)
			continue
		}

		filename := fmt.Sprintf("mount_%d.json", mount.ID)
		err = ioutil.WriteFile(filename, mountWithMediaJSON, 0644)
		if err != nil {
			fmt.Printf("Erreur lors de l'écriture du fichier %s : %v\n", filename, err)
			continue
		}

		fmt.Printf("Fichier %s créé avec succès\n", filename)
	}
}

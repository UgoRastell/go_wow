package commande

import (
    "encoding/json"
    "fmt"
    "net/http"
	"os"
	"io"
)

// Les données des montures sont stockées dans la propriété "mounts" de la réponse de l'API.

type Mount struct {
    Key  MountKey `json:"key"`
    Name string   `json:"name"`
    ID   int      `json:"id"`
}

type MountKey struct {
    Href string `json:"href"`
}

type MountsResponse struct {
    Mounts []Mount `json:"mounts"`
}




func MountIndex() {
    // Définir l'URL de l'API Blizzard
    url := "https://eu.api.blizzard.com/data/wow/mount/index?namespace=static-eu&locale=fr_FR&access_token=EUj0pcNc29UIRkc3IE5D5drxZRKxNENcyC"

    // Envoyer une requête GET à l'API Blizzard
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Erreur :", err)
        return
    }
    defer resp.Body.Close()

    // Décoder la réponse JSON en une structure Go
    var mountsResponse MountsResponse
    err = json.NewDecoder(resp.Body).Decode(&mountsResponse)
    if err != nil {
        fmt.Println("Erreur :", err)
        return
    }

    // Convertir la structure de données en JSON
    jsonData, err := json.Marshal(mountsResponse)
    if err != nil {
        fmt.Println("Erreur :", err)
        return
    }

	// Open a file for writing
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = io.WriteString(file, string(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

    fmt.Printf("Data to data.json")

    
}

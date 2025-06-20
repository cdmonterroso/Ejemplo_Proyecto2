package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Estructura Partition
type Partition struct {
	Name   string `json:"name"`
	SizeKB int    `json:"sizeKB"`
	Type   string `json:"type"` // P o E
	Fit    string `json:"fit"`
}

// Estructura Disk
type Disk struct {
	Letter     string      `json:"letter"`
	SizeMB     int         `json:"sizeMB"`
	Partitions []Partition `json:"partitions"`
}

// Base de datos simulada en memoria
var disks = []Disk{
	{
		Letter: "A",
		SizeMB: 10,
		Partitions: []Partition{
			{"A1", 1000, "P", "FF"},
			{"A2", 1500, "P", "WF"},
			{"A3", 1200, "P", "BF"},
			{"A4", 800, "P", "FF"},
		},
	},
	{
		Letter: "B",
		SizeMB: 15,
		Partitions: []Partition{
			{"B1", 2000, "P", "BF"},
			{"B2", 1000, "P", "FF"},
			{"B3", 500, "P", "WF"},
			{"B4", 2500, "P", "FF"},
		},
	},
	{
		Letter: "C",
		SizeMB: 20,
		Partitions: []Partition{
			{"C1", 3000, "P", "FF"},
			{"C2", 1500, "P", "WF"},
			{"CEXT", 4000, "E", "BF"},
		},
	},
	{
		Letter: "D",
		SizeMB: 25,
		Partitions: []Partition{
			{"D1", 3500, "P", "BF"},
			{"D2", 2000, "P", "WF"},
			{"DEXT", 5000, "E", "FF"},
		},
	},
}

func getDisks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(disks)
}

func getPartitionsByDisk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	letter := params["letter"]

	//Recorre todos los discos
	for _, d := range disks {
		if d.Letter == letter {
			json.NewEncoder(w).Encode(d.Partitions)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Disco no encontrado"})
}

func main() {
	//Crea un nuevo router
	r := mux.NewRouter()

	//Define petición GET
	r.HandleFunc("/api/discos", getDisks).Methods("GET")
	r.HandleFunc("/api/discos/{letter}/particiones", getPartitionsByDisk).Methods("GET")

	// Habilitar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}

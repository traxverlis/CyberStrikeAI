package main

import (
	"cyberstrike-ai/internal/app"
	"cyberstrike-ai/internal/config"
	"cyberstrike-ai/internal/logger"
	"flag"
	"fmt"
)

func main() {
	var configPath = flag.String("config", "config.yaml", "chemin du fichier de configuration")
	flag.Parse()

	// Charger la configuration
	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Printf("Échec du chargement de la configuration: %v\n", err)
		return
	}

	// Initialiser le journal
	log := logger.New(cfg.Log.Level, cfg.Log.Output)

	// Créer l'application
	application, err := app.New(cfg, log)
	if err != nil {
		log.Fatal("Échec de l'initialisation de l'application", "error", err)
	}

	// Démarrer le serveur
	if err := application.Run(); err != nil {
		log.Fatal("Échec du démarrage du serveur", "error", err)
	}
}


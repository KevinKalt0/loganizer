# Loganizer

## Description

**Loganizer** est un outil en ligne de commande (CLI) écrit en Go, permettant d'analyser efficacement plusieurs fichiers de logs en parallèle, de centraliser les résultats et de générer des rapports détaillés.  
Il est conçu pour les administrateurs système et développeurs souhaitant automatiser l'analyse de journaux applicatifs ou serveurs.

---

## Fonctionnalités

- **Analyse concurrente** de plusieurs fichiers de logs (utilisation des goroutines et WaitGroup)
- **Gestion robuste des erreurs** (fichier introuvable, parsing, etc.)
- **Affichage d'un résumé clair** pour chaque log analysé
- **Export des résultats** au format JSON (création automatique des dossiers si besoin)
- **Extensible** : architecture modulaire, facile à enrichir avec de nouvelles commandes

---

## Structure du projet

---

## Utilisation

### 1. **Analyser des logs**

```sh
Tester l'export``go run main.go analyze -c config.json -o report.json``
```

- `--config` ou `-c` : chemin vers le fichier de configuration JSON (obligatoire)
- `--output` ou `-o` : chemin du rapport JSON à générer (optionnel)

### 2. **Fichier de configuration**

```json
[
  {
    "id": "web-server-1",
    "path": "test_logs/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "test_logs/errors.log",
    "type": "custom-app"
  }
]
```

---

## report JSON généré

```json
[
  {
    "log_id": "db-server-3",
    "file_path": "test_logs/mysql_error.log",
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "CreateFile test_logs/mysql_error.log: The system cannot find the file specified."
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log",
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "CreateFile /non/existent/log.log: The system cannot find the path specified."
  }
]
```

---

## 🚀 Installation rapide

1. **Cloner le dépôt**
   ```sh
   git clone https://github.com/KevinKalt0/loganizer
   cd loganizer
   ```

2. **Installer les dépendances**
   ```sh
   go mod tidy
   ```

3. **Compiler le projet**
   ```sh
   go build -o loganizer .
   ```

---

## Architecture technique

- **cmd/** : commandes CLI (Cobra)
- **internal/config/** : lecture du fichier de configuration JSON
- **internal/analyzer/** : logique d'analyse, erreurs personnalisées
- **internal/reporter/** : export des résultats
- **test_logs/** : exemples de logs pour tests

---

## Membres de l'équipe

- DYLAN LERAY 
- RYAN HO 
- NATHANIEL ANTON HILLARY
- KEVIN KALTENIS 

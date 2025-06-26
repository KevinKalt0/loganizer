# Loganalyzer

## Description

**Loganalyzer** est un outil en ligne de commande (CLI) écrit en Go, permettant d’analyser efficacement plusieurs fichiers de logs en parallèle, de centraliser les résultats et de générer des rapports détaillés.  
Il est conçu pour les administrateurs système et développeurs souhaitant automatiser l’analyse de journaux applicatifs ou serveurs.

---

## Fonctionnalités

- **Analyse concurrente** de plusieurs fichiers de logs (utilisation des goroutines et WaitGroup)
- **Gestion robuste des erreurs** (fichier introuvable, parsing, etc.)
- **Affichage d’un résumé clair** pour chaque log analysé
- **Export des résultats** au format JSON (création automatique des dossiers si besoin)
- **Extensible** : architecture modulaire, facile à enrichir avec de nouvelles commandes

---

## Structure du projet

---

## Utilisation

### 1. **Analyser des logs**

```sh
loganalyzer analyze --config config.json
```

- `--config` ou `-c` : chemin vers le fichier de configuration JSON (obligatoire)
- `--output` ou `-o` : chemin du rapport JSON à générer (optionnel)

### 2. **Exemple de fichier de configuration**

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

## Exemple de rapport JSON généré

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "test_logs/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "test_logs/absent.log",
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "open test_logs/absent.log: no such file or directory"
  }
]
```

---

## Installation

1. **Cloner le dépôt**
   ```sh
   git clone <url_du_repo>
   cd loganizer
   ```

2. **Installer les dépendances**
   ```sh
   go mod tidy
   ```

3. **Compiler**
   ```sh
   go build -o loganalyzer .
   ```

---

## Architecture technique

- **cmd/** : commandes CLI (Cobra)
- **internal/config/** : lecture du fichier de configuration JSON
- **internal/analyzer/** : logique d’analyse, erreurs personnalisées
- **internal/reporter/** : export des résultats
- **test_logs/** : exemples de logs pour tests

---

## Membres de l’équipe

- *À compléter avec les noms/prénoms de l’équipe*

---

## Pour aller plus loin (bonus possibles)

- Ajout d’une sous-commande `add-log` pour enrichir le fichier de configuration
- Filtrage des résultats par statut (`--status OK/FAILED`)
- Ajout d’un horodatage automatique dans le nom du rapport exporté

---

## Licence

MIT

---

## Aide

Pour afficher l’aide :
```sh
loganalyzer --help
loganalyzer analyze --help
```
```

---

Si tu veux un README encore plus détaillé (badges, screenshots, etc.), ou une section sur la contribution, dis-le-moi !
<div align="center">
  <img src="web/static/logo.png" alt="Logo CyberStrikeAI" width="200">
</div>

# CyberStrikeAI

[中文](README_CN.md) | [English](README.md) | [Français](README_FR.md)

CyberStrikeAI est une **plateforme de test de sécurité native IA** construite en Go, intégrant 100+ outils de sécurité, un moteur d'orchestration intelligent, des tests basés sur les rôles avec des rôles de test de sécurité prédéfinis, un système de compétences (Skills) avec des compétences de test spécialisées, ainsi que des capacités complètes de gestion du cycle de vie des tests. Grâce au protocole MCP natif et aux agents IA, elle permet l'automatisation de bout en bout, de la commande conversationnelle à la découverte de vulnérabilités, l'analyse de la chaîne d'attaque, la récupération de connaissances et la visualisation des résultats, offrant un environnement de test professionnel auditable, traçable et collaboratif pour les équipes de sécurité.


## Aperçu de l'interface et de l'intégration

<div align="center">

### Vue d'ensemble du tableau de bord système

<img src="./images/dashboard.png" alt="Tableau de bord système" width="100%">

*Le tableau de bord fournit une vue d'ensemble complète de l'état d'exécution du système, des vulnérabilités de sécurité, de l'utilisation des outils et de la base de connaissances, aidant les utilisateurs à comprendre rapidement les fonctionnalités principales de la plateforme et son état actuel.*

### Vue d'ensemble des fonctionnalités principales

<table>
<tr>
<td width="33.33%" align="center">
<strong>Console Web</strong><br/>
<img src="./images/web-console.png" alt="Console Web" width="100%">
</td>
<td width="33.33%" align="center">
<strong>Visualisation de la chaîne d'attaque</strong><br/>
<img src="./images/attack-chain.png" alt="Chaîne d'attaque" width="100%">
</td>
<td width="33.33%" align="center">
<strong>Gestion des tâches</strong><br/>
<img src="./images/task-management.png" alt="Gestion des tâches" width="100%">
</td>
</tr>
<tr>
<td width="33.33%" align="center">
<strong>Gestion des vulnérabilités</strong><br/>
<img src="./images/vulnerability-management.png" alt="Gestion des vulnérabilités" width="100%">
</td>
<td width="33.33%" align="center">
<strong>Gestion MCP</strong><br/>
<img src="./images/mcp-management.png" alt="Gestion MCP" width="100%">
</td>
<td width="33.33%" align="center">
<strong>Mode MCP stdio</strong><br/>
<img src="./images/mcp-stdio2.png" alt="Mode MCP stdio" width="100%">
</td>
</tr>
<tr>
<td width="33.33%" align="center">
<strong>Base de connaissances</strong><br/>
<img src="./images/knowledge-base.png" alt="Base de connaissances" width="100%">
</td>
<td width="33.33%" align="center">
<strong>Gestion des Skills</strong><br/>
<img src="./images/skills.png" alt="Gestion des Skills" width="100%">
</td>
<td width="33.33%" align="center">
<strong>Gestion des rôles</strong><br/>
<img src="./images/role-management.png" alt="Gestion des rôles" width="100%">
</td>
</tr>
</table>

</div>

## Aperçu des fonctionnalités

- 🤖 Moteur de décision IA compatible avec les modèles OpenAI/DeepSeek/Claude, etc.
- 🔌 Protocole MCP natif, prenant en charge les modes de transport HTTP / stdio / SSE ainsi que l'intégration MCP externe
- 🧰 100+ modèles d'outils prêts à l'emploi + capacité d'extension YAML
- 📄 Pagination, compression et recherche plein texte pour les gros résultats
- 🔗 Visualisation de la chaîne d'attaque, notation des risques et relecture étape par étape
- 🔒 Protection de connexion Web, journaux d'audit, persistance SQLite
- 📚 Fonctionnalité de base de connaissances : recherche vectorielle et recherche hybride, fournissant des connaissances spécialisées en sécurité pour l'IA
- 📁 Gestion des groupes de conversation : prise en charge de la création de groupes, épinglage, renommage, suppression, etc.
- 🛡️ Fonctionnalité de gestion des vulnérabilités : opérations CRUD complètes sur les vulnérabilités, prise en charge de la classification par gravité, flux de statut, filtrage par conversation/gravité/statut, ainsi qu'un tableau de bord statistique
- 📋 Gestion des tâches par lots : créer des files d'attente de tâches, ajouter des tâches en masse, les exécuter séquentiellement, avec prise en charge de l'édition et du suivi de l'état des tâches
- 🎭 Tests basés sur les rôles : rôles de test de sécurité prédéfinis (test de pénétration, CTF, scan d'applications Web, etc.), avec prise en charge de prompts personnalisés et de restrictions d'outils
- 🎯 Système de compétences (Skills) : 20+ compétences de test de sécurité prédéfinies (injection SQL, XSS, sécurité API, etc.), pouvant être attachées à des rôles ou appelées à la demande par l'IA
- 📱 **Robot** : prise en charge des connexions longues DingTalk et Lark (Feishu), pour dialoguer avec CyberStrikeAI depuis un téléphone mobile (voir [Guide d'utilisation du robot](docs/robot.md) pour la configuration et les commandes)

## Vue d'ensemble des outils

Le système est préchargé avec 100+ outils de pénétration/défense, couvrant la chaîne d'attaque complète :

- **Scan réseau** : nmap, masscan, rustscan, arp-scan, nbtscan
- **Scan d'applications Web** : sqlmap, nikto, dirb, gobuster, feroxbuster, ffuf, httpx
- **Scan de vulnérabilités** : nuclei, wpscan, wafw00f, dalfox, xsser
- **Énumération de sous-domaines** : subfinder, amass, findomain, dnsenum, fierce
- **Moteurs de recherche d'espace réseau** : fofa_search, zoomeye_search
- **Sécurité API** : graphql-scanner, arjun, api-fuzzer, api-schema-analyzer
- **Sécurité des conteneurs** : trivy, clair, docker-bench-security, kube-bench, kube-hunter
- **Sécurité cloud** : prowler, scout-suite, cloudmapper, pacu, terrascan, checkov
- **Analyse binaire** : gdb, radare2, ghidra, objdump, strings, binwalk
- **Exploitation de vulnérabilités** : metasploit, msfvenom, pwntools, ropper, ropgadget
- **Cassage de mots de passe** : hashcat, john, hashpump
- **Analyse forensique** : volatility, volatility3, foremost, steghide, exiftool
- **Post-exploitation** : linpeas, winpeas, mimikatz, bloodhound, impacket, responder
- **Outils utilitaires CTF** : stegsolve, zsteg, hash-identifier, fcrackzip, pdfcrack, cyberchef
- **Auxiliaires système** : exec, create-file, delete-file, list-files, modify-file

## Utilisation de base

### Démarrage rapide (déploiement en une commande)

**Prérequis :**
- Go 1.21+ ([Télécharger](https://go.dev/dl/))
- Python 3.10+ ([Télécharger](https://www.python.org/downloads/))

**Déploiement en une commande :**
```bash
git clone https://github.com/Ed1s0nZ/CyberStrikeAI.git
cd CyberStrikeAI-main
chmod +x run.sh && ./run.sh
```

Le script `run.sh` effectuera automatiquement :
- ✅ Vérification et validation des environnements Go et Python
- ✅ Création d'un environnement virtuel Python
- ✅ Installation des dépendances Python
- ✅ Téléchargement des modules de dépendances Go
- ✅ Compilation et construction du projet
- ✅ Démarrage du serveur

**Configuration initiale :**
1. **Configurer l'API du modèle IA** (obligatoire avant la première utilisation)
   - Après le démarrage, accédez à http://localhost:8080
   - Allez dans `Paramètres` → Remplissez les informations de configuration de l'API :
     ```yaml
     openai:
       api_key: "sk-your-key"
       base_url: "https://api.openai.com/v1"  # ou https://api.deepseek.com/v1
       model: "gpt-4o"  # ou deepseek-chat, claude-3-opus, etc.
     ```
   - Ou modifiez directement le fichier `config.yaml` avant le démarrage
2. **Connexion au système** - Utilisez le mot de passe généré automatiquement affiché dans la console (ou définissez `auth.password` dans `config.yaml`)
3. **Installer les outils de sécurité (facultatif)** - Installez les outils nécessaires selon vos besoins :
   ```bash
   # macOS
   brew install nmap sqlmap nuclei httpx gobuster feroxbuster subfinder amass
   # Ubuntu/Debian
   sudo apt-get install nmap sqlmap nuclei httpx gobuster feroxbuster
   ```
   Les outils non installés seront automatiquement ignorés ou remplacés par des alternatives.

**Autres méthodes de démarrage :**
```bash
# Exécution directe (nécessite une configuration manuelle de l'environnement)
go run cmd/server/main.go

# Compilation manuelle
go build -o cyberstrike-ai cmd/server/main.go
./cyberstrike-ai
```

**Note :** L'environnement virtuel Python (`venv/`) est automatiquement créé et géré par `run.sh`. Les outils nécessitant Python (comme `api-fuzzer`, `http-framework-test`, etc.) utiliseront automatiquement cet environnement.

### Flux de travail courants
- **Test conversationnel** : Déclencher une orchestration d'outils en plusieurs étapes via un langage naturel, avec sortie en temps réel SSE.
- **Tests basés sur les rôles** : Sélectionner parmi les rôles de test de sécurité prédéfinis (test de pénétration, CTF, scan d'applications Web, test de sécurité API, etc.) pour personnaliser le comportement de l'IA et les outils disponibles. Chaque rôle peut appliquer des prompts système personnalisés et limiter la liste des outils disponibles pour réaliser des scénarios de test ciblés.
- **Surveillance des outils** : Consulter les files d'attente de tâches, les journaux d'exécution, les pièces jointes de gros fichiers.
- **Historique des sessions** : Toutes les conversations et appels d'outils sont sauvegardés dans SQLite et peuvent être relus à tout moment.
- **Groupes de conversation** : Organiser les conversations par projet ou thème dans différents groupes, avec prise en charge de l'épinglage, du renommage, de la suppression, etc., toutes les données étant stockées de manière persistante.
- **Gestion des vulnérabilités** : Créer, mettre à jour et suivre les vulnérabilités découvertes lors des tests. Prise en charge du filtrage par gravité (critique/élevé/moyen/faible/information), statut (en attente de confirmation/confirmé/corrigé/faux positif) et par conversation, visualisation des statistiques et exportation des découvertes.
- **Gestion des tâches par lots** : Créer des files d'attente de tâches, ajouter plusieurs tâches en masse, les éditer ou les supprimer avant exécution, puis les exécuter séquentiellement. Chaque tâche sera exécutée comme une conversation indépendante, avec prise en charge complète du suivi de l'état (en attente/en cours/terminé/échoué/annulé) et de l'historique d'exécution.
- **Configuration visuelle** : Changer de modèle, activer/désactiver des outils, définir le nombre d'itérations, etc. depuis l'interface.

### Mesures de sécurité par défaut
- Validation obligatoire intégrée dans le panneau de configuration pour éviter les clés API/Base URL/modèle manquants.
- Génération automatique d'un mot de passe fort de 24 caractères et écriture dans `config.yaml` lorsque `auth.password` est vide.
- Toutes les API (sauf la connexion) nécessitent un Bearer Token, avec middleware d'authentification unifié.
- Chaque exécution d'outil dispose d'un timeout, de journaux et d'une isolation des erreurs.

## Utilisation avancée

### Tests basés sur les rôles
- **Rôles prédéfinis** : Le système intègre 12+ rôles de test de sécurité prédéfinis (test de pénétration, CTF, scan d'applications Web, test de sécurité API, analyse binaire, audit de sécurité cloud, etc.), situés dans le répertoire `roles/`.
- **Prompts personnalisés** : Chaque rôle peut définir un `user_prompt`, qui sera automatiquement ajouté avant le message de l'utilisateur, guidant l'IA pour adopter des méthodes de test spécifiques et des points d'attention.
- **Restriction des outils** : Les rôles peuvent spécifier une liste `tools`, limitant les outils disponibles pour réaliser des flux de test ciblés (par exemple, le rôle CTF limite aux outils dédiés CTF).
- **Intégration des Skills** : Les rôles peuvent attacher des compétences de test de sécurité. Les noms des compétences sont ajoutés au prompt système en tant qu'indications, et l'agent IA peut obtenir le contenu des compétences à la demande via l'outil `read_skill`.
- **Création facile de rôles** : Créer des rôles personnalisés en ajoutant des fichiers YAML dans le répertoire `roles/`. Chaque rôle définit les champs `name`, `description`, `user_prompt`, `icon`, `tools`, `skills`, `enabled`.
- **Intégration à l'interface Web** : Sélectionner un rôle via le menu déroulant dans l'interface de chat. La sélection du rôle affecte le comportement de l'IA et les suggestions d'outils disponibles.

**Exemple de création d'un rôle personnalisé :**
1. Créer un fichier YAML dans le répertoire `roles/` (par exemple `roles/role-personnalise.yaml`) :
   ```yaml
   name: Rôle personnalisé
   description: Scénario de test spécialisé
   user_prompt: Vous êtes un testeur de sécurité spécialisé dans la sécurité API...
   icon: "\U0001F4E1"
   tools:
     - api-fuzzer
     - arjun
     - graphql-scanner
   skills:
     - api-security-testing
     - sql-injection-testing
   enabled: true
   ```
2. Redémarrer le service ou recharger la configuration ; le rôle apparaîtra dans le menu déroulant de sélection des rôles.

### Système de compétences (Skills)
- **Compétences prédéfinies** : Le système intègre 20+ compétences de test de sécurité prédéfinies (injection SQL, XSS, sécurité API, sécurité cloud, sécurité des conteneurs, etc.), situées dans le répertoire `skills/`.
- **Indications de compétences dans les prompts** : Lors de la sélection d'un rôle, les noms des compétences attachées à ce rôle sont ajoutés au prompt système en tant que recommandations. Le contenu des compétences n'est pas automatiquement injecté ; l'agent IA doit utiliser l'outil `read_skill` pour accéder aux détails de la compétence lorsque nécessaire.
- **Accès à la demande** : L'agent IA peut également accéder aux compétences à la demande via des outils intégrés (`list_skills`, `read_skill`), permettant une récupération dynamique des compétences pendant l'exécution des tâches.
- **Format structuré** : Chaque compétence est un répertoire contenant un fichier `SKILL.md`, décrivant en détail les méthodes de test, l'utilisation des outils, les meilleures pratiques et des exemples. Les compétences prennent en charge le format YAML front matter pour les métadonnées.
- **Compétences personnalisées** : Créer des compétences personnalisées en ajoutant des répertoires dans le répertoire `skills/`. Chaque répertoire de compétence doit contenir un fichier `SKILL.md`.

**Créer une compétence personnalisée :**
1. Créer un répertoire dans `skills/` (par exemple `skills/ma-competence/`)
2. Créer un fichier `SKILL.md` dans ce répertoire avec le contenu de la compétence
3. Dans le fichier YAML du rôle, ajouter la compétence en l'ajoutant au champ `skills`

### Orchestration et extension des outils
- `tools/*.yaml` définit les commandes, paramètres, prompts et métadonnées, avec rechargement à chaud possible.
- `security.tools_dir` pointe vers un répertoire pour activer en masse ; la définition inline dans la configuration principale reste prise en charge.
- **Pagination des gros résultats** : Les sorties dépassant 200 Ko sont sauvegardées en tant que pièces jointes, accessibles via l'outil `query_execution_result` avec pagination, filtres et recherche par regex.
- **Compression/résumé des résultats** : Les journaux de plusieurs mégaoctets peuvent être compressés ou résumés avant d'être écrits dans SQLite, réduisant ainsi la taille de l'archive.

**Étapes générales pour créer un outil personnalisé**
1. Copier un exemple existant dans `tools/` (par exemple `tools/sample.yaml`).
2. Modifier `name`, `command`, `args`, `short_description` et autres informations de base.
3. Déclarer les paramètres positionnels ou avec flags dans `parameters[]`, facilitant l'assemblage automatique des commandes par l'agent.
4. Compléter si nécessaire `description` ou `notes`, fournissant un contexte supplémentaire ou des conseils d'interprétation des résultats pour l'IA.
5. Redémarrer le service ou recharger la configuration dans l'interface, le nouvel outil pourra alors être activé/désactivé dans le panneau Settings.

### Analyse de la chaîne d'attaque
- L'agent analyse chaque conversation, extrayant les cibles, outils, vulnérabilités et relations de causalité.
- L'interface Web permet de visualiser de manière interactive les nœuds de la chaîne, le niveau de risque et la chronologie, avec prise en charge de l'exportation de rapports.

### MCP dans tous les scénarios
- **Mode Web** : Intègre un service MCP HTTP pour les appels front-end.
- **Mode MCP stdio** : `go run cmd/mcp-stdio/main.go` peut s'intégrer avec Cursor/ligne de commande.
- **Fédération MCP externe** : Enregistrer des MCP tiers (HTTP/stdio/SSE) dans les paramètres, activer/désactiver selon les besoins et visualiser en temps réel les statistiques d'appel et la santé.

#### Intégration rapide MCP stdio
1. **Compiler le fichier exécutable** (exécuter depuis la racine du projet) :
   ```bash
   go build -o cyberstrike-ai-mcp cmd/mcp-stdio/main.go
   ```
2. **Configurer dans Cursor**
   Ouvrir `Settings → Tools & MCP → Add Custom MCP`, sélectionner **Command**, spécifier le programme compilé et le fichier de configuration :
   ```json
   {
     "mcpServers": {
       "cyberstrike-ai": {
         "command": "/chemin/absolu/vers/cyberstrike-ai-mcp",
         "args": [
           "--config",
           "/chemin/absolu/vers/config.yaml"
         ]
       }
     }
   }
   ```
   Remplacer les chemins par vos adresses locales réelles, Cursor démarrera automatiquement la version stdio du MCP.

#### Intégration rapide MCP HTTP
1. Confirmer que `mcp.enabled: true` dans `config.yaml`, ajuster `mcp.host` / `mcp.port` selon les besoins (recommandé `127.0.0.1:8081` en local).
2. Démarrer le service principal (`./run.sh` ou `go run cmd/server/main.go`), le point de terminaison MCP est exposé par défaut sur `http://<host>:<port>/mcp`.
3. Dans Cursor `Add Custom MCP → HTTP`, définir `Base URL` sur `http://127.0.0.1:8081/mcp`.
4. Vous pouvez également créer `.cursor/mcp.json` à la racine du projet pour le partage d'équipe :
   ```json
   {
     "mcpServers": {
       "cyberstrike-ai-http": {
         "transport": "http",
         "url": "http://127.0.0.1:8081/mcp"
       }
     }
   }
   ```

#### Fédération MCP externe (HTTP/stdio/SSE)
CyberStrikeAI prend en charge la connexion à des serveurs MCP externes via trois modes de transport :
- **Mode HTTP** – Communication requête/réponse traditionnelle via HTTP POST
- **Mode stdio** – Communication inter-processus via entrée/sortie standard
- **Mode SSE** – Communication en temps réel via Server-Sent Events

Ajouter un serveur MCP externe :
1. Ouvrir l'interface Web, aller dans **Paramètres → MCP externe**.
2. Cliquer sur **Ajouter MCP externe**, fournir la configuration au format JSON :

   **Exemple mode HTTP :**
   ```json
   {
     "mon-mcp-http": {
       "transport": "http",
       "url": "http://127.0.0.1:8081/mcp",
       "description": "Serveur MCP HTTP",
       "timeout": 30
     }
   }
   ```

   **Exemple mode stdio :**
   ```json
   {
     "mon-mcp-stdio": {
       "command": "python3",
       "args": ["/chemin/vers/mcp-server.py"],
       "description": "Serveur MCP stdio",
       "timeout": 30
     }
   }
   ```

   **Exemple mode SSE :**
   ```json
   {
     "mon-mcp-sse": {
       "transport": "sse",
       "url": "http://127.0.0.1:8082/sse",
       "description": "Serveur MCP SSE",
       "timeout": 30
     }
   }
   ```

3. Cliquer sur **Enregistrer**, puis cliquer sur **Démarrer** pour se connecter au serveur.
4. Surveiller en temps réel l'état de connexion, le nombre d'outils et la santé.

**Avantages du mode SSE :**
- Communication bidirectionnelle en temps réel via Server-Sent Events
- Adapté aux scénarios nécessitant des flux de données continus
- Latence plus faible pour les notifications basées sur push

Un serveur MCP SSE de test est disponible dans le répertoire `cmd/test-sse-mcp-server/` pour validation.


### Fonctionnalité de base de connaissances
- **Recherche vectorielle** : L'agent IA peut automatiquement appeler l'outil `search_knowledge_base` pour rechercher des connaissances de sécurité dans la base de connaissances lors des conversations.
- **Recherche hybride** : Combine la recherche par similarité vectorielle et la correspondance de mots-clés pour améliorer la précision de la recherche.
- **Indexation automatique** : Scanne les fichiers Markdown du répertoire `knowledge_base/`, construit automatiquement un index d'embeddings vectoriels.
- **Gestion Web** : Créer, mettre à jour, supprimer des éléments de connaissance via l'interface Web, avec gestion par catégorie.
- **Journaux de recherche** : Enregistre toutes les opérations de recherche de connaissances, facilitant l'audit et le débogage.

**Démarrage rapide (utilisation de la base de connaissances pré-construite) :**
1. **Télécharger la base de données de connaissances** : Depuis [GitHub Releases](https://github.com/Ed1s0nZ/CyberStrikeAI/releases), télécharger le fichier de base de données de connaissances pré-construit.
2. **Extraire et placer** : Extraire le fichier de base de données de connaissances téléchargé (`knowledge.db`) et le placer dans le répertoire `data/` du projet.
3. **Redémarrer le service** : Redémarrer le service CyberStrikeAI, la base de connaissances sera prête à l'emploi immédiatement, sans nécessité de reconstruire l'index.

**Étapes de configuration de la base de connaissances :**
1. **Activer la fonctionnalité** : Dans `config.yaml`, définir `knowledge.enabled: true` :
   ```yaml
   knowledge:
     enabled: true
     base_path: knowledge_base
     embedding:
       provider: openai
       model: text-embedding-v4
       base_url: "https://api.openai.com/v1"  # ou votre API de modèle d'embedding
       api_key: "sk-xxx"
     retrieval:
       top_k: 5
       similarity_threshold: 0.7
       hybrid_weight: 0.7
   ```
2. **Ajouter des fichiers de connaissances** : Placer les fichiers Markdown dans le répertoire `knowledge_base/`, organisés par catégorie (par exemple `knowledge_base/Injection SQL/README.md`).
3. **Scanner l'index** : Dans l'interface Web, cliquer sur "Scanner la base de connaissances", le système importera automatiquement les fichiers et construira l'index vectoriel.
4. **Utiliser dans les conversations** : L'agent IA appellera automatiquement l'outil de recherche de connaissances lorsqu'il aura besoin de connaissances de sécurité. Vous pouvez également demander explicitement : "Rechercher dans la base de connaissances les techniques d'injection SQL".

**Description de la structure de la base de connaissances :**
- Les fichiers sont organisés par catégorie (le nom du répertoire devient la catégorie).
- Chaque fichier Markdown est automatiquement découpé en morceaux et génère des embeddings vectoriels.
- Prise en charge des mises à jour incrémentales, les fichiers modifiés sont automatiquement réindexés.


### Automatisation et sécurité
- **REST API** : Authentification, sessions, tâches, surveillance, gestion des vulnérabilités, gestion des rôles et toutes les autres interfaces sont ouvertes, intégration possible avec CI/CD.
- **API de gestion des rôles** : Gérer les rôles de test de sécurité via le point de terminaison `/api/roles` : `GET /api/roles` (liste), `GET /api/roles/:name` (obtenir un rôle), `POST /api/roles` (créer un rôle), `PUT /api/roles/:name` (mettre à jour un rôle), `DELETE /api/roles/:name` (supprimer un rôle). Les rôles sont stockés sous forme de fichiers YAML dans le répertoire `roles/`, avec prise en charge du rechargement à chaud.
- **API de gestion des vulnérabilités** : Gérer les vulnérabilités via le point de terminaison `/api/vulnerabilities` : `GET /api/vulnerabilities` (liste, avec filtres), `POST /api/vulnerabilities` (créer), `GET /api/vulnerabilities/:id` (obtenir), `PUT /api/vulnerabilities/:id` (mettre à jour), `DELETE /api/vulnerabilities/:id` (supprimer), `GET /api/vulnerabilities/stats` (statistiques).
- **API de tâches par lots** : Gérer les files d'attente de tâches par lots via le point de terminaison `/api/batch-tasks` : `POST /api/batch-tasks` (créer une file d'attente), `GET /api/batch-tasks` (liste), `GET /api/batch-tasks/:queueId` (obtenir une file d'attente), `POST /api/batch-tasks/:queueId/start` (démarrer l'exécution), `POST /api/batch-tasks/:queueId/cancel` (annuler), `DELETE /api/batch-tasks/:queueId` (supprimer une file d'attente), `POST /api/batch-tasks/:queueId/tasks` (ajouter une tâche), `PUT /api/batch-tasks/:queueId/tasks/:taskId` (mettre à jour une tâche), `DELETE /api/batch-tasks/:queueId/tasks/:taskId` (supprimer une tâche). Les tâches sont exécutées séquentiellement, chaque tâche créant une conversation indépendante, avec prise en charge complète du suivi de l'état.
- **Contrôle des tâches** : Prise en charge de la pause/arrêt des tâches longues, de la modification des paramètres et de la ré-exécution, de l'obtention en streaming des journaux.
- **Gestion de la sécurité** : `/api/auth/change-password` permet de changer immédiatement le mot de passe ; il est recommandé de combiner avec une ACL au niveau du réseau lors de l'exposition du port MCP.

## Référence de configuration

```yaml
auth:
  password: "change-me"
  session_duration_hours: 12
server:
  host: "0.0.0.0"
  port: 8080
log:
  level: "info"
  output: "stdout"
mcp:
  enabled: true
  host: "0.0.0.0"
  port: 8081
openai:
  api_key: "sk-xxx"
  base_url: "https://api.deepseek.com/v1"
  model: "deepseek-chat"
database:
  path: "data/conversations.db"
  knowledge_db_path: "data/knowledge.db"  # Facultatif : base de données indépendante pour la base de connaissances
security:
  tools_dir: "tools"
knowledge:
  enabled: false  # Activer ou non la fonctionnalité de base de connaissances
  base_path: "knowledge_base"  # Chemin du répertoire de la base de connaissances
  embedding:
    provider: "openai"  # Fournisseur de modèle d'embedding (actuellement seul "openai" est pris en charge)
    model: "text-embedding-v4"  # Nom du modèle d'embedding
    base_url: ""  # Laisser vide pour utiliser le base_url de la configuration OpenAI
    api_key: ""  # Laisser vide pour utiliser l'api_key de la configuration OpenAI
  retrieval:
    top_k: 5  # Nombre de résultats Top-K retournés par la recherche
    similarity_threshold: 0.7  # Seuil de similarité (0-1), les résultats en dessous de ce seuil seront filtrés
    hybrid_weight: 0.7  # Poids de la recherche hybride (0-1), poids de la recherche vectorielle, 1.0 signifie recherche purement vectorielle, 0.0 signifie recherche purement par mots-clés
roles_dir: "roles"  # Répertoire des fichiers de configuration des rôles (relatif au répertoire du fichier de configuration)
skills_dir: "skills"  # Répertoire des Skills (relatif au répertoire du fichier de configuration)
```

### Exemple de modèle d'outil (`tools/nmap.yaml`)

```yaml
name: "nmap"
command: "nmap"
args: ["-sT", "-sV", "-sC"]
enabled: true
short_description: "Scan d'actifs réseau et identification d'empreintes de service"
parameters:
  - name: "target"
    type: "string"
    description: "IP ou nom de domaine"
    required: true
    position: 0
  - name: "ports"
    type: "string"
    flag: "-p"
    description: "Plage de ports, par exemple 1-1000"
```

### Exemple de configuration de rôle (`roles/test-penetration.yaml`)

```yaml
name: Test de pénétration
description: Expert en test de pénétration professionnel, détection complète et approfondie des vulnérabilités
user_prompt: Vous êtes un expert professionnel en test de pénétration de sécurité réseau. Veuillez utiliser des méthodes et outils de test de pénétration professionnels pour effectuer un test de sécurité complet sur la cible, incluant mais sans s'y limiter l'injection SQL, XSS, CSRF, inclusion de fichiers, exécution de commandes et autres vulnérabilités courantes.
icon: "\U0001F3AF"
tools:
  - nmap
  - sqlmap
  - nuclei
  - burpsuite
  - metasploit
  - httpx
  - record_vulnerability
  - list_knowledge_risk_types
  - search_knowledge_base
enabled: true
```

## Documents connexes

- [Guide d'utilisation du robot (DingTalk / Lark)](docs/robot.md) : Étapes de configuration complètes, commandes et instructions de dépannage pour dialoguer avec CyberStrikeAI depuis un téléphone mobile via DingTalk et Lark (Feishu), **il est recommandé de suivre ce document pour éviter les détours**.

## Structure du projet

```
CyberStrikeAI/
├── cmd/                 # Point d'entrée du service Web, MCP stdio et outils auxiliaires
├── internal/            # Agent, noyau MCP, routage et exécuteur
├── web/                 # Ressources statiques et modèles front-end
├── tools/               # Répertoire d'outils YAML (contient 100+ exemples)
├── roles/               # Répertoire des fichiers de configuration de rôles (contient 12+ rôles de test de sécurité prédéfinis)
├── skills/              # Répertoire des Skills (contient 20+ compétences de test de sécurité prédéfinies)
├── docs/                # Documentation (par exemple guide d'utilisation du robot)
├── images/              # Illustrations de la documentation
├── config.yaml          # Configuration d'exécution
├── run.sh               # Script de démarrage
└── README*.md
```

## Exemples d'utilisation de base

```
Scanner les ports ouverts de 192.168.1.1
Effectuer un scan ciblé de 192.168.1.1 sur les ports 80/443/22
Vérifier si https://example.com/page?id=1 présente une injection SQL
Énumérer les répertoires cachés et les vulnérabilités de composants de https://example.com
Obtenir les sous-domaines de example.com et exécuter nuclei en masse
```

## Exemples de scénarios avancés

```
Charger le scénario de reconnaissance : d'abord amass/subfinder, puis forçage de répertoires sur les hôtes actifs.
Monter un MCP externe basé sur Burp, effectuer une relecture du trafic authentifié et retransmettre à la chaîne d'attaque.
Compresser le rapport nuclei de 5 Mo et générer un résumé, l'attacher à l'enregistrement de la conversation.
Construire la chaîne d'attaque du dernier test, exporter uniquement la liste des nœuds avec risque >= élevé.
```

## Plan 404 Starlink
<img src="./images/404StarLinkLogo.png" width="30%">

CyberStrikeAI a maintenant rejoint le [Plan 404 Starlink](https://github.com/knownsec/404StarLink)

## Projet de pentest intelligent classé premier TCH
<div align="left">
  <a href="https://zc.tencent.com/competition/competitionHackathon?code=cha004" target="_blank">
    <img src="./images/tch.png" alt="Projet de pentest intelligent classé premier TCH" width="30%">
  </a>
</div>

## Stargazers au fil du temps
![Stargazers au fil du temps](https://starchart.cc/Ed1s0nZ/CyberStrikeAI.svg)

---

## ⚠️ Clause de non-responsabilité

**Cet outil est uniquement destiné à l'éducation et aux tests autorisés !**

CyberStrikeAI est une plateforme de test de sécurité professionnelle, conçue pour aider les chercheurs en sécurité, les testeurs de pénétration et les professionnels de l'informatique à effectuer des évaluations de sécurité et des recherches sur les vulnérabilités **dans le cadre d'une autorisation explicite**.

**En utilisant cet outil, vous acceptez :**
- D'utiliser cet outil uniquement sur des systèmes pour lesquels vous disposez d'une autorisation écrite explicite
- De respecter toutes les lois, réglementations et normes éthiques applicables
- D'assumer l'entière responsabilité de toute utilisation non autorisée ou abusive
- De ne pas utiliser cet outil à des fins illégales ou malveillantes

**Les développeurs ne sont pas responsables de toute utilisation abusive !** Veuillez vous assurer que votre utilisation est conforme aux lois et réglementations locales et que vous avez obtenu l'autorisation explicite du propriétaire du système cible.

---

N'hésitez pas à soumettre des Issue/PR pour contribuer de nouveaux modèles d'outils ou des suggestions d'amélioration !

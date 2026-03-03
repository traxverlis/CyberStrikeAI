# Guide d'utilisation du robot CyberStrikeAI

[English](robot_en.md) | [中文](robot.md) | [Français](robot_fr.md)

Ce document explique comment dialoguer avec CyberStrikeAI via **DingTalk** ou **Feishu (Lark)** (mode connexion longue), utilisable depuis un téléphone mobile, sans avoir besoin d'ouvrir la page web sur le serveur. Suivez les étapes ci-dessous pour éviter les erreurs courantes.

---

## 1. Où configurer dans CyberStrikeAI

1. Connectez-vous à l'interface Web de CyberStrikeAI
2. Dans la navigation de gauche, allez dans **Paramètres système**
3. Dans les catégories de paramètres à gauche, cliquez sur **Paramètres du robot** (situé entre « Paramètres de base » et « Paramètres de sécurité »)
4. Cochez la case selon la plateforme et remplissez les champs (pour DingTalk : Client ID / Client Secret, pour Feishu : App ID / App Secret)
5. Cliquez sur **Appliquer la configuration** pour enregistrer
6. **Redémarrez l'application CyberStrikeAI** (si vous enregistrez seulement sans redémarrer, le robot ne se connectera pas)

La configuration sera écrite dans la section `robots` du fichier `config.yaml`, elle peut également être éditée directement dans le fichier de configuration. **Après modification de la configuration DingTalk/Feishu, vous devez impérativement redémarrer pour que la connexion longue prenne effet.**

---

## 2. Plateformes supportées (connexion longue)

| Plateforme | Description |
|------|------|
| DingTalk | Utilise une connexion Stream longue, le programme se connecte activement à DingTalk pour recevoir les messages |
| Feishu (Lark) | Utilise une connexion longue, le programme se connecte activement à Feishu pour recevoir les messages |

La section trois ci-dessous détaillera pour chaque plateforme : ce qu'il faut faire sur la plateforme ouverte, quels champs copier, où les coller dans CyberStrikeAI.

---

## 3. Éléments de configuration et étapes détaillées par plateforme

### 3.1 DingTalk

**Clarification importante : deux types de robots DingTalk différents**

| Type | Où le créer | Peut-il faire « utilisateur envoie message → robot répond » | Ce programme le supporte-t-il |
|------|------------|----------------------------------|----------------|
| **Robot personnalisé** | Dans un groupe DingTalk : paramètres du groupe → Ajouter un robot → Personnalisé (Webhook) | ❌ Non, peut seulement envoyer des messages au groupe | ❌ Non supporté |
| **Robot d'application interne d'entreprise** | Créer une application sur la [plateforme ouverte DingTalk](https://open.dingtalk.com) et activer le robot | ✅ Oui | ✅ Supporté |

Si vous avez une adresse Webhook de « robot personnalisé » (`oapi.dingtalk.com/robot/send?access_token=xxx`) et une clé de signature (`SEC...`), **vous ne pouvez pas les utiliser directement dans ce programme**. Vous devez suivre les étapes ci-dessous pour créer une « application interne d'entreprise » sur la plateforme ouverte et obtenir le **Client ID** et le **Client Secret**.

---

**Étapes complètes de configuration DingTalk (dans l'ordre)**

1. **Ouvrir la plateforme ouverte DingTalk**
   Accédez à [https://open.dingtalk.com](https://open.dingtalk.com) dans votre navigateur, connectez-vous avec un compte **administrateur d'entreprise**.

2. **Entrer dans le développement d'applications**
   À gauche, sélectionnez **Développement d'applications** → **Développement interne d'entreprise** → Cliquez sur **Créer une application** (ou sélectionnez une application existante). Remplissez les informations de base comme le nom de l'application et créez-la.

3. **Obtenir le Client ID et le Client Secret**
   - À gauche, cliquez sur **Informations d'identification et informations de base** (sous « Informations de base »).
   - Sur la page, vous trouverez le **Client ID (anciennement AppKey)** et le **Client Secret (anciennement AppSecret)**.
   - Cliquez pour copier, **ne tapez pas manuellement**, attention : le chiffre **0** et la lettre **o**, le chiffre **1** et la lettre **l** sont faciles à confondre (par exemple, dans `ding9gf9tiozuc504aer`, au milieu c'est le nombre **504** et non 5o4).

4. **Activer le robot et sélectionner le mode Stream**
   - À gauche, **Capacités de l'application** → **Robot**.
   - Activez l'interrupteur « Configuration du robot ».
   - Remplissez le nom du robot, la description, etc. (remplir les champs obligatoires selon les indications).
   - **Important** : pour la méthode de réception des messages, sélectionnez **« Mode Stream »** (accès en streaming). Si seul « Rappel HTTP » est disponible ou si Stream n'est pas sélectionné, ce programme ne recevra pas les messages.
   - Enregistrez.

5. **Permissions et publication**
   - À gauche, **Gestion des permissions** : recherchez « robot », « message », etc., cochez les permissions liées au robot comme **recevoir des messages**, **envoyer des messages**, et confirmez l'autorisation.
   - À gauche, **Gestion et publication des versions** : si des configurations ne sont pas publiées, cliquez sur **Publier une nouvelle version** / **Mettre en ligne**, sinon les modifications ne prendront pas effet.

6. **Remplir dans CyberStrikeAI**
   - Retournez dans CyberStrikeAI → Paramètres système → Paramètres du robot → DingTalk.
   - Cochez « Activer le robot DingTalk ».
   - **Client ID (AppKey)** collez le Client ID copié à l'étape 3.
   - **Client Secret** collez le Client Secret copié à l'étape 3.
   - Cliquez sur **Appliquer la configuration**, puis **redémarrez CyberStrikeAI**.

---

**Correspondance des champs DingTalk dans CyberStrikeAI**

| Champ à remplir dans CyberStrikeAI | Source sur la plateforme ouverte DingTalk |
|------------------------|------------------------|
| Activer le robot DingTalk | Cocher pour activer |
| Client ID (AppKey) | Informations d'identification et informations de base → **Client ID (anciennement AppKey)** |
| Client Secret | Informations d'identification et informations de base → **Client Secret (anciennement AppSecret)** |

---

### 3.2 Feishu (Lark)

| Élément de configuration | Description |
|--------|------|
| Activer le robot Feishu | Cocher pour démarrer la connexion longue Feishu |
| App ID | App ID dans les informations d'identification de l'application sur la plateforme ouverte Feishu |
| App Secret | App Secret dans les informations d'identification de l'application sur la plateforme ouverte Feishu |
| Verify Token | Pour l'abonnement aux événements (facultatif) |

**Étapes de configuration Feishu en bref** : Connectez-vous à la [plateforme ouverte Feishu](https://open.feishu.cn) → Créez une application auto-construite d'entreprise → Dans « Informations d'identification et informations de base », obtenez **App ID**, **App Secret** → Dans « Capacités de l'application », activez le **Robot** et activez les permissions correspondantes → Publiez l'application → Remplissez l'App ID et l'App Secret dans les paramètres du robot de CyberStrikeAI → Enregistrez et **redémarrez l'application**.

---

## 4. Commandes du robot

Envoyez les **commandes textuelles** suivantes au robot dans DingTalk/Feishu (seul le texte est supporté) :

| Commande | Description |
|------|------|
| **aide** | Affiche l'aide et les explications des commandes |
| **liste** ou **liste de conversations** | Liste toutes les conversations avec leur titre et ID |
| **basculer \<ID_conversation\>** ou **continuer \<ID_conversation\>** | Spécifie un ID de conversation, les messages suivants continueront dans cette conversation |
| **nouvelle conversation** | Démarre une nouvelle conversation, les messages suivants seront dans la nouvelle conversation |
| **effacer** | Efface le contexte de la conversation actuelle (équivalent à « nouvelle conversation ») |
| **actuel** | Affiche l'ID et le titre de la conversation actuelle |
| **arrêter** | Interrompt la tâche en cours d'exécution |
| **rôle** ou **liste des rôles** | Liste tous les rôles disponibles (test de pénétration, CTF, scan d'applications Web, etc.) |
| **rôle \<nom_rôle\>** ou **basculer rôle \<nom_rôle\>** | Change le rôle actuellement utilisé |
| **supprimer \<ID_conversation\>** | Supprime la conversation spécifiée |
| **version** | Affiche le numéro de version actuel de CyberStrikeAI |

En dehors de ces commandes, **tout texte saisi directement** sera envoyé à l'IA comme message utilisateur, avec la même logique de conversation que sur l'interface Web (test de pénétration/analyse de sécurité, etc.).

---

## 5. Comment utiliser (faut-il mentionner le robot avec @?)

- **Conversation privée (recommandé)** : Dans DingTalk/Feishu, **recherchez et ouvrez le robot**, entrez dans le **chat privé** avec le robot, tapez directement « aide » ou n'importe quel texte, **pas besoin de @**.
- **Conversation de groupe** : Si le robot est ajouté à un groupe, dans le groupe, seuls les messages envoyés après **@ le robot** seront reçus et traités par le robot ; les messages du groupe sans @ ne déclencheront pas le robot.

En résumé : en **chat privé avec le robot, envoyez directement** ; dans un **groupe, vous devez @le robot** avant d'envoyer le contenu.

---

## 6. Flux d'utilisation recommandé (pour éviter d'oublier des étapes)

1. **Sur la plateforme ouverte** : Complétez selon la section trois la création de l'application DingTalk ou Feishu, la copie des informations d'identification, l'activation du robot (pour DingTalk, sélectionnez impérativement le **mode Stream**), les permissions et la publication.
2. **Dans CyberStrikeAI** : Paramètres système → Paramètres du robot → Cochez la plateforme correspondante, collez le Client ID/App ID, Client Secret/App Secret → Cliquez sur **Appliquer la configuration**.
3. **Redémarrez le processus CyberStrikeAI** (sinon la connexion longue ne sera pas établie).
4. **Sur votre téléphone DingTalk/Feishu** : Trouvez le robot (en chat privé envoyez directement, en groupe vous devez @le robot), envoyez « aide » ou n'importe quel contenu pour tester.

Si l'envoi de messages ne donne aucune réponse, consultez d'abord la **section neuf sur le dépannage** et la **section dix sur les erreurs courantes**.

---

## 7. Exemple de fichier de configuration

Extrait du fichier `config.yaml` concernant le robot :

```yaml
robots:
  dingtalk:
    enabled: true
    client_id: "your_dingtalk_app_key"
    client_secret: "your_dingtalk_app_secret"
  lark:
    enabled: true
    app_id: "your_lark_app_id"
    app_secret: "your_lark_app_secret"
    verify_token: ""
```

Après modification, vous devez **redémarrer l'application**, la connexion longue s'établit au démarrage de l'application.

---

## 8. Comment vérifier si c'est fonctionnel (sans client DingTalk/Feishu)

Lorsque DingTalk ou Feishu n'est pas installé, vous pouvez utiliser l'**interface de test** pour vérifier si la logique du robot fonctionne normalement :

1. Connectez-vous d'abord à l'interface Web de CyberStrikeAI (assurez-vous d'être connecté).
2. Utilisez curl pour appeler l'interface de test (vous devez fournir le Cookie après connexion) :

```bash
# Remplacez YOUR_COOKIE par le Cookie obtenu après connexion (navigateur F12 → Réseau → n'importe quelle requête → en-têtes de requête → Cookie)
curl -X POST "http://localhost:8080/api/robot/test" \
  -H "Content-Type: application/json" \
  -H "Cookie: YOUR_COOKIE" \
  -d '{"platform":"dingtalk","user_id":"test_user","text":"aide"}'
```

Si le JSON retourné contient `"reply":"【Commandes du robot CyberStrikeAI】..."`  cela signifie que le traitement des commandes fonctionne normalement. Vous pouvez également essayer `"text":"liste"`, `"text":"actuel"`, etc.

Description de l'interface : `POST /api/robot/test` (connexion requise), corps de la requête `{"platform":"facultatif","user_id":"facultatif","text":"obligatoire"}`, réponse `{"reply":"contenu de la réponse"}`.

---

## 9. Dépannage quand DingTalk ne répond pas aux messages

Vérifiez dans l'ordre :

0. **Après mise en veille d'un ordinateur portable / perte de connexion réseau**
   DingTalk et Feishu utilisent tous deux des connexions longues pour recevoir les messages. Après mise en veille ou perte de réseau, la connexion sera interrompue. Le programme se **reconnectera automatiquement** (retry dans environ 5 secondes à 60 secondes). Après réveil ou restauration du réseau, attendez un moment avant d'envoyer un message ; si toujours aucune réponse, vous pouvez redémarrer le processus CyberStrikeAI.

1. **Le Client ID / Client Secret correspondent-ils exactement à ceux de la plateforme ouverte**
   **Copiez-collez** depuis « Informations d'identification et informations de base », ne tapez pas manuellement. Attention au chiffre **0** et à la lettre **o**, au chiffre **1** et à la lettre **l** (par exemple, dans `ding9gf9tiozuc504aer`, au milieu c'est **504** et non 5o4).

2. **Avez-vous redémarré l'application après avoir enregistré la configuration**
   La connexion longue du robot s'établit **au démarrage de l'application**. Cliquer sur « Appliquer la configuration » dans l'interface Web écrit seulement le fichier de configuration, **vous devez redémarrer le processus CyberStrikeAI** pour que la connexion DingTalk prenne effet.

3. **Consultez les journaux du programme**
   - Après le démarrage, vous devriez voir : `DingTalk Stream se connecte...`, `DingTalk Stream démarré (pas besoin de réseau public), en attente de messages`.
   - Si vous voyez `DingTalk Stream connexion longue terminée` avec un message d'erreur, c'est généralement dû à un **Client ID / Client Secret incorrect** ou à **l'accès en streaming non activé sur la plateforme ouverte**.
   - Après avoir envoyé un message dans DingTalk, s'il est reçu, vous devriez avoir un journal : `Message reçu de DingTalk` ; sinon, cela signifie que DingTalk n'a pas envoyé le message à ce programme (vérifiez si le « Robot » est activé sur la plateforme ouverte, si le **mode Stream** est sélectionné).

4. **Du côté de la plateforme ouverte**
   L'application doit être **publiée** ; dans la capacité « Robot », vous devez activer **l'accès en streaming (Stream)** pour recevoir les messages (le rappel HTTP seul ne suffit pas) ; dans la gestion des permissions, vous devez avoir les permissions de réception et d'envoi de messages du robot, etc.

---

## 10. Erreurs courantes (éviter les pièges)

- **Mauvais type de robot utilisé** : Le robot « personnalisé » ajouté dans un **groupe** DingTalk (Webhook + signature) **ne peut pas** être utilisé pour des conversations, ce programme ne supporte que les robots de la **« application interne d'entreprise » sur la plateforme ouverte**.
- **Seulement enregistré sans redémarrer** : Après avoir modifié la configuration du robot dans CyberStrikeAI, vous devez **redémarrer l'application**, sinon la connexion longue ne sera pas établie.
- **Client ID mal copié** : Si la plateforme ouverte indique `504`, remplissez `504`, ne remplissez pas `5o4` ; utilisez de préférence copier-coller.
- **DingTalk a seulement activé le rappel HTTP sans Stream** : Ce programme reçoit les messages via **connexion longue Stream**, la méthode de réception des messages du robot sur la plateforme ouverte doit sélectionner le **mode Stream**.
- **Application non publiée** : Après avoir modifié le robot ou les permissions sur la plateforme ouverte, vous devez **publier une nouvelle version** dans « Gestion et publication des versions », sinon cela ne prendra pas effet.

---

## 11. Notes importantes

- DingTalk et Feishu **traitent uniquement les messages textuels** ; les autres types (images, audio) afficheront un message indiquant qu'ils ne sont pas supportés ou seront ignorés.
- Les sessions partagent les mêmes données de conversation que l'interface Web : les conversations créées dans le robot seront visibles dans la liste « Conversations » de l'interface Web, et vice versa.
- La logique d'exécution du robot est identique à **`/api/agent-loop/stream`** (inclut les rappels de progression, les détails du processus écrits dans la base de données), sauf qu'il ne pousse pas de SSE au client, et envoie finalement la réponse complète en une fois à DingTalk/Feishu/WeChat Entreprise.

# MLOps Platform (Front Vue + Backend Go)

Ce projet simule un **MLOps Control Center** :
- un **backend Go (Gin)** expose une API REST pour gérer des *jobs* de training,
- un **frontend Vue 3 (Vite)** permet de créer, afficher, filtrer, redémarrer et supprimer ces jobs.

Les jobs sont stockés localement dans `backend/data/jobs.json`.

---

## Prérequis

- **Go** (ex: Go 1.26.3)
- **Node.js** (recommandé conforme à `frontend/package.json`, ex Node >= 20)

---

## Backend (Go)

### Lancer le backend

1. Ouvrir un terminal dans `backend/`
2. Installer les dépendances (si besoin) :

```bash
go mod tidy
```

3. Démarrer le serveur :

```bash
go run main.go
```

Le backend écoute sur :
- `http://localhost:8080`

### API exposée

Base URL : `http://localhost:8080`

#### `GET /jobs`
Retourne la liste des jobs.

#### `POST /jobs`
Crée un nouveau job.
Corps JSON :

```json
{ "name": "training-job" }
```

Lors de la création, le job passe en `running` et une tâche “simulation” met à jour :
- `accuracy` (progression 0 -> 100 par pas de 10)
- `status` passe à `completed` quand `accuracy == 100`

#### `PATCH /jobs/:id`
Met à jour partiellement un job.
Corps JSON (exemple) :

```json
{ "status": "running", "accuracy": 0 }
```

Le backend met à jour seulement les champs non vides / non nuls.

#### `DELETE /jobs/:id`
Supprime un job.

---

## Frontend (Vue 3)

Le frontend appelle l’API du backend via Axios.
- URL de base : `http://localhost:8080` (voir `frontend/src/services/api.js`)

### Lancer le frontend

1. Ouvrir un terminal dans `frontend/`
2. Installer les dépendances :

```bash
npm install
```

3. Démarrer le serveur de dev :

```bash
npm run dev
```

Le frontend sera accessible (généralement) sur :
- `http://localhost:5173`

---

## Utilisation (workflow)

1. Démarrer d’abord le backend Go.
2. Démarrer ensuite le frontend Vue.
3. Dans l’UI :
   - saisir un nom et cliquer sur **Create Job**
   - suivre la progression : `status` + `accuracy`
   - cliquer **Restart** pour repasser le job à `running`
   - cliquer **Delete** pour supprimer le job

> Le frontend rafraîchit automatiquement la liste des jobs toutes les ~2 secondes.

---

## Fichiers importants

- `backend/main.go` : implémentation de l’API REST + simulation de training
- `backend/data/jobs.json` : stockage local des jobs
- `frontend/src/services/api.js` : configuration Axios (baseURL)
- `frontend/src/stores/jobs.js` : store Pinia (fetch/create/delete/update)
- `frontend/src/App.vue` : interface principale

---

## Commandes rapides (copier/coller)

### Backend

```bash
cd backend
go run main.go
```

### Frontend

```bash
cd frontend
npm install
npm run dev
```


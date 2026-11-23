# Forum-Project

Forum web inspiré de Reddit, développé en Go en équipe dans le cadre de la formation B1 Informatique d'Ynov.

> **La version finale du projet se trouve dans [pomaxime/Forum-Project-complete](https://github.com/pomaxime/Forum-Project-complete).**
> Ce dépôt-ci garde le travail en cours de route : la branche `main` ne contient que ce README, et le code est réparti sur les branches de développement.

## Le projet

- Inscription, connexion et sessions par cookie, mots de passe hachés avec bcrypt
- Publications classées par catégories, avec filtres
- Commentaires sous chaque publication
- Likes et dislikes sur les publications et les commentaires
- Page de profil avec les publications de l'utilisateur
- Pages rendues côté serveur avec les templates Go, CSS responsive
- Base de données SQLite

**Technologies :** Go (`net/http`, `html/template`), SQLite (`mattn/go-sqlite3`), bcrypt (`golang.org/x/crypto`), HTML, CSS, JavaScript.

## Branches

| Branche | Contenu |
|---|---|
| `auth` | La plus avancée : inscription, connexion, validation des formulaires, templates et styles |
| `dev` | Intégration : structure complète du projet et templates |
| `templates`, `static` | Pages HTML (accueil, publication, profil, catégories…) et feuilles de style |
| `server` | Premier serveur HTTP |
| `arbo`, `database`, `models`, `utils` | Mise en place de l'arborescence |

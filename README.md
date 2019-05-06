# Justify

API Rest qui justifie un texte sur l'endpoint api/justify (POST) en ne dépassant pas les 80 caractères.
Si le texte dépasse 80 000 mots par jour pour un utilisateur, un message d'erreur s'affiche

Systeme d'authentification par token unique avec un email préalablement enregistré sur l'endpoint api/token (POST)

Pour compiler le projet : go build -o bin/justify src/*.go


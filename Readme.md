
# Téléchargeur de Fichiers CLI en Golang

Un téléchargeur de fichiers en ligne de commande écrit en Golang avec une barre de progression.

## Fonctionnalités

- Téléchargement de fichiers depuis une URL spécifiée.
- Affichage de la progression du téléchargement.
- Sauvegarde du fichier téléchargé à l'emplacement spécifié.

## Utilisation

Compilez le programme et exécutez-le avec les arguments de ligne de commande `-url` et `-output` :

```sh
go run main.go -url <file_url> -output <output_path>
```

Par exemple :

```sh
go run main.go -url https://example.com/file.zip -output file.zip
```

## Exemple de sortie

```
[==================================================] 100.00%
Download completed!
```

## Image

![Golang CLI](https://blog.golang.org/go-brand/Go-Logo/PNG/Go-Logo_Blue.png)

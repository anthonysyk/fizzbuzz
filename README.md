# fizzbuzz

## Contexte
Ce repository contient une implementation d'un serveur HTTP REST en exposant la réponse au jeu de mot FizzBuzz. 

Le service exposé permet de jouer à FizzBuzz en envoyant une requête HTTP avec plusieurs paramètres. 

```go
type FizzBuzzQuery struct {
	int1  int    `form:"int1"`
	int2  int    `form:"int2"`
	limit int    `form:"limit"`
	str1  string `form:"str1"`
	str2  string `form:"str2"`
}
```

Le serveur retournera la chaîne de caractères "Fizz" si le nombre est divisible par 3, "Buzz" si le nombre est divisible par 5 et "FizzBuzz" si le nombre est divisible à la fois par 3 et par 5. 

Si le nombre n'est divisible ni par 3 ni par 5, alors le serveur retournera simplement le nombre lui-même en chaîne de caractères. 

## Déploiement

### Lancer le serveur en local

```make run```

### Lancer le serveur avec Docker

```
make docker
```

### Lancer les tests

```make test```

```
?       github.com/anthonysyk/fizzbuzz  [no test files]
ok      github.com/anthonysyk/fizzbuzz/api      0.180s
ok      github.com/anthonysyk/fizzbuzz/middleware       0.116s
?       github.com/anthonysyk/fizzbuzz/model    [no test files]
?       github.com/anthonysyk/fizzbuzz/store    [no test files]
ok      github.com/anthonysyk/fizzbuzz/test/integration 0.234s
ok      github.com/anthonysyk/fizzbuzz/test/load        2.234s
```

### Build et lancer en executable

```
make build
./fizzbuzz
```

## Routes

- `/health` retourne le statut du serveur
- `/fizzbuzz` retourne le résultat du service fizzbuzz
- `/stats` retourne des données aggrégées sur l'utilisation du endpoint `/fizzbuzz` comme la requête la plus fréquence

## Structure
```
.
├── api
├── middleware
├── model
├── store
└── test
    ├── integration
    └── load
```

- `api` contient les fonctionnalités de l'API
- `middleware` contient les middlewares et permet de séparer le code métier et technique
- `model` contient les data structures
- `store` contient une interface ainsi que les implémentations liées au stockage
- `test` contient les tests d'intégration et les tests de charge

## Pistes d'amélioration

- Ajouter une infra pour le stockage du nombre de hits par requêtes
  - redis ou memorystore (GCP)
- Ajouter des variables d'environnement
- Découper `api` si nécessaire (ajout de nouvelles features)

```
.
├── api
│   ├── handler
│   └── service
```

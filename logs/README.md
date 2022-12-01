## This exercice aims to make me practice Runes concept 
- Exercise link : https://exercism.org/tracks/go/exercises/logs-logs-logs

### Theory 
- "Rune" est un alias pour le type int32
- Cependant, l'integer stocké dans une rune represente le code unique d'un caractère 
- Ex : 
    - Unicode Character : 0 
    - Unicode Code Point : U+0030
    - Decimal Value : 48

- le code source des fichiers en Go sont encodés en UTF-8 
- pour déclarer une rune, on place un caractère dans un single quote : 
` myRune := '&' `
=> print le type renverra `int32`
=> print la valeur renverra sa valeur décimale et non un string 
=> il faut utilise %c pour afficher le Unicode character
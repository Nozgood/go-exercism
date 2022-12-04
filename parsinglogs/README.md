## This exercice aims to make me practice Regexp concept 

- Exercise link : https://exercism.org/tracks/go/exercises/parsing-log-files

### Theory

- pour gérer l'utilisation de regexp en go, il existe une librairie (package) standard:  `regexp`
- La syntaxe des regexp est la même que pour les autres langages (comme j'ai pu pratiquer en JS)
- Le pattern de recherche comme l'input sont interprétés en UTF-8
- quand j'utlise les "``" (accent grave, backticks en anglais), les \ ne marquent pas le début de caractères spéciaux comme \n par exemple (passage à la ligne)
- grâce à ça, il est préférable d'utiliser les `` pour faire une regexp
- pour pouvoir manipuler un string en regexp il faut d'abord le formater en utilisant la method compile de l'objet regexp => renvoie la rep et une erreur (nil si pas)
- => mustCompile permet de ne pas avoir à gérer l'erreur 
- Warning : utiliser mustCompile seulement quand je suis sûr à 100% que la gestion d'erreur ne sera pas nécéssaire 
- 
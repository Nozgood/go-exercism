## This exercise aims to make me practice interfaces' concept
- The exercise is about an election system, here is the link : https://exercism.org/tracks/go/exercises/the-farm

### Theory 
- regroupe toutes les fonctions qu'une variable qui posède le type donnée, doit avoir 
- S'écrit de cette façon : 
- `type InterfaceName interface { MethodOne () MethodOneReturnType}`
- Ex : 
    - type Animal Interface {
        Noise() string
        NumberOflegs() string
    }
- Ainsi, lorsque je déclare une struct, et que je souhaite utiliser une méthode de l'interface choisie sur un objet de ce struct, je dois au préalable implémenter la méthode en passant le struct en receveur (pointé la plupart du temps par convention => rapidité d'éxécution du code puisque pas de copie ?)

- Je vois un peu l'interface comme le struct des méthodes : il prépare et distribue les fonctions qui pouront être utilisés par les struct
- Un exemple est dispo dans le fichier "farm_example"


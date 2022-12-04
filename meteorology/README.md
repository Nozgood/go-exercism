## This exercise aims to make me practice Stringer concept 
- Exercise link : https://exercism.org/tracks/go/exercises/meteorology

### Theory 
- Stringer est une interface pour définir la "string format" des valeurs auxquels elle est appliquée : 
    type Stringer interface {
        String  string
    }
- les types qui veulent implémenter ce genre d'interface doivent obligatoirement avoir une methode `String()` qui retourne une valeur lisible pour l'Homme
- les packages comme `fmt` se servent de cette methode pour afficher le resultat attendu 
- un exemple tres parlant du concept de stringer est consultable ici : https://exercism.org/tracks/go/concepts/stringers

## This exercise aims to make me practice pointer's concept
- The exercise is about an election system, here is the link : https://exercism.org/tracks/go/exercises/elons-toy

- une methode est une fonction que l'on applique sur un élement donné (le receveur)
- le receveur est indiqué entre le nom de la méthode est le mot clé "func"
- par exemple si j'ai un struct "Person", je peux écrire une méthode qui concerne ce struct comme ceci : 
- `func (p Person) Greetings() string {}`
- Ensuite, je peux librement utiliser cette méthode sur toutes les variables qui sont du type Person
- avec la même notation (dot) que lorsque je veux récupérer un élement du struct 
- 2 types de receveurs: de type valeur et de type pointeur 
- type valeur : les modifs apportés dans la méthode ne sont pas visible par la valeur appelante
- type pointeur : les modifs apportés sont directement appliqués à l'appelant 




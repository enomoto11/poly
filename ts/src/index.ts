import { newAnimal } from "./domain/animal";

function main() {
  // Create a new dog
  const dog = newAnimal("dog", "koro");
  console.log(dog.speak());

  // Create a loyal dog
  const loyalDog = newAnimal("dog", "pochi");
  console.log(loyalDog.speak());

  // Create a new cat
  const cat = newAnimal("cat", "tama");
  console.log(cat.speak());
}

main();

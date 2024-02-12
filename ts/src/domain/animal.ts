import { Cat } from "./cat";
import { Dog } from "./dog";

export interface Animal {
  speak: () => string;
}

// Factory method
// Whether retuning Dog or Cat, the type of the return value is Animal
// Because both Dog and Cat implement the Animal interface
export function newAnimal(kind: string, name: string): Animal {
  switch (kind) {
    case "dog":
      return new Dog(name);
    case "cat":
      return new Cat(name);
    default:
      throw new Error("unknown animal");
  }
}

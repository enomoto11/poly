import { UUID, randomUUID } from "crypto";

export class Dog {
  private id: UUID;
  private name: string;

  constructor(name: string) {
    this.id = randomUUID(); // We assign uuid to the id field so that the data store does not need to generate it
    this.name = name;

    this.validate();
  }

  private validate(): void {
    if (this.name === "") {
      throw new Error("name is required");
    }

    if (this.name.length > 50) {
      throw new Error("name must be less than 50 characters");
    }
  }

  // Implicitly implements the Animal interface
  public speak(): string {
    if (this.name === "pochi") {
      return "I'm loyal! I'm " + this.name + "!";
    }

    return "Woof!";
  }

  // Even if this class has a method that is not defined in the Animal interface, it is not a problem
  // Because the type cast is done in the factory method
  public getID(): UUID {
    return this.id;
  }

  // If the return type is different from the Animal interface, it is a problem
  // And the client cannot assume that Dog implements the Animal interface
  // public speak(): number {
  //   return 1;
  // }
}

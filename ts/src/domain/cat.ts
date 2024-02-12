import { UUID, randomUUID } from "crypto";

export class Cat {
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
    return "Meow!";
  }
}

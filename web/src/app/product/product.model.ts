export class Product {
    constructor(
        private id: string,
        public name
    ) { }

    public getId(): string {
        return this.id;
    }
}
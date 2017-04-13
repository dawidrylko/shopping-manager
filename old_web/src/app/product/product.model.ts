import { Price } from './price.model'

export class Product {
    constructor(
        private id: string,

        public name: string,
        public price: Price[],

        public manufacturer?: string,
        public ean?: string,
        public energy?: string,
        public netWeight?: string,
        public unitOfMeasure?: string
    ) { }

    public getId(): string {
        return this.id;
    }
}
import { Injectable } from '@angular/core';

import { Product } from './product.model';

@Injectable()
export class ProductService {

  constructor() { }

  getProducts(): Promise<Product> {

  }

}

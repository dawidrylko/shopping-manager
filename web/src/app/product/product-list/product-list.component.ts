import { Component, OnInit } from '@angular/core';

import { ProductService } from './../product.service'
import { Product } from './../product.model'

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.scss']
})
export class ProductListComponent implements OnInit {

  public products: any;

  constructor(
    private productService: ProductService
  ) { }

  private getProducts(): any {
    return this.productService.list()
      .subscribe((products) => {
        return this.products = products;
      });
  }

  ngOnInit() {
    this.getProducts();
  }

}

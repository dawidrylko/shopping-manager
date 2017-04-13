import { Component, OnInit } from '@angular/core';

import { ProductService } from './../product.service'
import { Product } from './../product.model'

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.scss']
})
export class ProductListComponent implements OnInit {

  public products: Product[];

  constructor(
    private productService: ProductService
  ) { }

  private getProducts(): void {
    this.productService.list()
    .subscribe((products: Product[]) => {
        this.products = products;
    });
  }

  ngOnInit() {
    this.getProducts();
  }

}

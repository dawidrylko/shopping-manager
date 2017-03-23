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
  ) { 
    this.products = this.productService.list();

    console.log(this.products);
  }

  ngOnInit() {
  }

}

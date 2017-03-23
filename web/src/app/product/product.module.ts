import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProductRoutes } from './product.routing';
import { ProductService } from './product.service';

import { ProductListComponent } from './product-list/product-list.component';


@NgModule({
  imports: [
    CommonModule,
    ProductRoutes
  ],
  providers: [ProductService],
  declarations: [ProductListComponent]
})
export class ProductModule { }

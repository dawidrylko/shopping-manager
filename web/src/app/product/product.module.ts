import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProductRoutes } from './product.routing'

import { ProductListComponent } from './product-list/product-list.component';


@NgModule({
  imports: [
    CommonModule,
    ProductRoutes
  ],
  declarations: [ProductListComponent]
})
export class ProductModule { }

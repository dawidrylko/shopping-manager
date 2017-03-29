import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProductRoutes } from './product.routing';
import { ProductService } from './product.service';

import { ProductListComponent } from './product-list/product-list.component';
import { ProductFormComponent } from './product-form/product-form.component';


@NgModule({
  imports: [
    CommonModule,
    ProductRoutes
  ],
  providers: [ProductService],
  declarations: [ProductListComponent, ProductFormComponent]
})
export class ProductModule { }

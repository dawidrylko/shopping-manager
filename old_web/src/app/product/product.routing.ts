import { RouterModule, Routes } from '@angular/router'

import { ProductListComponent } from './product-list/product-list.component'
import { ProductFormComponent } from './product-form/product-form.component'

const routes: Routes = [
  { path: 'product',

    children: [
      { path: 'list',
        component: ProductListComponent,
        data: { title: 'Products list' }
      },
      { path: 'form',
        component: ProductFormComponent,
        data: { title: 'Product form' }
      }
    ]
  }
];

export const ProductRoutes = RouterModule.forChild(routes);

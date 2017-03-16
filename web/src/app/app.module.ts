import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { RouterModule, Routes } from '@angular/router';

import { AppComponent } from './app.component';

import { ProductComponent } from './product/product.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { DashboardComponent } from './dashboard/dashboard.component';

const appRoutes: Routes = [
  { path: '',
    redirectTo: '/dashboard',
    pathMatch: 'full'
  },
  { path: 'dashboard',
    component: DashboardComponent,
    data: { title: 'Dashboard' }
  },
  { path: 'product-list',
    component: DashboardComponent,
    data: { title: 'Product list' }
  },
  { path: 'product/:id', 
    component: ProductComponent,
    data: { title: 'Product' }
  },
  { path: '**', 
    component: PageNotFoundComponent 
  }
];

@NgModule({
  declarations: [
    AppComponent,
    ProductComponent,
    PageNotFoundComponent,
    DashboardComponent
  ],
  imports: [
    RouterModule.forRoot(appRoutes),
    BrowserModule,
    FormsModule,
    HttpModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

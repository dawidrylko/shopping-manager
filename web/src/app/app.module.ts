import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { RouterModule, Routes } from '@angular/router';
import { MaterialModule } from '@angular/material';

import { AppComponent } from './app.component';

import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { SidebarComponent } from './sidebar/sidebar.component';

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
    data: { title: 'Product list' }
  },
  { path: 'product/:id', 
    data: { title: 'Product' }
  },
  { path: '**', 
    component: PageNotFoundComponent 
  }
];

@NgModule({
  declarations: [
    AppComponent,
    PageNotFoundComponent,
    DashboardComponent,
    SidebarComponent
  ],
  imports: [
    RouterModule.forRoot(appRoutes),
    BrowserModule,
    FormsModule,
    HttpModule,
    MaterialModule.forRoot()
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

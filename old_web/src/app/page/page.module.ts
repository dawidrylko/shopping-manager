import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DashboardComponent } from './dashboard/dashboard.component'
import { PageNotFoundComponent } from './page-not-found/page-not-found.component'

@NgModule({
  imports: [
    CommonModule
  ],
  declarations: [
    DashboardComponent,
    PageNotFoundComponent
  ]
})
export class PageModule { }

import { RouterModule, Routes } from '@angular/router';

import { PageNotFoundComponent } from './page/page-not-found/page-not-found.component';
import { DashboardComponent } from './page/dashboard/dashboard.component';

const routes: Routes = [
  { path: '',
    redirectTo: '/dashboard',
    pathMatch: 'full'
  },
  { path: 'dashboard',
    component: DashboardComponent
  },
  { path: '**',
    component: PageNotFoundComponent
  }
];

export const AppRoutes = RouterModule.forRoot(routes);

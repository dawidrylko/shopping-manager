import { RouterModule, Routes } from '@angular/router';

import { PageNotFoundComponent } from './ui';

const routes: Routes = [
{
    path: '',
    redirectTo: '/dashboard',
  },
  {
    path: '**',
    component: PageNotFoundComponent
  }
];

export const AppRoutes = RouterModule.forRoot(routes);

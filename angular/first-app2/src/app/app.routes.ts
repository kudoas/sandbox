import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { DetailsComponent } from './details/details.component';

export const routes: Routes = [
  { path: '', component: HomeComponent, title: 'Home page' },
  { path: 'details/:id', component: DetailsComponent, title: 'Details page' },
  {
    path: 'search',
    loadComponent: () => import('./sample-search/sample-search.component'),
    title: 'Search page',
  },
];

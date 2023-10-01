import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { DetailsComponent } from './details/details.component';
import { SampleSearchComponent } from './sample-search/sample-search.component';

export const routes: Routes = [
  { path: '', component: HomeComponent, title: 'Home page' },
  { path: 'details/:id', component: DetailsComponent, title: 'Details page' },
  { path: 'search', component: SampleSearchComponent, title: 'Search page' },
];

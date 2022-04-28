import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Route } from '@angular/router';
import { HomeComponent } from './components/home/home.component';

export const pagesHomeRoutes: Route[] = [
  {
    path: '',
    component: HomeComponent,
  },
];

@NgModule({
  imports: [CommonModule, RouterModule.forChild(pagesHomeRoutes)],
  declarations: [HomeComponent],
})
export class PagesHomeModule {}

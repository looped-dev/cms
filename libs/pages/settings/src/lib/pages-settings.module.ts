import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Route } from '@angular/router';
import { HomeComponent } from './components/home/home.component';

export const pagesSettingsRoutes: Route[] = [
  {
    path: '',
    component: HomeComponent,
  },
];

@NgModule({
  imports: [CommonModule, RouterModule.forChild(pagesSettingsRoutes)],
  declarations: [HomeComponent],
})
export class PagesSettingsModule {}

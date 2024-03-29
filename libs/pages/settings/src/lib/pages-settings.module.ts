import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Route } from '@angular/router';
import { HomeComponent } from './components/home/home.component';
import { FormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { CdkAccordionModule } from '@angular/cdk/accordion';

export const pagesSettingsRoutes: Route[] = [
  {
    path: '',
    component: HomeComponent,
  },
];

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forChild(pagesSettingsRoutes),
    FormsModule,
    FontAwesomeModule,
    CdkAccordionModule,
  ],
  declarations: [HomeComponent],
})
export class PagesSettingsModule {}

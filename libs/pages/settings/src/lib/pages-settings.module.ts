import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Route } from '@angular/router';
import { HomeComponent } from './components/home/home.component';
import { FormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { GeneralSettingsComponent } from './components/general-settings/general-settings.component';
import { SeoComponent } from './components/seo/seo.component';
import { TwitterComponent } from './components/twitter/twitter.component';
import { FacebookComponent } from './components/facebook/facebook.component';
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
  declarations: [
    HomeComponent,
    GeneralSettingsComponent,
    SeoComponent,
    TwitterComponent,
    FacebookComponent,
  ],
})
export class PagesSettingsModule {}

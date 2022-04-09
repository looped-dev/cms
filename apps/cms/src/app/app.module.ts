import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { RouterModule, Routes } from '@angular/router';
import { ApolloModule, APOLLO_OPTIONS } from 'apollo-angular';
import { HttpLink } from 'apollo-angular/http';
import { HttpClientModule } from '@angular/common/http';
import { AkitaNgDevtools } from '@datorama/akita-ngdevtools';
import { AkitaNgRouterStoreModule } from '@datorama/akita-ng-router-store';
import { environment } from '../environments/environment';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { MainComponent } from './main/main.component';
import {
  SessionQuery,
  StaffMustBeLoggedInGuard,
  StaffMustNotBeLoggedInGuard,
} from '@looped-cms/auth';
import { createApollo } from './utils/createApollo';

const routes: Routes = [
  {
    path: '',
    component: MainComponent,
    children: [
      {
        path: '',
        loadChildren: () =>
          import('@looped-cms/dashboard').then((m) => m.DashboardModule),
        canActivate: [StaffMustBeLoggedInGuard],
      },
      {
        path: 'auth',
        loadChildren: () =>
          import('@looped-cms/auth').then((m) => m.AuthModule),
        canActivate: [StaffMustNotBeLoggedInGuard],
      },
    ],
  },
  {
    path: 'setup',
    loadChildren: () => import('@looped-cms/setup').then((m) => m.SetupModule),
  },
];

@NgModule({
  declarations: [AppComponent, MainComponent],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot(routes, { initialNavigation: 'enabledBlocking' }),
    ApolloModule,
    environment.production ? [] : AkitaNgDevtools.forRoot(),
    AkitaNgRouterStoreModule,
    FontAwesomeModule,
  ],
  providers: [
    {
      provide: APOLLO_OPTIONS,
      useFactory: createApollo,
      deps: [HttpLink, SessionQuery],
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}

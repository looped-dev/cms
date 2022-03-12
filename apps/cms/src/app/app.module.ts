import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { RouterModule, Routes } from '@angular/router';
import { ApolloModule, APOLLO_OPTIONS } from 'apollo-angular';
import { InMemoryCache } from '@apollo/client/core';
import { HttpLink } from 'apollo-angular/http';
import { AuthModule } from '@looped-cms/auth';
import { HttpClientModule } from '@angular/common/http';
import { AkitaNgDevtools } from '@datorama/akita-ngdevtools';
import { AkitaNgRouterStoreModule } from '@datorama/akita-ng-router-store';
import { environment } from '../environments/environment';
import { CanViewSetupGuard } from '@looped-cms/setup';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

const routes: Routes = [
  {
    path: '',
    component: AppComponent,
    children: [
      {
        path: 'auth',
        loadChildren: () =>
          import('@looped-cms/auth').then((m) => m.AuthModule),
      },
    ],
  },
  {
    path: 'setup',
    loadChildren: () => import('@looped-cms/setup').then((m) => m.SetupModule),
    canActivate: [CanViewSetupGuard],
  },
];

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot(routes, { initialNavigation: 'enabledBlocking' }),
    AuthModule,
    ApolloModule,
    environment.production ? [] : AkitaNgDevtools.forRoot(),
    AkitaNgRouterStoreModule,
    FontAwesomeModule,
  ],
  providers: [
    {
      provide: APOLLO_OPTIONS,
      useFactory: (httpLink: HttpLink) => {
        return {
          cache: new InMemoryCache(),
          link: httpLink.create({
            uri: environment.graphql.endpoint,
          }),
        };
      },
      deps: [HttpLink],
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}

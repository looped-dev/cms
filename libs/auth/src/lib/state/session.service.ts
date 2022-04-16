import { Inject, Injectable, InjectionToken } from '@angular/core';
import { map, tap } from 'rxjs/operators';
import { SessionStore } from './session.store';
import { StaffLoginDocument, StaffLoginMutation } from '@looped-cms/graphql';
import { Apollo } from 'apollo-angular';
import { LocalStorageService } from '../services/local-storage.service';

export interface ApolloImport {
  apollo: Apollo;
}

export const apolloClientInjectionToken = new InjectionToken<ApolloImport>(
  'Apollo Client Injection Token'
);

@Injectable({ providedIn: 'root' })
export class SessionService {
  constructor(
    private sessionStore: SessionStore,
    private localStorage: LocalStorageService,
    @Inject(apolloClientInjectionToken) private apollo: Apollo
  ) {}

  login(email: string, password: string) {
    return this.apollo
      .mutate<StaffLoginMutation>({
        mutation: StaffLoginDocument,
        variables: {
          input: {
            email: email,
            password: password,
          },
        },
      })
      .pipe(
        tap(({ data }) => {
          this.localStorage.setTokens(
            data?.staffLogin.accessToken ?? '',
            data?.staffLogin.refreshToken ?? ''
          );
          this.sessionStore.update({
            accessToken: data?.staffLogin.accessToken,
            refreshToken: data?.staffLogin.refreshToken,
            staff: data?.staffLogin.staff,
          });
        }),
        map(({ data }) => {
          return data?.staffLogin;
        })
      );
  }
}

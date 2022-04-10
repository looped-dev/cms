import { Injectable } from '@angular/core';
import { map, tap } from 'rxjs/operators';
import { SessionStore } from './session.store';
import {
  RefreshStaffTokenDocument,
  RefreshStaffTokenMutation,
  StaffLoginDocument,
  StaffLoginMutation,
} from '@looped-cms/graphql';
import { Apollo } from 'apollo-angular';
import { LocalStorageService } from '../services/local-storage.service';
import { Observable } from 'rxjs';
import { SessionQuery } from './session.query';

@Injectable({ providedIn: 'root' })
export class SessionService {
  constructor(
    private sessionStore: SessionStore,
    private localStorage: LocalStorageService,
    private sessionQuery: SessionQuery,
    private apollo: Apollo
  ) {}

  refreshToken(): Observable<
    RefreshStaffTokenMutation['staffRefreshToken'] | undefined
  > {
    return this.apollo
      .mutate<RefreshStaffTokenMutation>({
        mutation: RefreshStaffTokenDocument,
        variables: {
          input: {
            accessToken: this.sessionQuery.getValue().accessToken,
            refreshToken: this.sessionQuery.getValue().refreshToken,
          },
        },
      })
      .pipe(
        tap(({ data }) => {
          this.localStorage.setTokens(
            data?.staffRefreshToken.accessToken ?? '',
            data?.staffRefreshToken.refreshToken ?? ''
          );
          this.sessionStore.update({
            accessToken: data?.staffRefreshToken.accessToken,
            refreshToken: data?.staffRefreshToken.refreshToken,
            staff: data?.staffRefreshToken.staff,
          });
        }),
        map(({ data }) => {
          return data?.staffRefreshToken;
        })
      );
  }

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

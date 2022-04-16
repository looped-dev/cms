import { Injectable, Injector } from '@angular/core';
import {
  RefreshStaffTokenMutation,
  RefreshStaffTokenDocument,
} from '@looped-cms/graphql';
import { Apollo } from 'apollo-angular';
import { Observable, tap, map } from 'rxjs';
import { SessionQuery, SessionStore } from '../state';
import { LocalStorageService } from './local-storage.service';

@Injectable({
  providedIn: 'root',
})
export class RefreshTokenService {
  constructor(
    private injector: Injector,
    private sessionStore: SessionStore,
    private sessionQuery: SessionQuery,
    private localStorage: LocalStorageService
  ) {}

  refreshToken(): Observable<
    RefreshStaffTokenMutation['staffRefreshToken'] | undefined
  > {
    return this.injector
      .get(Apollo)
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
}

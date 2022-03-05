import { Injectable } from '@angular/core';
import { map, tap } from 'rxjs/operators';
import { SessionStore } from './session.store';
import { StaffLoginDocument, StaffLoginMutation } from '@looped-cms/graphql';
import { Apollo } from 'apollo-angular';

@Injectable({ providedIn: 'root' })
export class SessionService {
  constructor(private sessionStore: SessionStore, private apollo: Apollo) {}

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
          console.log(data);
          this.sessionStore.add({
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

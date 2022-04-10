import { Injectable } from '@angular/core';
import {
  InitialSetupInput,
  IsSiteSetupDocument,
  IsSiteSetupQuery,
  SetupSiteDocument,
  SetupSiteMutation,
} from '@looped-cms/graphql';
import { Apollo } from 'apollo-angular';
import { map, Observable } from 'rxjs';
import { SetupRegisterStore } from './setup-register.store';

@Injectable({ providedIn: 'root' })
export class SetupRegisterService {
  constructor(
    private setupRegisterStore: SetupRegisterStore,
    private apollo: Apollo
  ) {}

  initialSetup(
    input: InitialSetupInput
  ): Observable<SetupSiteMutation['initialSetup'] | undefined> {
    return this.apollo
      .mutate<SetupSiteMutation>({
        mutation: SetupSiteDocument,
        variables: { input },
      })
      .pipe(map((result) => result.data?.initialSetup));
  }

  isSiteSetup = () =>
    this.apollo
      .query<IsSiteSetupQuery>({
        query: IsSiteSetupDocument,
      })
      .pipe(map((result) => result.data?.isSiteSetup));
}

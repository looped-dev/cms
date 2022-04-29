import { Injectable } from '@angular/core';
import {
  FetchSettingsDocument,
  FetchSettingsQuery,
  UpdateSiteSettingsDocument,
  UpdateSiteSettingsInput,
  UpdateSiteSettingsMutation,
  UpdateSiteSettingsMutationVariables,
} from '@looped-cms/graphql';
import { Apollo } from 'apollo-angular';
import { map, Observable, tap } from 'rxjs';
import { SettingsStore } from './settings.store';

@Injectable({ providedIn: 'root' })
export class SettingsService {
  constructor(private settingsStore: SettingsStore, private apollo: Apollo) {}

  get(): Observable<FetchSettingsQuery['settings']> {
    return this.apollo
      .query<FetchSettingsQuery>({
        query: FetchSettingsDocument,
      })
      .pipe(
        map((data) => data?.data.settings),
        tap((settings) => {
          this.settingsStore.update(settings);
        })
      );
  }

  update(
    settings: UpdateSiteSettingsInput
  ): Observable<UpdateSiteSettingsMutation['updateSiteSettings'] | undefined> {
    return this.apollo
      .mutate<UpdateSiteSettingsMutation, UpdateSiteSettingsMutationVariables>({
        mutation: UpdateSiteSettingsDocument,
        variables: {
          input: settings,
        },
      })
      .pipe(
        map(({ data }) => {
          if (!data) {
            throw 'Error updating settings';
          }
          return data.updateSiteSettings;
        }),
        tap((settings) => {
          this.settingsStore.update(settings);
        })
      );
  }
}

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
      .query<FetchSettingsQuery, null>({
        query: FetchSettingsDocument,
      })
      .pipe(
        map(({ data, error }) => {
          if (!data || error) {
            this.settingsStore.setError(error || 'Error fetching settings');
            throw error || 'Error updating settings';
          }
          return data.settings;
        }),
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
        map(({ data, errors }) => {
          if (!data || errors) {
            throw errors || 'Error updating settings';
          }
          return data.updateSiteSettings;
        }),
        tap((settings) => {
          this.settingsStore.update(settings);
        })
      );
  }
}

import { Injectable } from '@angular/core';
import { Query } from '@datorama/akita';
import { Setting } from './setting.model';
import { SettingsStore } from './settings.store';

@Injectable({ providedIn: 'root' })
export class SettingsQuery extends Query<Setting> {
  constructor(protected override store: SettingsStore) {
    super(store);
  }
}

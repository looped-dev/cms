import { Injectable } from '@angular/core';
import { Store, StoreConfig } from '@datorama/akita';
import { createSetting, Setting } from './setting.model';

@Injectable({ providedIn: 'root' })
@StoreConfig({ name: 'settings' })
export class SettingsStore extends Store<Setting> {
  constructor() {
    super(createSetting({}));
  }
}

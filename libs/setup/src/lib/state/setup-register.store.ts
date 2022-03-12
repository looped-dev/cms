import { Injectable } from '@angular/core';
import { EntityState, EntityStore, StoreConfig } from '@datorama/akita';
import { SetupRegister } from './setup-register.model';

export type SetupRegisterState = EntityState<SetupRegister>;

@Injectable({ providedIn: 'root' })
@StoreConfig({ name: 'setup-register' })
export class SetupRegisterStore extends EntityStore<SetupRegisterState> {
  constructor() {
    super();
  }
}

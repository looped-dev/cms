import { Injectable } from '@angular/core';
import { QueryEntity } from '@datorama/akita';
import { SetupRegisterStore, SetupRegisterState } from './setup-register.store';

@Injectable({ providedIn: 'root' })
export class SetupRegisterQuery extends QueryEntity<SetupRegisterState> {
  isSetup$ = this.select('isSetup');

  constructor(protected override store: SetupRegisterStore) {
    super(store);
  }
}

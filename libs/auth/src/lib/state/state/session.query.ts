import { Injectable } from '@angular/core';
import { QueryEntity } from '@datorama/akita';
import { SessionStore, SessionState } from './session.store';

@Injectable({ providedIn: 'root' })
export class SessionQuery extends QueryEntity<SessionState> {
  selectIsLogin$ = this.select('token');
  selectStaff = this.select('staff');

  constructor(protected override store: SessionStore) {
    super(store);
  }
}

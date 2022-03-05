import { Injectable } from '@angular/core';
import { EntityState, EntityStore, StoreConfig } from '@datorama/akita';
import { Session } from './session.model';

export type SessionState = EntityState<Session>;

@Injectable({ providedIn: 'root' })
@StoreConfig({ name: 'session' })
export class SessionStore extends EntityStore<SessionState> {
  constructor() {
    super();
  }
}

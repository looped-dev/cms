import { Injectable } from '@angular/core';
import { Store, StoreConfig } from '@datorama/akita';
import { createSession, Session } from './session.model';

@Injectable({ providedIn: 'root' })
@StoreConfig({ name: 'session' })
export class SessionStore extends Store<Session> {
  constructor() {
    super(createSession());
  }
}

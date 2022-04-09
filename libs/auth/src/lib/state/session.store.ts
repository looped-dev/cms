import { Injectable } from '@angular/core';
import { Store, StoreConfig } from '@datorama/akita';
import { LocalStorageService } from '../../services/local-storage.service';
import { createSession, Session } from './session.model';

@Injectable({ providedIn: 'root' })
@StoreConfig({ name: 'session' })
export class SessionStore extends Store<Session> {
  constructor(localStorage: LocalStorageService) {
    const accessToken = localStorage.getAccessToken();
    const refreshToken = localStorage.getRefreshToken();

    super(
      createSession({
        accessToken,
        refreshToken,
      })
    );
  }
}

import { Injectable } from '@angular/core';

const ACCESS_TOKEN_KEY = 'accessToken';
const REFRESH_TOKEN_KEY = 'refreshToken';

@Injectable({
  providedIn: 'root',
})
export class LocalStorageService {
  getAccessToken() {
    return localStorage.getItem(ACCESS_TOKEN_KEY) ?? '';
  }

  getRefreshToken() {
    return localStorage.getItem(REFRESH_TOKEN_KEY) ?? '';
  }

  setTokens(accessToken: string, refreshToken: string) {
    localStorage.setItem(ACCESS_TOKEN_KEY, accessToken);
    localStorage.setItem(REFRESH_TOKEN_KEY, refreshToken);
  }

  removeTokens() {
    localStorage.removeItem(ACCESS_TOKEN_KEY);
    localStorage.removeItem(REFRESH_TOKEN_KEY);
  }
}

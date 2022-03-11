import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivate,
  RouterStateSnapshot,
  UrlTree,
} from '@angular/router';
import { map, Observable } from 'rxjs';
import { SetupRegisterQuery } from '../state/state/setup-register.query';

@Injectable({
  providedIn: 'root',
})
export class CanViewSetupGuard implements CanActivate {
  constructor(private setupRegisterQuery: SetupRegisterQuery) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<boolean | UrlTree> {
    // if the site is setup, then we don't allow the user to view the setup page
    return this.setupRegisterQuery.isSetup$.pipe(map((isSetup) => !isSetup));
  }
}

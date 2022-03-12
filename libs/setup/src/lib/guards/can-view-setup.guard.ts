import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivate,
  RouterStateSnapshot,
  UrlTree,
} from '@angular/router';
import { map, Observable, tap } from 'rxjs';
import { SetupRegisterQuery } from '../state/state/setup-register.query';
import { SetupRegisterService } from '../state/state/setup-register.service';

@Injectable({
  providedIn: 'root',
})
export class CanViewSetupGuard implements CanActivate {
  constructor(
    private setupRegisterQuery: SetupRegisterQuery,
    private setupRegisterService: SetupRegisterService
  ) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<boolean | UrlTree> {
    // if the site is setup, then we don't allow the user to view the setup page
    return this.setupRegisterQuery.isSetup$.pipe(
      tap((v) => console.log(v)),
      map((isSetup) => !isSetup)
    );
  }
}

import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { map, Observable } from 'rxjs';
import { SessionQuery } from './state/session.query';

@Injectable({
  providedIn: 'root',
})
export class StaffMustBeLoggedInGuard implements CanActivate {
  constructor(private router: Router, private sessionQuery: SessionQuery) {}

  canActivate(): Observable<boolean> {
    return this.sessionQuery.isLoggedIn$.pipe(
      map((isLoggedIn) => {
        if (isLoggedIn) {
          return true;
        }
        this.router.navigateByUrl(`auth/signin?redirectTo=${this.router.url}`);
        return false;
      })
    );
  }
}

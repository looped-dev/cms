import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { map, Observable } from 'rxjs';
import { SessionQuery } from './state/state/session.query';

@Injectable({
  providedIn: 'root',
})
export class StaffMustNotBeLoggedInGuard implements CanActivate {
  constructor(private router: Router, private sessionQuery: SessionQuery) {}

  canActivate(): Observable<boolean> {
    return this.sessionQuery.isLoggedIn$.pipe(
      map((isLoggedIn) => {
        // if user is not logged in, allow access
        if (!isLoggedIn) {
          return true;
        }
        this.router.navigateByUrl('/');
        return false;
      })
    );
  }
}

import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { faSpinner } from '@fortawesome/free-solid-svg-icons';
import { map, Observable, tap } from 'rxjs';
import { SetupRegisterService } from '../../state/setup-register.service';

@Component({
  selector: 'looped-cms-layout',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.scss'],
})
export class LayoutComponent {
  loadingIcon = faSpinner;

  isSetup$: Observable<{ isSetup: boolean }> = this.setupRegisterService
    .isSiteSetup()
    .pipe(
      tap((isSetup) => {
        // if cms is setup, then redirect to the home page
        if (isSetup) {
          this.router.navigate(['/']);
        }
      }),
      map((isSetup) => ({ isSetup: isSetup })),
      tap((isSetup) => console.log(isSetup))
    );

  constructor(
    private setupRegisterService: SetupRegisterService,
    private router: Router
  ) {}
}

import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { faSpinner } from '@fortawesome/free-solid-svg-icons';
import { SetupRegisterService } from '@looped-cms/setup';
import { delay, Observable, tap } from 'rxjs';

@Component({
  selector: 'looped-cms-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  loadingIcon = faSpinner;

  isSetup$: Observable<boolean> = this.setupRegisterService.isSiteSetup().pipe(
    tap((isSetup) => {
      if (!isSetup) {
        this.router.navigate(['/setup']);
      }
    })
  );

  constructor(
    private setupRegisterService: SetupRegisterService,
    private router: Router
  ) {}
}

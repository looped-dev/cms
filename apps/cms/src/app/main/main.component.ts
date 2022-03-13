import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { faSpinner } from '@fortawesome/free-solid-svg-icons';
import { SetupRegisterService } from '@looped-cms/setup';
import { map, Observable, tap } from 'rxjs';

@Component({
  selector: 'looped-cms-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss'],
})
export class MainComponent {
  loadingIcon = faSpinner;

  isSetup$: Observable<{ isSetup: boolean }> = this.setupRegisterService
    .isSiteSetup()
    .pipe(
      map((isSetup) => ({ isSetup })),
      tap(({ isSetup }) => {
        if (!isSetup) {
          this.router.navigate(['/setup']);
        }
      })
    );

  constructor(
    private setupRegisterService: SetupRegisterService,
    private router: Router
  ) {
    console.log('MainComponent');
  }
}

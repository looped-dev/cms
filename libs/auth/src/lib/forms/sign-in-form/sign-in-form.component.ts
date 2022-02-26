import { Component } from '@angular/core';
import { faUnlock } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'looped-cms-sign-in-form',
  templateUrl: './sign-in-form.component.html',
  styleUrls: ['./sign-in-form.component.scss'],
})
export class SignInFormComponent {
  unlockIcon = faUnlock;

  constructor() {
    console.log('Hello World');
  }
}

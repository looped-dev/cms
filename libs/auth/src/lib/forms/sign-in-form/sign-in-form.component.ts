import { Component, ViewChild } from '@angular/core';
import { faUnlock } from '@fortawesome/free-solid-svg-icons';
import { NgForm } from '@angular/forms';

type SignInFormData = {
  email: string;
  password: string;
};

@Component({
  selector: 'looped-cms-sign-in-form',
  templateUrl: './sign-in-form.component.html',
  styleUrls: ['./sign-in-form.component.scss'],
})
export class SignInFormComponent {
  unlockIcon = faUnlock;

  @ViewChild(NgForm) signInForm!: NgForm;

  signInFormModel: SignInFormData = {
    email: '',
    password: '',
  };

  constructor() {
    console.log('Hello World');
  }

  onSubmit() {
    console.log(this.signInForm.value);
  }
}

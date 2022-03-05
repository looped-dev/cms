import { Component, ViewChild } from '@angular/core';
import { faUnlock } from '@fortawesome/free-solid-svg-icons';
import { NgForm } from '@angular/forms';
import { StaffLoginDocument } from '@looped-cms/graphql';
import { Apollo } from 'apollo-angular';
import { SessionService } from '../../state/state/session.service';

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

  errorMessage = '';

  signInFormModel: SignInFormData = {
    email: '',
    password: '',
  };

  constructor(private sessionService: SessionService) {}

  onSubmit() {
    this.sessionService
      .login(this.signInFormModel.email, this.signInFormModel.password)
      .subscribe({
        next: (data) => {
          console.log({ data });
        },
        error: (error) => (this.errorMessage = error?.message),
      });
  }
}

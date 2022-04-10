import { Component, ViewChild } from '@angular/core';
import { faUnlock } from '@fortawesome/free-solid-svg-icons';
import { NgForm } from '@angular/forms';
import { SessionService } from '../../state/session.service';
import { ActivatedRoute, Router } from '@angular/router';

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

  constructor(
    private sessionService: SessionService,
    private router: Router,
    private activatedRoute: ActivatedRoute
  ) {}

  onSubmit() {
    this.sessionService
      .login(this.signInFormModel.email, this.signInFormModel.password)
      .subscribe({
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        next: (_) => {
          const redirectTo =
            this.activatedRoute.snapshot.queryParams['redirectTo'];
          this.router.navigateByUrl(redirectTo ?? '/');
        },
        error: (error) => (this.errorMessage = error?.message),
      });
  }
}

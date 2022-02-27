import { Component, ViewChild } from '@angular/core';
import { faUnlock } from '@fortawesome/free-solid-svg-icons';
import { NgForm } from '@angular/forms';
import { StaffLoginDocument } from '@looped-cms/graphql';
import { Apollo } from 'apollo-angular';

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

  constructor(private apollo: Apollo) {}

  onSubmit() {
    this.apollo
      .mutate({
        mutation: StaffLoginDocument,
        variables: {
          input: {
            email: this.signInFormModel.email,
            password: this.signInFormModel.password,
          },
        },
      })
      .subscribe({
        next: ({ data }) => {
          console.log({ data });
        },
        error: (error) => console.log(error),
      });
  }
}

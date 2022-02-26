import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
import { SignInFormComponent } from './forms/sign-in-form/sign-in-form.component';
import { LayoutComponent } from './containers/layout/layout.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

const routes: Routes = [
  {
    path: '',
    component: LayoutComponent,
    children: [
      {
        path: 'signin',
        component: SignInFormComponent,
      },
    ],
  },
];

@NgModule({
  imports: [CommonModule, FontAwesomeModule, RouterModule.forChild(routes)],
  declarations: [SignInFormComponent, LayoutComponent],
})
export class AuthModule {}

import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
import { SignInFormComponent } from './forms/sign-in-form/sign-in-form.component';
import { LayoutComponent } from './containers/layout/layout.component';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { FormsModule } from '@angular/forms';
import { StaffMustNotBeLoggedInGuard } from './staff-must-not-be-logged-in.guard';

const routes: Routes = [
  {
    path: '',
    component: LayoutComponent,
    canActivate: [StaffMustNotBeLoggedInGuard],
    children: [
      {
        path: 'signin',
        component: SignInFormComponent,
      },
    ],
  },
];

@NgModule({
  imports: [
    CommonModule,
    FontAwesomeModule,
    RouterModule.forChild(routes),
    FormsModule,
  ],
  declarations: [SignInFormComponent, LayoutComponent],
})
export class AuthModule {}

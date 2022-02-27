import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Route, RouterModule } from '@angular/router';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { LayoutComponent } from './containers/layout/layout.component';
import { StepOneComponent } from './forms/step-one/step-one.component';
import { StepTwoComponent } from './forms/step-two/step-two.component';

export const setupRoutes: Route[] = [
  {
    path: '',
    component: LayoutComponent,
    children: [
      {
        path: 'steps/one',
        component: StepOneComponent,
      },
      {
        path: 'steps/two',
        component: StepTwoComponent,
      },
    ],
  },
];

@NgModule({
  imports: [CommonModule, RouterModule, FormsModule, FontAwesomeModule],
  declarations: [LayoutComponent, StepOneComponent, StepTwoComponent],
})
export class SetupModule {}

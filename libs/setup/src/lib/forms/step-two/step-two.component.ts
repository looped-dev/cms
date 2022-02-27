import { Component, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { faSave } from '@fortawesome/free-solid-svg-icons';

// todo: replace this with schema input for the same input
type SetupSiteModel = {
  siteTitle: string;
  name: string;
  email: string;
  password: string;
};

@Component({
  selector: 'looped-cms-step-two',
  templateUrl: './step-two.component.html',
  styleUrls: ['./step-two.component.scss'],
})
export class StepTwoComponent {
  errorMessage = '';

  saveIcon = faSave;

  @ViewChild(NgForm) setupSiteForm!: NgForm;

  setupSiteModel: SetupSiteModel = {
    siteTitle: '',
    name: '',
    email: '',
    password: '',
  };

  constructor() {
    console.log({ setupSiteModel: this.setupSiteModel });
  }

  onSubmit() {
    console.log();
  }
}

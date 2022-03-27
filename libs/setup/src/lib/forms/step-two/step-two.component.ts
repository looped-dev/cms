import { Component, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';
import { faSave } from '@fortawesome/free-solid-svg-icons';
import { InitialSetupInput } from '@looped-cms/graphql';
import { SetupRegisterService } from '../../state/setup-register.service';

@Component({
  selector: 'looped-cms-step-two',
  templateUrl: './step-two.component.html',
  styleUrls: ['./step-two.component.scss'],
})
export class StepTwoComponent {
  errorMessage = '';

  saveIcon = faSave;

  @ViewChild(NgForm) setupSiteForm!: NgForm;

  setupSiteModel: InitialSetupInput = {
    siteName: '',
    name: '',
    email: '',
    password: '',
  };

  constructor(
    private setupRegisterService: SetupRegisterService,
    private router: Router
  ) {}

  onSubmit() {
    this.setupRegisterService.initialSetup(this.setupSiteModel).subscribe({
      next: (result) => {
        if (result) {
          this.router.navigate(['/']);
        }
      },
      error: (error) => {
        this.errorMessage = error?.message ?? 'An error occurred';
      },
    });
  }
}

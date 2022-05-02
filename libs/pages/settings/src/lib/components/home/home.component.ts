import { Component, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { faSave } from '@fortawesome/free-solid-svg-icons';
import { UpdateSiteSettingsInput } from '@looped-cms/graphql';
import { SettingsQuery } from '../../state/settings.query';
import { SettingsService } from '../../state/settings.service';

@Component({
  selector: 'looped-cms-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
})
export class HomeComponent {
  @ViewChild(NgForm) settingsForm!: NgForm;

  settings$ = this.settingsQuery.select();
  errorMessage = '';

  saveIcon = faSave;

  settingsFormModel: Partial<UpdateSiteSettingsInput> = {};

  constructor(
    private settingsQuery: SettingsQuery,
    private settingsService: SettingsService
  ) {
    this.settingsService.get().subscribe({
      next: (data) => {
        this.settingsFormModel = data;
      },
      error: (err) => (this.errorMessage = err.message),
    });
  }

  onSubmit() {
    console.log(this.settingsForm.value);
  }
}

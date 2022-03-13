import { Component } from '@angular/core';

@Component({
  selector: 'looped-cms-layout',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.scss'],
})
export class LayoutComponent {
  constructor() {
    console.log('Auth Layout');
  }
}

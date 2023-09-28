import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';

import { Housinglocation } from '../housinglocation';

@Component({
  selector: 'app-housing-location',
  standalone: true,
  imports: [CommonModule],
  template: ` <p>housing-location works!</p> `,
  styleUrls: ['./housing-location.component.scss'],
})
export class HousingLocationComponent {
  @Input({ required: true }) housingLocation!: Housinglocation;
}

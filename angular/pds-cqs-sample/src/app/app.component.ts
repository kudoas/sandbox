import { Component } from '@angular/core';
import { PdsCqsTodoListComponent } from './pds-cqs.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  imports: [PdsCqsTodoListComponent],
  standalone: true,
})
export class AppComponent {
  title = 'pds-cqs-sample';
}

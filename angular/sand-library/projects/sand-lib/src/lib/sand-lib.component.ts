import { Component, OnInit, inject } from '@angular/core';
import { EndpointService } from './endpoint.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'lib-sand-lib',
  standalone: true,
  imports: [CommonModule],
  template: ` <p>{{ apiUrl }}sand-lib works!</p> `,
  styles: ``
})
export class SandLibComponent implements OnInit {
  private endpoint = inject(EndpointService);
  apiUrl: string = '';

  ngOnInit() {
    this.apiUrl = this.endpoint.apiUrl();
  }
}

import { Component, OnInit, inject } from '@angular/core';
import { EndpointService } from './endpoint.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'lib-sand-lib',
  standalone: true,
  imports: [CommonModule],
  template: `<p>URL: {{ apiUrl }}</p> `
})
export class SandLibComponent implements OnInit {
  private endpointService = inject(EndpointService);
  apiUrl: string = '';

  ngOnInit() {
    this.apiUrl = this.endpointService.getApiUrl();
  }
}

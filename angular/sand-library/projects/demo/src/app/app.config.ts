import { ApplicationConfig } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideEndpointService } from '@lib/sandbox';

import { routes } from './app.routes';

export const appConfig: ApplicationConfig = {
  providers: [provideRouter(routes), provideEndpointService({ apiUrl: 'http://localhost:3000' })]
};

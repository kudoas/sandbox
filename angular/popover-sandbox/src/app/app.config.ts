import { ApplicationConfig, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideAnimations } from '@angular/platform-browser/animations';
import { CDK_CONNECTED_OVERLAY_DEFAULT_CONFIG, OVERLAY_DEFAULT_CONFIG } from '@angular/cdk/overlay';

import { routes } from './app.routes';

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideRouter(routes),
    provideAnimations(),
    {
      provide: CDK_CONNECTED_OVERLAY_DEFAULT_CONFIG,
      useValue: {
        usePopover: 'global',
      },
    },
    {
      provide: OVERLAY_DEFAULT_CONFIG,
      useValue: {
        usePopover: false,
      },
    },
  ],
};

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
        usePopover: 'inline',
      },
    },
    // {
    // provide: OVERLAY_DEFAULT_CONFIG, // overlay.create 時のデフォルト設定
    //   useValue: {
    //     usePopover: true,
    //   },
    // },
  ],
};

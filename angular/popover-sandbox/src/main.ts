import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { App } from './app/app';
import { patchMatTooltipPopoverInline } from './app/patch-mat-tooltip-popover';

// patchMatTooltipPopoverInline();
bootstrapApplication(App, appConfig).catch((err) => console.error(err));

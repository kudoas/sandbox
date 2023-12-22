import { Injectable, Inject, Provider } from '@angular/core';
import { InjectionToken } from '@angular/core';

export interface EndpointServiceParams {
  apiUrl: string;
}

export const HERO_SERVICE_PARAMS = new InjectionToken<EndpointServiceParams>('endpoint');

@Injectable()
export class EndpointService {
  constructor(@Inject(HERO_SERVICE_PARAMS) private params: EndpointServiceParams) {}

  apiUrl(): string {
    return this.params.apiUrl;
  }
}

export function provideEndpointService(params: EndpointServiceParams): Provider[] {
  return [EndpointService, { provide: HERO_SERVICE_PARAMS, useValue: params }];
}

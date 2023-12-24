import { Injectable, Inject, Provider, Optional } from '@angular/core';
import { InjectionToken } from '@angular/core';

export interface EndpointServiceParams {
  isProduction: boolean;
}

export const APP_ENVIRONMENT = new InjectionToken<EndpointServiceParams>('environmentConfig');

@Injectable()
export class EndpointService {
  private apiUrl: string;

  constructor(@Inject(APP_ENVIRONMENT) private params: EndpointServiceParams) {
    const isProduction = this.params?.isProduction ?? false;
    this.apiUrl = isProduction ? 'https://production.api.endpoint' : 'https://development.api.endpoint';
  }

  getApiUrl(): string {
    return this.apiUrl;
  }
}

export function provideEndpointService(params: EndpointServiceParams): Provider[] {
  return [EndpointService, { provide: APP_ENVIRONMENT, useValue: params }];
}

import { Signal } from '@angular/core';

export type SignalState<T> = {
  asReadonly(): ReadonlyState<T>;
};

export type ReadonlyState<T> = {
  [K in keyof T]: Signal<T[K]>;
};

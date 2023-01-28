import { createAction, props } from '@ngrx/store';

/**
 * The createAction function returns a function,
 * that when called returns an object in the shape of the Action interface.
 * The props method is used to define any additional metadata needed for the handling of the action.
 * Action creators provide a consistent,
 * type-safe way to construct an action that is being dispatched.
 */
export const login = createAction(
  '[Login Page] Login',
  props<{ username: string; password: string }>()
);

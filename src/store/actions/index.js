import * as actionTypes from "./actionTypes";

export const numClick = number => ({
  type: actionTypes.INPUT_NUMBER,
  num: number
});

export const plusClick = () => ({
  type: actionTypes.PLUS
});

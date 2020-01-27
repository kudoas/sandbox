import * as actionTypes from "../actions/actionTypes";
import { updateObject } from "../../utils/updateObject";

const initialState = {
  inputValue: 0,
  resultValue: 0,
  showingResult: false
};

const calculator = (state = initialState, action) => {
  switch (action.type) {
    case actionTypes.INPUT_NUMBER:
      return updateObject(state, {
        inputValue: state.inputValue * 10 + action.num,
        showingResult: false
      });
    case actionTypes.PLUS:
      return updateObject(state, {
        inputValue: 0,
        resultValue: state.inputValue + state.resultValue,
        showingResult: true
      });
    default:
      return state;
  }
};

export default calculator;

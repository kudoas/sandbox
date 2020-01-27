import React, { Component } from "react";
import { connect } from "react-redux";

import * as actionCreater from "../store/actions/index";
import NumBtn from "../components/NumBtn";
import PlusBtn from "../components/PlusBtn";
import Result from "../components/Result";

class CalculatorContainer extends Component {
  render() {
    return (
      <div>
        <div>
          <NumBtn num={1} clicked={() => this.props.onNumClick(1)} />
          <NumBtn num={2} clicked={() => this.props.onNumClick(2)} />
          <NumBtn num={3} clicked={() => this.props.onNumClick(3)} />
        </div>
        <div>
          <NumBtn num={4} clicked={() => this.props.onNumClick(4)} />
          <NumBtn num={5} clicked={() => this.props.onNumClick(5)} />
          <NumBtn num={6} clicked={() => this.props.onNumClick(6)} />
        </div>
        <div>
          <NumBtn num={7} clicked={() => this.props.onNumClick(7)} />
          <NumBtn num={8} clicked={() => this.props.onNumClick(8)} />
          <NumBtn num={9} clicked={() => this.props.onNumClick(9)} />
        </div>
        <div>
          <NumBtn num={0} clicked={() => this.props.onNumClick(0)} />
          <PlusBtn clicked={() => this.props.onPlusClick()} />
        </div>
        <Result
          result={
            this.props.calculator.showingResult
              ? this.props.calculator.resultValue
              : this.props.calculator.inputValue
          }
        />
      </div>
    );
  }
}

const mapStateToProps = state => ({
  calculator: state.calculator
});

const mapDispatchToProps = dispatch => {
  return {
    onNumClick: n => dispatch(actionCreater.numClick(n)),
    onPlusClick: () => dispatch(actionCreater.plusClick())
  };
};

export default connect(mapStateToProps, mapDispatchToProps)(CalculatorContainer);

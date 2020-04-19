import React, { Component } from "react";
import { connect } from "react-redux";

import styled from "@emotion/styled";

import * as actionCreater from "../store/actions/index";
import NumBtn from "../components/NumBtn";
import CalBtn from "../components/CalBtn";
import Result from "../components/Result";

class Calculator extends Component {
  render() {
    return (
      <Wrapper>
        <Result
          result={
            this.props.calculator.showingResult
              ? this.props.calculator.resultValue
              : this.props.calculator.inputValue
          }
        />
        <div style={{ display: "flex" }}>
          <NumBtn />
          <NumBtn />
          <NumBtn />
          <NumBtn />
        </div>
        <div style={{ display: "flex" }}>
          <NumBtn num={1} clicked={() => this.props.onNumClick(1)} />
          <NumBtn num={2} clicked={() => this.props.onNumClick(2)} />
          <NumBtn num={3} clicked={() => this.props.onNumClick(3)} />
          <NumBtn />
        </div>
        <div style={{ display: "flex" }}>
          <NumBtn num={4} clicked={() => this.props.onNumClick(4)} />
          <NumBtn num={5} clicked={() => this.props.onNumClick(5)} />
          <NumBtn num={6} clicked={() => this.props.onNumClick(6)} />
          <NumBtn />
        </div>
        <div style={{ display: "flex" }}>
          <NumBtn num={7} clicked={() => this.props.onNumClick(7)} />
          <NumBtn num={8} clicked={() => this.props.onNumClick(8)} />
          <NumBtn num={9} clicked={() => this.props.onNumClick(9)} />
          <CalBtn clicked={() => this.props.onPlusClick()}>+</CalBtn>
        </div>
        <div style={{ display: "flex" }}>
          <NumBtn />
          <NumBtn num={0} clicked={() => this.props.onNumClick(0)} />
          <NumBtn />
          <NumBtn />
        </div>
      </Wrapper>
    );
  }
}

const Wrapper = styled.div`
  width: 200px;
  margin: 50px auto;
  border: 2px solid black;
  border-radius: 5px;
  background-color: #e5f9dd;
`;

const mapStateToProps = (state) => ({
  calculator: state.calculator,
});

const mapDispatchToProps = (dispatch) => {
  return {
    onNumClick: (n) => dispatch(actionCreater.numClick(n)),
    onPlusClick: () => dispatch(actionCreater.plusClick()),
  };
};

export default connect(mapStateToProps, mapDispatchToProps)(Calculator);

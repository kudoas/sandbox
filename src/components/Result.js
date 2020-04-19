import React from "react";

import styled from "@emotion/styled";

const Result = (props) => (
  <Wrapper>
    <span>{props.result}</span>
  </Wrapper>
);

const Wrapper = styled.div`
  text-align: center;
  margin: 5px;
  border-radius: 5px;
  border: solid 1px black;
`;

export default Result;

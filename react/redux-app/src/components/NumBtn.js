import React from "react";

import styled from "@emotion/styled";

const NumBtn = (props) => <Button onClick={props.clicked}>{props.num}</Button>;

const Button = styled.button`
  width: 100px;
  padding: 0.8em;
  text-align: center;
  text-decoration: none;
  color: #12250e;
  border-radius: 5px;
  background: #f1ddf9;
`;

export default NumBtn;

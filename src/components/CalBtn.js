import React from "react";

import styled from "@emotion/styled";

const CalBtn = (props) => <Button onClick={props.clicked}>{props.children}</Button>;

const Button = styled.button`
  width: 100px;
  padding: 0.8em;
  text-align: center;
  text-decoration: none;
  color: #12250e;
  border-radius: 5px;
  background: #f9dde5;
`;

export default CalBtn;

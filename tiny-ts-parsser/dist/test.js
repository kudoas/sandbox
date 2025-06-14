"use strict";
exports.__esModule = true;
var arith_ts_1 = require("./arith.ts");
var tiny_ts_parser_ts_1 = require("./tiny-ts-parser.ts");
var ast = tiny_ts_parser_ts_1.parseArith("false ? true : false");
console.log(ast);
console.log(arith_ts_1.typecheck(ast));

import { typecheck } from "./arith.ts";
import { parse, parseArith } from "./tiny-ts-parser.ts";

const ast = parseArith("false ? true : false")
console.log(ast)
console.log(typecheck(ast));

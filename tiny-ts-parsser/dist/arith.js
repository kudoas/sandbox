"use strict";
exports.__esModule = true;
exports.typecheck = void 0;
function typecheck(t) {
    switch (t.tag) {
        case "true":
            return { tag: "Boolean" };
        case "false":
            return { tag: "Boolean" };
        case "if": {
            var condTy = typecheck(t.cond);
            if (condTy.tag === "Unknown") {
                throw "should not use unknown type";
            }
            var thnTy = typecheck(t.thn);
            var elsTy = typecheck(t.els);
            if (thnTy.tag !== elsTy.tag) {
                throw "then and else have different types";
            }
            return thnTy;
        }
        case "number":
            return { tag: "Number" };
        case "add": {
            var leftTy = typecheck(t.left);
            if (leftTy.tag !== "Number")
                throw "number expected";
            var rightTy = typecheck(t.right);
            if (rightTy.tag !== "Number")
                throw "number expected";
            return { tag: "Number" };
        }
        default:
            return { tag: "Unknown" };
    }
}
exports.typecheck = typecheck;

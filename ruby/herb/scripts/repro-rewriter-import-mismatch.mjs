import { Herb } from "@herb-tools/node-wasm"
import { rewriteString } from "@herb-tools/rewriter"
import { tailwindClassSorter } from "@herb-tools/rewriter/loader"

await Herb.load()

const rewriter = await tailwindClassSorter({
  baseDir: process.cwd(),
  tailwindStylesheet: "app/assets/tailwind/application.css"
})

const output = rewriteString(
  Herb,
  '<div class="text-red-500 p-4 mt-2"></div>\n',
  [rewriter],
  { baseDir: process.cwd() }
)

console.log("rewriteString:", output)

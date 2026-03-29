import { TailwindClassSorter } from "@herb-tools/tailwind-class-sorter"
import { Herb } from "@herb-tools/node-wasm"
import { rewriteString, rewrite, tailwindClassSorter } from "@herb-tools/rewriter/loader"

const directSorter = await TailwindClassSorter.fromConfig({
  baseDir: process.cwd(),
  tailwindStylesheet: "app/assets/tailwind/application.css"
})

console.log("direct context:", directSorter.getContext())
console.log("direct sort:", directSorter.sortClasses("text-red-500 p-4 mt-2"))

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

const actionViewTemplate = '<%= tag.div "記事を書く", class: ["text-4xl", "uppercase", "font-semibold", "tracking-tight"] %>\n'
const actionViewParseResult = Herb.parse(actionViewTemplate, {
  track_whitespace: true,
  action_view_helpers: true
})

console.log("action_view_helpers parsed:", actionViewParseResult.value.children[0]?.type)

if (!actionViewParseResult.failed) {
  const { output: actionViewOutput } = rewrite(
    actionViewParseResult.value,
    [rewriter],
    { baseDir: process.cwd() }
  )

  console.log("action_view_helpers rewrite:", actionViewOutput)
}

import { Herb } from "@herb-tools/node-wasm"
import { IdentityPrinter } from "@herb-tools/printer"
import {
  rewrite,
  tailwindClassSorter
} from "@herb-tools/rewriter/loader"
import {
  ActionViewTagHelperToHTMLRewriter,
  HTMLToActionViewTagHelperRewriter
} from "@herb-tools/rewriter"

const baseDir = process.cwd()

const samples = [
  {
    name: "tag.div with inline content",
    template: '<%= tag.div "記事を書く", class: ["text-4xl", "uppercase", "font-semibold", "tracking-tight"] %>\n'
  },
  {
    name: "link_to with class string",
    template: '<%= link_to "一覧へ戻る", posts_path, class: "py-2 px-4 text-blue-600 font-bold" %>\n'
  },
  {
    name: "tag.div block with class string",
    template: '<%= tag.div class: "py-2 px-4 text-blue-600 font-bold" do %>\n  記事を書く\n<% end %>\n'
  }
]

await Herb.load()

const sorter = await tailwindClassSorter({
  baseDir,
  tailwindStylesheet: "app/assets/tailwind/application.css"
})

for (const sample of samples) {
  console.log(`\n=== ${sample.name} ===`)
  console.log("input:")
  console.log(sample.template)

  const parsed = Herb.parse(sample.template, {
    track_whitespace: true,
    action_view_helpers: true
  })

  if (parsed.failed) {
    console.log("parse failed")
    continue
  }

  console.log("parsed child type:", parsed.value.children[0]?.type)
  console.log("identity:")
  console.log(IdentityPrinter.print(parsed.value))

  const htmlNode = Herb.parse(sample.template, {
    track_whitespace: true,
    action_view_helpers: true
  }).value
  const helperToHtml = new ActionViewTagHelperToHTMLRewriter()
  const htmlResult = rewrite(htmlNode, [helperToHtml], { baseDir })

  console.log("after actionViewTagHelperToHTML:")
  console.log(htmlResult.output)

  const sortedNode = Herb.parse(sample.template, {
    track_whitespace: true,
    action_view_helpers: true
  }).value
  const helperToHtmlForSort = new ActionViewTagHelperToHTMLRewriter()
  const backToHelper = new HTMLToActionViewTagHelperRewriter()
  const sortedResult = rewrite(sortedNode, [helperToHtmlForSort, sorter, backToHelper], { baseDir })

  console.log("after helper -> html -> sort -> helper:")
  console.log(sortedResult.output)

  const sortOnlyNode = Herb.parse(sample.template, {
    track_whitespace: true,
    action_view_helpers: true
  }).value
  const sortOnlyResult = rewrite(sortOnlyNode, [sorter], { baseDir })

  console.log("after sort only:")
  console.log(sortOnlyResult.output)
}

const { TailwindClassSorter } = require("@herb-tools/tailwind-class-sorter")
const { Herb } = require("@herb-tools/node-wasm")
const { rewriteString, tailwindClassSorter } = require("@herb-tools/rewriter/loader")

async function main() {
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
}

main().catch((error) => {
  console.error(error)
  process.exitCode = 1
})

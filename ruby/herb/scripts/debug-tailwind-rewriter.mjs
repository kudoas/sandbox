import fs from "node:fs/promises"
import path from "node:path"
import process from "node:process"
import { pathToFileURL } from "node:url"
import { TailwindClassSorter } from "@herb-tools/tailwind-class-sorter"

async function main() {
  const baseDir = process.cwd()
  const stylesheet = path.resolve("app/assets/tailwind/application.css")
  const jsPath = path.resolve("node_modules/tailwindcss/dist/lib.mjs")
  const cssPath = path.resolve("node_modules/tailwindcss/theme.css")
  const css = await fs.readFile(stylesheet, "utf8")
  const tw = await import(pathToFileURL(jsPath).toString())

  console.log("baseDir:", baseDir)
  console.log("stylesheet:", stylesheet)
  console.log("tailwind js:", jsPath)
  console.log("tailwind css:", cssPath)
  console.log("tailwind exports:", Object.keys(tw))
  console.log("has __unstable__loadDesignSystem:", "__unstable__loadDesignSystem" in tw)

  try {
    const supportsImportsProbe = { value: false }

    try {
      await tw.__unstable__loadDesignSystem('@import "./empty";', {
        loadStylesheet: () => {
          supportsImportsProbe.value = true
          return { base: path.dirname(stylesheet), content: "" }
        }
      })
    } catch (error) {
      console.log("supportsImports probe error:", error?.message || String(error))
    }

    console.log("supportsImports:", supportsImportsProbe.value)

    const tailwindThemeCss = await fs.readFile(cssPath, "utf8")
    const expandedCss = css.replace('@import "tailwindcss";', tailwindThemeCss)
    console.log("expandedCss length:", expandedCss.length)

    const design = await tw.__unstable__loadDesignSystem(expandedCss, {
      base: path.dirname(stylesheet),
      loadStylesheet: async (id, base) => {
        const resolved = id === "tailwindcss" ? cssPath : path.resolve(base, id)
        console.log("loadStylesheet:", { id, base, resolved })

        return {
          base: path.dirname(resolved),
          content: await fs.readFile(resolved, "utf8")
        }
      },
      loadModule: async (id, base, resourceType) => {
        console.log("loadModule:", { id, base, resourceType })
        return { base, module: {} }
      },
      loadPlugin: async (id) => {
        console.log("loadPlugin:", id)
        return () => {}
      },
      loadConfig: async (id) => {
        console.log("loadConfig:", id)
        return {}
      }
    })

    console.log("class order:", design.getClassOrder(["rounded", "bg-blue-500", "px-4", "py-2", "text-white"]))
  } catch (error) {
    console.error("loadDesignSystem failed:")
    console.error(error)
  }

  const sorter = await TailwindClassSorter.fromConfig({
    baseDir,
    tailwindStylesheet: "app/assets/tailwind/application.css"
  })

  console.log("sorter context:", sorter.getContext())
  console.log("sort sample:", sorter.sortClasses("rounded bg-blue-500 px-4 py-2 text-white"))
  console.log("sort sample 2:", sorter.sortClasses("px-4 bg-blue-500 text-white rounded py-2"))
}

main().catch((error) => {
  console.error(error)
  process.exitCode = 1
})

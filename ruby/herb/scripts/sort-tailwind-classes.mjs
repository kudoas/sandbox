import fs from "node:fs/promises"
import path from "node:path"

import { Herb } from "@herb-tools/node-wasm"
import { rewriteString, tailwindClassSorter } from "@herb-tools/rewriter/loader"

const args = process.argv.slice(2)
const check = args.includes("--check")
const targets = args.filter(arg => arg !== "--check")
const baseDir = process.cwd()
const defaultTargets = ["app/views"]
const ignoredDirectories = new Set([".git", "node_modules", "tmp", "log", "storage", "vendor"])
const templatePattern = /(?:\.html(?:\+[^/]+)?\.erb|\.turbo_stream\.erb|\.rhtml)$/
const tailwindStylesheet = "app/assets/tailwind/application.css"

function isTemplateFile(filePath) {
  return templatePattern.test(filePath)
}

async function collectTemplateFiles(target) {
  const resolvedTarget = path.resolve(baseDir, target)
  const stat = await fs.stat(resolvedTarget)

  if (stat.isFile()) {
    return isTemplateFile(resolvedTarget) ? [resolvedTarget] : []
  }

  if (!stat.isDirectory()) {
    return []
  }

  const entries = await fs.readdir(resolvedTarget, { withFileTypes: true })
  const files = await Promise.all(entries.map(async entry => {
    const entryPath = path.join(resolvedTarget, entry.name)

    if (entry.isDirectory()) {
      if (ignoredDirectories.has(entry.name)) {
        return []
      }

      return collectTemplateFiles(entryPath)
    }

    return isTemplateFile(entryPath) ? [entryPath] : []
  }))

  return files.flat()
}

async function main() {
  const fileSet = new Set()

  for (const target of targets.length > 0 ? targets : defaultTargets) {
    const files = await collectTemplateFiles(target)

    for (const file of files) {
      fileSet.add(file)
    }
  }

  const files = [...fileSet].sort()

  if (files.length === 0) {
    console.log("No ERB templates matched.")
    return
  }

  await Herb.load()

  const rewriter = await tailwindClassSorter({
    baseDir,
    tailwindStylesheet
  })
  const changedFiles = []

  for (const filePath of files) {
    const input = await fs.readFile(filePath, "utf8")
    const output = rewriteString(Herb, input, [rewriter], { baseDir, filePath })

    if (output === input) {
      continue
    }

    changedFiles.push(path.relative(baseDir, filePath))

    if (!check) {
      await fs.writeFile(filePath, output, "utf8")
    }
  }

  if (changedFiles.length === 0) {
    console.log(check ? "All ERB templates already have sorted Tailwind classes." : "No Tailwind class changes were needed.")
    return
  }

  const action = check ? "Would update" : "Updated"
  console.log(`${action} ${changedFiles.length} file(s):`)

  for (const filePath of changedFiles) {
    console.log(`- ${filePath}`)
  }

  if (check) {
    process.exitCode = 1
  }
}

main().catch(error => {
  console.error(error)
  process.exit(1)
})

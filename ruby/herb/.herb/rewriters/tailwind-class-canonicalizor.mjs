import { execFileSync } from "node:child_process"
import { existsSync } from "node:fs"
import path from "node:path"

import {
  Visitor,
  getStaticAttributeName,
  isHTMLAttributeNode,
  isHTMLOpenTagNode,
  isLiteralNode
} from "@herb-tools/core"
import { ASTRewriter, asMutable } from "@herb-tools/rewriter"

function parseCanonicalizeOutput(output) {
  const lines = output
    .split(/\r?\n/)
    .filter(Boolean)

  const jsonLines = lines[0]?.startsWith("{")
    ? lines
    : lines.slice(1)

  return jsonLines
    .map(line => JSON.parse(line))
}

class CandidateCollector extends Visitor {
  constructor () {
    super()
    this.candidates = []
  }

  visitHTMLElementNode(node) {
    const openTag = node.open_tag

    if (isHTMLOpenTagNode(openTag)) {
      for (const child of openTag.children) {
        if (!isHTMLAttributeNode(child) || !child.name || !child.value) {
          continue
        }

        const attributeName = getStaticAttributeName(child.name)

        if (attributeName !== "class") {
          continue
        }

        const valueNode = child.value

        if (valueNode.children.length !== 1) {
          continue
        }

        const valueChild = valueNode.children[0]

        if (isHTMLOpenTagNode(openTag) && isLiteralNode(valueChild)) {
          const input = valueChild.content.trim()

          if (!input) {
            continue
          }

          this.candidates.push({
            input,
            apply(output) {
              asMutable(valueChild).content = output
            }
          })
        }
      }
    }

    this.visitChildNodes(node)
  }
}

export default class TailwindClassCanonicalizor extends ASTRewriter {
  get name() {
    return "tailwind-class-canonicalizor"
  }

  get description() {
    return "Canonicalizes static Tailwind classes in HTML class attributes using tailwindcss-ruby"
  }

  async initialize(context) {
    this.baseDir = context.baseDir
    this.command = "bundle"
    this.commandArgs = ["exec", "tailwindcss"]

    const stylesheet = path.join(context.baseDir, "app/assets/tailwind/application.css")
    this.stylesheet = existsSync(stylesheet) ? stylesheet : null
  }

  rewrite(node, context) {
    if (!this.command) {
      return node
    }

    const collector = new CandidateCollector()
    collector.visit(node)

    if (collector.candidates.length === 0) {
      return node
    }

    const args = [...this.commandArgs, "canonicalize", "--stream", "--format", "jsonl"]

    if (this.stylesheet) {
      args.push("--css", this.stylesheet)
    }

    const input = collector.candidates.map(candidate => candidate.input).join("\n")

    let results

    try {
      const stdout = execFileSync(this.command, args, {
        cwd: context.baseDir ?? this.baseDir,
        encoding: "utf8",
        input: `${input}\n`,
        stdio: ["pipe", "pipe", "pipe"]
      })

      results = parseCanonicalizeOutput(stdout)
    } catch {
      return node
    }

    if (results.length !== collector.candidates.length) {
      return node
    }

    for (const [index, candidate] of collector.candidates.entries()) {
      candidate.apply(results[index]?.output ?? candidate.input)
    }

    return node
  }
}

import * as vscode from "vscode";

type TypeScriptModule = typeof import("typescript");

type BlockRange = {
  start: number;
  end: number;
};

type QuickInfoResult = {
  display: string;
  documentation: string;
  textSpan: import("typescript").TextSpan;
};

type CompletionItemData = {
  name: string;
  offset: number;
  source?: string;
};

class TypeScriptCompletionService {
  private readonly ts: TypeScriptModule;
  private readonly fileName = "/virtual/erb-javascript-tag.js";
  private content = "";
  private version = 0;
  private readonly compilerOptions: import("typescript").CompilerOptions;
  private readonly service: import("typescript").LanguageService;

  constructor(ts: TypeScriptModule) {
    this.ts = ts;
    this.compilerOptions = {
      allowJs: true,
      checkJs: false,
      target: ts.ScriptTarget.ES2020,
    };
    this.service = ts.createLanguageService(this.createHost());
  }

  updateContent(content: string): void {
    if (content === this.content) {
      return;
    }
    this.content = content;
    this.version += 1;
  }

  getCompletions(
    offset: number,
    triggerCharacter?: string
  ): import("typescript").CompletionInfo | undefined {
    return this.service.getCompletionsAtPosition(this.fileName, offset, {
      includeCompletionsWithInsertText: true,
      includeAutomaticOptionalChainCompletions: true,
      includeCompletionsForModuleExports: true,
      triggerCharacter: triggerCharacter as any,
    });
  }

  getCompletionDetails(
    name: string,
    offset: number,
    source?: string
  ): import("typescript").CompletionEntryDetails | undefined {
    return this.service.getCompletionEntryDetails(
      this.fileName,
      offset,
      name,
      undefined,
      source,
      undefined,
      undefined
    );
  }

  getQuickInfo(offset: number): QuickInfoResult | null {
    const info = this.service.getQuickInfoAtPosition(this.fileName, offset);
    if (!info) {
      return null;
    }
    const display = this.ts.displayPartsToString(info.displayParts || []);
    const documentation = this.ts.displayPartsToString(
      info.documentation || []
    );
    return { display, documentation, textSpan: info.textSpan };
  }

  getDefinitions(
    offset: number
  ): readonly import("typescript").DefinitionInfo[] | undefined {
    return this.service.getDefinitionAtPosition(this.fileName, offset);
  }

  getVirtualFileName(): string {
    return this.fileName;
  }

  private createHost(): import("typescript").LanguageServiceHost {
    const ts = this.ts;
    return {
      getScriptFileNames: () => [this.fileName],
      getScriptVersion: () => String(this.version),
      getScriptSnapshot: (fileName) => {
        if (fileName === this.fileName) {
          return ts.ScriptSnapshot.fromString(this.content);
        }
        const fileText = ts.sys.readFile(fileName);
        if (fileText === undefined) {
          return undefined;
        }
        return ts.ScriptSnapshot.fromString(fileText);
      },
      getCurrentDirectory: () => process.cwd(),
      getCompilationSettings: () => this.compilerOptions,
      getDefaultLibFileName: (options) => ts.getDefaultLibFilePath(options),
      fileExists: (fileName) =>
        fileName === this.fileName || ts.sys.fileExists(fileName),
      readFile: (fileName) =>
        fileName === this.fileName ? this.content : ts.sys.readFile(fileName),
      readDirectory: ts.sys.readDirectory,
    };
  }
}

class JavaScriptTagCompletionProvider implements vscode.CompletionItemProvider {
  private readonly itemData = new WeakMap<
    vscode.CompletionItem,
    CompletionItemData
  >();

  constructor(private readonly tsService: TypeScriptCompletionService) {}

  provideCompletionItems(
    document: vscode.TextDocument,
    position: vscode.Position,
    token: vscode.CancellationToken
  ): vscode.ProviderResult<vscode.CompletionItem[] | vscode.CompletionList> {
    const text = document.getText();
    const offset = document.offsetAt(position);
    const block = findJavaScriptBlock(text, offset);
    if (!block || token.isCancellationRequested) {
      return undefined;
    }

    const jsContent = text.slice(block.start, block.end);
    const jsOffset = offset - block.start;
    const lastChar = jsOffset > 0 ? jsContent[jsOffset - 1] : undefined;
    const triggerCharacter = lastChar === "." ? "." : undefined;

    this.tsService.updateContent(jsContent);
    const completions = this.tsService.getCompletions(
      jsOffset,
      triggerCharacter
    );
    if (!completions || token.isCancellationRequested) {
      return undefined;
    }

    const items = completions.entries.map((entry) =>
      mapCompletionEntry(entry, jsOffset, this.itemData)
    );
    return new vscode.CompletionList(items, completions.isIncomplete);
  }

  resolveCompletionItem(
    item: vscode.CompletionItem,
    token: vscode.CancellationToken
  ): vscode.ProviderResult<vscode.CompletionItem> {
    const data = this.itemData.get(item);
    if (!data || token.isCancellationRequested) {
      return item;
    }

    const details = this.tsService.getCompletionDetails(
      data.name,
      data.offset,
      data.source
    );
    if (!details || token.isCancellationRequested) {
      return item;
    }

    const detail = details.displayParts
      ? details.displayParts.map((part) => part.text).join("")
      : "";
    const documentation = details.documentation
      ? details.documentation.map((part) => part.text).join("")
      : "";

    if (detail) {
      item.detail = detail;
    }
    if (documentation) {
      item.documentation = new vscode.MarkdownString(documentation);
    }

    return item;
  }
}

class JavaScriptTagHoverProvider implements vscode.HoverProvider {
  constructor(private readonly tsService: TypeScriptCompletionService) {}

  provideHover(
    document: vscode.TextDocument,
    position: vscode.Position,
    token: vscode.CancellationToken
  ): vscode.ProviderResult<vscode.Hover> {
    const text = document.getText();
    const offset = document.offsetAt(position);
    const block = findJavaScriptBlock(text, offset);
    if (!block || token.isCancellationRequested) {
      return undefined;
    }

    const jsContent = text.slice(block.start, block.end);
    const jsOffset = offset - block.start;
    this.tsService.updateContent(jsContent);

    const quickInfo = this.tsService.getQuickInfo(jsOffset);
    if (!quickInfo || token.isCancellationRequested) {
      return undefined;
    }

    const rangeStart = document.positionAt(
      block.start + quickInfo.textSpan.start
    );
    const rangeEnd = document.positionAt(
      block.start + quickInfo.textSpan.start + quickInfo.textSpan.length
    );
    const range = new vscode.Range(rangeStart, rangeEnd);

    const markdown = new vscode.MarkdownString();
    if (quickInfo.display) {
      markdown.appendCodeblock(quickInfo.display, "javascript");
    }
    if (quickInfo.documentation) {
      markdown.appendMarkdown(`\n\n${quickInfo.documentation}`);
    }

    return new vscode.Hover(markdown, range);
  }
}

class JavaScriptTagDefinitionProvider implements vscode.DefinitionProvider {
  constructor(private readonly tsService: TypeScriptCompletionService) {}

  async provideDefinition(
    document: vscode.TextDocument,
    position: vscode.Position,
    token: vscode.CancellationToken
  ): Promise<vscode.Definition | undefined> {
    const text = document.getText();
    const offset = document.offsetAt(position);
    const block = findJavaScriptBlock(text, offset);
    if (!block || token.isCancellationRequested) {
      return undefined;
    }

    const jsContent = text.slice(block.start, block.end);
    const jsOffset = offset - block.start;
    this.tsService.updateContent(jsContent);

    const definitions = this.tsService.getDefinitions(jsOffset);
    if (!definitions || token.isCancellationRequested) {
      return undefined;
    }

    const virtualFileName = this.tsService.getVirtualFileName();
    const locations = await Promise.all(
      definitions.map(async (definition) => {
        if (definition.fileName === virtualFileName) {
          const range = new vscode.Range(
            document.positionAt(block.start + definition.textSpan.start),
            document.positionAt(
              block.start +
                definition.textSpan.start +
                definition.textSpan.length
            )
          );
          return new vscode.Location(document.uri, range);
        }

        const targetUri = vscode.Uri.file(definition.fileName);
        const targetDoc = await vscode.workspace.openTextDocument(targetUri);
        const range = new vscode.Range(
          targetDoc.positionAt(definition.textSpan.start),
          targetDoc.positionAt(
            definition.textSpan.start + definition.textSpan.length
          )
        );
        return new vscode.Location(targetUri, range);
      })
    );

    return locations;
  }
}

function findJavaScriptBlock(text: string, offset: number): BlockRange | null {
  const erbBlock = findJavascriptTagBlock(text, offset);
  if (erbBlock) {
    return erbBlock;
  }

  return findScriptTagBlock(text, offset);
}

function findJavascriptTagBlock(
  text: string,
  offset: number
): BlockRange | null {
  const tagPattern =
    /<%[-=]?\s*javascript_tag\b[^%]*\bdo\b[^%]*-?%>|<%\s*end\s*-?%>/g;
  const stack: Array<{ index: number; length: number }> = [];
  let match: RegExpExecArray | null;

  while ((match = tagPattern.exec(text))) {
    const token = match[0];
    if (token.includes("javascript_tag")) {
      stack.push({ index: match.index, length: token.length });
      continue;
    }

    if (!stack.length) {
      continue;
    }

    const begin = stack.pop();
    if (!begin) {
      continue;
    }

    const start = begin.index + begin.length;
    const end = match.index;

    if (offset >= start && offset <= end) {
      return { start, end };
    }
  }

  return null;
}

function findScriptTagBlock(text: string, offset: number): BlockRange | null {
  const tagPattern = /<script\b[^>]*>|<\/script\s*>/gi;
  const stack: Array<{ index: number; length: number; isJavaScript: boolean }> =
    [];
  let match: RegExpExecArray | null;

  while ((match = tagPattern.exec(text))) {
    const token = match[0];
    const isOpenTag = /^<script\b/i.test(token) && !/^<\/script/i.test(token);
    if (isOpenTag) {
      stack.push({
        index: match.index,
        length: token.length,
        isJavaScript: isJavaScriptScriptTag(token),
      });
      continue;
    }

    if (!stack.length) {
      continue;
    }

    const begin = stack.pop();
    if (!begin || !begin.isJavaScript) {
      continue;
    }

    const start = begin.index + begin.length;
    const end = match.index;

    if (offset >= start && offset <= end) {
      return { start, end };
    }
  }

  return null;
}

function isJavaScriptScriptTag(tag: string): boolean {
  const typeMatch = tag.match(
    /\btype\s*=\s*(\"([^\"]*)\"|'([^']*)'|([^\s>]+))/i
  );
  if (!typeMatch) {
    return true;
  }

  const value = (typeMatch[2] || typeMatch[3] || typeMatch[4] || "")
    .trim()
    .toLowerCase();
  if (!value) {
    return true;
  }

  return (
    value === "module" ||
    value === "text/js" ||
    value.includes("javascript") ||
    value.includes("ecmascript")
  );
}

function mapCompletionEntry(
  entry: import("typescript").CompletionEntry,
  offset: number,
  itemData: WeakMap<vscode.CompletionItem, CompletionItemData>
): vscode.CompletionItem {
  const item = new vscode.CompletionItem(entry.name);
  item.sortText = entry.sortText;
  item.filterText = entry.name;
  itemData.set(item, {
    name: entry.name,
    offset,
    source: entry.source,
  });

  if (entry.insertText) {
    if (entry.isSnippet) {
      item.insertText = new vscode.SnippetString(entry.insertText);
    } else {
      item.insertText = entry.insertText;
    }
  }

  return item;
}

export function activate(context: vscode.ExtensionContext): void {
  let ts: TypeScriptModule;
  try {
    ts = require("typescript");
  } catch (error) {
    const message = error instanceof Error ? error.message : String(error);
    vscode.window.showErrorMessage(
      `erb-javascript-tag: typescript not found. Run npm install. (${message})`
    );
    return;
  }

  const tsService = new TypeScriptCompletionService(ts);
  const completionProvider = new JavaScriptTagCompletionProvider(tsService);
  const hoverProvider = new JavaScriptTagHoverProvider(tsService);
  const definitionProvider = new JavaScriptTagDefinitionProvider(tsService);
  context.subscriptions.push(
    vscode.languages.registerCompletionItemProvider(
      { language: "erb" },
      completionProvider,
      "."
    )
  );
  context.subscriptions.push(
    vscode.languages.registerHoverProvider({ language: "erb" }, hoverProvider)
  );
  context.subscriptions.push(
    vscode.languages.registerDefinitionProvider(
      { language: "erb" },
      definitionProvider
    )
  );
}

export function deactivate(): void {}

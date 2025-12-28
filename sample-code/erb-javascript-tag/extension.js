const vscode = require('vscode');

class TypeScriptCompletionService {
  constructor (ts) {
    this._ts = ts;
    this._fileName = '/virtual/erb-javascript-tag.js';
    this._content = '';
    this._version = 0;
    this._compilerOptions = {
      allowJs: true,
      checkJs: false,
      target: ts.ScriptTarget.ES2020
    };
    this._service = ts.createLanguageService(this._createHost());
  }

  updateContent(content) {
    if (content === this._content) {
      return;
    }
    this._content = content;
    this._version += 1;
  }

  getCompletions(offset, triggerCharacter) {
    return this._service.getCompletionsAtPosition(
      this._fileName,
      offset,
      {
        includeCompletionsWithInsertText: true,
        includeAutomaticOptionalChainCompletions: true,
        includeCompletionsForModuleExports: true,
        triggerCharacter
      }
    );
  }

  _createHost() {
    const ts = this._ts;
    return {
      getScriptFileNames: () => [this._fileName],
      getScriptVersion: () => String(this._version),
      getScriptSnapshot: (fileName) => {
        if (fileName === this._fileName) {
          return ts.ScriptSnapshot.fromString(this._content);
        }
        const fileText = ts.sys.readFile(fileName);
        if (fileText === undefined) {
          return undefined;
        }
        return ts.ScriptSnapshot.fromString(fileText);
      },
      getCurrentDirectory: () => process.cwd(),
      getCompilationSettings: () => this._compilerOptions,
      getDefaultLibFileName: (options) => ts.getDefaultLibFilePath(options),
      fileExists: (fileName) => fileName === this._fileName || ts.sys.fileExists(fileName),
      readFile: (fileName) => (fileName === this._fileName ? this._content : ts.sys.readFile(fileName)),
      readDirectory: ts.sys.readDirectory
    };
  }
}

class JavaScriptTagCompletionProvider {
  constructor (tsService) {
    this._tsService = tsService;
  }

  provideCompletionItems(document, position, token) {
    const text = document.getText();
    const offset = document.offsetAt(position);
    const block = findJavascriptTagBlock(text, offset);
    if (!block || token.isCancellationRequested) {
      return undefined;
    }

    const jsContent = text.slice(block.start, block.end);
    const jsOffset = offset - block.start;
    const triggerCharacter = jsOffset > 0 ? jsContent[jsOffset - 1] : undefined;

    this._tsService.updateContent(jsContent);
    const completions = this._tsService.getCompletions(jsOffset, triggerCharacter);
    if (!completions || token.isCancellationRequested) {
      return undefined;
    }

    const items = completions.entries.map((entry) => mapCompletionEntry(entry));
    return new vscode.CompletionList(items, completions.isIncomplete);
  }
}

function findJavascriptTagBlock(text, offset) {
  const tagPattern = /<%[-=]?\s*javascript_tag\b[^%]*\bdo\b[^%]*-?%>|<%\s*end\s*-?%>/g;
  const stack = [];
  let match;

  while ((match = tagPattern.exec(text))) {
    const token = match[0];
    if (token.includes('javascript_tag')) {
      stack.push({ index: match.index, length: token.length });
      continue;
    }

    if (!stack.length) {
      continue;
    }

    const begin = stack.pop();
    const start = begin.index + begin.length;
    const end = match.index;

    if (offset >= start && offset <= end) {
      return { start, end };
    }
  }

  return null;
}

function mapCompletionEntry(entry) {
  const item = new vscode.CompletionItem(entry.name, mapCompletionKind(entry.kind));
  item.sortText = entry.sortText;
  item.filterText = entry.name;

  if (entry.insertText) {
    if (entry.isSnippet) {
      item.insertText = new vscode.SnippetString(entry.insertText);
    } else {
      item.insertText = entry.insertText;
    }
  }

  return item;
}

function mapCompletionKind(kind) {
  switch (kind) {
    case 'primitive type':
    case 'keyword':
      return vscode.CompletionItemKind.Keyword;
    case 'variable':
    case 'local var':
    case 'member var':
    case 'const':
    case 'let':
      return vscode.CompletionItemKind.Variable;
    case 'function':
      return vscode.CompletionItemKind.Function;
    case 'member function':
    case 'method':
      return vscode.CompletionItemKind.Method;
    case 'class':
      return vscode.CompletionItemKind.Class;
    case 'interface':
      return vscode.CompletionItemKind.Interface;
    case 'enum':
      return vscode.CompletionItemKind.Enum;
    case 'module':
      return vscode.CompletionItemKind.Module;
    case 'property':
      return vscode.CompletionItemKind.Property;
    default:
      return vscode.CompletionItemKind.Text;
  }
}

function activate(context) {
  let ts;
  try {
    ts = require('typescript');
  } catch (error) {
    vscode.window.showErrorMessage('erb-javascript-tag: typescript not found. Run npm install.');
    return;
  }

  const tsService = new TypeScriptCompletionService(ts);
  const completionProvider = new JavaScriptTagCompletionProvider(tsService);
  context.subscriptions.push(
    vscode.languages.registerCompletionItemProvider(
      { language: 'erb' },
      completionProvider,
      '.'
    )
  );
}

function deactivate() { }

module.exports = { activate, deactivate };

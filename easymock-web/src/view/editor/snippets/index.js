/* global ace */
import mockSnippets from 'mock.snippets'
import javascriptSnippets from './javascript.snippets'
ace.define('ace/snippets/javascript', ['require', 'exports', 'module'], function (e, t) {
  t.snippetText = javascriptSnippets + '\n' + mockSnippets
  t.scope = 'javascript'
})

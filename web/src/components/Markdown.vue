<template lang="pug">
  .markdown( v-html="replaceNewlines(md($slots.default[0].text.replace(/\\n/g, ':NEWLINE:')))" )
</template>

<script>
import { parse as Marked } from 'marked'

import { 
  languages,
  highlight,
} from 'prismjs'

Marked.setOptions({
  highlight (string, language) {
    return language && languages[language]  
      ? highlight(string, languages[language])
      : string
  },
})

export default {
  methods: {
    md: string => Marked(string),
    replaceNewlines: string => string.replace(/:NEWLINE:/g, '\<span class=\"newline\"\>\<\/span\>'),
  },
}
</script>

<style lang="scss" scoped>
  @import '~css/prism.css';
</style>

<style lang="scss">
  .markdown {
    width: 100%;

    * {
      white-space: pre-wrap;
    }

    .newline {
      display: block;
      height: 7.5px;

      // &:last-child {
      //   display: none;
      // }
    }

    pre, code {
      * {
        font-family: 'Fira Code';
      }
    }

    a {
      color: var(--accent);
      text-decoration: underline;

      &:hover {
        color: rgba(white, 0.8)
      }
    }
  }
</style>

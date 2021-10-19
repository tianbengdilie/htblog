<template>
  <div>
    <markdown-it-vue class="md-body" :content="content" />
  </div>
</template>

<script>
import MarkdownItVue from "markdown-it-vue";
import "markdown-it-vue/dist/markdown-it-vue.css";
import text from "./text";
export default {
  components: {
    MarkdownItVue
  },
  mounted() {
    var MarkdownIt = require("markdown-it");
    let md = new MarkdownIt();
    var result = md.render("# markdown-it rulezz!");
    console.log(result);
  },
  data() {
    return {
      content: text
    };
  }
};
</script>
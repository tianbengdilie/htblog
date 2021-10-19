<template>
  <div>
    <div v-html="renderContent"></div>
  </div>
</template>

<script>
import "markdown-it-vue/dist/markdown-it-vue.css";
import md from "./md-render";
import text from "./text";
export default {
  mounted() {},
  data() {
    return {
      content: text
    };
  },
  computed: {
    renderContent: function() {
      console.log(md);
      return md.render(this.content);
    }
  }
};
</script>
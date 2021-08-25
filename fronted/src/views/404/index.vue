<template>
  <editor
    ref="toastuiEditor"
    :initialValue="editorText"
    previewStyle="vertical"
    height="100%"
    @change="onEditorChange"
  />
</template>

<script>
import "@toast-ui/editor/dist/toastui-editor.css";
import "highlight.js/styles/github.css";
import { Editor, Viewer } from "@toast-ui/vue-editor";

export default {
  name: "dashboard",
  components: { editor: Editor, viewer: Viewer },
  data: () => {
    return {
      editorText: "# title 11111"
    };
  },
  methods: {
    onEditorChange: function() {
      let markdown = this.$refs.toastuiEditor.invoke("getMarkdown");
      console.log(markdown);
    }
  }
};
</script>

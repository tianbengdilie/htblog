<template>
  <v-container fluid class="flex-column overflow-y-auto">
    <v-row class="feature-panel align-self-stretch overflow-y-auto rounded-lg">
      <v-col cols="6" v-show="editMode">
        <quill-editor @change="changeText" ref="editor"></quill-editor>
      </v-col>
      <v-col :cols="editMode ? 6 : 12">
        <markdown-it-vue class="md-body" :content="nativeContent" />
      </v-col>
    </v-row>
    <v-row justify="space-around">
      <v-col>
        <v-layout justify-space-around>
          <template v-if="editMode">
            <v-btn depressed color="primary" @click.stop="exitEditMode(true)">
              保存
            </v-btn>
            <v-btn depressed color="error" @click.stop="exitEditMode(false)">
              退出
            </v-btn>
          </template>
          <template v-else>
            <v-btn depressed color="error" @click.stop="changeEdit">
              编辑
            </v-btn>
          </template>
        </v-layout>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import MarkdownItVue from "markdown-it-vue";
import "markdown-it-vue/dist/markdown-it-vue.css";
import text from "@/plugins/md-render/example";
import "quill/dist/quill.core.css";
import "quill/dist/quill.snow.css";
import "quill/dist/quill.bubble.css";

import { quillEditor } from "vue-quill-editor";

export default {
  components: {
    MarkdownItVue,
    quillEditor
  },
  data() {
    return {
      content: text,
      nativeContent: text,
      editMode: false
    };
  },
  mounted() {
    this.initEditor();
  },
  methods: {
    uploadImage(p1) {
      console.log(p1);
    },
    changeText(delta) {
      this.nativeContent = delta.text;
    },
    changeEdit() {
      this.editMode = true;
    },
    initEditor() {
      this.$refs.editor.quill.setText(this.content);
    },
    exitEditMode(saved) {
      this.editMode = false;
      if (saved) {
        console.log("TODO saved opera");
      }
    }
  }
};
</script>

<style lang="less" >
.feature-panel {
  border: 1px solid rgb(152, 175, 149);
  .ql-toolbar {
    white-space: nowrap;
    overflow-y: auto;
  }

  .md-body {
    border: 1px solid #ccc;
    padding: 2px 8px;
  }
}

.control-panel {
  width: 60px;
}
</style>
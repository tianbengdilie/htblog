<template>
  <v-container
    fluid
    class="content-display grow d-flex flex-column flex-nowrap"
  >
    <v-row class="grow full-width">
      <v-col cols="12" class="rounded-lg content-area fill-height">
        <template v-if="editMode">
          <editor
            ref="toastuiEditor"
            :initialValue="text"
            previewStyle="vertical"
            height="100%"
            @change="onEditorChange"
          />
        </template>
        <template v-else>
          <viewer
            ref="toastuiViewer"
            :initialValue="text"
            previewStyle="vertical"
            height="100%"
          />
        </template>
      </v-col>
    </v-row>
    <v-row class="shrink" justify="space-around">
      <v-col>
        <v-layout justify-space-around>
          <template v-if="editMode">
            <v-btn depressed color="primary"> 保存 </v-btn>
            <v-btn depressed color="error" @click.stop="editMode = false">
              退出
            </v-btn>
          </template>
          <template v-else>
            <v-btn depressed color="error" @click.stop="editMode = true">
              编辑
            </v-btn>
          </template>
        </v-layout>
      </v-col>
    </v-row>
  </v-container>
</template>

<style lang="less">
.content-display {
  //   .row {
  //     margin: 0, 10px;
  //   }

  .content-area {
    border: black solid 1px;
  }
}
</style>

<script>
import "@toast-ui/editor/dist/toastui-editor.css";
import "highlight.js/styles/github.css";
import { Editor, Viewer } from "@toast-ui/vue-editor";

export default {
  name: "dashboard",
  components: { editor: Editor, viewer: Viewer },
  data: () => {
    return {
      text: "# title 11111fjo",
      editMode: false
    };
  },
  methods: {
    onEditorChange: function() {
      this.text = this.$refs.toastuiEditor.invoke("getMarkdown");
    }
  }
};
</script>

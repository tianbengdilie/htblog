<template>
  <v-treeview
    hoverable
    activatable
    rounded
    dense
    class="layout-drawer"
    :items="dirs"
    @update:active="test"
  >
  </v-treeview>
</template>

<script>
import { resolve, join } from "path";
import { dirs as data } from "@/api/directory";

export default {
  name: "TheLayoutDrawerList",
  props: {
    routes: {
      type: Array,
      default: () => {}
    },
    basePath: {
      type: String,
      default: "/blog"
    }
  },
  data() {
    return {};
  },
  computed: {
    dirs: () => data
  },
  methods: {
    test(item) {
      if (item.length < 1) return;
      let nextPath = join(this.basePath, "/id");
      if (nextPath === this.$router.currentRoute.path) return;
      this.$router.push({ path: nextPath });

      // console.log(resolve(this.$router.));
    },
    isExternal(path) {
      return /^(https?:|mailto:|tel:)/.test(path);
    },
    isVisibleItem(item) {
      return (
        this.hasOneVisibleChild(item.children, item) &&
        (!this.onlyOneChild.children || this.onlyOneChild.noVisibleChildren) &&
        !item.alwaysShow
      );
    },
    hasOneVisibleChild(children = [], parent) {
      const visibleChildren = children.filter(item => {
        if (item.hidden) return false;
        // Temp set(will be used if only has one visible child)
        this.onlyOneChild = item;
        return true;
      });

      // When there is only one child router, the child router is displayed by default
      if (visibleChildren.length === 1) {
        this.onlyOneChild.path = resolve(parent.path, this.onlyOneChild.path);
        this.onlyOneChild.meta.icon =
          this.onlyOneChild.meta.icon || parent.meta.icon || "";

        return true;
      }

      // Show parent if there are no child router to display
      if (visibleChildren.length === 0) {
        this.onlyOneChild = { ...parent, noVisibleChildren: true };
        return true;
      }

      return false;
    },
    resolvePath(path) {
      if (this.isExternal(path)) {
        return path;
      }
      return resolve(this.basePath, path);
    },
    getListIcon(item) {
      return this.iconShow && item.meta ? item.meta.icon : " ";
    },
    getListTitle(item) {
      return item.meta ? this.$t(item.meta.title) : "";
    }
  }
};
</script>

<style>
.layout-drawer {
  padding-bottom: 0px;
  padding-top: 0px;
}
.layout-drawer__icon {
  margin-bottom: 8px;
  margin-top: 8px;
}
</style>

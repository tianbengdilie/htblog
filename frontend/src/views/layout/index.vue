<template>
  <div class="v-application--wrap">
    <the-layout-drawer />
    <the-layout-toolbar />
    <v-main>
      <transition name="fade-transform" mode="out-in">
        <keep-alive>
          <router-view />
        </keep-alive>
      </transition>
    </v-main>
    <!-- <the-layout-fab /> -->
    <!-- <the-layout-to-top-fab /> -->
    <the-layout-footer />
  </div>
</template>

<script>
import TheLayoutDrawer from "./TheLayoutDrawer.vue";
import TheLayoutToolbar from "./TheLayoutToolbar.vue";
// import TheLayoutFab from "./TheLayoutFab.vue";
// import TheLayoutToTopFab from "./TheLayoutToTopFab.vue";
import TheLayoutFooter from "./TheLayoutFooter.vue";

export default {
  name: "TheLayout",
  components: {
    TheLayoutDrawer,
    TheLayoutToolbar,
    // TheLayoutFab,
    // TheLayoutToTopFab,
    TheLayoutFooter
  },
  data: () => ({})
};
</script>

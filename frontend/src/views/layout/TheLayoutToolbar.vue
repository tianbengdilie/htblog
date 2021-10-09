<template>
  <v-app-bar app dense>
    <v-app-bar-nav-icon @click.stop="toggleNavbar">
      <v-icon>{{ toggleNavbarIcon }}</v-icon>
    </v-app-bar-nav-icon>
    <breadcrumbs />
    <v-spacer />
    <error-log />
    <full-screen />
    <notification />
    <profile />
  </v-app-bar>
</template>

<script>
import { mapGetters } from "vuex";
import Breadcrumbs from "../widget/AppBreadcrumbs.vue";
import ErrorLog from "../widget/AppErrorLog.vue";
import FullScreen from "../widget/AppFullScreen.vue";
import Notification from "../widget/AppNotification.vue";
import Profile from "../widget/AppProfile.vue";

export default {
  name: "TheLayoutToolbar",
  components: {
    Breadcrumbs,
    ErrorLog,
    FullScreen,
    Notification,
    Profile
  },
  data: () => ({}),
  computed: {
    ...mapGetters(["navbarShow"]),
    toggleNavbarIcon() {
      return this.navbarShow
        ? "mdi-format-indent-decrease"
        : "mdi-format-indent-increase";
    }
  },
  methods: {
    toggleNavbar() {
      this.$store.dispatch("NavbarToggle");
    }
  }
};
</script>

<template>
  <v-container class="container--fluid fill-height primary">
    <v-row no-gutters align="center" justify="center">
      <v-col cols="12" sm="8" md="4" lg="4">
        <v-card class="elevation-5 pa-3">
          <v-card-text>
            <div class="layout column align-center">
              <img src="img/icons/error.png" width="120" height="120" />
              <h1 class="display-2 primary--text">
                {{ `${$t("errors.whoops")}, ${errorCode}` }}
              </h1>
              <p>{{ $t(`errors.${errorCode}`) }}</p>
              <v-btn outlined color="primary" @click="$router.go(-1)">
                {{ $t("errors.back") }}
              </v-btn>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "error404",
  components: {},
  props: {
    errorCode: {
      type: Number,
      default: 404
    }
  }
};
</script>
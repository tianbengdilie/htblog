<template>
  <v-container fuild>
    <v-row class="fill-height" justify="center" align="center">
      <v-col col="12" class="fill-height">
        <v-card flat>
          <v-card-title primary-title>
            <h4>Login</h4>
          </v-card-title>
          <v-form ref="loginForm">
            <v-text-field
              label="Account"
              v-model="loginForm.account"
              :counter="10"
            ></v-text-field>
            <v-text-field
              label="Password"
              type="password"
              v-model="loginForm.password"
              :counter="10"
            ></v-text-field>
            <v-card-actions class="justify-space-around">
              <v-btn color="primary" large @click.stop="login">Login</v-btn>
              <v-btn large @click.stop="registerDialog = true">
                Register
              </v-btn>
              <v-dialog v-model="registerDialog">
                <register-card @completed="registerDialog = false" />
              </v-dialog>
            </v-card-actions>
          </v-form>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapActions } from "vuex";
import { register } from "@/api/user";
import registerCard from "@/components/register-card";

export default {
  name: "loginpage",
  components: { registerCard },
  data: () => {
    return {
      loginForm: {
        account: "dev",
        password: "dev"
      },
      registerDialog: false
    };
  },
  methods: {
    ...mapActions("user", ["LoginIn"]),
    login() {
      console.log("here");
      this.LoginIn(this.loginForm).then(data => {
        console.log("login success : " + data);
      });
    },
    onClickRegister() {
      register(this.loginForm).then(data => {
        console.log("register success : " + data);
      });
    }
  }
};
</script>